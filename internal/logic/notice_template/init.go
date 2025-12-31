package notice_template

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/internal/consts"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/internal/model/do"
	"github.com/iimeta/fastapi-admin/v2/internal/model/entity"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"github.com/iimeta/fastapi-admin/v2/utility/util"
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
			Name:     consts.SCENE[consts.SCENE_CODE],
			Scenes:   []string{consts.SCENE_CODE},
			Title:    consts.SCENE[consts.SCENE_CODE],
			Content:  `<div style="max-width:600px;margin:0 auto;background:#fff;border-radius:12px;overflow:hidden;box-shadow:0 4px 12px rgba(0,0,0,.05);font-family:'Helvetica Neue',Arial,sans-serif;line-height:1.6;color:#333"><div style="background:linear-gradient(135deg,#6366f1,#8b5cf6);padding:30px;text-align:center;color:#fff"><span style="margin:0;font-size:24px;font-weight:700">安全验证</span></div><div style="padding:30px"><p style="margin:0 0 15px">尊敬的用户：</p><p style="margin:0 0 20px">您好！您正在尝试 <strong>{{.scene}}</strong>，请使用以下验证码完成验证：</p><div style="background:#f0f5ff;border-radius:8px;padding:20px;margin:25px 0;text-align:center;border:1px dashed #c7d2fe"><p style="margin:0 0 8px;color:#64748b;font-size:14px">您的验证码</p><div style="font-size:36px;font-weight:700;letter-spacing:4px;color:#4f46e5;margin:10px 0">{{.code}}</div><p style="margin:8px 0 0;color:#64748b;font-size:14px">有效期：15分钟</p></div><div style="background:#fef2f2;border-radius:8px;padding:16px;margin:20px 0;border-left:4px solid #ef4444"><p style="margin:0 0 10px 0;color:#b91c1c;font-weight:700">⚠️ 安全提示</p><ul style="margin:0;padding-left:22px;color:#b91c1c"><li style="margin-bottom:6px">请勿向任何人透露此验证码，包括自称客服的人员。</li><li style="margin-bottom:6px">如非本人操作，请忽略此邮件。</li></ul></div></div><div style="padding:20px;text-align:center;font-size:12px;color:#9ca3af;background:#f9fafb"><p style="margin:0 0 8px 0">此为系统邮件，请勿直接回复</p><p style="margin:0"><a href="{{.site.jump_url}}" style="color:#4f46e5;text-decoration:none">{{.site.copyright}}</a></p></div></div>`,
			Channels: []string{consts.NOTICE_CHANNEL_WEB, consts.NOTICE_CHANNEL_EMAIL},
			IsPublic: true,
			Status:   1,
			UserId:   1,
			Remark:   "系统默认模板",
		},
		{
			Name:     consts.SCENE[consts.SCENE_QUOTA_RECHARGE],
			Scenes:   []string{consts.SCENE_QUOTA_RECHARGE},
			Title:    consts.SCENE[consts.SCENE_QUOTA_RECHARGE],
			Content:  `<div style="max-width:600px;margin:0 auto;background:#fff;border-radius:12px;overflow:hidden;box-shadow:0 4px 12px rgba(0,0,0,.05);font-family:'Helvetica Neue',Arial,sans-serif;line-height:1.6;color:#333"><div style="background:linear-gradient(135deg,#6366f1,#8b5cf6);padding:30px;text-align:center;color:#fff"><span style="margin:0;font-size:24px;font-weight:700">{{.quota_type}}额度</span></div><div style="padding:30px"><p style="margin:0 0 15px">尊敬的 <strong>{{.name}}</strong>：</p><div style="background:#f0f5ff;border-radius:8px;padding:20px;margin:25px 0;border:1px dashed #c7d2fe"><div style="margin-bottom:10px"><p style="margin:0 0 5px;color:#64748b;font-size:16px">{{.quota_type}}额度</p><div style="font-size:28px;font-weight:700;color:#4f46e5;font-family:monospace">{{.recharge_quota}}</div></div><div style="height:1px;background:linear-gradient(90deg,transparent,#e0e7ff,transparent);margin:10px 0"></div><div style="margin-bottom:10px"><p style="margin:0 0 5px;color:#64748b;font-size:16px">当前额度</p><div style="font-size:28px;font-weight:700;color:#4f46e5;font-family:monospace">{{.quota}}</div></div><div style="height:1px;background:linear-gradient(90deg,transparent,#e0e7ff,transparent);margin:10px 0"></div><div><p style="margin:0 0 5px;color:#64748b;font-size:16px">过期时间</p><div style="font-size:28px;font-weight:700;color:#4f46e5;font-family:monospace">{{.quota_expires_at}}</div></div></div><div style="background:#fff7ed;border-radius:8px;padding:16px;margin:20px 0;border-left:4px solid #f97316"><p style="margin:0 0 10px 0;color:#ea580c;font-weight:700">⚠️ 温馨提示</p><ul style="margin:0;padding-left:22px;color:#ea580c"><li style="margin-bottom:6px">请在过期前使用完额度，避免额度过期后无法使用造成损失！</li></ul></div><div style="background:#f0fdf4;border-radius:8px;padding:16px;margin:20px 0;border-left:4px solid #10b981"><p style="margin:0 0 10px 0;color:#065f46;font-weight:600">💡 如何延长使用期限？</p><ul style="margin:0;padding-left:22px;color:#4b5563"><li style="margin-bottom:6px">如需充值，请联系管理员</li><li style="margin-bottom:6px">充值额度可延长过期时间</li></ul></div></div><div style="padding:20px;text-align:center;font-size:12px;color:#9ca3af;background:#f9fafb"><p style="margin:0 0 8px 0">此为系统邮件，请勿直接回复</p><p style="margin:0"><a href="{{.site.jump_url}}" style="color:#4f46e5;text-decoration:none">{{.site.copyright}}</a></p></div></div>`,
			Channels: []string{consts.NOTICE_CHANNEL_WEB, consts.NOTICE_CHANNEL_EMAIL},
			IsPublic: true,
			Status:   1,
			UserId:   1,
			Remark:   "系统默认模板",
		},
		{
			Name:     consts.SCENE[consts.SCENE_QUOTA_WARNING],
			Scenes:   []string{consts.SCENE_QUOTA_WARNING},
			Title:    consts.SCENE[consts.SCENE_QUOTA_WARNING],
			Content:  `<div style="max-width:600px;margin:0 auto;background:#fff;border-radius:12px;overflow:hidden;box-shadow:0 4px 12px rgba(0,0,0,.05);font-family:'Helvetica Neue',Arial,sans-serif;line-height:1.6;color:#333"><div style="background:linear-gradient(135deg,#6366f1,#8b5cf6);padding:30px;text-align:center;color:#fff"><span style="margin:0;font-size:24px;font-weight:700">额度不足提醒</span></div><div style="padding:30px"><p style="margin:0 0 15px">尊敬的 <strong>{{.name}}</strong>：</p><div style="background:#f0f5ff;border-radius:8px;padding:20px;margin:25px 0;border:1px dashed #c7d2fe"><div style="margin-bottom:10px"><p style="margin:0 0 5px;color:#64748b;font-size:16px">剩余额度</p><div style="font-size:28px;font-weight:700;color:#f97316;font-family:monospace">{{.quota}}</div></div><div style="height:1px;background:linear-gradient(90deg,transparent,#e0e7ff,transparent);margin:10px 0"></div><div><p style="margin:0 0 5px;color:#64748b;font-size:16px">预警阈值</p><div style="font-size:28px;font-weight:700;color:#f97316;font-family:monospace">{{.warning_threshold}}</div></div></div><div style="background:#fff7ed;border-radius:8px;padding:16px;margin:20px 0;border-left:4px solid #f97316"><p style="margin:0 0 10px 0;color:#ea580c;font-weight:700">⚠️ 重要提醒</p><ul style="margin:0;padding-left:22px;color:#ea580c"><li style="margin-bottom:6px">您当前额度 <strong>{{.quota}}</strong> 已低于预警阈值<strong>{{.warning_threshold}}</strong>。</li><li style="margin-bottom:6px">请关注额度使用情况！</li></ul></div><div style="background:#f0fdf4;border-radius:8px;padding:16px;margin:20px 0;border-left:4px solid #10b981"><p style="margin:0 0 10px 0;color:#065f46;font-weight:600">💡 如何保持正常使用？</p><ul style="margin:0;padding-left:22px;color:#4b5563"><li style="margin-bottom:6px">如需充值，请联系管理员</li><li style="margin-bottom:6px">避免额度耗尽影响使用！</li></ul></div></div><div style="padding:20px;text-align:center;font-size:12px;color:#9ca3af;background:#f9fafb"><p style="margin:0 0 8px 0">此为系统邮件，请勿直接回复</p><p style="margin:0"><a href="{{.site.jump_url}}" style="color:#4f46e5;text-decoration:none">{{.site.copyright}}</a></p></div></div>`,
			Channels: []string{consts.NOTICE_CHANNEL_WEB, consts.NOTICE_CHANNEL_EMAIL},
			IsPublic: true,
			Status:   1,
			UserId:   1,
			Remark:   "系统默认模板",
		},
		{
			Name:     consts.SCENE[consts.SCENE_QUOTA_EXHAUSTION],
			Scenes:   []string{consts.SCENE_QUOTA_EXHAUSTION},
			Title:    consts.SCENE[consts.SCENE_QUOTA_EXHAUSTION],
			Content:  `<div style="max-width:600px;margin:0 auto;background:#fff;border-radius:12px;overflow:hidden;box-shadow:0 4px 12px rgba(0,0,0,.05);font-family:'Helvetica Neue',Arial,sans-serif;line-height:1.6;color:#333"><div style="background:linear-gradient(135deg,#6366f1,#8b5cf6);padding:30px;text-align:center;color:#fff"><span style="margin:0;font-size:24px;font-weight:700">额度耗尽通知</span></div><div style="padding:30px"><p style="margin:0 0 15px">尊敬的 <strong>{{.name}}</strong>：</p><div style="background:#f0f5ff;border-radius:8px;padding:20px;margin:25px 0;border:1px dashed #c7d2fe;text-align:center"><div style="display:inline-block;text-align:center"><p style="margin:0 0 5px;color:#64748b;font-size:16px">当前额度</p><div style="font-size:32px;font-weight:800;color:#dc2626;font-family:monospace;letter-spacing:1px">{{.quota}}</div><div style="margin-top:8px;padding:4px 12px;background:#ef4444;color:#fff;border-radius:15px;font-size:13px;display:inline-block">额度已耗尽</div></div></div><div style="background:#fef2f2;border-radius:8px;padding:16px;margin:20px 0;border-left:4px solid #ef4444"><p style="margin:0 0 10px 0;color:#b91c1c;font-weight:700">⚠️ 重要通知</p><ul style="margin:0;padding-left:22px;color:#b91c1c"><li style="margin-bottom:6px">您账户额度已耗尽，部分功能暂时无法使用，请充值即可恢复服务！</li></ul></div><div style="background:#f0fdf4;border-radius:8px;padding:16px;margin:20px 0;border-left:4px solid #10b981"><p style="margin:0 0 10px 0;color:#065f46;font-weight:600">💡 如何恢复使用？</p><ul style="margin:0;padding-left:22px;color:#4b5563"><li style="margin-bottom:6px">请及时联系管理员进行充值</li><li style="margin-bottom:6px">充值成功后将自动恢复服务</li></ul></div></div><div style="padding:20px;text-align:center;font-size:12px;color:#9ca3af;background:#f9fafb"><p style="margin:0 0 8px 0">此为系统邮件，请勿直接回复</p><p style="margin:0"><a href="{{.site.jump_url}}" style="color:#4f46e5;text-decoration:none">{{.site.copyright}}</a></p></div></div>`,
			Channels: []string{consts.NOTICE_CHANNEL_WEB, consts.NOTICE_CHANNEL_EMAIL},
			IsPublic: true,
			Status:   1,
			UserId:   1,
			Remark:   "系统默认模板",
		},
		{
			Name:     consts.SCENE[consts.SCENE_QUOTA_EXPIRE_WARNING],
			Scenes:   []string{consts.SCENE_QUOTA_EXPIRE_WARNING},
			Title:    consts.SCENE[consts.SCENE_QUOTA_EXPIRE_WARNING],
			Content:  `<div style="max-width:600px;margin:0 auto;background:#fff;border-radius:12px;overflow:hidden;box-shadow:0 4px 12px rgba(0,0,0,.05);font-family:'Helvetica Neue',Arial,sans-serif;line-height:1.6;color:#333"><div style="background:linear-gradient(135deg,#6366f1,#8b5cf6);padding:30px;text-align:center;color:#fff"><span style="margin:0;font-size:24px;font-weight:700">额度过期提醒</span></div><div style="padding:30px"><p style="margin:0 0 15px">尊敬的 <strong>{{.name}}</strong>：</p><div style="background:#f0f5ff;border-radius:8px;padding:20px;margin:25px 0;border:1px dashed #c7d2fe"><div style="margin-bottom:10px"><p style="margin:0 0 5px;color:#64748b;font-size:16px">剩余额度</p><div style="font-size:28px;font-weight:700;color:#f97316;font-family:monospace">{{.quota}}</div></div><div style="height:1px;background:linear-gradient(90deg,transparent,#e0e7ff,transparent);margin:10px 0"></div><div><p style="margin:0 0 5px;color:#64748b;font-size:16px">过期时间</p><div style="font-size:28px;font-weight:700;color:#f97316;font-family:monospace">{{.quota_expires_at}}</div></div></div><div style="background:#fff7ed;border-radius:8px;padding:16px;margin:20px 0;border-left:4px solid #f97316"><p style="margin:0 0 10px 0;color:#ea580c;font-weight:700">⚠️ 重要提醒</p><ul style="margin:0;padding-left:22px;color:#ea580c"><li style="margin-bottom:6px">您的 <strong>{{.quota}}</strong> 额度将在 <strong>{{.quota_expires_at}}</strong> 过期。</li><li style="margin-bottom:6px">请在过期前使用完额度，避免额度过期后无法使用造成损失！</li></ul></div><div style="background:#f0fdf4;border-radius:8px;padding:16px;margin:20px 0;border-left:4px solid #10b981"><p style="margin:0 0 10px 0;color:#065f46;font-weight:600">💡 如何延长使用期限？</p><ul style="margin:0;padding-left:22px;color:#4b5563"><li style="margin-bottom:6px">如需充值，请联系管理员</li><li style="margin-bottom:6px">充值额度可延长过期时间</li></ul></div></div><div style="padding:20px;text-align:center;font-size:12px;color:#9ca3af;background:#f9fafb"><p style="margin:0 0 8px 0">此为系统邮件，请勿直接回复</p><p style="margin:0"><a href="{{.site.jump_url}}" style="color:#4f46e5;text-decoration:none">{{.site.copyright}}</a></p></div></div>`,
			Channels: []string{consts.NOTICE_CHANNEL_WEB, consts.NOTICE_CHANNEL_EMAIL},
			IsPublic: true,
			Status:   1,
			UserId:   1,
			Remark:   "系统默认模板",
		},
		{
			Name:     consts.SCENE[consts.SCENE_QUOTA_EXPIRE],
			Scenes:   []string{consts.SCENE_QUOTA_EXPIRE},
			Title:    consts.SCENE[consts.SCENE_QUOTA_EXPIRE],
			Content:  `<div style="max-width:600px;margin:0 auto;background:#fff;border-radius:12px;overflow:hidden;box-shadow:0 4px 12px rgba(0,0,0,.05);font-family:'Helvetica Neue',Arial,sans-serif;line-height:1.6;color:#333"><div style="background:linear-gradient(135deg,#6366f1,#8b5cf6);padding:30px;text-align:center;color:#fff"><span style="margin:0;font-size:24px;font-weight:700">额度过期通知</span></div><div style="padding:30px"><p style="margin:0 0 15px">尊敬的 <strong>{{.name}}</strong>：</p><div style="background:#f0f5ff;border-radius:8px;padding:20px;margin:25px 0;border:1px dashed #c7d2fe"><div style="margin-bottom:10px"><p style="margin:0 0 5px;color:#64748b;font-size:16px">剩余额度</p><div style="font-size:28px;font-weight:700;color:#dc2626;font-family:monospace">{{.quota}}</div></div><div style="height:1px;background:linear-gradient(90deg,transparent,#e0e7ff,transparent);margin:10px 0"></div><div><p style="margin:0 0 5px;color:#64748b;font-size:16px">过期时间</p><div style="font-size:28px;font-weight:700;color:#dc2626;font-family:monospace">{{.quota_expires_at}}</div></div></div><div style="background:#fef2f2;border-radius:8px;padding:16px;margin:20px 0;border-left:4px solid #ef4444"><p style="margin:0 0 10px 0;color:#b91c1c;font-weight:700">⚠️ 重要通知</p><ul style="margin:0;padding-left:22px;color:#b91c1c"><li style="margin-bottom:6px">您的 <strong>{{.quota}}</strong> 额度已于 <strong>{{.quota_expires_at}}</strong> 过期。</li></ul></div><div style="background:#f0fdf4;border-radius:8px;padding:16px;margin:20px 0;border-left:4px solid #10b981"><p style="margin:0 0 10px 0;color:#065f46;font-weight:600">💡 如何延长使用期限？</p><ul style="margin:0;padding-left:22px;color:#4b5563"><li style="margin-bottom:6px">如需充值，请联系管理员</li><li style="margin-bottom:6px">充值额度可延长过期时间</li></ul></div></div><div style="padding:20px;text-align:center;font-size:12px;color:#9ca3af;background:#f9fafb"><p style="margin:0 0 8px 0">此为系统邮件，请勿直接回复</p><p style="margin:0"><a href="{{.site.jump_url}}" style="color:#4f46e5;text-decoration:none">{{.site.copyright}}</a></p></div></div>`,
			Channels: []string{consts.NOTICE_CHANNEL_WEB, consts.NOTICE_CHANNEL_EMAIL},
			IsPublic: true,
			Status:   1,
			UserId:   1,
			Remark:   "系统默认模板",
		},
		{
			Name:     consts.SCENE[consts.SCENE_NOTICE_REGISTER],
			Scenes:   []string{consts.SCENE_NOTICE_REGISTER},
			Title:    "欢迎来到 {{.site.title}}",
			Content:  `<div style="max-width:600px;margin:0 auto;background:#fff;border-radius:12px;overflow:hidden;box-shadow:0 4px 12px rgba(0,0,0,.05);font-family:'Helvetica Neue',Arial,sans-serif;line-height:1.6;color:#333"><div style="background:linear-gradient(135deg,#6366f1,#8b5cf6);padding:30px;text-align:center;color:#fff"><span style="margin:0;font-size:24px;font-weight:700">🎉 欢迎来到 {{.site.title}}</span></div><div style="padding:30px"><p style="margin:0 0 15px">尊敬的 <strong>{{.name}}</strong>：</p><div style="background:#f0f5ff;border-radius:8px;padding:20px;margin:25px 0;border:1px dashed #c7d2fe"><p style="margin:0 0 10px;font-size:18px;font-weight:600;color:#4f46e5">🎉 您的账户已成功开通</p><ul style="margin:0;padding-left:22px;color:#4b5563"><li style="margin:0;color:#4b5563">登录账号：<span style="font-size:16px;font-weight:600;color:#4f46e5">{{.account}}</span></li><li style="margin:0;color:#4b5563">账户额度：<span style="font-size:16px;font-weight:600;color:#4f46e5">{{.quota}}</span></li><li style="margin:0;color:#4b5563">额度过期：<span style="font-size:16px;font-weight:600;color:#4f46e5">{{.quota_expires_at}}</span></li></ul></div>{{.site.register_welcome}}</div><div style="padding:20px;text-align:center;font-size:12px;color:#9ca3af;background:#f9fafb"><p style="margin:0 0 8px 0">此为系统邮件，请勿直接回复</p><p style="margin:0"><a href="{{.site.jump_url}}" style="color:#4f46e5;text-decoration:none">{{.site.copyright}}</a></p></div></div>`,
			Channels: []string{consts.NOTICE_CHANNEL_WEB, consts.NOTICE_CHANNEL_EMAIL},
			IsPublic: true,
			Status:   1,
			UserId:   1,
			Remark:   "系统默认模板",
		},
	}

	for _, noticeTemplate := range noticeTemplates {
		noticeTemplate.Variables = util.GetTemplateVariables(noticeTemplate.Title, noticeTemplate.Content)
	}

	return noticeTemplates
}
