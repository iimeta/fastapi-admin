package site_config

import (
	"context"
	"regexp"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/fastapi-admin/v2/internal/consts"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/internal/errors"
	"github.com/iimeta/fastapi-admin/v2/internal/logic/common"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	mcommon "github.com/iimeta/fastapi-admin/v2/internal/model/common"
	"github.com/iimeta/fastapi-admin/v2/internal/model/do"
	"github.com/iimeta/fastapi-admin/v2/internal/model/entity"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
	"github.com/iimeta/fastapi-admin/v2/utility/db"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"github.com/iimeta/fastapi-admin/v2/utility/util"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type sSiteConfig struct{}

func init() {
	service.RegisterSiteConfig(New())
}

func New() service.ISiteConfig {
	return &sSiteConfig{}
}

// 新建站点配置
func (s *sSiteConfig) Create(ctx context.Context, params model.SiteConfigCreateReq) error {

	if existDomains := s.GetExistDomains(ctx, params.Domains); len(existDomains) > 0 {
		return errors.Newf("域名 \"%s\" 已存在", gstr.Join(existDomains, ", "))
	}

	siteConfig := &do.SiteConfig{
		Domains:            trimDomains(params.Domains),
		Title:              params.Title,
		Logo:               params.Logo,
		Favicon:            params.Favicon,
		Avatar:             params.Avatar,
		BgImg:              params.BgImg,
		Copyright:          params.Copyright,
		JumpUrl:            params.JumpUrl,
		Keywords:           params.Keywords,
		Description:        params.Description,
		IcpBeian:           params.IcpBeian,
		GaBeian:            params.GaBeian,
		RegisterTips:       params.RegisterTips,
		GrantQuota:         common.ConvQuotaUnit(params.GrantQuota),
		InviteEnabled:      params.InviteEnabled,
		InviteCodeRequired: params.InviteCodeRequired,
		InviteConfig: mcommon.InviteConfig{
			RewardQuota:                 common.ConvQuotaUnit(params.InviteConfig.RewardQuota),
			GrantQuota:                  common.ConvQuotaUnit(params.InviteConfig.GrantQuota),
			MinApplyQuota:               common.ConvQuotaUnit(params.InviteConfig.MinApplyQuota),
			DailyLimit:                  params.InviteConfig.DailyLimit,
			TotalLimit:                  params.InviteConfig.TotalLimit,
			IpDailyLimit:                params.InviteConfig.IpDailyLimit,
			IpTotalLimit:                params.InviteConfig.IpTotalLimit,
			IpPerInviterLimit:           params.InviteConfig.IpPerInviterLimit,
			IpLimitAction:               params.InviteConfig.IpLimitAction,
			RuleText:                    params.InviteConfig.RuleText,
			InvalidCodeAction:           params.InviteConfig.InvalidCodeAction,
			RechargeRebateEnabled:       params.InviteConfig.RechargeRebateEnabled,
			RechargeRebateFirstEnabled:  params.InviteConfig.RechargeRebateFirstEnabled,
			RechargeRebateFirstType:     params.InviteConfig.RechargeRebateFirstType,
			RechargeRebateFirstRate:     params.InviteConfig.RechargeRebateFirstRate,
			RechargeRebateFirstQuota:    common.ConvQuotaUnit(params.InviteConfig.RechargeRebateFirstQuota),
			RechargeRebateSecondEnabled: params.InviteConfig.RechargeRebateSecondEnabled,
			RechargeRebateSecondType:    params.InviteConfig.RechargeRebateSecondType,
			RechargeRebateSecondRate:    params.InviteConfig.RechargeRebateSecondRate,
			RechargeRebateSecondQuota:   common.ConvQuotaUnit(params.InviteConfig.RechargeRebateSecondQuota),
		},
		QuotaExpiresAt:      params.QuotaExpiresAt,
		SupportEmailSuffix:  params.SupportEmailSuffix,
		RegisterWelcome:     params.RegisterWelcome,
		DefaultLanguage:     params.DefaultLanguage,
		CurrencySymbol:      params.CurrencySymbol,
		Host:                params.Host,
		Port:                params.Port,
		UserName:            params.UserName,
		Password:            params.Password,
		FromName:            params.FromName,
		Carousel1Title:      params.Carousel1Title,
		Carousels1:          params.Carousels1,
		Carousel2Title:      params.Carousel2Title,
		Carousels2:          params.Carousels2,
		AnnouncementTitle:   params.AnnouncementTitle,
		AnnouncementMoreUrl: params.AnnouncementMoreUrl,
		Announcements:       params.Announcements,
		DocumentTitle:       params.DocumentTitle,
		DocumentMoreUrl:     params.DocumentMoreUrl,
		Documents:           params.Documents,
		Apis:                params.Apis,
		RechargeTips:        params.RechargeTips,
		Remark:              params.Remark,
		Status:              params.Status,
		UserId:              service.Session().GetUserId(ctx),
	}

	if _, err := dao.SiteConfig.Insert(ctx, siteConfig); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更新站点配置
func (s *sSiteConfig) Update(ctx context.Context, params model.SiteConfigUpdateReq) error {

	if existDomains := s.GetExistDomains(ctx, params.Domains, params.Id); len(existDomains) > 0 {
		return errors.Newf("域名 \"%s\" 已存在", gstr.Join(existDomains, ", "))
	}

	siteConfig := &do.SiteConfig{
		Domains:            trimDomains(params.Domains),
		Title:              params.Title,
		Logo:               params.Logo,
		Favicon:            params.Favicon,
		Avatar:             params.Avatar,
		BgImg:              params.BgImg,
		Copyright:          params.Copyright,
		JumpUrl:            params.JumpUrl,
		Keywords:           params.Keywords,
		Description:        params.Description,
		IcpBeian:           params.IcpBeian,
		GaBeian:            params.GaBeian,
		RegisterTips:       params.RegisterTips,
		GrantQuota:         common.ConvQuotaUnit(params.GrantQuota),
		InviteEnabled:      params.InviteEnabled,
		InviteCodeRequired: params.InviteCodeRequired,
		InviteConfig: mcommon.InviteConfig{
			RewardQuota:                 common.ConvQuotaUnit(params.InviteConfig.RewardQuota),
			GrantQuota:                  common.ConvQuotaUnit(params.InviteConfig.GrantQuota),
			MinApplyQuota:               common.ConvQuotaUnit(params.InviteConfig.MinApplyQuota),
			DailyLimit:                  params.InviteConfig.DailyLimit,
			TotalLimit:                  params.InviteConfig.TotalLimit,
			IpDailyLimit:                params.InviteConfig.IpDailyLimit,
			IpTotalLimit:                params.InviteConfig.IpTotalLimit,
			IpPerInviterLimit:           params.InviteConfig.IpPerInviterLimit,
			IpLimitAction:               params.InviteConfig.IpLimitAction,
			RuleText:                    params.InviteConfig.RuleText,
			InvalidCodeAction:           params.InviteConfig.InvalidCodeAction,
			RechargeRebateEnabled:       params.InviteConfig.RechargeRebateEnabled,
			RechargeRebateFirstEnabled:  params.InviteConfig.RechargeRebateFirstEnabled,
			RechargeRebateFirstType:     params.InviteConfig.RechargeRebateFirstType,
			RechargeRebateFirstRate:     params.InviteConfig.RechargeRebateFirstRate,
			RechargeRebateFirstQuota:    common.ConvQuotaUnit(params.InviteConfig.RechargeRebateFirstQuota),
			RechargeRebateSecondEnabled: params.InviteConfig.RechargeRebateSecondEnabled,
			RechargeRebateSecondType:    params.InviteConfig.RechargeRebateSecondType,
			RechargeRebateSecondRate:    params.InviteConfig.RechargeRebateSecondRate,
			RechargeRebateSecondQuota:   common.ConvQuotaUnit(params.InviteConfig.RechargeRebateSecondQuota),
		},
		QuotaExpiresAt:      params.QuotaExpiresAt,
		SupportEmailSuffix:  params.SupportEmailSuffix,
		RegisterWelcome:     params.RegisterWelcome,
		DefaultLanguage:     params.DefaultLanguage,
		CurrencySymbol:      params.CurrencySymbol,
		Host:                params.Host,
		Port:                params.Port,
		UserName:            params.UserName,
		Password:            params.Password,
		FromName:            params.FromName,
		Carousel1Title:      params.Carousel1Title,
		Carousels1:          params.Carousels1,
		Carousel2Title:      params.Carousel2Title,
		Carousels2:          params.Carousels2,
		AnnouncementTitle:   params.AnnouncementTitle,
		AnnouncementMoreUrl: params.AnnouncementMoreUrl,
		Announcements:       params.Announcements,
		DocumentTitle:       params.DocumentTitle,
		DocumentMoreUrl:     params.DocumentMoreUrl,
		Documents:           params.Documents,
		Apis:                params.Apis,
		RechargeTips:        params.RechargeTips,
		Remark:              params.Remark,
		Status:              params.Status,
	}

	if err := dao.SiteConfig.UpdateById(ctx, params.Id, siteConfig); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更改站点配置状态
func (s *sSiteConfig) ChangeStatus(ctx context.Context, params model.SiteConfigChangeStatusReq) error {

	if service.Session().IsResellerRole(ctx) {

		siteConfig, err := dao.SiteConfig.FindById(ctx, params.Id)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if siteConfig.Rid != service.Session().GetRid(ctx) {
			return errors.New("Unauthorized")
		}
	}

	if err := dao.SiteConfig.UpdateById(ctx, params.Id, bson.M{
		"status": params.Status,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 删除站点配置
func (s *sSiteConfig) Delete(ctx context.Context, id string) error {

	if service.Session().IsResellerRole(ctx) {

		siteConfig, err := dao.SiteConfig.FindById(ctx, id)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if siteConfig.Rid != service.Session().GetRid(ctx) {
			return errors.New("Unauthorized")
		}
	}

	if _, err := dao.SiteConfig.DeleteById(ctx, id); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 站点配置详情
func (s *sSiteConfig) Detail(ctx context.Context, params model.SiteConfigDetailReq) (*model.SiteConfig, error) {

	siteConfig, err := dao.SiteConfig.FindById(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if service.Session().IsResellerRole(ctx) && siteConfig.Rid != service.Session().GetRid(ctx) {
		return nil, errors.New("Unauthorized")
	}

	return &model.SiteConfig{
		Id:                 siteConfig.Id,
		Domains:            siteConfig.Domains,
		Title:              siteConfig.Title,
		Logo:               siteConfig.Logo,
		Favicon:            siteConfig.Favicon,
		Avatar:             siteConfig.Avatar,
		BgImg:              siteConfig.BgImg,
		Copyright:          siteConfig.Copyright,
		JumpUrl:            siteConfig.JumpUrl,
		Keywords:           siteConfig.Keywords,
		Description:        siteConfig.Description,
		IcpBeian:           siteConfig.IcpBeian,
		GaBeian:            siteConfig.GaBeian,
		RegisterTips:       siteConfig.RegisterTips,
		GrantQuota:         common.ConvQuotaUnitReverse(siteConfig.GrantQuota),
		InviteEnabled:      siteConfig.InviteEnabled,
		InviteCodeRequired: siteConfig.InviteCodeRequired,
		InviteConfig: model.InviteConfig{
			RewardQuota:                 common.ConvQuotaUnitReverse(siteConfig.InviteConfig.RewardQuota),
			GrantQuota:                  common.ConvQuotaUnitReverse(siteConfig.InviteConfig.GrantQuota),
			MinApplyQuota:               common.ConvQuotaUnitReverse(siteConfig.InviteConfig.MinApplyQuota),
			DailyLimit:                  siteConfig.InviteConfig.DailyLimit,
			TotalLimit:                  siteConfig.InviteConfig.TotalLimit,
			IpDailyLimit:                siteConfig.InviteConfig.IpDailyLimit,
			IpTotalLimit:                siteConfig.InviteConfig.IpTotalLimit,
			IpPerInviterLimit:           siteConfig.InviteConfig.IpPerInviterLimit,
			IpLimitAction:               siteConfig.InviteConfig.IpLimitAction,
			RuleText:                    siteConfig.InviteConfig.RuleText,
			InvalidCodeAction:           siteConfig.InviteConfig.InvalidCodeAction,
			RechargeRebateEnabled:       siteConfig.InviteConfig.RechargeRebateEnabled,
			RechargeRebateFirstEnabled:  siteConfig.InviteConfig.RechargeRebateFirstEnabled,
			RechargeRebateFirstType:     siteConfig.InviteConfig.RechargeRebateFirstType,
			RechargeRebateFirstRate:     siteConfig.InviteConfig.RechargeRebateFirstRate,
			RechargeRebateFirstQuota:    common.ConvQuotaUnitReverse(siteConfig.InviteConfig.RechargeRebateFirstQuota),
			RechargeRebateSecondEnabled: siteConfig.InviteConfig.RechargeRebateSecondEnabled,
			RechargeRebateSecondType:    siteConfig.InviteConfig.RechargeRebateSecondType,
			RechargeRebateSecondRate:    siteConfig.InviteConfig.RechargeRebateSecondRate,
			RechargeRebateSecondQuota:   common.ConvQuotaUnitReverse(siteConfig.InviteConfig.RechargeRebateSecondQuota),
		},
		QuotaExpiresAt:      siteConfig.QuotaExpiresAt,
		SupportEmailSuffix:  siteConfig.SupportEmailSuffix,
		RegisterWelcome:     siteConfig.RegisterWelcome,
		DefaultLanguage:     siteConfig.DefaultLanguage,
		CurrencySymbol:      siteConfig.CurrencySymbol,
		Host:                siteConfig.Host,
		Port:                siteConfig.Port,
		UserName:            siteConfig.UserName,
		Password:            siteConfig.Password,
		FromName:            siteConfig.FromName,
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
		Apis:                siteConfig.Apis,
		RechargeTips:        siteConfig.RechargeTips,
		Remark:              siteConfig.Remark,
		Status:              siteConfig.Status,
		UserId:              siteConfig.UserId,
		Creator:             siteConfig.Creator,
		Updater:             siteConfig.Updater,
		CreatedAt:           util.FormatDateTime(siteConfig.CreatedAt),
		UpdatedAt:           util.FormatDateTime(siteConfig.UpdatedAt),
	}, nil
}

// 站点配置分页列表
func (s *sSiteConfig) Page(ctx context.Context, params model.SiteConfigPageReq) (*model.SiteConfigPageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

	if service.Session().IsResellerRole(ctx) {
		filter["rid"] = service.Session().GetRid(ctx)
	}

	if params.UserId != 0 {
		filter["user_id"] = params.UserId
	}

	if params.Domain != "" {
		filter["domains"] = bson.M{
			"$regex": regexp.QuoteMeta(params.Domain),
		}
	}

	if params.Title != "" {
		filter["title"] = bson.M{
			"$regex": regexp.QuoteMeta(params.Title),
		}
	}

	if params.RegisterTips != "" {
		filter["register_tips"] = bson.M{
			"$regex": regexp.QuoteMeta(params.RegisterTips),
		}
	}

	if params.Logo != "" {
		filter["logo"] = bson.M{
			"$regex": regexp.QuoteMeta(params.Logo),
		}
	}

	if params.Status != 0 {
		filter["status"] = params.Status
	}

	results, err := dao.SiteConfig.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"status", "-updated_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.SiteConfig, 0)
	for _, result := range results {
		items = append(items, &model.SiteConfig{
			Id:           result.Id,
			Domains:      result.Domains,
			Title:        result.Title,
			RegisterTips: result.RegisterTips,
			Remark:       result.Remark,
			Status:       result.Status,
			UserId:       result.UserId,
			Creator:      result.Creator,
			Updater:      result.Updater,
			CreatedAt:    util.FormatDateTimeMonth(result.CreatedAt),
			UpdatedAt:    util.FormatDateTimeMonth(result.UpdatedAt),
		})
	}

	return &model.SiteConfigPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}

// 站点配置批量操作
func (s *sSiteConfig) BatchOperate(ctx context.Context, params model.SiteConfigBatchOperateReq) error {

	switch params.Action {
	case consts.ACTION_STATUS:
		for _, id := range params.Ids {
			if err := s.ChangeStatus(ctx, model.SiteConfigChangeStatusReq{
				Id:     id,
				Status: gconv.Int(params.Value),
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	case consts.ACTION_DELETE:
		for _, id := range params.Ids {
			if err := s.Delete(ctx, id); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	}

	return nil
}

// 根据域名获取站点配置
func (s *sSiteConfig) GetSiteConfigByDomain(ctx context.Context, domain string) *entity.SiteConfig {

	if domain == "" {
		return nil
	}

	siteConfig, _ := dao.SiteConfig.FindOne(ctx, bson.M{"domains": domain, "status": 1})
	return siteConfig
}

// 获取已存在的域名
func (s *sSiteConfig) GetExistDomains(ctx context.Context, domains []string, id ...string) []string {

	trimmed := trimDomains(domains)
	if len(trimmed) == 0 {
		return nil
	}

	filter := bson.M{"domains": bson.M{"$in": trimmed}}
	if len(id) > 0 && id[0] != "" {
		filter["_id"] = bson.M{"$ne": id[0]}
	}

	siteConfigs, err := dao.SiteConfig.Find(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil
		}
		logger.Error(ctx, err)
		return trimmed
	}

	existMap := make(map[string]struct{})
	for _, sc := range siteConfigs {
		for _, d := range sc.Domains {
			existMap[d] = struct{}{}
		}
	}

	var exist []string
	for _, d := range trimmed {
		if _, ok := existMap[d]; ok {
			exist = append(exist, d)
		}
	}

	return exist
}

func trimDomains(domains []string) []string {

	result := make([]string, 0, len(domains))
	for _, d := range domains {
		if t := gstr.Trim(d); t != "" {
			result = append(result, t)
		}
	}

	return result
}

// 根据代理商ID获取站点配置列表
func (s *sSiteConfig) GetSiteConfigsByRid(ctx context.Context, rid int) []*entity.SiteConfig {

	if rid == 0 {
		return []*entity.SiteConfig{}
	}

	siteConfigs, err := dao.SiteConfig.Find(ctx, bson.M{"user_id": rid, "status": 1}, &dao.FindOptions{SortFields: []string{"-updated_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return []*entity.SiteConfig{}
	}

	return siteConfigs
}
