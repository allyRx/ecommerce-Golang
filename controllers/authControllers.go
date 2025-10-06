package controllers

import (
	"github.com/allyRx/ecommerce-Golang/database"
	"github.com/allyRx/ecommerce-Golang/models"
	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
)



func Register(c fiber.Ctx) error {
  
	data := new(models.Customer)

	// Bind les données JSON du corps dans la struct
	if err := c.Bind().Body(data); err != nil{
		return  c.Status(500).JSON(fiber.Map{"message" : "Bad Request" + err.Error(),
	})
	}
	

	// Hasher le mot de passe seulement après avoir reçu la donnée
	 passwordCrypter , err := bcrypt.GenerateFromPassword([]byte(data.Password) , 12)

	 if err != nil {
		return  c.Status(500).JSON(fiber.Map{"Message" : "Erreur cryptage" + err.Error()})
	}


	// Stocker le mot de passe hashé en tant que []byte
	data.Password = string(passwordCrypter)
	
	
 	if err := database.DB.Create(data).Error;  err != nil{
		return c.Status(500).JSON(fiber.Map{"Message" : "Il y a une erreur critique" + err.Error(),})
	}

	return c.Status(201).JSON(fiber.Map{
		"message" : "User create succesfully",
		"customers" : data,
	})
}
