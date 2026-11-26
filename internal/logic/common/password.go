package common

import "github.com/gogf/gf/v2/os/gtime"

// 计算密码到期时间, 返回0表示不限制有效期
// passwordUpdatedAt 为密码最后修改时间, 历史数据可能为空, 为空时按账号创建时间计算
// expireDays 为有效天数, 小于等于0表示不限制有效期
func PasswordExpireAt(passwordUpdatedAt, createdAt int64, expireDays int) int64 {

	if expireDays <= 0 {
		return 0
	}

	base := passwordUpdatedAt
	if base <= 0 {
		base = createdAt
	}

	if base <= 0 {
		return 0
	}

	return base + int64(expireDays)*24*3600*1000
}

// 判断密码是否已过期, expireAt 为0表示不限制有效期
func IsPasswordExpired(expireAt int64) bool {
	return expireAt > 0 && gtime.TimestampMilli() > expireAt
}
