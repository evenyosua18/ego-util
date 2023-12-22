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

func ShowTraceLog() {
	tracing.ShowTraceLog(true)
}

func StartParent(ctx interface{}) interface{} {
	tracing.PrintLog("hit exported function of StartParent")

	if tracing.tracer != nil {
		tracing.PrintLog("call StartParent")
		return tracing.StartParent(ctx)
	}

	tracing.PrintLog("tracer is nil")

	return nil
}

func StartChild(ctx context.Context, request ...interface{}) interface{} {
	tracing.PrintLog("hit exported function of StartParent")

	if tracing.tracer != nil {
		tracing.PrintLog("call StartChild")
		return tracing.StartChild(ctx, request)
	}

	tracing.PrintLog("tracer is nil")

	return nil
}

func Close(span interface{}) {
	tracing.PrintLog("hit exported function of Close")

	if tracing.tracer != nil {
		tracing.PrintLog("call Close")
		tracing.Close(span)
	}

	tracing.PrintLog("tracer is nil")
}

func Context(span interface{}) context.Context {
	tracing.PrintLog("hit exported function of Context")

	if tracing.tracer != nil {
		tracing.PrintLog("call Context")
		return tracing.Context(span)
	}

	tracing.PrintLog("tracer is nil")
	return nil
}

func LogObject(span interface{}, name string, obj interface{}) {
	tracing.PrintLog("hit exported function of LogObject")

	if tracing.tracer != nil {
		tracing.PrintLog("call LogObject")
		tracing.LogObject(span, name, obj)
	}

	tracing.PrintLog("tracer is nil")

	//logging
	if tracing.showLog {
		tracing.log.Printf(name, obj)
	}
}

func LogError(span interface{}, err error) error {
	tracing.PrintLog("hit exported function of LogError")

	//tracing for error
	if tracing.tracer != nil {
		tracing.PrintLog("call LogError")
		tracing.LogError(span, err)
	}

	tracing.PrintLog("tracer is nil")

	//logging
	if tracing.showLog {
		tracing.log.Println(err)
	}

	return err
}

func LogResponse(span, response interface{}) {
	tracing.PrintLog("hit exported function of LogResponse")

	if tracing.tracer != nil {
		tracing.PrintLog("call LogResponse")
		tracing.LogObject(span, "response", response)
	}

	tracing.PrintLog("tracer is nil")

	if tracing.showLog {
		tracing.log.Println(response)
	}
}

func LogRequest(span, request interface{}) {
	tracing.PrintLog("hit exported function of LogRequest")

	if tracing.tracer != nil {
		tracing.PrintLog("call LogRequest")
		tracing.LogObject(span, "request", request)
	}

	tracing.PrintLog("tracer is nil")

	if tracing.showLog {
		tracing.log.Println(request)
	}
}

func ResponseError(span, ctx, code interface{}, err error) error {
	tracing.PrintLog("hit exported function of ResponseError")

	//tracing & logging
	LogError(span, err)

	if tracing.res == nil {
		return fmt.Errorf("response model is empty")
	}

	//return tracing..ResponseFailedHTTP(ctx, err)
	return tracing.res.ResponseFailed(ctx, code)
}

func ResponseSuccess(span, ctx, response interface{}, code ...interface{}) error {
	tracing.PrintLog("hit exported function of ResponseSuccess")

	//tracing & logging
	LogResponse(span, response)

	if tracing.res == nil {
		return fmt.Errorf("response model is empty")
	}

	return tracing.res.ResponseSuccess(ctx, response, code...)
}

func GetTraceID(span interface{}) string {
	tracing.PrintLog("hit exported function of GetTraceID")

	if tracing.tracer != nil {
		tracing.PrintLog("call GetTraceID")
		return tracing.GetTraceID(span)
	}

	tracing.PrintLog("tracer is nil")

	return ""
}

func AddContextValue(span interface{}, key string, value interface{}) context.Context {
	tracing.PrintLog("hit exported function of AddContextValue")

	if tracing.tracer != nil {
		tracing.PrintLog("call AddContextValue")
		return context.WithValue(tracing.Context(span), key, value)
	}

	tracing.PrintLog("tracer is nil")

	return nil
}
