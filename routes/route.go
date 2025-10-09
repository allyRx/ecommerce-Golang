package routes

import (
	middleware "github.com/allyRx/ecommerce-Golang/Middleware"
	"github.com/allyRx/ecommerce-Golang/controllers"
	"github.com/gofiber/fiber/v3"
)


func Route(app  *fiber.App){

	//public routes 
	 app.Post("/api/login" ,controllers.Login)
	 app.Post("/api/register", controllers.Register)
	 
	 app.Use(middleware.IsAuthenticated)
	 
	 
	 app.Get("/api/user" , controllers.User)
	 app.Post("/api/logout" , controllers.Logout)
	 
}