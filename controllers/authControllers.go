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

func Login(c fiber.Ctx) error {
	var data models.LoginInput

	//on recupere les donnes via de l'user et on remplis le varibale data 
	if err := c.Bind().Body(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Rechercher l'utilisateur
	var user models.Customer
	if err := database.DB.Where("email = ?", data.Email).First(&user).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	//  Vérifier le mot de passe
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Password incorrect",
		})
	}

	//  Réponse (ne jamais renvoyer le hash)
	return c.JSON(user)
}
