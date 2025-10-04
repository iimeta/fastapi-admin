package model

import (
	"github.com/iimeta/fastapi-admin/internal/model/common"
	"github.com/iimeta/fastapi-admin/utility/util"
)

// 转换定价成倍率
func (s *sModel) ConvPricingToRatio(pricing common.Pricing) common.Pricing {

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

// 转换定价成价格
func (s *sModel) ConvPricingToPrice(pricing common.Pricing) common.Pricing {

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
