package common

import (
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/model/common"
	"github.com/iimeta/fastapi-admin/utility/util"
)

// 价格转倍率
func ConvRatio(price float64) float64 {

	if price == 0 {
		return 0.0
	}

	return price * consts.QUOTA_DEFAULT_UNIT / 1000000
}

// 倍率转价格
func ConvPrice(ratio float64) float64 {

	if ratio == 0.0 {
		return ratio
	}

	return util.Round(ratio*1000000/consts.QUOTA_DEFAULT_UNIT, 6)
}

// 转换额度单位
func ConvQuotaUnit(quota float64) int {

	if quota != 0 {
		return int(quota * consts.QUOTA_DEFAULT_UNIT)
	}

	return 0
}

// 转换额度单位(反向)
func ConvQuotaUnitReverse(quota int, n ...int) float64 {

	if quota == 0 {
		return 0.0
	}

	if len(n) == 0 {
		n = []int{6}
	}

	return util.Round(float64(quota)/consts.QUOTA_DEFAULT_UNIT, n[0])
}

// 转换模型定价成倍率
func ConvModelPricingToRatio(pricing common.Pricing) common.Pricing {

	// 文本
	if pricing.Text != nil {
		for i, text := range pricing.Text {
			pricing.Text[i].InputRatio = ConvRatio(text.InputRatio)
			pricing.Text[i].OutputRatio = ConvRatio(text.OutputRatio)
		}
	}

	// 文本缓存
	if pricing.TextCache != nil {
		for i, textCache := range pricing.TextCache {
			pricing.TextCache[i].ReadRatio = ConvRatio(textCache.ReadRatio)
			pricing.TextCache[i].WriteRatio = ConvRatio(textCache.WriteRatio)
		}
	}

	// 阶梯文本
	if pricing.TieredText != nil {
		for i, tieredText := range pricing.TieredText {
			pricing.TieredText[i].Gt *= 1000
			pricing.TieredText[i].Lte *= 1000
			pricing.TieredText[i].InputRatio = ConvRatio(tieredText.InputRatio)
			pricing.TieredText[i].OutputRatio = ConvRatio(tieredText.OutputRatio)
		}
	}

	// 阶梯文本缓存
	if pricing.TieredTextCache != nil {
		for i, tieredTextCache := range pricing.TieredTextCache {
			pricing.TieredTextCache[i].Gt *= 1000
			pricing.TieredTextCache[i].Lte *= 1000
			pricing.TieredTextCache[i].ReadRatio = ConvRatio(tieredTextCache.ReadRatio)
			pricing.TieredTextCache[i].WriteRatio = ConvRatio(tieredTextCache.WriteRatio)
		}
	}

	// 图像
	if pricing.Image != nil {
		pricing.Image.InputRatio = ConvRatio(pricing.Image.InputRatio)
		pricing.Image.OutputRatio = ConvRatio(pricing.Image.OutputRatio)
	}

	// 图像缓存
	if pricing.ImageCache != nil {
		pricing.ImageCache.ReadRatio = ConvRatio(pricing.ImageCache.ReadRatio)
		pricing.ImageCache.WriteRatio = ConvRatio(pricing.ImageCache.WriteRatio)
	}

	// 音频
	if pricing.Audio != nil {
		pricing.Audio.InputRatio = ConvRatio(pricing.Audio.InputRatio)
		pricing.Audio.OutputRatio = ConvRatio(pricing.Audio.OutputRatio)
	}

	// 音频缓存
	if pricing.AudioCache != nil {
		pricing.AudioCache.ReadRatio = ConvRatio(pricing.AudioCache.ReadRatio)
		pricing.AudioCache.WriteRatio = ConvRatio(pricing.AudioCache.WriteRatio)
	}

	// 视频
	if pricing.Video != nil {
		pricing.Video.InputRatio = ConvRatio(pricing.Video.InputRatio)
		pricing.Video.OutputRatio = ConvRatio(pricing.Video.OutputRatio)
	}

	// 视频缓存
	if pricing.VideoCache != nil {
		pricing.VideoCache.ReadRatio = ConvRatio(pricing.VideoCache.ReadRatio)
		pricing.VideoCache.WriteRatio = ConvRatio(pricing.VideoCache.WriteRatio)
	}

	return pricing
}

// 转换模型定价成价格
func ConvModelPricingToPrice(pricing common.Pricing) common.Pricing {

	// 文本
	if pricing.Text != nil {
		for i, text := range pricing.Text {
			pricing.Text[i].InputRatio = ConvPrice(text.InputRatio)
			pricing.Text[i].OutputRatio = ConvPrice(text.OutputRatio)
		}
	}

	// 文本缓存
	if pricing.TextCache != nil {
		for i, textCache := range pricing.TextCache {
			pricing.TextCache[i].ReadRatio = ConvPrice(textCache.ReadRatio)
			pricing.TextCache[i].WriteRatio = ConvPrice(textCache.WriteRatio)
		}
	}

	// 阶梯文本
	if pricing.TieredText != nil {
		for i, tieredText := range pricing.TieredText {
			pricing.TieredText[i].Gt /= 1000
			pricing.TieredText[i].Lte /= 1000
			pricing.TieredText[i].InputRatio = ConvPrice(tieredText.InputRatio)
			pricing.TieredText[i].OutputRatio = ConvPrice(tieredText.OutputRatio)
		}
	}

	// 阶梯文本缓存
	if pricing.TieredTextCache != nil {
		for i, tieredTextCache := range pricing.TieredTextCache {
			pricing.TieredTextCache[i].Gt /= 1000
			pricing.TieredTextCache[i].Lte /= 1000
			pricing.TieredTextCache[i].ReadRatio = ConvPrice(tieredTextCache.ReadRatio)
			pricing.TieredTextCache[i].WriteRatio = ConvPrice(tieredTextCache.WriteRatio)
		}
	}

	// 图像
	if pricing.Image != nil {
		pricing.Image.InputRatio = ConvPrice(pricing.Image.InputRatio)
		pricing.Image.OutputRatio = ConvPrice(pricing.Image.OutputRatio)
	}

	// 图像缓存
	if pricing.ImageCache != nil {
		pricing.ImageCache.ReadRatio = ConvPrice(pricing.ImageCache.ReadRatio)
		pricing.ImageCache.WriteRatio = ConvPrice(pricing.ImageCache.WriteRatio)
	}

	// 音频
	if pricing.Audio != nil {
		pricing.Audio.InputRatio = ConvPrice(pricing.Audio.InputRatio)
		pricing.Audio.OutputRatio = ConvPrice(pricing.Audio.OutputRatio)
	}

	// 音频缓存
	if pricing.AudioCache != nil {
		pricing.AudioCache.ReadRatio = ConvPrice(pricing.AudioCache.ReadRatio)
		pricing.AudioCache.WriteRatio = ConvPrice(pricing.AudioCache.WriteRatio)
	}

	// 视频
	if pricing.Video != nil {
		pricing.Video.InputRatio = ConvPrice(pricing.Video.InputRatio)
		pricing.Video.OutputRatio = ConvPrice(pricing.Video.OutputRatio)
	}

	// 视频缓存
	if pricing.VideoCache != nil {
		pricing.VideoCache.ReadRatio = ConvPrice(pricing.VideoCache.ReadRatio)
		pricing.VideoCache.WriteRatio = ConvPrice(pricing.VideoCache.WriteRatio)
	}

	return pricing
}

// 转换花费
func ConvSpend(spend common.Spend) common.Spend {

	// 文本
	if spend.Text != nil {

		if spend.Text.Pricing != nil {
			spend.Text.Pricing.InputRatio = ConvPrice(spend.Text.Pricing.InputRatio)
			spend.Text.Pricing.OutputRatio = ConvPrice(spend.Text.Pricing.OutputRatio)
		}

		spend.Text.SpendTokens = ConvQuotaUnitReverse(int(spend.Text.SpendTokens))
	}

	// 文本缓存
	if spend.TextCache != nil {

		if spend.TextCache.Pricing != nil {
			spend.TextCache.Pricing.ReadRatio = ConvPrice(spend.TextCache.Pricing.ReadRatio)
			spend.TextCache.Pricing.WriteRatio = ConvPrice(spend.TextCache.Pricing.WriteRatio)
		}

		spend.TextCache.SpendTokens = ConvQuotaUnitReverse(int(spend.TextCache.SpendTokens))
	}

	// 阶梯文本
	if spend.TieredText != nil {

		if spend.TieredText.Pricing != nil {
			spend.TieredText.Pricing.Gt /= 1000
			spend.TieredText.Pricing.Lte /= 1000
			spend.TieredText.Pricing.InputRatio = ConvPrice(spend.TieredText.Pricing.InputRatio)
			spend.TieredText.Pricing.OutputRatio = ConvPrice(spend.TieredText.Pricing.OutputRatio)
		}

		spend.TieredText.SpendTokens = ConvQuotaUnitReverse(int(spend.TieredText.SpendTokens))
	}

	// 阶梯文本缓存
	if spend.TieredTextCache != nil {

		if spend.TieredTextCache.Pricing != nil {
			spend.TieredTextCache.Pricing.Gt /= 1000
			spend.TieredTextCache.Pricing.Lte /= 1000
			spend.TieredTextCache.Pricing.ReadRatio = ConvPrice(spend.TieredTextCache.Pricing.ReadRatio)
			spend.TieredTextCache.Pricing.WriteRatio = ConvPrice(spend.TieredTextCache.Pricing.WriteRatio)
		}

		spend.TieredTextCache.SpendTokens = ConvQuotaUnitReverse(int(spend.TieredTextCache.SpendTokens))
	}

	// 图像
	if spend.Image != nil {

		if spend.Image.Pricing != nil {
			spend.Image.Pricing.InputRatio = ConvPrice(spend.Image.Pricing.InputRatio)
			spend.Image.Pricing.OutputRatio = ConvPrice(spend.Image.Pricing.OutputRatio)
		}

		spend.Image.SpendTokens = ConvQuotaUnitReverse(int(spend.Image.SpendTokens))
	}

	// 图像生成
	if spend.ImageGeneration != nil {
		spend.ImageGeneration.SpendTokens = ConvQuotaUnitReverse(int(spend.ImageGeneration.SpendTokens))
	}

	// 图像缓存
	if spend.ImageCache != nil {

		if spend.ImageCache.Pricing != nil {
			spend.ImageCache.Pricing.ReadRatio = ConvPrice(spend.ImageCache.Pricing.ReadRatio)
			spend.ImageCache.Pricing.WriteRatio = ConvPrice(spend.ImageCache.Pricing.WriteRatio)
		}

		spend.ImageCache.SpendTokens = ConvQuotaUnitReverse(int(spend.ImageCache.SpendTokens))
	}

	// 识图
	if spend.Vision != nil {
		spend.Vision.SpendTokens = ConvQuotaUnitReverse(int(spend.Vision.SpendTokens))
	}

	// 音频
	if spend.Audio != nil {

		if spend.Audio.Pricing != nil {
			spend.Audio.Pricing.InputRatio = ConvPrice(spend.Audio.Pricing.InputRatio)
			spend.Audio.Pricing.OutputRatio = ConvPrice(spend.Audio.Pricing.OutputRatio)
		}

		spend.Audio.SpendTokens = ConvQuotaUnitReverse(int(spend.Audio.SpendTokens))
	}

	// 音频缓存
	if spend.AudioCache != nil {

		if spend.AudioCache.Pricing != nil {
			spend.AudioCache.Pricing.ReadRatio = ConvPrice(spend.AudioCache.Pricing.ReadRatio)
			spend.AudioCache.Pricing.WriteRatio = ConvPrice(spend.AudioCache.Pricing.WriteRatio)
		}

		spend.AudioCache.SpendTokens = ConvQuotaUnitReverse(int(spend.AudioCache.SpendTokens))
	}

	// 视频
	if spend.Video != nil {

		if spend.Video.Pricing != nil {
			spend.Video.Pricing.InputRatio = ConvPrice(spend.Video.Pricing.InputRatio)
			spend.Video.Pricing.OutputRatio = ConvPrice(spend.Video.Pricing.OutputRatio)
		}

		spend.Video.SpendTokens = ConvQuotaUnitReverse(int(spend.Video.SpendTokens))
	}

	// 视频生成
	if spend.VideoGeneration != nil {
		spend.VideoGeneration.SpendTokens = ConvQuotaUnitReverse(int(spend.VideoGeneration.SpendTokens))
	}

	// 视频缓存
	if spend.VideoCache != nil {

		if spend.VideoCache.Pricing != nil {
			spend.VideoCache.Pricing.ReadRatio = ConvPrice(spend.VideoCache.Pricing.ReadRatio)
			spend.VideoCache.Pricing.WriteRatio = ConvPrice(spend.VideoCache.Pricing.WriteRatio)
		}

		spend.VideoCache.SpendTokens = ConvQuotaUnitReverse(int(spend.VideoCache.SpendTokens))
	}

	// 搜索
	if spend.Search != nil {
		spend.Search.SpendTokens = ConvQuotaUnitReverse(int(spend.Search.SpendTokens))
	}

	// Midjourney
	if spend.Midjourney != nil {
		spend.Midjourney.SpendTokens = ConvQuotaUnitReverse(int(spend.Midjourney.SpendTokens))
	}

	// 一次
	if spend.Once != nil {
		spend.Once.SpendTokens = ConvQuotaUnitReverse(int(spend.Once.SpendTokens))
	}

	// 总花费
	spend.TotalSpendTokens = ConvQuotaUnitReverse(int(spend.TotalSpendTokens))

	return spend
}
