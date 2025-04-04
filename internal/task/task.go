package task

import "context"

// 初始化任务
func Init(ctx context.Context) {

	// 统计任务
	statisticsTask(ctx)

	// 错误检查任务
	errorCheckTask(ctx)

	// 日志删除任务
	logDelTask(ctx)

	// 通知任务
	noticeTask(ctx)
}
