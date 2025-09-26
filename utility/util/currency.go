package util

import (
	"math"

	"github.com/iimeta/fastapi-admin/internal/consts"
)

// 价格单位M换算成k的倍率
func ConvRatio(price float64) float64 {

	if price == 0 {
		return 0.0
	}

	return price / 1000 * 500
}

// 倍率单位k换算成M的价格
func ConvPrice(ratio float64) float64 {

	if ratio == 0.0 {
		return ratio
	}

	return Round(1000*1000/(consts.QUOTA_USD_UNIT/ratio), 6)
}

func ConvQuota(quota int) float64 {

	if quota == 0 {
		return 0.0
	}

	return Round(float64(quota)/consts.QUOTA_USD_UNIT, 6)
}

func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}
