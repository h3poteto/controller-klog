package ctrklog

import "context"

type contextKey string

const RequestIDKey contextKey = "requestID"
const ControllerKey contextKey = "controller"
const ObjectKey contextKey = "object"

func SetRequestID(parent context.Context, requestID string) context.Context {
	return context.WithValue(parent, RequestIDKey, requestID)
}

func SetController(parent context.Context, controller string) context.Context {
	return context.WithValue(parent, ControllerKey, controller)
}

func SetObject(parent context.Context, object string) context.Context {
	return context.WithValue(parent, ObjectKey, object)
}
