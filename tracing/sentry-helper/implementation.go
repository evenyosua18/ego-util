package sentry_helper

import (
	"context"
	"github.com/getsentry/sentry-go"
)

func (h *Helper) LogError(span any, err error, status int) {
	sp := span.(*sentry.Span)
	LogError(sp, err, status)
}

func (h *Helper) LogObject(span any, name string, obj any) {
	sp := span.(*sentry.Span)
	LogObject(sp, name, obj)
}

func (h *Helper) StartParent(ctx interface{}) interface{} {
	return StartParent(ctx)
}

func (h *Helper) StartChild(ctx context.Context, request ...interface{}) interface{} {
	return StartChild(ctx, request)
}

func (h *Helper) Close(span interface{}) {
	sp := span.(*sentry.Span)
	sp.Finish()
}

func (h *Helper) Context(span interface{}) context.Context {
	sp := span.(*sentry.Span)
	return sp.Context()
}

//	Close(span interface{})
//	Context() context.Context
