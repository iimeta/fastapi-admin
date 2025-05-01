package consts

const (
	LOCK_LOGIN           = "lock:login:account:%s"
	LOCK_CHANGE_PASSWORD = "lock:change_password:user:%d"
	LOCK_CODE            = "lock:code:%s"
	METHOD_ACCOUNT       = "account"
	METHOD_CODE          = "code"
)

const (
	RESELLER_TOKEN_PREFIX = "R"
	RESELLER_SESSION      = "reseller:session:%s"
	RESELLER_CHANNEL      = "reseller"

	USER_TOKEN_PREFIX = "U"
	USER_SESSION      = "user:session:%s"
	USER_CHANNEL      = "user"

	ADMIN_TOKEN_PREFIX = "A"
	ADMIN_SESSION      = "admin:session:%s"
	ADMIN_CHANNEL      = "admin"
)

const (
	SESSION_HOST     = "host"
	SESSION_TOKEN    = "token"
	SESSION_UID      = "uid"
	SESSION_RID      = "rid"
	SESSION_USER_ID  = "user_id"
	SESSION_RESELLER = "reseller"
	SESSION_USER     = "user"
	SESSION_ADMIN    = "admin"
	SESSION_ROLE     = "role"
	SESSION_CREATOR  = "creator"
)

const (
	TASK_STATISTICS_LOCK_KEY        = "task:statistics:lock"
	TASK_STATISTICS_END_TIME_KEY    = "task:statistics:end_time"
	TASK_CHECK_LOCK_KEY             = "task:check:lock:%s"
	TASK_CHECK_END_TIME_KEY         = "task:check:end_time:%s"
	TASK_LOG_LOCK_KEY               = "task:log:lock"
	TASK_LOG_END_TIME_KEY           = "task:log:end_time"
	TASK_QUOTA_WARNING_LOCK_KEY     = "task:quota:warning:lock"
	TASK_QUOTA_WARNING_END_TIME_KEY = "task:quota:warning:end_time"
)

const (
	STATISTICS_CHAT_LAST_ID_KEY    = "statistics:chat:last_id"
	STATISTICS_CHAT_LAST_TIME_KEY  = "statistics:chat:last_time"
	STATISTICS_IMAGE_LAST_ID_KEY   = "statistics:image:last_id"
	STATISTICS_IMAGE_LAST_TIME_KEY = "statistics:image:last_time"
	STATISTICS_AUDIO_LAST_ID_KEY   = "statistics:audio:last_id"
	STATISTICS_AUDIO_LAST_TIME_KEY = "statistics:audio:last_time"
)

// 模型类型[1:文生文, 2:文生图, 3:图生文, 4:图生图, 5:文生语音, 6:语音生文, 100:多模态, 101:多模态实时, 102:多模态语音]
var MODEL_TYPES = []int{
	1,
	2,
	3,
	4,
	5,
	6,
	100,
	101,
	102,
}

// 模型类型[1:文生文, 2:文生图, 3:图生文, 4:图生图, 5:文生语音, 6:语音生文, 100:多模态, 101:多模态实时, 102:多模态语音]
var MODEL_TYPE = map[int]string{
	1:   "文生文",
	2:   "文生图",
	3:   "图生文",
	4:   "图生图",
	5:   "文生语音",
	6:   "语音生文",
	100: "多模态",
	101: "多模态实时",
	102: "多模态语音",
}
