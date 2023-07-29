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
}

type response interface {
	ResponseFailed(ctx interface{}, err error) error
	ResponseSuccess(ctx, response, message interface{}) error
}

type Trace struct {
	tracer
	log     logger
	res     response
	showLog bool
}

func New() *Trace {
	return &Trace{
		log:     &DefaultLogger{},
		showLog: false,
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
