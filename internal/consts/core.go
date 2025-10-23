package consts

const (
	API_RESELLER_USAGE_KEY = "api:reseller:%d:usage"
	API_USER_USAGE_KEY     = "api:user:%d:usage"
	API_GROUP_USAGE_KEY    = "api:group:usage"

	RESELLER_QUOTA_FIELD = "reseller.quota"
	USER_QUOTA_FIELD     = "user.quota"

	APP_QUOTA_FIELD          = "app.%d.quota"
	APP_IS_LIMIT_QUOTA_FIELD = "app.%d.is_limit_quota"

	KEY_QUOTA_FIELD          = "key.%d.%s.quota"
	KEY_IS_LIMIT_QUOTA_FIELD = "key.%d.%s.is_limit_quota"

	QUOTA_DEFAULT_UNIT = 1000000.0 // $1 = 1M Tokens
)
