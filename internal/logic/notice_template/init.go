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

	noticeTemplates := []*do.NoticeTemplate{
		{
			Name:     "验证码",
			Scenes:   []string{consts.SCENE_CODE},
			Title:    "验证码",
			Content:  `<p>尊敬的用户：您好！ 您本次操作验证码：<span style="color: #ff8926; font-size: 18px"><strong>{{.code}}</strong></span>，验证码有效期 <span style="color: #ff8926; font-size: 18px"><strong>15</strong></span> 分钟。</p><p>注意：此操作如非本人操作，请及时登录并修改密码以保证帐户安全。<br>（工作人员不会向你索取此验证码，请勿向他人泄漏！！！)</p><p>此为系统邮件，请勿回复<br>请保管好您的邮箱，避免账号被他人盗用</p><p><a href="{{.site.jump_url}}">{{.site.copyright}}</a></p>`,
			Channels: []string{consts.NOTICE_CHANNEL_EMAIL},
			IsPublic: true,
			Status:   1,
			UserId:   1,
		},
		{
			Name:     "额度不足提醒",
			Scenes:   []string{consts.SCENE_QUOTA_WARNING},
			Title:    "额度不足提醒",
			Content:  `<p><strong>尊敬的用户：您好！</strong> <strong>您当前额度：</strong><span style="color: red"><strong>{{.quota}}</strong></span><strong> 已低于预警阈值：</strong><span style="color: rgb(255, 102, 0)"><strong>${{.warning_threshold}}</strong></span><strong>，请关注额度使用情况。<br>如需充值，请及时联系管理员，避免额度耗尽影响使用！</strong></p><p>此为系统邮件，请勿回复<br>请保管好您的邮箱，避免账号被他人盗用</p>`,
			Channels: []string{consts.NOTICE_CHANNEL_EMAIL},
			IsPublic: true,
			Status:   1,
			UserId:   1,
		},
		{
			Name:     "额度耗尽通知",
			Scenes:   []string{consts.SCENE_QUOTA_EXHAUSTION},
			Title:    "额度耗尽通知",
			Content:  `<p><strong>尊敬的用户：您好！</strong> <strong>您的额度已耗尽，当前额度：</strong><span style="color: red"><strong>{{.quota}}</strong></span><strong>，请关注额度使用情况。<br>如需充值，请及时联系管理员，充值成功后即可恢复正常使用！</strong></p><p>此为系统邮件，请勿回复<br>请保管好您的邮箱，避免账号被他人盗用</p>`,
			Channels: []string{consts.NOTICE_CHANNEL_EMAIL},
			IsPublic: true,
			Status:   1,
			UserId:   1,
		},
		{
			Name:     "额度过期提醒",
			Scenes:   []string{consts.SCENE_QUOTA_EXPIRE_WARNING},
			Title:    "额度过期提醒",
			Content:  `<p><strong>尊敬的用户：您好！</strong> <strong>您当前额度：</strong><span style="color: red"><strong>{{.quota}}</strong></span><strong> 即将过期，过期时间：</strong><span style="color: rgb(255, 102, 0)"><strong>${{.quota_expires_at}}</strong></span><strong><br>请在过期前使用完额度，避免额度过期后无法使用造成损失。<br>如需充值，请及时联系管理员，充值额度可延长过期时间！</strong></p><p>此为系统邮件，请勿回复<br>请保管好您的邮箱，避免账号被他人盗用</p>`,
			Channels: []string{consts.NOTICE_CHANNEL_EMAIL},
			IsPublic: true,
			Status:   1,
			UserId:   1,
		},
		{
			Name:     "额度过期通知",
			Scenes:   []string{consts.SCENE_QUOTA_EXPIRE},
			Title:    "额度过期通知",
			Content:  `<p><strong>尊敬的用户：您好！</strong> <strong>您当前额度：</strong><span style="color: red"><strong>{{.quota}}</strong></span><strong> 已于 </strong><span style="color: red"><strong>${{.quota_expires_at}}</strong></span><strong> 过期。<br>如需充值，请及时联系管理员，充值额度可延长过期时间！</strong></p><p>此为系统邮件，请勿回复<br>请保管好您的邮箱，避免账号被他人盗用</p>`,
			Channels: []string{consts.NOTICE_CHANNEL_EMAIL},
			IsPublic: true,
			Status:   1,
			UserId:   1,
		},
	}

	for _, noticeTemplate := range noticeTemplates {
		noticeTemplate.Variables = util.GetTemplateVariables(noticeTemplate.Title, noticeTemplate.Content)
	}

	return noticeTemplates
}
