package fiber_helper

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
)

type FiberImpl struct{}

func (h *FiberImpl) GetContextName(ctx interface{}) (context.Context, string) {
	//check nil
	if ctx == nil {
		return nil, ""
	}

	//to fiber
	f, ok := ctx.(*fiber.Ctx)

	if !ok {
		return nil, ""
	}

	//check transaction name, use it if exist
	transactionName, ok := f.Locals("transaction_name").(string)

	if !ok {
		transactionName = f.Route().Name
	}

	return f.UserContext(), transactionName
}

func (h *FiberImpl) GetInfo(ctx interface{}) (info map[string]interface{}) {
	//check nil
	if ctx == nil {
		return
	}

	//to fiber
	f, ok := ctx.(*fiber.Ctx)

	if !ok {
		return
	}

	//set info
	info = make(map[string]interface{})
	info["header"] = f.GetReqHeaders()
	info["endpoint"] = f.OriginalURL()
	info["context"] = f.Context().String()

	return
}

type FiberResponseImpl struct{}

func (h *FiberImpl) ResponseSuccess(ctx, response interface{}, statusCode int) error {
	f := ctx.(*fiber.Ctx)

	//set http status code
	msg := Response{}
	if err := mapstructure.Decode(response, &msg); err != nil {
		return err
	}

	f.Status(msg.CustomCode)

	//set response
	return f.JSON(HttpResponse{
		Code:         msg.CustomCode,
		Message:      msg.ResponseMessage,
		ErrorMessage: "",
		Data:         response,
	})
}
