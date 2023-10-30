package tracing

import (
	"context"
	"fmt"
)

var (
	tracing = New()
)

func SetLogger(log logger) {
	tracing.SetLogger(log)
}

func SetTracer(tracerInterface tracer) {
	tracing.SetTracer(tracerInterface)
}

func SetResponse(responseInterface response) {
	tracing.SetResponse(responseInterface)
}

func ShowLog() {
	tracing.ShowLog(true)
}

func StartParent(ctx interface{}) interface{} {
	if tracing.tracer != nil {
		return tracing.StartParent(ctx)
	}

	return nil
}

func StartChild(ctx context.Context, request ...interface{}) interface{} {
	if tracing.tracer != nil {
		return tracing.StartChild(ctx, request)
	}

	return nil
}

func Close(span interface{}) {
	if tracing.tracer != nil {
		tracing.Close(span)
	}
}

func Context(span interface{}) context.Context {
	if tracing.tracer != nil {
		return tracing.Context(span)
	}

	return nil
}

func LogObject(span interface{}, name string, obj interface{}) {
	if tracing.tracer != nil {
		tracing.LogObject(span, name, obj)
	}

	//logging
	if tracing.showLog {
		tracing.log.Printf(name, obj)
	}
}

func LogError(span interface{}, err error) error {
	//tracing for error
	if tracing.tracer != nil {
		tracing.LogError(span, err)
	}

	//logging
	if tracing.showLog {
		tracing.log.Println(err)
	}

	return err
}

func LogResponse(span, response interface{}) {
	if tracing.tracer != nil {
		tracing.LogObject(span, "response", response)
	}

	if tracing.showLog {
		tracing.log.Println(response)
	}
}

func LogRequest(span, request interface{}) {
	if tracing.tracer != nil {
		tracing.LogObject(span, "request", request)
	}

	if tracing.showLog {
		tracing.log.Println(request)
	}
}

func ResponseError(span, ctx, code interface{}, err error) error {
	//tracing & logging
	LogError(span, err)

	if tracing.res == nil {
		return fmt.Errorf("response model is empty")
	}

	//return tracing..ResponseFailedHTTP(ctx, err)
	return tracing.res.ResponseFailed(ctx, code)
}

func ResponseSuccess(span, ctx, response interface{}, code ...interface{}) error {
	//tracing & logging
	LogResponse(span, response)

	if tracing.res == nil {
		return fmt.Errorf("response model is empty")
	}

	return tracing.res.ResponseSuccess(ctx, response, code...)
}

func GetTraceID(span interface{}) string {
	if tracing.tracer != nil {
		return tracing.GetTraceID(span)
	}

	return ""
}

func AddContextValue(span interface{}, key string, value interface{}) context.Context {
	if tracing.tracer != nil {
		return context.WithValue(tracing.Context(span), key, value)
	}

	return nil
}
