package consts

const (
	API_USAGE_KEY = "api:user:%d:usage"

	USER_QUOTA_FIELD = "user.quota"

	APP_QUOTA_FIELD          = "app.%d.quota"
	APP_IS_LIMIT_QUOTA_FIELD = "app.%d.is_limit_quota"

	KEY_QUOTA_FIELD          = "key.%d.%s.quota"
	KEY_IS_LIMIT_QUOTA_FIELD = "key.%d.%s.is_limit_quota"

	QUOTA_USD_UNIT = 500000.0 // $1 = 50万tokens
)
