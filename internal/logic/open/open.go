package open

import (
	"context"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/iimeta/fastapi-admin/v2/internal/consts"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/internal/errors"
	"github.com/iimeta/fastapi-admin/v2/internal/logic/common"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/model/entity"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"github.com/iimeta/fastapi-admin/v2/utility/util"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type sOpen struct{}

func init() {
	service.RegisterOpen(New())
}

func New() service.IOpen {
	return &sOpen{}
}

// 站点配置
func (s *sOpen) Site(ctx context.Context, params model.SiteConfigDetailReq) *model.SiteConfig {

	var siteConfig *entity.SiteConfig

	if params.Domain != "" {
		siteConfig, _ = dao.SiteConfig.FindOne(ctx, bson.M{"domain": params.Domain, "status": 1})
	}

	if siteConfig == nil {
		siteConfig, _ = dao.SiteConfig.FindOne(ctx, bson.M{"user_id": 1, "status": 1}, &dao.FindOptions{SortFields: []string{"-updated_at"}})
	}

	if siteConfig == nil {
		return nil
	}

	return &model.SiteConfig{
		Domain:              siteConfig.Domain,
		Title:               siteConfig.Title,
		Logo:                siteConfig.Logo,
		Favicon:             siteConfig.Favicon,
		Avatar:              siteConfig.Avatar,
		BgImg:               siteConfig.BgImg,
		Copyright:           siteConfig.Copyright,
		JumpUrl:             siteConfig.JumpUrl,
		Keywords:            siteConfig.Keywords,
		Description:         siteConfig.Description,
		IcpBeian:            siteConfig.IcpBeian,
		GaBeian:             siteConfig.GaBeian,
		RegisterTips:        siteConfig.RegisterTips,
		DefaultLanguage:     siteConfig.DefaultLanguage,
		CurrencySymbol:      siteConfig.CurrencySymbol,
		Carousel1Title:      siteConfig.Carousel1Title,
		Carousels1:          siteConfig.Carousels1,
		Carousel2Title:      siteConfig.Carousel2Title,
		Carousels2:          siteConfig.Carousels2,
		AnnouncementTitle:   siteConfig.AnnouncementTitle,
		AnnouncementMoreUrl: siteConfig.AnnouncementMoreUrl,
		Announcements:       siteConfig.Announcements,
		DocumentTitle:       siteConfig.DocumentTitle,
		DocumentMoreUrl:     siteConfig.DocumentMoreUrl,
		Documents:           siteConfig.Documents,
		RechargeTips:        siteConfig.RechargeTips,
	}
}

// 系统配置
func (s *sOpen) Config(ctx context.Context, params model.SysConfigReq) (*model.SysConfig, error) {

	sysConfig, err := dao.SysConfig.FindOne(ctx, bson.M{})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if sysConfig.AdminLogin.Path == "" {
		sysConfig.AdminLogin.Path = "admin"
	} else if sysConfig.AdminLogin.Path != gstr.TrimLeft(params.Path, "/") {
		sysConfig.AdminLogin = nil
	}

	return &model.SysConfig{
		UserLoginRegister:     sysConfig.UserLoginRegister,
		ResellerLoginRegister: sysConfig.ResellerLoginRegister,
		AdminLogin:            sysConfig.AdminLogin,
	}, nil
}

// 视频文件
func (s *sOpen) Video(ctx context.Context, fileName string) (string, error) {

	taskVideo, err := dao.TaskVideo.FindOne(ctx, bson.M{"file_name": fileName})
	if err != nil {
		logger.Error(ctx, err)
		return "", errors.New("视频文件未找到")
	}

	if taskVideo == nil || taskVideo.FilePath == "" {
		return "", errors.New("视频文件未找到")
	}

	return taskVideo.FilePath, nil
}

// 文件
func (s *sOpen) File(ctx context.Context, fileName string) (string, error) {

	taskFile, err := dao.TaskFile.FindOne(ctx, bson.M{"file_id": gfile.Name(fileName)})
	if err != nil {
		logger.Error(ctx, err)
		return "", errors.New("文件未找到")
	}

	if taskFile == nil || taskFile.FilePath == "" {
		return "", errors.New("文件未找到")
	}

	return taskFile.FilePath, nil
}

// 用户协议
func (s *sOpen) UserAgreement(ctx context.Context, params model.SysConfigReq) (string, error) {

	template, err := service.NoticeTemplate().GetNoticeTemplateByScene(ctx, consts.SCENE_USER_AGREEMENT, nil)
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	var siteConfig *entity.SiteConfig

	if params.Domain != "" {
		siteConfig, _ = dao.SiteConfig.FindOne(ctx, bson.M{"domain": params.Domain, "status": 1})
	}

	if siteConfig == nil {
		siteConfig, _ = dao.SiteConfig.FindOne(ctx, bson.M{"user_id": 1, "status": 1}, &dao.FindOptions{SortFields: []string{"-updated_at"}})
	}

	_, content, err := util.RenderTemplate(template.Title, template.Content, common.GetVariableData(ctx, nil, nil, siteConfig, template.Variables))
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	return content, nil
}

// 隐私政策
func (s *sOpen) PrivacyPolicy(ctx context.Context, params model.SysConfigReq) (string, error) {

	template, err := service.NoticeTemplate().GetNoticeTemplateByScene(ctx, consts.SCENE_PRIVACY_POLICY, nil)
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	var siteConfig *entity.SiteConfig

	if params.Domain != "" {
		siteConfig, _ = dao.SiteConfig.FindOne(ctx, bson.M{"domain": params.Domain, "status": 1})
	}

	if siteConfig == nil {
		siteConfig, _ = dao.SiteConfig.FindOne(ctx, bson.M{"user_id": 1, "status": 1}, &dao.FindOptions{SortFields: []string{"-updated_at"}})
	}

	_, content, err := util.RenderTemplate(template.Title, template.Content, common.GetVariableData(ctx, nil, nil, siteConfig, template.Variables))
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	return content, nil
}
