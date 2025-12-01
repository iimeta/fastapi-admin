package task_video

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/db"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/util"
	"go.mongodb.org/mongo-driver/bson"
)

type sTaskVideo struct{}

func init() {
	service.RegisterTaskVideo(New())
}

func New() service.ITaskVideo {
	return &sTaskVideo{}
}

// 视频任务分页列表
func (s *sTaskVideo) Page(ctx context.Context, params model.TaskVideoPageReq) (*model.TaskVideoPageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

	if params.TraceId != "" {
		filter["trace_id"] = gstr.Trim(params.TraceId)
	}

	if service.Session().IsResellerRole(ctx) {
		filter["rid"] = service.Session().GetRid(ctx)
	}

	if service.Session().IsUserRole(ctx) {
		filter["user_id"] = service.Session().GetUserId(ctx)
	} else if params.UserId != 0 {
		filter["user_id"] = params.UserId
	}

	if params.AppId != 0 {
		filter["app_id"] = params.AppId
	}

	if params.TaskId != "" {
		filter["task_id"] = params.TaskId
	}

	if params.Status != 0 {
		filter["status"] = params.Status
	}

	if len(params.CreatedAt) > 0 {
		gte := gtime.NewFromStrFormat(params.CreatedAt[0], time.DateTime).TimestampMilli()
		lte := gtime.NewFromStrLayout(params.CreatedAt[1], time.DateTime).TimestampMilli() + 999
		filter["created_at"] = bson.M{
			"$gte": gte,
			"$lte": lte,
		}
	}

	results, err := dao.TaskVideo.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"-created_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.TaskVideo, 0)
	for _, result := range results {

		audio := &model.TaskVideo{
			Id:        result.Id,
			TraceId:   result.TraceId,
			UserId:    result.UserId,
			AppId:     result.AppId,
			TaskId:    result.TaskId,
			VideoUrl:  result.VideoUrl,
			VideoTime: result.VideoTime,
			Status:    result.Status,
			CreatedAt: util.FormatDateTimeMonth(result.CreatedAt),
		}

		items = append(items, audio)
	}

	return &model.TaskVideoPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}
