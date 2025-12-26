package task

import "context"

// 初始化任务
func Init(ctx context.Context) {

	// 统计任务
	statisticsTask(ctx)

	// 额度任务
	quotaTask(ctx)

	// 视频任务
	videoTask(ctx)

	// 文件任务
	fileTask(ctx)

	// 批处理任务
	batchTask(ctx)

	// 错误检查任务
	errorCheckTask(ctx)

	// 日志删除任务
	logDelTask(ctx)

	// 通知任务
	noticeTask(ctx)
}
