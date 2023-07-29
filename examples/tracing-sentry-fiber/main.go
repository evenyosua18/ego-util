package main

import (
	"context"
	fiber_helper "github.com/evenyosua18/ego-util/routing/fiber-helper"
	sentry "github.com/evenyosua18/ego-util/tracing/sentry-helper"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"net/http"
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

	//set fiber
	api := fiber.New(fiber.Config{BodyLimit: 50 * 1024 * 1024})

	//set test api
	api.Get("/test", parentFunction).Name("Test")

	//listen route
	if err := api.Listen(":8080"); err != nil {
		panic(err)
	}
}

func parentFunction(f *fiber.Ctx) error {
	//start sentry
	sp := sentry.StartParent(f)
	defer sp.Finish()

	//call first child
	firstChildFunction(sp.Context())

	//call second child
	secondChildFunction(sp.Context())

	return f.Status(http.StatusOK).JSON(fiber.Map{"trace_id": sp.TraceID})
}

func firstChildFunction(ctx context.Context) {
	//start sentry
	sp := sentry.StartChild(ctx)
	defer sp.Finish()

	log.Println("call first child")
}

func secondChildFunction(ctx context.Context) {
	//start sentry
	sp := sentry.StartChild(ctx)
	defer sp.Finish()

	//call grand child function
	grandChildFunction(sp.Context())

	log.Println("call second child")
}

func grandChildFunction(ctx context.Context) {
	//start sentry
	sp := sentry.StartChild(ctx)
	defer sp.Finish()

	log.Println("call grand child")
}
