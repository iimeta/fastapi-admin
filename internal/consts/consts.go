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
)

const (
	ADMIN_TOKEN_PREFIX = "A"
	ADMIN_SESSION      = "admin:session:%s"
	ADMIN_CHANNEL      = "admin"
)

const (
	SESSION_TOKEN   = "token"
	SESSION_UID     = "uid"
	SESSION_USER_ID = "user_id"
	SESSION_USER    = "user"
	SESSION_ADMIN   = "admin"
	SESSION_ROLE    = "role"
	SESSION_CREATOR = "creator"
)

const (
	CORP_OPENAI     = "OpenAI"
	CORP_BAIDU      = "Baidu"
	CORP_XFYUN      = "Xfyun"
	CORP_ALIYUN     = "Aliyun"
	CORP_ZHIPUAI    = "ZhipuAI"
	CORP_MIDJOURNEY = "Midjourney"
)

const (
	API_USAGE_KEY = "api:%d:usage"

	USER_TOTAL_TOKENS_FIELD = "user.total_tokens"

	APP_TOTAL_TOKENS_FIELD   = "app.%d.total_tokens"
	APP_IS_LIMIT_QUOTA_FIELD = "app.%d.is_limit_quota"

	KEY_TOTAL_TOKENS_FIELD   = "key.%d.%s.total_tokens"
	KEY_IS_LIMIT_QUOTA_FIELD = "key.%d.%s.is_limit_quota"
)
