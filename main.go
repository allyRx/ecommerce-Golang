package main

import (
	"log"

	"github.com/allyRx/ecommerce-Golang/database"
	"github.com/allyRx/ecommerce-Golang/routes"

	"github.com/gofiber/fiber/v3"
)

func main() {

	//import database
	database.ConnectDb()



	// Initialize a new Fiber app
	app := fiber.New()

		//import routes
	routes.Route(app)

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
