package notice_template

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/util"
	"go.mongodb.org/mongo-driver/bson"
)

// 初始化通知模板
func (s *sNoticeTemplate) Init(ctx context.Context) {

	noticeTemplates, err := dao.NoticeTemplate.Find(ctx, bson.M{"rid": bson.M{"$exists": false}})
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	noticeTemplateMap := util.ToMap(noticeTemplates, func(t *entity.NoticeTemplate) string {
		return t.Name
	})

	for _, defaultNoticeTemplate := range s.Default() {
		if _, ok := noticeTemplateMap[defaultNoticeTemplate.Name]; !ok {
			if _, err = dao.NoticeTemplate.Insert(ctx, defaultNoticeTemplate); err != nil {
				logger.Error(ctx, err)
			}
		}
	}
}

// 默认通知模板
func (s *sNoticeTemplate) Default() []*do.NoticeTemplate {
	return []*do.NoticeTemplate{
		{
			Name:      "验证码模板",
			Scenes:    []string{consts.SCENE_CODE},
			Title:     "智元 Fast API",
			Content:   "验证码模板",
			Channels:  []string{consts.NOTICE_CHANNEL_EMAIL},
			IsPublic:  true,
			Status:    1,
			Variables: []string{consts.ATTRIBUTE_CODE},
			UserId:    1,
		},
		{
			Name:      "注册验证码模板",
			Scenes:    []string{consts.SCENE_CODE},
			Title:     "智元 Fast API",
			Content:   "注册验证码模板",
			Channels:  []string{consts.NOTICE_CHANNEL_EMAIL},
			IsPublic:  true,
			Status:    1,
			Variables: []string{consts.ATTRIBUTE_CODE},
			UserId:    1,
		},
	}
}
