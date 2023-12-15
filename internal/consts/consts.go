package consts

const (
	LOCK_LOGIN           = "lock:login:account:%s"
	LOCK_CHANGE_PASSWORD = "lock:change_password:user:%d"
	LOCK_CODE            = "lock:code:%s"
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
	CORP_OPENAI     = "OpenAI"
	CORP_BAIDU      = "Baidu"
	CORP_XFYUN      = "Xfyun"
	CORP_ALIYUN     = "Aliyun"
	CORP_MIDJOURNEY = "Midjourney"
)
