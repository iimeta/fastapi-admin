// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/internal/model"
)

type (
	IMonitor interface {
		// 监控中心全局实时指标
		Global(ctx context.Context, params model.MonitorGlobalReq) (*model.MonitorGlobalRes, error)
		// 监控中心实时性能维度分析
		PerfBreakdown(ctx context.Context, params model.MonitorPerfBreakdownReq) (*model.MonitorPerfBreakdownRes, error)
		// 监控中心历史性能数据
		PerfHistory(ctx context.Context, params model.MonitorPerfHistoryReq) (*model.MonitorPerfHistoryRes, error)
	}
)

var (
	localMonitor IMonitor
)

func Monitor() IMonitor {
	if localMonitor == nil {
		panic("implement not found for interface IMonitor, forgot register?")
	}
	return localMonitor
}

func RegisterMonitor(i IMonitor) {
	localMonitor = i
}
