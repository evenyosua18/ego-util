package main

import (
	"github.com/evenyosua18/ego-util/codes"
	"github.com/evenyosua18/ego-util/examples/tracing-sentry-fiber-error/controller"
	fiber_helper "github.com/evenyosua18/ego-util/routing/fiber-helper"
	"github.com/evenyosua18/ego-util/tracing"
	sentry "github.com/evenyosua18/ego-util/tracing/sentry-helper"
	"github.com/gofiber/fiber/v2"
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

	//setup tracer
	sentry.SetRouter(&fiber_helper.FiberImpl{})
	sentry.SetSkippedCaller(5)
	tracing.SetResponse(&fiber_helper.FiberResponseImpl{})
	tracing.SetTracer(sentry.Get())

	//set error
	codes.RegisterCode("./codes.yaml")

	//set fiber
	api := fiber.New(fiber.Config{BodyLimit: 50 * 1024 * 1024})

	//set api
	api.Get("/test/:id", controller.TestController).Name("Test Controller")

	//listen route
	if err := api.Listen(":8080"); err != nil {
		panic(err)
	}
}
