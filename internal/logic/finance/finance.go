package finance

import (
	"cmp"
	"context"
	"fmt"
	"regexp"
	"slices"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/internal/errors"
	"github.com/iimeta/fastapi-admin/v2/internal/logic/common"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	mcommon "github.com/iimeta/fastapi-admin/v2/internal/model/common"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
	"github.com/iimeta/fastapi-admin/v2/utility/db"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"github.com/iimeta/fastapi-admin/v2/utility/util"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type sFinance struct{}

func init() {
	service.RegisterFinance(New())
}

func New() service.IFinance {
	return &sFinance{}
}

// 账单明细详情
func (s *sFinance) BillDetail(ctx context.Context, id string) (*model.StatisticsUser, error) {

	statisticsUser, err := dao.StatisticsUser.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if service.Session().IsUserRole(ctx) && statisticsUser.UserId != service.Session().GetUserId(ctx) {
		return nil, errors.New("Unauthorized")
	}

	slices.SortFunc(statisticsUser.ModelStats, func(s1, s2 *mcommon.ModelStat) int {
		return cmp.Compare(s2.Tokens, s1.Tokens)
	})

	for _, modelStat := range statisticsUser.ModelStats {
		modelStat.Tokens = common.ConvQuotaUnitReverse(int(modelStat.Tokens))
		modelStat.AbnormalTokens = common.ConvQuotaUnitReverse(int(modelStat.AbnormalTokens))
	}

	return &model.StatisticsUser{
		Id:             statisticsUser.Id,
		UserId:         statisticsUser.UserId,
		StatDate:       statisticsUser.StatDate,
		StatTime:       statisticsUser.StatTime,
		Total:          statisticsUser.Total,
		Tokens:         common.ConvQuotaUnitReverse(statisticsUser.Tokens),
		Abnormal:       statisticsUser.Abnormal,
		AbnormalTokens: common.ConvQuotaUnitReverse(statisticsUser.AbnormalTokens),
		ModelStats:     statisticsUser.ModelStats,
	}, nil
}

// 账单明细分页列表
func (s *sFinance) BillPage(ctx context.Context, params model.FinanceBillPageReq) (*model.FinanceBillPageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

	if service.Session().IsResellerRole(ctx) {
		filter["rid"] = service.Session().GetRid(ctx)
	}

	if service.Session().IsUserRole(ctx) {
		filter["user_id"] = service.Session().GetUserId(ctx)
	} else if params.UserId != 0 {
		filter["user_id"] = params.UserId
	}

	if len(params.StatDate) > 0 {
		filter["stat_date"] = bson.M{
			"$gte": params.StatDate[0],
			"$lte": params.StatDate[1],
		}
	}

	results, err := dao.StatisticsUser.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"-stat_date", "-tokens"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.Bill, 0)
	for _, result := range results {
		items = append(items, &model.Bill{
			Id:       result.Id,
			UserId:   result.UserId,
			Total:    result.Total,
			Tokens:   common.ConvQuotaUnitReverse(result.Tokens),
			Models:   len(result.ModelStats),
			StatDate: result.StatDate,
		})
	}

	return &model.FinanceBillPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}

// 账单明细导出
func (s *sFinance) BillExport(ctx context.Context, params model.FinanceBillExportReq) (string, error) {

	filter := bson.M{}
	if len(params.Ids) > 0 {

		selectedResults, err := dao.StatisticsUser.Find(ctx, bson.M{"_id": bson.M{"$in": params.Ids}})
		if err != nil {
			logger.Error(ctx, err)
			return "", err
		}

		orFilters := make([]bson.M, 0)
		for _, result := range selectedResults {
			orFilters = append(orFilters, bson.M{"user_id": result.UserId, "stat_date": result.StatDate})
		}

		if len(orFilters) == 0 {
			return "", errors.New("selected bill data not found")
		}

		filter["$or"] = orFilters

	} else {
		filter = bson.M{
			"stat_date": bson.M{
				"$gte": params.StatDate[0],
				"$lte": params.StatDate[1],
			},
		}
	}

	if service.Session().IsResellerRole(ctx) {
		filter["rid"] = service.Session().GetRid(ctx)
	}

	if service.Session().IsUserRole(ctx) {
		filter["user_id"] = service.Session().GetUserId(ctx)
	} else if len(params.Ids) == 0 && params.UserId != 0 {
		filter["user_id"] = params.UserId
	}

	colFieldMap := make(map[string]string)
	colFieldMap["账单日期"] = "StatDate"
	colFieldMap["模型"] = "Model"
	colFieldMap["调用数"] = "Total"
	colFieldMap["花费"] = "Tokens"

	titleCols := []string{"账单日期"}

	if !service.Session().IsUserRole(ctx) {
		titleCols = append(titleCols, "用户ID")
		colFieldMap["用户ID"] = "UserId"
	}

	if params.DataType == "app" || params.DataType == "app_key" {
		titleCols = append(titleCols, "应用ID")
		colFieldMap["应用ID"] = "AppId"
	}

	if params.DataType == "app_key" {
		titleCols = append(titleCols, "应用密钥")
		colFieldMap["应用密钥"] = "AppKey"
	}

	titleCols = append(titleCols, "模型", "调用数", "花费")

	filePath := fmt.Sprintf("./resource/export/bill_%d.xlsx", gtime.TimestampMilli())

	values := make([]any, 0)
	appendValues := func(statDate string, userId int, appId int, appKey string, modelStats []*mcommon.ModelStat) {

		slices.SortFunc(modelStats, func(s1, s2 *mcommon.ModelStat) int {
			return cmp.Compare(s2.Tokens, s1.Tokens)
		})

		for _, modelStat := range modelStats {
			values = append(values, &model.BillExport{
				StatDate: statDate,
				UserId:   userId,
				AppId:    appId,
				AppKey:   appKey,
				Model:    modelStat.Model,
				Total:    modelStat.Total,
				Tokens:   common.ConvQuotaUnitReverse(int(modelStat.Tokens)),
			})
		}
	}

	findOptions := &dao.FindOptions{SortFields: []string{"-stat_date", "-tokens"}}

	switch params.DataType {
	case "app":

		results, err := dao.StatisticsApp.Find(ctx, filter, findOptions)
		if err != nil {
			logger.Error(ctx, err)
			return "", err
		}

		for _, result := range results {
			appendValues(result.StatDate, result.UserId, result.AppId, "", result.ModelStats)
		}

	case "app_key":

		results, err := dao.StatisticsAppKey.Find(ctx, filter, findOptions)
		if err != nil {
			logger.Error(ctx, err)
			return "", err
		}

		for _, result := range results {
			appendValues(result.StatDate, result.UserId, result.AppId, result.AppKey, result.ModelStats)
		}

	default:

		results, err := dao.StatisticsUser.Find(ctx, filter, findOptions)
		if err != nil {
			logger.Error(ctx, err)
			return "", err
		}

		for _, result := range results {
			appendValues(result.StatDate, result.UserId, 0, "", result.ModelStats)
		}
	}

	if err := util.ExcelExport("账单明细", titleCols, colFieldMap, values, filePath); err != nil {
		return "", err
	}

	return filePath, nil
}

// 交易记录分页列表
func (s *sFinance) DealRecordPage(ctx context.Context, params model.FinanceDealRecordPageReq) (*model.FinanceDealRecordPageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

	if service.Session().IsResellerRole(ctx) {
		filter["rid"] = service.Session().GetRid(ctx)
	}

	if service.Session().IsUserRole(ctx) {
		filter["user_id"] = service.Session().GetUserId(ctx)
	} else if params.UserId != 0 {
		filter["user_id"] = params.UserId
	}

	if params.Type != 0 {
		filter["type"] = params.Type
	}

	if params.Status != 0 {
		filter["status"] = params.Status
	}

	if params.Remark != "" {
		filter["remark"] = bson.M{
			"$regex": regexp.QuoteMeta(params.Remark),
		}
	}

	if len(params.CreatedAt) > 0 {
		gte := gtime.NewFromStrFormat(params.CreatedAt[0], time.DateOnly).TimestampMilli()
		lte := gtime.NewFromStrLayout(params.CreatedAt[1], time.DateOnly).EndOfDay(true).TimestampMilli()
		filter["created_at"] = bson.M{
			"$gte": gte,
			"$lte": lte,
		}
	}

	results, err := dao.DealRecord.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"-updated_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.DealRecord, 0)
	for _, result := range results {
		items = append(items, &model.DealRecord{
			Id:        result.Id,
			UserId:    result.UserId,
			Quota:     common.ConvQuotaUnitReverse(result.Quota),
			Type:      result.Type,
			Remark:    result.Remark,
			Status:    result.Status,
			CreatedAt: util.FormatDateTime(result.CreatedAt),
			UpdatedAt: util.FormatDateTime(result.UpdatedAt),
		})
	}

	return &model.FinanceDealRecordPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}
