package task_image

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	stdimage "image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"time"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/fastapi-admin/v2/internal/config"
	"github.com/iimeta/fastapi-admin/v2/internal/consts"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/internal/errors"
	"github.com/iimeta/fastapi-admin/v2/internal/model/entity"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	smodel "github.com/iimeta/fastapi-sdk/v2/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func loadImageSizeCheckGroup(ctx context.Context, logImage *entity.LogImage) *entity.Group {

	if logImage == nil || logImage.Spend.GroupId == "" {
		return nil
	}

	group, err := dao.Group.FindById(ctx, logImage.Spend.GroupId)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(ctx, err)
		}
		return nil
	}

	return group
}

func imageSizeCheckEnabled(group *entity.Group) bool {
	return group != nil && group.IsEnableImageSizeCheck
}

func imageSizeCheckMaxRetry(group *entity.Group, agentTotal int) int {

	if group != nil && group.ImageSizeCheckRetry > 0 {
		return group.ImageSizeCheckRetry
	}

	errRetry := 0
	if config.Cfg != nil && config.Cfg.Base != nil {
		errRetry = config.Cfg.Base.ErrRetry
	}

	if errRetry > 0 {
		return errRetry
	}
	if errRetry < 0 {
		return agentTotal
	}

	return 0
}

func expectedSizeFromTask(taskImage *entity.TaskImage) (width, height int, ok bool) {

	if taskImage == nil {
		return 0, 0, false
	}

	if w, h := parseSizeWH(taskImage.Size); w > 0 && h > 0 {
		return w, h, true
	}

	quality, aspectRatio := splitQualityRatio(taskImage.Size)
	if quality == "" {
		quality = taskImage.Quality
	}
	if taskImage.RequestData != nil {
		if quality == "" {
			if v, _ := taskImage.RequestData["quality"].(string); v != "" {
				quality = v
			}
			if v, _ := taskImage.RequestData["size"].(string); gstr.HasSuffix(v, "K") {
				quality = v
			}
		}
		if aspectRatio == "" {
			if v, _ := taskImage.RequestData["aspect_ratio"].(string); v != "" {
				aspectRatio = v
			}
		}
	}

	return expectedGoogleImageSize(quality, aspectRatio)
}

func expectedGoogleImageSize(imageSize, aspectRatio string) (width, height int, ok bool) {

	imageSize = gstr.Trim(imageSize)
	aspectRatio = gstr.Trim(aspectRatio)

	if imageSize == "" && aspectRatio == "" {
		return 0, 0, false
	}

	if aspectRatio == "" {
		aspectRatio = "1:1"
	}
	if imageSize == "" {
		imageSize = "1K"
	}
	if !gstr.HasSuffix(imageSize, "K") {
		return 0, 0, false
	}

	size := consts.RESOLUTION_ASPECT_RATIO[imageSize+aspectRatio]
	width, height = parseSizeWH(size)
	if width > 0 && height > 0 {
		return width, height, true
	}

	return 0, 0, false
}

func splitQualityRatio(size string) (quality, aspectRatio string) {

	size = gstr.Trim(size)
	if size == "" {
		return "", ""
	}

	parts := gstr.Split(size, " ")
	if len(parts) == 2 {
		return gstr.Trim(parts[0]), gstr.Trim(parts[1])
	}

	if gstr.HasSuffix(size, "K") {
		return size, ""
	}

	return "", ""
}

func parseSizeWH(size string) (width, height int) {

	size = gstr.Trim(size)
	if size == "" {
		return 0, 0
	}

	size = gstr.ReplaceByMap(size, map[string]string{
		"×": "x",
		"X": "x",
		"*": "x",
	})

	parts := gstr.Split(size, "x")
	if len(parts) == 2 {
		return gconv.Int(gstr.Trim(parts[0])), gconv.Int(gstr.Trim(parts[1]))
	}

	return 0, 0
}

func imageSizeMatches(expectedW, expectedH, actualW, actualH int) bool {
	return (actualW == expectedW && actualH == expectedH) || (actualW == expectedH && actualH == expectedW)
}

func (s *sTaskImage) checkGeneratedImageSize(ctx context.Context, taskImage *entity.TaskImage, response smodel.ImageResponse, timeout time.Duration) error {

	expectedW, expectedH, ok := expectedSizeFromTask(taskImage)
	if !ok {
		return nil
	}

	data := imageDataForSizeCheck(response)
	if len(data) == 0 {
		return errors.New("Generated image size mismatch: no image in response.")
	}

	for i, item := range data {

		imageBytes, err := s.imageBytesFromData(ctx, item, timeout)
		if err != nil {
			logger.Errorf(ctx, "sTaskImage checkGeneratedImageSize image[%d] read error: %v", i, err)
			return errors.Newf("Generated image size mismatch: failed to read image %d: %s.", i, err.Error())
		}

		actualW, actualH, err := decodeImageDimension(imageBytes)
		if err != nil {
			logger.Errorf(ctx, "sTaskImage checkGeneratedImageSize image[%d] decode error: %v", i, err)
			return errors.Newf("Generated image size mismatch: failed to decode image %d: %s.", i, err.Error())
		}

		if imageSizeMatches(expectedW, expectedH, actualW, actualH) {
			continue
		}

		logger.Infof(ctx, "sTaskImage checkGeneratedImageSize image[%d] mismatch: expected %dx%d, got %dx%d", i, expectedW, expectedH, actualW, actualH)
		return errors.Newf("Generated image size mismatch: expected %dx%d, got %dx%d.", expectedW, expectedH, actualW, actualH)
	}

	return nil
}

func imageDataForSizeCheck(response smodel.ImageResponse) []smodel.ImageResponseData {

	if len(response.Data) > 0 {
		return response.Data
	}

	if len(response.ResponseBytes) == 0 {
		return nil
	}

	var job smodel.ImageJobResponse
	if err := json.Unmarshal(response.ResponseBytes, &job); err != nil {
		return nil
	}

	if len(job.Data) > 0 {
		return job.Data
	}

	var data []smodel.ImageResponseData
	if len(job.ImageUrls) > 0 {
		for _, u := range job.ImageUrls {
			if u != "" {
				data = append(data, smodel.ImageResponseData{Url: u})
			}
		}
	} else if job.ImageUrl != "" {
		data = append(data, smodel.ImageResponseData{Url: job.ImageUrl})
	}

	return data
}

func (s *sTaskImage) imageBytesFromData(ctx context.Context, item smodel.ImageResponseData, timeout time.Duration) ([]byte, error) {

	if payload := gstr.Trim(item.B64Json); payload != "" {
		if isHTTPURL(payload) {
			body, _, err := s.downloadImage(ctx, payload, timeout)
			return body, err
		}
		return decodeImagePayload(payload)
	}

	if payload := gstr.Trim(item.Url); payload != "" {
		if gstr.HasPrefix(payload, "data:") {
			return decodeImagePayload(payload)
		}
		if isHTTPURL(payload) {
			body, _, err := s.downloadImage(ctx, payload, timeout)
			return body, err
		}
		return decodeImagePayload(payload)
	}

	return nil, errors.New("image url and b64_json are empty")
}

func decodeImagePayload(payload string) ([]byte, error) {

	if gstr.HasPrefix(payload, "data:") {
		return decodeDataURI(payload)
	}

	decoded, err := base64.StdEncoding.DecodeString(payload)
	if err == nil {
		return decoded, nil
	}

	compact := gstr.ReplaceByMap(payload, map[string]string{
		"\n": "",
		"\r": "",
		" ":  "",
	})
	return base64.StdEncoding.DecodeString(compact)
}

func isHTTPURL(s string) bool {
	return gstr.HasPrefix(s, "http://") || gstr.HasPrefix(s, "https://")
}

func decodeImageDimension(imageBytes []byte) (width, height int, err error) {

	cfg, _, err := stdimage.DecodeConfig(bytes.NewReader(imageBytes))
	if err == nil {
		return cfg.Width, cfg.Height, nil
	}

	if w, h, ok := decodeWebPConfig(imageBytes); ok {
		return w, h, nil
	}

	return 0, 0, err
}

func decodeWebPConfig(b []byte) (width, height int, ok bool) {

	if len(b) < 30 {
		return 0, 0, false
	}

	if string(b[0:4]) != "RIFF" || string(b[8:12]) != "WEBP" {
		return 0, 0, false
	}

	switch string(b[12:16]) {
	case "VP8X":
		width = 1 + int(b[24]) + int(b[25])<<8 + int(b[26])<<16
		height = 1 + int(b[27]) + int(b[28])<<8 + int(b[29])<<16
		return width, height, width > 0 && height > 0
	case "VP8 ":
		width = int(b[26]) | int(b[27]&0x3f)<<8
		height = int(b[28]) | int(b[29]&0x3f)<<8
		return width, height, width > 0 && height > 0
	case "VP8L":
		if b[20] != 0x2f {
			return 0, 0, false
		}
		bits := uint32(b[21]) | uint32(b[22])<<8 | uint32(b[23])<<16 | uint32(b[24])<<24
		width = int(bits&0x3fff) + 1
		height = int((bits>>14)&0x3fff) + 1
		return width, height, width > 0 && height > 0
	default:
		return 0, 0, false
	}
}
