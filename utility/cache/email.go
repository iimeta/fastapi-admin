package cache

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"time"

	"github.com/redis/go-redis/v9"
)

type EmailStorage struct {
	redis *redis.Client
}

func NewEmailStorage(redis *redis.Client) *EmailStorage {
	return &EmailStorage{redis}
}

func (s *EmailStorage) Set(ctx context.Context, channel string, email string, code string, exp time.Duration) error {
	_, err := s.redis.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.Del(ctx, s.failName(channel, email))
		pipe.Set(ctx, s.name(channel, email), code, exp)
		return nil
	})
	return err
}

func (s *EmailStorage) Get(ctx context.Context, channel string, email string) (string, error) {
	return s.redis.Get(ctx, s.name(channel, email)).Result()
}

func (s *EmailStorage) Del(ctx context.Context, channel string, email string) error {
	return s.redis.Del(ctx, s.name(channel, email)).Err()
}

func (s *EmailStorage) Verify(ctx context.Context, channel string, email string, code string) (pass bool) {

	defer func() {
		if !pass {
			// 3分钟内同一个邮件验证码错误次数超过5次, 删除验证码
			num := s.redis.Incr(ctx, s.failName(channel, email)).Val()
			if num >= 5 {
				_, _ = s.redis.Pipelined(ctx, func(pipe redis.Pipeliner) error {
					pipe.Del(ctx, s.name(channel, email))
					pipe.Del(ctx, s.failName(channel, email))
					return nil
				})
			} else if num == 1 {
				s.redis.Expire(ctx, s.failName(channel, email), 3*time.Minute)
			}
		}
	}()

	value, err := s.Get(ctx, channel, email)
	if err != nil || len(value) == 0 {
		return false
	}

	if value == code {
		return true
	}

	return false
}

func (s *EmailStorage) name(channel string, email string) string {
	return fmt.Sprintf("im:auth:email:%s:%s", channel, gmd5.MustEncryptString(email))
}

func (s *EmailStorage) failName(channel string, email string) string {
	return fmt.Sprintf("im:auth:email_fail:%s:%s", channel, gmd5.MustEncryptString(email))
}
