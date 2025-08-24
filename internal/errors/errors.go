package errors

import (
	"errors"
	"fmt"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	ERR_SERVER_ERROR    = gerror.NewCode(gcode.New(-1, "Server errors", nil))
	ERR_UNAUTHORIZED    = gerror.NewCode(gcode.New(999, "Unauthorized", nil))
	ERR_INVALID_ACCOUNT = gerror.NewCode(gcode.New(10001, "Invalid Phone", nil))
	ERR_INVALID_API_KEY = gerror.NewCode(gcode.New(10002, "无效的应用密钥", nil))
)

func New(text string) error {
	return errors.New(text)
}

func Newf(format string, args ...interface{}) error {
	return errors.New(fmt.Sprintf(format, args...))
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}
