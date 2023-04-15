package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/nwochaadim/go-crm-fiber/database"
	"github.com/nwochaadim/go-crm-fiber/routes"
)

var (
	app *fiber.App
)

func main() {
	defer database.Close()
	app = fiber.New()
	routes.SetupRoutes(app)
	fmt.Println("Listening on port 8080...")
	app.Listen(":8080")
}
