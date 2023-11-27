package main

import (
	"fmt"
	"log"

	"github.com/ekart/user/database"
	"github.com/ekart/user/handler"
	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	fmt.Println("Server started")
	dbconf := mysql.NewConfig()
	dbconf.Addr = "database-1.cbxusm6jod58.eu-north-1.rds.amazonaws.com"
	dbconf.DBName = "ekart"
	dbconf.Net = "tcp"
	dbconf.User = "admin"
	dbconf.Passwd = "root12345"
	dbconf.AllowNativePasswords = true
	dbconf.ParseTime = true
	dsn := dbconf.FormatDSN()

	database.DbConnection(dsn)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))

	handler.UserHandler(app)

	log.Fatal(app.Listen(":9000"))
}
