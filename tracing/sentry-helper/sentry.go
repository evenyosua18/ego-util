package sentry_helper

var (
	helper = New()
)

func New() *Helper {
	return &Helper{}
}

type Helper struct {
	dsn       string
	env       string
	flushTime string
}

func (h *Helper) SetDSN(dsn string) *Helper {
	h.dsn = dsn
	return h
}
