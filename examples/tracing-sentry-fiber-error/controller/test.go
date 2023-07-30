package controller

import (
	"github.com/evenyosua18/ego-util/codes"
	"github.com/evenyosua18/ego-util/examples/tracing-sentry-fiber-error/service"
	"github.com/evenyosua18/ego-util/tracing"
	"github.com/gofiber/fiber/v2"
)

func TestController(f *fiber.Ctx) error {
	//start tracer
	sp := tracing.StartParent(f)
	defer tracing.Close(sp)

	//get param
	id := f.Params("id")

	if id == "" {
		code := codes.Create(401)
		return tracing.ResponseError(sp, f, code, code.Error())
	}

	//call service
	if err := service.TestService(tracing.Context(sp), id); err != nil {
		return tracing.ResponseError(sp, f, codes.Extract(err), err)
	}

	return tracing.ResponseSuccess(sp, f, nil)
}
