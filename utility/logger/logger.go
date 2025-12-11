package logger

import (
	"context"
	"encoding/json"

	"github.com/gogf/gf/v2/frame/g"
)

func Debug(ctx context.Context, v ...any) {
	var compactV []any
	for _, item := range v {
		switch val := item.(type) {
		case string:
			var jsonObj any
			if err := json.Unmarshal([]byte(val), &jsonObj); err == nil {
				if compactJSON, err := json.Marshal(jsonObj); err == nil {
					compactV = append(compactV, string(compactJSON))
				} else {
					compactV = append(compactV, val)
				}
			} else {
				compactV = append(compactV, val)
			}
		case []byte:
			var jsonObj any
			if err := json.Unmarshal(val, &jsonObj); err == nil {
				if compactJSON, err := json.Marshal(jsonObj); err == nil {
					compactV = append(compactV, string(compactJSON))
				} else {
					compactV = append(compactV, val)
				}
			} else {
				compactV = append(compactV, val)
			}
		default:
			compactV = append(compactV, item)
		}
	}
	g.Log().Debug(ctx, compactV...)
}

func Info(ctx context.Context, v ...any) {
	g.Log().Info(ctx, v...)
}

func Error(ctx context.Context, v ...any) {
	g.Log().Error(ctx, v...)
}

func Debugf(ctx context.Context, format string, v ...any) {
	var compactV []any
	for _, item := range v {
		switch val := item.(type) {
		case string:
			var jsonObj any
			if err := json.Unmarshal([]byte(val), &jsonObj); err == nil {
				if compactJSON, err := json.Marshal(jsonObj); err == nil {
					compactV = append(compactV, string(compactJSON))
				} else {
					compactV = append(compactV, val)
				}
			} else {
				compactV = append(compactV, val)
			}
		case []byte:
			var jsonObj any
			if err := json.Unmarshal(val, &jsonObj); err == nil {
				if compactJSON, err := json.Marshal(jsonObj); err == nil {
					compactV = append(compactV, string(compactJSON))
				} else {
					compactV = append(compactV, val)
				}
			} else {
				compactV = append(compactV, val)
			}
		default:
			compactV = append(compactV, item)
		}
	}
	g.Log().Debugf(ctx, format, compactV...)
}

func Infof(ctx context.Context, format string, v ...any) {
	g.Log().Infof(ctx, format, v...)
}

func Errorf(ctx context.Context, format string, v ...any) {
	g.Log().Errorf(ctx, format, v...)
}
