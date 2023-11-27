package handler

import (
	"github.com/ekart/user/service"
	"github.com/gofiber/fiber/v2"
)

func UserHandler(app *fiber.App) {

	api := app.Group("/user")
	v1 := api.Group("/v1")

	v1.Post("/createuser", service.CreateUser)
	v1.Get("/getuserbyid/:id", service.GetUserById)
	v1.Post("/login", service.LoginUser)

	v1.Post("/createorder", service.CreateOrder)
	v1.Get("/getorder/:userid", service.GetOrderByUserId)
}
