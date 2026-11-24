package common

import (
	"fmt"
	"sort"
	"strings"

	"github.com/iimeta/fastapi-admin/v2/internal/config"
	"github.com/iimeta/fastapi-admin/v2/internal/consts"
)

func NormalizePositiveInts(values []int) []int {

	if len(values) == 0 {
		return []int{}
	}

	seen := make(map[int]struct{}, len(values))
	result := make([]int, 0, len(values))

	for _, value := range values {

		if value <= 0 {
			continue
		}

		if _, ok := seen[value]; ok {
			continue
		}

		seen[value] = struct{}{}
		result = append(result, value)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(result)))

	return result
}

func ConvQuotaUnits(values []int) []int {

	result := make([]int, 0, len(values))
	for _, value := range NormalizePositiveInts(values) {
		result = append(result, value*consts.QUOTA_DEFAULT_UNIT)
	}

	return result
}

func ConvQuotaUnitsReverse(values []int) []int {

	result := make([]int, 0, len(values))
	for _, value := range values {

		if value <= 0 {
			continue
		}

		result = append(result, value/consts.QUOTA_DEFAULT_UNIT)
	}

	return NormalizePositiveInts(result)
}

func QuotaWarningThresholds(values []int) []int {

	if len(values) > 0 {
		return values
	}

	if config.Cfg.Quota == nil {
		return []int{}
	}

	return config.Cfg.Quota.Thresholds
}

func QuotaExpireWarningThresholds(values []int) []int {

	if len(values) > 0 {
		return values
	}

	if config.Cfg.Quota == nil {
		return []int{}
	}

	return config.Cfg.Quota.ExpiredThresholds
}

func FormatQuotaThresholds(values []int, currencySymbol string) string {

	items := ConvQuotaUnitsReverse(QuotaWarningThresholds(values))
	if len(items) == 0 {
		return "-"
	}

	parts := make([]string, 0, len(items))
	for _, item := range items {
		parts = append(parts, fmt.Sprintf("%s%d", currencySymbol, item))
	}

	return strings.Join(parts, ", ")
}

func FormatExpireWarningThresholds(values []int) string {

	items := QuotaExpireWarningThresholds(values)
	if len(items) == 0 {
		return "-"
	}

	parts := make([]string, 0, len(items))
	for _, item := range items {
		parts = append(parts, fmt.Sprintf("%d", item))
	}

	return strings.Join(parts, ", ")
}
