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
	SESSION_TOKEN   = "token"
	SESSION_UID     = "uid"
	SESSION_USER_ID = "user_id"
	SESSION_USER    = "user"
	SESSION_ADMIN   = "admin"
	SESSION_ROLE    = "role"
	SESSION_CREATOR = "creator"
)
