package controllers

import (
	"github.com/allyRx/ecommerce-Golang/database"
	"github.com/allyRx/ecommerce-Golang/models"
	"github.com/gofiber/fiber/v3"
)



func Register(c fiber.Ctx) error {
  
	data := new(models.Customer)

	if err := c.Bind().Body(data); err != nil{
		return  c.Status(500).JSON(fiber.Map{"message" : "Bad Request" + err.Error(),
	})
	}


	
 	if err := database.DB.Create(data).Error;  err != nil{
		return c.Status(500).JSON(fiber.Map{"Message" : "Il y a une erreur critique" + err.Error(),})
	}

	return c.Status(201).JSON(fiber.Map{
		"message" : "User create succesfully",
		"customers" : data,
	})
}
