package common

import (
	"context"
	"math"

	"github.com/iimeta/fastapi-admin/v2/internal/consts"
	"github.com/iimeta/fastapi-admin/v2/internal/model/common"
	smodel "github.com/iimeta/fastapi-sdk/v2/model"
)

// 计算花费
func Billing(ctx context.Context, usage smodel.Usage, spend *common.Spend, isBatch ...bool) {

	for _, billingItem := range spend.BillingItems {
		switch billingItem {
		case "text":
			text(ctx, usage, spend)
		case "text_cache":
			textCache(ctx, usage, spend)
		case "image":
			image(ctx, usage, spend)
		case "image_generation":
			imageGeneration(ctx, usage, spend)
		case "image_cache":
			imageCache(ctx, usage, spend)
		case "video_generation":
			videoGeneration(ctx, usage, spend)
		case "once":
			once(ctx, usage, spend)
		}
	}

	if spend.Text != nil {
		if len(isBatch) > 0 && isBatch[0] {
			spend.Text.SpendTokens *= 0.5
		}
		spend.TotalSpendTokens += spend.Text.SpendTokens
	}

	if spend.TextCache != nil {
		if len(isBatch) > 0 && isBatch[0] {
			spend.TextCache.SpendTokens *= 0.5
		}
		spend.TotalSpendTokens += spend.TextCache.SpendTokens
	}

	if spend.Image != nil {
		spend.TotalSpendTokens += spend.Image.SpendTokens
	}

	if spend.ImageCache != nil {
		spend.TotalSpendTokens += spend.ImageCache.SpendTokens
	}

	if spend.ImageGeneration != nil {
		spend.TotalSpendTokens = spend.ImageGeneration.SpendTokens
	}

	if spend.VideoGeneration != nil {
		spend.TotalSpendTokens += spend.VideoGeneration.SpendTokens
	}

	if spend.Once != nil {
		spend.TotalSpendTokens = spend.Once.SpendTokens
	}

	// 模型时段折扣
	if spend.ModelTimeRule != nil {
		spend.TotalSpendTokens = math.Ceil(spend.TotalSpendTokens * spend.ModelTimeRule.Discount)
	}

	// 分组时段折扣
	if spend.GroupId != "" && spend.GroupTimeRule != nil {
		spend.TotalSpendTokens = math.Ceil(spend.TotalSpendTokens * spend.GroupTimeRule.Discount)
	}
}

// 文本
func text(ctx context.Context, usage smodel.Usage, spend *common.Spend) {
	spend.Text.InputTokens = usage.InputTokensDetails.TextTokens
	spend.Text.OutputTokens = usage.OutputTokensDetails.TextTokens
	spend.Text.ReasoningTokens = usage.OutputTokensDetails.ReasoningTokens
	spend.Text.SpendTokens = math.Ceil(float64(spend.Text.InputTokens)*spend.Text.Pricing.InputRatio) + math.Ceil(float64(spend.Text.OutputTokens)*spend.Text.Pricing.OutputRatio) + math.Ceil(float64(spend.Text.ReasoningTokens)*spend.Text.Pricing.ReasoningRatio)
}

// 文本缓存
func textCache(ctx context.Context, usage smodel.Usage, spend *common.Spend) {

	if usage.InputTokensDetails.CachedTokens > 0 {
		spend.TextCache.ReadTokens += usage.InputTokensDetails.CachedTokens
	}

	spend.TextCache.SpendTokens = math.Ceil(float64(spend.TextCache.ReadTokens) * spend.TextCache.Pricing.ReadRatio)
}

// 图像
func image(ctx context.Context, usage smodel.Usage, spend *common.Spend) {

	if usage.InputTokensDetails.ImageTokens > 0 {
		spend.Image.InputTokens += usage.InputTokensDetails.ImageTokens
	}

	if usage.OutputTokensDetails.ImageTokens > 0 {
		spend.Image.OutputTokens += usage.OutputTokensDetails.ImageTokens
	} else if usage.CompletionTokensDetails.ImageTokens > 0 {
		spend.Image.OutputTokens += usage.CompletionTokensDetails.ImageTokens
	}

	spend.Image.SpendTokens = math.Ceil(float64(spend.Image.InputTokens)*spend.Image.Pricing.InputRatio) + math.Ceil(float64(spend.Image.OutputTokens)*spend.Image.Pricing.OutputRatio)
}

// 图像生成
func imageGeneration(ctx context.Context, usage smodel.Usage, spend *common.Spend) {
	spend.ImageGeneration.SpendTokens = math.Ceil(consts.QUOTA_DEFAULT_UNIT*spend.ImageGeneration.Pricing.OnceRatio) * float64(spend.ImageGeneration.N)
}

// 图像缓存
func imageCache(ctx context.Context, usage smodel.Usage, spend *common.Spend) {
	spend.ImageCache.ReadTokens = usage.InputTokensDetails.CachedTokens
	spend.ImageCache.SpendTokens = math.Ceil(float64(spend.ImageCache.ReadTokens) * spend.ImageCache.Pricing.ReadRatio)
}

// 视频生成
func videoGeneration(ctx context.Context, usage smodel.Usage, spend *common.Spend) {
	spend.VideoGeneration.InputTokens = usage.CompletionTokens
	spend.VideoGeneration.SpendTokens = math.Ceil(float64(spend.VideoGeneration.InputTokens) * ConvRatio(spend.VideoGeneration.Pricing.OnceRatio))
}

// 一次
func once(ctx context.Context, usage smodel.Usage, spend *common.Spend) {
	spend.Once.SpendTokens = math.Ceil(consts.QUOTA_DEFAULT_UNIT * spend.Once.Pricing.OnceRatio)
}
