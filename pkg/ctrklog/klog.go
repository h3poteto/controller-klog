package ctrklog

import (
	"context"
	"fmt"
	"strings"

	"k8s.io/klog/v2"
)

func Info(ctx context.Context, args ...interface{}) {
	prefix := []interface{}{}
	if requestID, ok := parseRequestID(ctx); ok {
		prefix = append(prefix, requestID)
	}
	if controller, ok := parseController(ctx); ok {
		prefix = append(prefix, controller)
	}
	if object, ok := parseObject(ctx); ok {
		prefix = append(prefix, object)
	}
	args = append(prefix, args...)
	klog.InfoDepth(1, args...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	prefix := []string{}
	if requestID, ok := parseRequestID(ctx); ok {
		prefix = append(prefix, requestID)
	}
	if controller, ok := parseController(ctx); ok {
		prefix = append(prefix, controller)
	}
	if object, ok := parseObject(ctx); ok {
		prefix = append(prefix, object)
	}
	klog.InfoDepth(1, fmt.Sprintf(strings.Join(prefix, " ")+format, args...))
}

func Warning(ctx context.Context, args ...interface{}) {
	prefix := []interface{}{}
	if requestID, ok := parseRequestID(ctx); ok {
		prefix = append(prefix, requestID)
	}
	if controller, ok := parseController(ctx); ok {
		prefix = append(prefix, controller)
	}
	if object, ok := parseObject(ctx); ok {
		prefix = append(prefix, object)
	}
	args = append(prefix, args...)
	klog.WarningDepth(1, args...)
}

func Warningf(ctx context.Context, format string, args ...interface{}) {
	prefix := []string{}
	if requestID, ok := parseRequestID(ctx); ok {
		prefix = append(prefix, requestID)
	}
	if controller, ok := parseController(ctx); ok {
		prefix = append(prefix, controller)
	}
	if object, ok := parseObject(ctx); ok {
		prefix = append(prefix, object)
	}
	klog.WarningDepth(1, fmt.Sprintf(strings.Join(prefix, " ")+format, args...))
}

func Error(ctx context.Context, args ...interface{}) {
	prefix := []interface{}{}
	if requestID, ok := parseRequestID(ctx); ok {
		prefix = append(prefix, requestID)
	}
	if controller, ok := parseController(ctx); ok {
		prefix = append(prefix, controller)
	}
	if object, ok := parseObject(ctx); ok {
		prefix = append(prefix, object)
	}
	args = append(prefix, args...)
	klog.ErrorDepth(1, args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	prefix := []string{}
	if requestID, ok := parseRequestID(ctx); ok {
		prefix = append(prefix, requestID)
	}
	if controller, ok := parseController(ctx); ok {
		prefix = append(prefix, controller)
	}
	if object, ok := parseObject(ctx); ok {
		prefix = append(prefix, object)
	}
	klog.ErrorDepth(1, fmt.Sprintf(strings.Join(prefix, " ")+format, args...))
}

func parseController(ctx context.Context) (string, bool) {
	return parseValue(ctx, ControllerKey)
}

func parseRequestID(ctx context.Context) (string, bool) {
	return parseValue(ctx, RequestIDKey)
}

func parseObject(ctx context.Context) (string, bool) {
	return parseValue(ctx, ObjectKey)
}

func parseValue(ctx context.Context, key contextKey) (string, bool) {
	cv := ctx.Value(key)
	value, ok := cv.(string)
	if !ok {
		return "", false
	}
	return fmt.Sprintf("{\"%s\": %s}", key, value), true
}
