package common

import (
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/iimeta/fastapi-admin/v2/internal/config"
)

func isImageUrlReplaceOpen() bool {
	return config.Cfg != nil && config.Cfg.ImageUrl != nil && config.Cfg.ImageUrl.Open && len(config.Cfg.ImageUrl.Urls) > 0
}

// 按配置顺序对图像 URL 做前缀替换, 命中后继续套后续规则
func ReplaceImageUrl(imageUrl string) string {

	if imageUrl == "" || !isImageUrlReplaceOpen() {
		return imageUrl
	}

	for _, item := range config.Cfg.ImageUrl.Urls {
		if item.ReplaceUrl == "" {
			continue
		}
		if gstr.HasPrefix(imageUrl, item.ReplaceUrl) {
			imageUrl = item.TargetUrl + imageUrl[len(item.ReplaceUrl):]
		}
	}

	return imageUrl
}
