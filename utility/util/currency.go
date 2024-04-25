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
	return Round(1000/(consts.QUOTA_USD_UNIT/ratio), 4)
}
