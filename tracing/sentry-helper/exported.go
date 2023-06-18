package sentry_helper

import (
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
