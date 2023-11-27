package service

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ekart/user/database"
	"github.com/ekart/user/model"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	if user.UserId >= 1 {
		updateUser := database.Db().Save(&user)
		if updateUser.RowsAffected == 1 {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusOK, "message": "User Update Successfully"})
		}
	} else {
		currentTime := time.Now().Format("02-01-2006 3:4:5 PM")
		user.CreateDate = currentTime
		saveUser := database.Db().Create(&user)
		if saveUser.RowsAffected == 1 {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusOK, "message": "User Save Successfully"})
		}

	}
	return nil
}
func GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println(id)
	user := new(model.User)

	database.Db().Find(&user, id)
	fmt.Println(&user.Email)
	if user.UserId >= 1 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"user": user})
	} else {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"msg": "User not found..."})
	}
}
func LoginUser(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	database.Db().Where("email = ? and password=?", user.Email, user.Password).Find(&user)

	if user.UserId >= 1 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusOK, "user": user})
	} else {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"message": "UserName and Password is incorreect"})
	}
}

func CreateOrder(c *fiber.Ctx) error {
	var orders []model.Orders

	if err := c.BodyParser(&orders); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	id := GenerateOrderId()

	for _, order := range orders {
		order.OderId = id
		order.CreateDate = time.Now().Format("02-01-2006 3:4:5 PM")
		result := database.Db().Create(&order)

		if result.RowsAffected > 1 {
			return c.Status(500).JSON(fiber.Map{"status": "500", "message": "Something went wrong to create order..."})
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusOK, "OrderId": id})
}
func GenerateOrderId() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(900000) + 100000
}
func GetOrderByUserId(c *fiber.Ctx) error {
	var order []model.Orders

	userid := c.Params("userid")
	database.Db().Where("userid = ?", userid).Find(&order)
	if order != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusOK, "order": order})
	} else {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"message": "Order is not found"})
	}

}
