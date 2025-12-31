// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/model/entity"
)

type (
	ISiteConfig interface {
		// 新建站点配置
		Create(ctx context.Context, params model.SiteConfigCreateReq) error
		// 更新站点配置
		Update(ctx context.Context, params model.SiteConfigUpdateReq) error
		// 更改站点配置状态
		ChangeStatus(ctx context.Context, params model.SiteConfigChangeStatusReq) error
		// 删除站点配置
		Delete(ctx context.Context, id string) error
		// 站点配置详情
		Detail(ctx context.Context, params model.SiteConfigDetailReq) (*model.SiteConfig, error)
		// 站点配置分页列表
		Page(ctx context.Context, params model.SiteConfigPageReq) (*model.SiteConfigPageRes, error)
		// 站点配置批量操作
		BatchOperate(ctx context.Context, params model.SiteConfigBatchOperateReq) error
		// 站点配置
		Site(ctx context.Context, params model.SiteConfigDetailReq) *model.SiteConfig
		// 根据域名获取站点配置
		GetSiteConfigByDomain(ctx context.Context, domain string) *entity.SiteConfig
		// 站点域名是否存在
		IsDomainExist(ctx context.Context, domain string, id ...string) bool
		// 根据代理商ID获取站点配置列表
		GetSiteConfigsByRid(ctx context.Context, rid int) []*entity.SiteConfig
	}
)

var (
	localSiteConfig ISiteConfig
)

func SiteConfig() ISiteConfig {
	if localSiteConfig == nil {
		panic("implement not found for interface ISiteConfig, forgot register?")
	}
	return localSiteConfig
}

func RegisterSiteConfig(i ISiteConfig) {
	localSiteConfig = i
}
