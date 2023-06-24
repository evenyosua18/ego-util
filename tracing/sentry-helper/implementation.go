package sentry_helper

import "github.com/getsentry/sentry-go"

func (h *Helper) LogResponse(span, response any) {
	sp := span.(*sentry.Span)
	LogResponse(sp, response)
}

func (h *Helper) LogError(span any, err error, status int) {
	sp := span.(*sentry.Span)
	LogError(sp, err, status)
}

func (h *Helper) LogObject(span any, name string, obj any) {
	sp := span.(*sentry.Span)
	LogObject(sp, name, obj)
}
