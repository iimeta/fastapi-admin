package util

import (
	"github.com/iimeta/fastapi-admin/internal/consts"
	"math"
)

func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

func PriceConv(ratio float64) float64 {

	if ratio == 0.0 {
		return ratio
	}

	return Round(1000/(consts.QUOTA_USD_UNIT/ratio), 6)
}

func QuotaConv(quota int) float64 {

	if quota == 0 {
		return 0.0
	}

	return Round(float64(quota)/consts.QUOTA_USD_UNIT, 6)
}
