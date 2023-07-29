package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func TestController(f *fiber.Ctx) error {
	fmt.Println("controller has been called")

	//call service

	return nil
}
