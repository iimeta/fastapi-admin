package common

import (
	"context"
	"math"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
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
		case "layer_decomp":
			layerDecomp(ctx, usage, spend)
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

	if spend.LayerDecomp != nil {
		spend.TotalSpendTokens += spend.LayerDecomp.SpendTokens
	}

	if spend.VideoGeneration != nil {
		spend.TotalSpendTokens += spend.VideoGeneration.SpendTokens
	}

	if spend.Once != nil {
		spend.TotalSpendTokens = spend.Once.SpendTokens
	}

	// 模型时段折扣
	if spend.ModelTimeRule != nil {
		spend.TotalSpendTokens = discountTokens(spend.TotalSpendTokens, spend.ModelTimeRule.Discount)
	}

	// 分组时段折扣
	if spend.GroupId != "" && spend.GroupTimeRule != nil {
		spend.TotalSpendTokens = discountTokens(spend.TotalSpendTokens, spend.GroupTimeRule.Discount)
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
	if spend.ImageGeneration == nil || spend.ImageGeneration.Pricing == nil {
		return
	}
	spend.ImageGeneration.SpendTokens = math.Ceil(consts.QUOTA_DEFAULT_UNIT*spend.ImageGeneration.Pricing.OnceRatio) * float64(spend.ImageGeneration.N)
}

// 图层拆分
func layerDecomp(ctx context.Context, usage smodel.Usage, spend *common.Spend) {
	if spend.LayerDecomp == nil || spend.LayerDecomp.Pricing == nil {
		return
	}
	spend.LayerDecomp.SpendTokens = math.Ceil(consts.QUOTA_DEFAULT_UNIT*spend.LayerDecomp.Pricing.OnceRatio) * float64(spend.LayerDecomp.N)
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

// 按折扣计算消费 token, 全程使用整数运算, 避免浮点精度误差
func discountTokens(tokens, discount float64) float64 {

	const scale = 1_000_000 // 折扣精度: 保留 6 位小数

	// 折扣先四舍五入成整数基数, 消除小数本身的表示误差
	basis := int64(math.Round(discount * scale))

	// 整数向上取整除法: ceil(a/b) = (a + b - 1) / b (a, b 均非负)
	return float64((int64(tokens)*basis + scale - 1) / scale)
}

func RecalcImageOutputSpend(spend *common.Spend, sizes []string, pricing common.Pricing) {

	if spend == nil || len(sizes) == 0 {
		return
	}

	quality := ""
	if spend.ImageGeneration != nil && spend.ImageGeneration.Pricing != nil {
		quality = spend.ImageGeneration.Pricing.Quality
	} else if spend.LayerDecomp != nil && spend.LayerDecomp.Pricing != nil {
		quality = spend.LayerDecomp.Pricing.Quality
	}

	if spend.ImageGeneration != nil && len(pricing.ImageGeneration) > 0 {

		total := 0.0
		var last *common.ImageGenerationPricing

		for _, sz := range sizes {

			w, h := parseAdminSizeWH(sz)

			p := pickAdminImageGenerationPricing(pricing.ImageGeneration, quality, w, h)
			if p != nil {
				last = p
				total += math.Ceil(consts.QUOTA_DEFAULT_UNIT * p.OnceRatio)
			}
		}

		spend.ImageGeneration.N = len(sizes)

		if last != nil {
			spend.ImageGeneration.Pricing = last
		}

		spend.ImageGeneration.SpendTokens = total
	}

	if spend.LayerDecomp != nil && len(pricing.LayerDecomp) > 0 {

		total := 0.0
		var last *common.LayerDecompPricing

		for _, sz := range sizes {

			w, h := parseAdminSizeWH(sz)

			p := pickAdminLayerDecompPricing(pricing.LayerDecomp, quality, w, h)
			if p != nil {
				last = p
				total += math.Ceil(consts.QUOTA_DEFAULT_UNIT * p.OnceRatio)
			}
		}

		spend.LayerDecomp.N = len(sizes)

		if last != nil {
			spend.LayerDecomp.Pricing = last
		}

		spend.LayerDecomp.SpendTokens = total
	}

	spend.TotalSpendTokens = 0

	if spend.Text != nil {
		spend.TotalSpendTokens += spend.Text.SpendTokens
	}

	if spend.TextCache != nil {
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

	if spend.LayerDecomp != nil {
		spend.TotalSpendTokens += spend.LayerDecomp.SpendTokens
	}

	if spend.VideoGeneration != nil {
		spend.TotalSpendTokens += spend.VideoGeneration.SpendTokens
	}

	if spend.Once != nil {
		spend.TotalSpendTokens = spend.Once.SpendTokens
	}

	if spend.ModelTimeRule != nil {
		spend.TotalSpendTokens = discountTokens(spend.TotalSpendTokens, spend.ModelTimeRule.Discount)
	}

	if spend.GroupId != "" && spend.GroupTimeRule != nil {
		spend.TotalSpendTokens = discountTokens(spend.TotalSpendTokens, spend.GroupTimeRule.Discount)
	}
}

func parseAdminSizeWH(size string) (width, height int) {
	size = gstr.Trim(size)
	if size == "" {
		return 0, 0
	}
	size = gstr.ReplaceByMap(size, map[string]string{"×": "x", "X": "x", "*": "x"})
	parts := gstr.Split(size, "x")
	if len(parts) == 2 {
		return gconv.Int(gstr.Trim(parts[0])), gconv.Int(gstr.Trim(parts[1]))
	}
	return 0, 0
}

func parseAdminPixelSize(size string) int {
	w, h := parseAdminSizeWH(size)
	if w > 0 && h > 0 {
		return w * h
	}
	return gconv.Int(gstr.Trim(size))
}

func matchAdminPixel(pixels int, gte, lte string) bool {
	if pixels <= 0 {
		return false
	}
	g := parseAdminPixelSize(gte)
	l := parseAdminPixelSize(lte)
	if g > 0 && pixels < g {
		return false
	}
	if l > 0 && pixels > l {
		return false
	}
	return true
}

func pickAdminImageGenerationPricing(pricings []*common.ImageGenerationPricing, quality string, width, height int) *common.ImageGenerationPricing {
	pixels := width * height
	var fallback *common.ImageGenerationPricing
	for _, item := range pricings {
		qualityMatch := item.Quality == quality || item.Quality == ""
		if item.Mode == "pixel" {
			if qualityMatch && matchAdminPixel(pixels, item.PixelGte, item.PixelLte) {
				return item
			}
		} else if qualityMatch && item.Width == width && item.Height == height {
			return item
		}
		if item.IsDefault {
			fallback = item
		}
	}
	return fallback
}

func pickAdminLayerDecompPricing(pricings []*common.LayerDecompPricing, quality string, width, height int) *common.LayerDecompPricing {
	pixels := width * height
	var fallback *common.LayerDecompPricing
	for _, item := range pricings {
		qualityMatch := item.Quality == quality || item.Quality == ""
		if item.Mode == "pixel" {
			if qualityMatch && matchAdminPixel(pixels, item.PixelGte, item.PixelLte) {
				return item
			}
		} else if qualityMatch && item.Width == width && item.Height == height {
			return item
		}
		if item.IsDefault {
			fallback = item
		}
	}
	return fallback
}
