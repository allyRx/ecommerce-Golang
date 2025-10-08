package routes

import (
	"github.com/allyRx/ecommerce-Golang/controllers"
	"github.com/gofiber/fiber/v3"
)


func Route(app  *fiber.App){
	 app.Post("/api/register", controllers.Register)
	 app.Post("/api/login" ,controllers.Login)
}