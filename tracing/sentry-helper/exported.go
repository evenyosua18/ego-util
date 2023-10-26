package sentry_helper

import (
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"
	"time"
)

func InitializeSentry(dsn, env string) (flush func(flushTime string), err error) {
	//check dsn
	if dsn == "" {
		err = fmt.Errorf("%s", "set dsn first")
		return
	}

	//set dsn to struct
	helper.dsn = dsn

	//set default environment
	if env == "" {
		helper.env = "test"
	}

	//setup sentry config
	if err = sentry.Init(sentry.ClientOptions{
		Dsn:              helper.dsn,
		Environment:      helper.env,
		Transport:        sentry.NewHTTPTransport(),
		TracesSampleRate: 1.0, //percentage chance of being sent to Sentry 1 = 100%
		EnableTracing:    true,
	}); err != nil {
		return
	}

	return flushSentry, nil
}

//default flush time is 1 second
func flushSentry(flushTime string) {
	timeout, err := time.ParseDuration(flushTime + "s")

	if err != nil {
		timeout = 1 * time.Second
	}

	sentry.Flush(timeout)
}

func StartParent(ctx interface{}) *sentry.Span {
	//start transaction
	sp := sentry.StartTransaction(helper.ctx.GetContextName(ctx))

	//get caller details
	caller, function := getCaller(helper.skippedCaller)

	//set operation
	sp.Description = getFunction(function)
	sp.Op = getFunction(function)

	//add information
	sp.Data = map[string]interface{}{}
	sp.Data["caller"] = caller

	information := helper.ctx.GetInfo(ctx)

	for key, value := range information {
		sp.Data[key] = value
	}

	return sp
}

func StartChild(ctx context.Context, request ...interface{}) *sentry.Span {
	//get caller details
	caller, function := getCaller(helper.skippedCaller)

	//sp := span.StartChild(function)
	sp := sentry.StartSpan(ctx, function)
	sp.Description = getFunction(function)

	sp.Data = map[string]interface{}{}
	sp.Data["caller"] = caller

	if len(request) == 1 {
		sp.Data["request"] = request[0]
	} else if len(request) > 1 {
		for idx, req := range request {
			sp.Data[fmt.Sprintf("%s-%d", "request", idx+1)] = req
		}
	}

	return sp
}

func LogError(sp *sentry.Span, err error, status int) {
	sp.Status = sentry.SpanStatus(status)

	if err != nil {
		sp.Data["error"] = err.Error()
	}
}

func LogResponse(sp *sentry.Span, response interface{}) {
	sp.Status = sentry.SpanStatusOK
	sp.Data["response"] = response
}

func LogObject(sp *sentry.Span, name string, obj interface{}) {
	sp.Data[name] = obj
}

func SetRouter(routeContext RouteContext) {
	helper.SetRouter(routeContext)
}

func Get() *Helper {
	return helper
}

func SetSkippedCaller(childSkipped, parentSkipped int) {
	helper.skippedCaller = childSkipped
	helper.parentSkippedCaller = parentSkipped
}

func AlertError(err error, modules map[string]string) {
	sentry.CaptureEvent(&sentry.Event{
		Environment: helper.env,
		Level:       sentry.LevelError,
		Message:     err.Error(),
		ServerName:  helper.name,
		Modules:     modules,
	})
}

func AlertPanic(err error, modules map[string]string) {
	sentry.CaptureEvent(&sentry.Event{
		Environment: helper.env,
		Level:       sentry.LevelFatal,
		Message:     err.Error(),
		ServerName:  helper.name,
		Modules:     modules,
	})
}
