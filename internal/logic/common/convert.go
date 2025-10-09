package common

import (
	"github.com/iimeta/fastapi-admin/internal/model/common"
	"github.com/iimeta/fastapi-admin/utility/util"
)

// 转换模型定价成倍率
func ConvModelPricingToRatio(pricing common.Pricing) common.Pricing {

	// 文本
	if pricing.Text != nil {
		pricing.Text.InputRatio = util.ConvRatio(pricing.Text.InputRatio)
		pricing.Text.OutputRatio = util.ConvRatio(pricing.Text.OutputRatio)
	}

	// 文本缓存
	if pricing.TextCache != nil {
		pricing.TextCache.ReadRatio = util.ConvRatio(pricing.TextCache.ReadRatio)
		pricing.TextCache.WriteRatio = util.ConvRatio(pricing.TextCache.WriteRatio)
	}

	// 阶梯文本
	if pricing.TieredText != nil {
		for i, tieredText := range pricing.TieredText {
			pricing.TieredText[i].Gt *= 1000
			pricing.TieredText[i].Lte *= 1000
			pricing.TieredText[i].InputRatio = util.ConvRatio(tieredText.InputRatio)
			pricing.TieredText[i].OutputRatio = util.ConvRatio(tieredText.OutputRatio)
		}
	}

	// 阶梯文本缓存
	if pricing.TieredTextCache != nil {
		for i, tieredTextCache := range pricing.TieredTextCache {
			pricing.TieredTextCache[i].Gt *= 1000
			pricing.TieredTextCache[i].Lte *= 1000
			pricing.TieredTextCache[i].ReadRatio = util.ConvRatio(tieredTextCache.ReadRatio)
			pricing.TieredTextCache[i].WriteRatio = util.ConvRatio(tieredTextCache.WriteRatio)
		}
	}

	// 图像
	if pricing.Image != nil {
		pricing.Image.InputRatio = util.ConvRatio(pricing.Image.InputRatio)
		pricing.Image.OutputRatio = util.ConvRatio(pricing.Image.OutputRatio)
	}

	// 图像缓存
	if pricing.ImageCache != nil {
		pricing.ImageCache.ReadRatio = util.ConvRatio(pricing.ImageCache.ReadRatio)
		pricing.ImageCache.WriteRatio = util.ConvRatio(pricing.ImageCache.WriteRatio)
	}

	// 音频
	if pricing.Audio != nil {
		pricing.Audio.InputRatio = util.ConvRatio(pricing.Audio.InputRatio)
		pricing.Audio.OutputRatio = util.ConvRatio(pricing.Audio.OutputRatio)
	}

	// 音频缓存
	if pricing.AudioCache != nil {
		pricing.AudioCache.ReadRatio = util.ConvRatio(pricing.AudioCache.ReadRatio)
		pricing.AudioCache.WriteRatio = util.ConvRatio(pricing.AudioCache.WriteRatio)
	}

	return pricing
}

// 转换模型定价成价格
func ConvModelPricingToPrice(pricing common.Pricing) common.Pricing {

	// 文本
	if pricing.Text != nil {
		pricing.Text.InputRatio = util.ConvPrice(pricing.Text.InputRatio)
		pricing.Text.OutputRatio = util.ConvPrice(pricing.Text.OutputRatio)
	}

	// 文本缓存
	if pricing.TextCache != nil {
		pricing.TextCache.ReadRatio = util.ConvPrice(pricing.TextCache.ReadRatio)
		pricing.TextCache.WriteRatio = util.ConvPrice(pricing.TextCache.WriteRatio)
	}

	// 阶梯文本
	if pricing.TieredText != nil {
		for i, tieredText := range pricing.TieredText {
			pricing.TieredText[i].Gt /= 1000
			pricing.TieredText[i].Lte /= 1000
			pricing.TieredText[i].InputRatio = util.ConvPrice(tieredText.InputRatio)
			pricing.TieredText[i].OutputRatio = util.ConvPrice(tieredText.OutputRatio)
		}
	}

	// 阶梯文本缓存
	if pricing.TieredTextCache != nil {
		for i, tieredTextCache := range pricing.TieredTextCache {
			pricing.TieredTextCache[i].Gt /= 1000
			pricing.TieredTextCache[i].Lte /= 1000
			pricing.TieredTextCache[i].ReadRatio = util.ConvPrice(tieredTextCache.ReadRatio)
			pricing.TieredTextCache[i].WriteRatio = util.ConvPrice(tieredTextCache.WriteRatio)
		}
	}

	// 图像
	if pricing.Image != nil {
		pricing.Image.InputRatio = util.ConvPrice(pricing.Image.InputRatio)
		pricing.Image.OutputRatio = util.ConvPrice(pricing.Image.OutputRatio)
	}

	// 图像缓存
	if pricing.ImageCache != nil {
		pricing.ImageCache.ReadRatio = util.ConvPrice(pricing.ImageCache.ReadRatio)
		pricing.ImageCache.WriteRatio = util.ConvPrice(pricing.ImageCache.WriteRatio)
	}

	// 音频
	if pricing.Audio != nil {
		pricing.Audio.InputRatio = util.ConvPrice(pricing.Audio.InputRatio)
		pricing.Audio.OutputRatio = util.ConvPrice(pricing.Audio.OutputRatio)
	}

	// 音频缓存
	if pricing.AudioCache != nil {
		pricing.AudioCache.ReadRatio = util.ConvPrice(pricing.AudioCache.ReadRatio)
		pricing.AudioCache.WriteRatio = util.ConvPrice(pricing.AudioCache.WriteRatio)
	}

	return pricing
}

// 转换花费定价成价格
func ConvSpendPricingToPrice(spend common.Spend) common.Spend {

	// 文本
	if spend.Text != nil && spend.Text.Pricing != nil {
		spend.Text.Pricing.InputRatio = util.ConvPrice(spend.Text.Pricing.InputRatio)
		spend.Text.Pricing.OutputRatio = util.ConvPrice(spend.Text.Pricing.OutputRatio)
	}

	// 文本缓存
	if spend.TextCache != nil && spend.TextCache.Pricing != nil {
		spend.TextCache.Pricing.ReadRatio = util.ConvPrice(spend.TextCache.Pricing.ReadRatio)
		spend.TextCache.Pricing.WriteRatio = util.ConvPrice(spend.TextCache.Pricing.WriteRatio)
	}

	// 阶梯文本
	if spend.TieredText != nil && spend.TieredText.Pricing != nil {
		spend.TieredText.Pricing.Gt /= 1000
		spend.TieredText.Pricing.Lte /= 1000
		spend.TieredText.Pricing.InputRatio = util.ConvPrice(spend.TieredText.Pricing.InputRatio)
		spend.TieredText.Pricing.OutputRatio = util.ConvPrice(spend.TieredText.Pricing.OutputRatio)
	}

	// 阶梯文本缓存
	if spend.TieredTextCache != nil && spend.TieredTextCache.Pricing != nil {
		spend.TieredTextCache.Pricing.Gt /= 1000
		spend.TieredTextCache.Pricing.Lte /= 1000
		spend.TieredTextCache.Pricing.ReadRatio = util.ConvPrice(spend.TieredTextCache.Pricing.ReadRatio)
		spend.TieredTextCache.Pricing.WriteRatio = util.ConvPrice(spend.TieredTextCache.Pricing.WriteRatio)
	}

	// 图像
	if spend.Image != nil && spend.Image.Pricing != nil {
		spend.Image.Pricing.InputRatio = util.ConvPrice(spend.Image.Pricing.InputRatio)
		spend.Image.Pricing.OutputRatio = util.ConvPrice(spend.Image.Pricing.OutputRatio)
	}

	// 图像缓存
	if spend.ImageCache != nil && spend.ImageCache.Pricing != nil {
		spend.ImageCache.Pricing.ReadRatio = util.ConvPrice(spend.ImageCache.Pricing.ReadRatio)
		spend.ImageCache.Pricing.WriteRatio = util.ConvPrice(spend.ImageCache.Pricing.WriteRatio)
	}

	// 音频
	if spend.Audio != nil && spend.Audio.Pricing != nil {
		spend.Audio.Pricing.InputRatio = util.ConvPrice(spend.Audio.Pricing.InputRatio)
		spend.Audio.Pricing.OutputRatio = util.ConvPrice(spend.Audio.Pricing.OutputRatio)
	}

	// 音频缓存
	if spend.AudioCache != nil && spend.AudioCache.Pricing != nil {
		spend.AudioCache.Pricing.ReadRatio = util.ConvPrice(spend.AudioCache.Pricing.ReadRatio)
		spend.AudioCache.Pricing.WriteRatio = util.ConvPrice(spend.AudioCache.Pricing.WriteRatio)
	}

	return spend
}
