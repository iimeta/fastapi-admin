package finance

import (
	"cmp"
	"context"
	"fmt"
	"regexp"
	"slices"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/errors"
	common2 "github.com/iimeta/fastapi-admin/internal/logic/common"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/common"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/db"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/util"
	"go.mongodb.org/mongo-driver/bson"
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

	slices.SortFunc(statisticsUser.ModelStats, func(s1, s2 *common.ModelStat) int {
		return cmp.Compare(s2.Tokens, s1.Tokens)
	})

	return &model.StatisticsUser{
		Id:             statisticsUser.Id,
		UserId:         statisticsUser.UserId,
		StatDate:       statisticsUser.StatDate,
		StatTime:       statisticsUser.StatTime,
		Total:          statisticsUser.Total,
		Tokens:         statisticsUser.Tokens,
		Abnormal:       statisticsUser.Abnormal,
		AbnormalTokens: statisticsUser.AbnormalTokens,
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
			Tokens:   result.Tokens,
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
		filter = bson.M{"_id": bson.M{"$in": params.Ids}}
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
	} else {
		if params.UserId != 0 {
			filter["user_id"] = params.UserId
		}
	}

	results, err := dao.StatisticsUser.Find(ctx, filter, &dao.FindOptions{SortFields: []string{"-stat_date", "-tokens"}})
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	colFieldMap := make(map[string]string)
	colFieldMap["账单日期"] = "StatDate"
	colFieldMap["模型"] = "Model"
	colFieldMap["调用数"] = "Total"
	colFieldMap["花费($)"] = "Tokens"

	var titleCols []string
	if service.Session().IsUserRole(ctx) {
		titleCols = append(titleCols, "账单日期", "模型", "调用数", "花费($)")
	} else {
		titleCols = append(titleCols, "账单日期", "用户ID", "模型", "调用数", "花费($)")
		colFieldMap["用户ID"] = "UserId"
	}

	filePath := fmt.Sprintf("./resource/export/bill_%d.xlsx", gtime.TimestampMilli())

	values := make([]interface{}, 0)
	for _, result := range results {

		slices.SortFunc(result.ModelStats, func(s1, s2 *common.ModelStat) int {
			return cmp.Compare(s2.Tokens, s1.Tokens)
		})

		for _, modelStat := range result.ModelStats {
			values = append(values, &model.BillExport{
				StatDate: result.StatDate,
				UserId:   result.UserId,
				Model:    modelStat.Model,
				Total:    modelStat.Total,
				Tokens:   gconv.String(common2.ConvQuota(modelStat.Tokens)),
			})
		}
	}

	if err = util.ExcelExport("账单明细", titleCols, colFieldMap, values, filePath); err != nil {
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
			Quota:     result.Quota,
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
