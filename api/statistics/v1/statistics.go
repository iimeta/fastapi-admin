package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 用户数据接口请求参数
type DataUserReq struct {
	g.Meta `path:"/data/user" method:"post" auth:"true" role:"admin" tags:"statistics" summary:"用户数据接口"`
	model.StatisticsDataReq
}

// 用户数据接口响应参数
type DataUserRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.StatisticsDataRes
}

// 应用数据接口请求参数
type DataAppReq struct {
	g.Meta `path:"/data/app" method:"post" auth:"true" role:"admin" tags:"statistics" summary:"应用数据接口"`
	model.StatisticsDataReq
}

// 应用数据接口响应参数
type DataAppRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.StatisticsDataRes
}

// 应用密钥数据接口请求参数
type DataAppKeyReq struct {
	g.Meta `path:"/data/app/key" method:"post" auth:"true" role:"admin" tags:"statistics" summary:"应用密钥数据接口"`
	model.StatisticsDataReq
}

// 应用密钥数据接口响应参数
type DataAppKeyRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.StatisticsDataRes
}
