package tracing

import "context"

type logger interface {
	Println(...interface{})
	Printf(string, ...interface{})
}

type tracer interface {
	StartParent(ctx interface{}) interface{}
	StartChild(ctx context.Context, request ...interface{}) interface{}
	Close(span interface{})
	Context(span interface{}) context.Context
	LogError(span interface{}, err error, status ...int)
	LogObject(span interface{}, name string, obj interface{})
	GetTraceID(span interface{}) string
}

type response interface {
	ResponseFailed(ctx, code interface{}) error
	ResponseSuccess(ctx, response interface{}, code ...interface{}) error
}

type Trace struct {
	tracer
	log          logger
	res          response
	showLog      bool
	showTraceLog bool
}

func New() *Trace {
	return &Trace{
		log:          &DefaultLogger{},
		showLog:      false,
		showTraceLog: false,
	}
}

func (t *Trace) SetLogger(logInterface logger) {
	t.log = logInterface
}

func (t *Trace) SetTracer(traceInterface tracer) {
	t.tracer = traceInterface
}

func (t *Trace) SetResponse(responseInterface response) {
	t.res = responseInterface
}

func (t *Trace) ShowLog(isShow bool) {
	t.showLog = isShow
}

func (t *Trace) ShowTraceLog(isShow bool) {
	t.showTraceLog = isShow
}

func (t *Trace) PrintLog(msg string) {
	if t.showLog && t.showTraceLog && t.log != nil {
		t.log.Println(msg)
	}
}
