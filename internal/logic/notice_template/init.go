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
			Name:     "安全验证",
			Scenes:   []string{consts.SCENE_CODE},
			Title:    "安全验证",
			Content:  `<div style="max-width:600px;margin:0 auto;background:#fff;border-radius:12px;overflow:hidden;box-shadow:0 4px 12px rgba(0,0,0,.05);font-family:'Helvetica Neue',Arial,sans-serif;line-height:1.6;color:#333"><div style="background:linear-gradient(135deg,#6366f1,#8b5cf6);padding:30px;text-align:center;color:#fff"><h1 style="margin:0;font-size:24px">安全验证</h1></div><div style="padding:30px"><p style="margin:0 0 15px">尊敬的用户：</p><p style="margin:0 0 20px">您好！您正在尝试 <strong>{{.scene}}</strong>，请使用以下验证码完成验证：</p><div style="background:#f0f5ff;border-radius:8px;padding:20px;margin:25px 0;text-align:center;border:1px dashed #c7d2fe"><p style="margin:0 0 8px;color:#64748b;font-size:14px">您的验证码</p><div style="font-size:36px;font-weight:700;letter-spacing:4px;color:#4f46e5;margin:10px 0">{{.code}}</div><p style="margin:8px 0 0;color:#64748b;font-size:14px">有效期：15分钟</p></div><div style="background:#fef2f2;border-left:4px solid #ef4444;padding:12px;border-radius:8px;margin:20px 0"><p style="margin:0;color:#b91c1c;font-weight:700">⚠️ 安全提示</p><p style="margin:8px 0 0;color:#b91c1c">请勿向任何人透露此验证码，包括自称客服的人员。<br>如非本人操作，请忽略此邮件。</p></div></div><div style="padding:20px;text-align:center;font-size:12px;color:#9ca3af;background:#f9fafb"><p style="margin:0 0 8px 0">此为系统邮件，请勿直接回复</p><p style="margin:0"><a href="{{.site.jump_url}}" style="color:#4f46e5;text-decoration:none">{{.site.copyright}}</a></p></div></div>`,
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
			Content:  `<div style="max-width:600px;margin:0 auto;background:#fff;border-radius:12px;overflow:hidden;box-shadow:0 4px 12px rgba(0,0,0,.05);font-family:'Helvetica Neue',Arial,sans-serif;line-height:1.6;color:#333"><div style="background:linear-gradient(135deg,#6366f1,#8b5cf6);padding:30px;text-align:center;color:#fff"><h1 style="margin:0;font-size:24px">额度过期提醒</h1></div><div style="padding:30px"><p style="margin:0 0 15px">尊敬的 <strong>{{.name}}</strong>：</p><div style="background:#f0f5ff;border-radius:8px;padding:20px;margin:25px 0;border:1px dashed #c7d2fe"><div style="margin-bottom:10px"><p style="margin:0 0 5px;color:#64748b;font-size:16px">剩余额度</p><div style="font-size:28px;font-weight:700;color:#4f46e5;font-family:monospace">{{.quota}}</div></div><div style="height:1px;background:linear-gradient(90deg,transparent,#e0e7ff,transparent);margin:10px 0"></div><div><p style="margin:0 0 5px;color:#64748b;font-size:16px">过期时间</p><div style="font-size:28px;font-weight:700;color:#f97316;font-family:monospace">{{.quota_expires_at}}</div></div></div><div style="background:#fff7ed;border-left:4px solid #f97316;padding:12px;border-radius:8px;margin:20px 0"><p style="margin:0;color:#ea580c;font-weight:700">⚠️ 重要提醒</p><p style="margin:8px 0 0;color:#ea580c">您的额度将在 <strong>{{.quota_expires_at}}</strong> 过期。<br>请在过期前使用完额度，避免额度过期后无法使用造成损失！</p></div><div style="background:#f0fdf4;border-radius:8px;padding:16px;margin:20px 0;border-left:4px solid #10b981"><p style="margin:0 0 10px 0;color:#065f46;font-weight:600">💡 如何延长使用期限？</p><ul style="margin:0;padding-left:20px;color:#4b5563"><li style="margin-bottom:6px">如需充值，请及时联系管理员</li><li style="margin-bottom:6px">充值额度可延长过期时间</li></ul></div></div><div style="padding:20px;text-align:center;font-size:12px;color:#9ca3af;background:#f9fafb"><p style="margin:0 0 8px 0">此为系统邮件，请勿直接回复</p><p style="margin:0"><a href="{{.site.jump_url}}" style="color:#4f46e5;text-decoration:none">{{.site.copyright}}</a></p></div></div>`,
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
