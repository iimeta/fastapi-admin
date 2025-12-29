package common

import (
	"context"
	"math"

	"github.com/iimeta/fastapi-admin/internal/model/common"
	smodel "github.com/iimeta/fastapi-sdk/model"
)

// 计算花费
func Billing(ctx context.Context, usage smodel.Usage, spend *common.Spend) {

	for _, billingItem := range spend.BillingItems {
		switch billingItem {
		case "text":
			text(ctx, usage, spend)
		case "text_cache":
			textCache(ctx, usage, spend)
		}
	}

	if spend.Text != nil {
		spend.Text.SpendTokens *= 0.5
		spend.TotalSpendTokens += spend.Text.SpendTokens
	}

	if spend.TextCache != nil {
		spend.TextCache.SpendTokens *= 0.5
		spend.TotalSpendTokens += spend.TextCache.SpendTokens
	}

	// 分组折扣
	if spend.GroupId != "" {
		spend.TotalSpendTokens = math.Ceil(spend.TotalSpendTokens * spend.GroupDiscount)
	}
}

// 文本
func text(ctx context.Context, usage smodel.Usage, spend *common.Spend) {

	spend.Text.InputTokens = usage.InputTokens
	spend.Text.OutputTokens = usage.OutputTokens
	spend.Text.ReasoningTokens = usage.OutputTokensDetails.ReasoningTokens
	spend.Text.SpendTokens = math.Ceil(float64(spend.Text.InputTokens)*spend.Text.Pricing.InputRatio) + math.Ceil(float64(spend.Text.OutputTokens)*spend.Text.Pricing.OutputRatio) + math.Ceil(float64(spend.Text.ReasoningTokens)*spend.Text.Pricing.ReasoningRatio)
}

// 文本缓存
func textCache(ctx context.Context, usage smodel.Usage, spend *common.Spend) {

	if usage.InputTokensDetails.CachedTokens == 0 {
		return
	}

	if spend.TextCache == nil {
		spend.TextCache = new(common.CacheSpend)
	}

	if usage.InputTokensDetails.CachedTokens > 0 {
		spend.TextCache.ReadTokens += usage.InputTokensDetails.CachedTokens
	}

	spend.TextCache.SpendTokens = math.Ceil(float64(spend.TextCache.ReadTokens) * spend.TextCache.Pricing.ReadRatio)
}
