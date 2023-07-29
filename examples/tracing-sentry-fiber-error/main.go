package main

import (
	fiber_helper "github.com/evenyosua18/ego-util/routing/fiber-helper"
	sentry "github.com/evenyosua18/ego-util/tracing/sentry-helper"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	//get env
	godotenv.Load(".env")

	//initialize sentry
	flushFunction, err := sentry.InitializeSentry(os.Getenv("SENTRY_DSN"), os.Getenv("APP_ENV"))
	if err != nil {
		panic(err)
	}
	defer flushFunction(os.Getenv("SENTRY_FLUSH_TIME"))

	//set router
	sentry.SetRouter(&fiber_helper.FiberImpl{})

	//set error
	//errors.RegisterError("./test.yaml")
}
