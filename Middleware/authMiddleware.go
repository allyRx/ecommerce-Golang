package middleware

import (
	"github.com/allyRx/ecommerce-Golang/utils"
	"github.com/gofiber/fiber/v3"
)


func IsAuthenticated(c fiber.Ctx) error{
	
	cookie := c.Cookies("jwt")

	if _ , err := utils.Parsejwt(cookie); err != nil{
		c.Status(fiber.StatusUnauthorized)
		return  c.JSON(fiber.Map{
			"message" : "unauthorized",
		})
	}
	
	return c.Next()
}