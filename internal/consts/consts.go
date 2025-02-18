package consts

const (
	LOCK_LOGIN           = "lock:login:account:%s"
	LOCK_CHANGE_PASSWORD = "lock:change_password:user:%d"
	LOCK_CODE            = "lock:code:%s"
	METHOD_ACCOUNT       = "account"
	METHOD_CODE          = "code"
)

const (
	USER_TOKEN_PREFIX = "U"
	USER_SESSION      = "user:session:%s"
	USER_CHANNEL      = "user"

	ADMIN_TOKEN_PREFIX = "A"
	ADMIN_SESSION      = "admin:session:%s"
	ADMIN_CHANNEL      = "admin"
)

const (
	SESSION_HOST    = "host"
	SESSION_TOKEN   = "token"
	SESSION_UID     = "uid"
	SESSION_USER_ID = "user_id"
	SESSION_USER    = "user"
	SESSION_ADMIN   = "admin"
	SESSION_ROLE    = "role"
	SESSION_CREATOR = "creator"
)

const (
	STATISTICS_LOCK_KEY            = "statistics:lock"
	STATISTICS_END_TIME_KEY        = "statistics:end_time"
	STATISTICS_CHAT_LAST_ID_KEY    = "statistics:chat:last_id"
	STATISTICS_CHAT_LAST_TIME_KEY  = "statistics:chat:last_time"
	STATISTICS_IMAGE_LAST_ID_KEY   = "statistics:image:last_id"
	STATISTICS_IMAGE_LAST_TIME_KEY = "statistics:image:last_time"
	STATISTICS_AUDIO_LAST_ID_KEY   = "statistics:audio:last_id"
	STATISTICS_AUDIO_LAST_TIME_KEY = "statistics:audio:last_time"
)

const (
	CHECK_LOCK_KEY     = "check:lock"
	CHECK_END_TIME_KEY = "check:end_time"
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
