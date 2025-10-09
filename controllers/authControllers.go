package controllers

import (
	"strconv"
	"time"
	"github.com/allyRx/ecommerce-Golang/database"
	"github.com/allyRx/ecommerce-Golang/models"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v4"
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

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer: strconv.Itoa(int(user.Id)),
	
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

	token , err := claims.SignedString([]byte("secret"))


	if err != nil {
		return  c.SendStatus(fiber.StatusInternalServerError)
	}


	cookies := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookies)

	//  Réponse (ne jamais renvoyer le hash)
	return c.JSON(fiber.Map{
		"message" : "succes",
	})
}

	type Claims struct {
		jwt.StandardClaims
	}

func User(c fiber.Ctx) error{
	
	cookie := c.Cookies("jwt") 

	token , err := jwt.ParseWithClaims(cookie , &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret") , nil
	} )

	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return  c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	claims := token.Claims.(*Claims)

	id_user := claims.Issuer
	
	var user models.Customer

	if err := database.DB.Where("id = ?" , id_user).First(&user);err != nil{
		c.Status(404).JSON(fiber.Map{
			"messge" : "User not found",
		})
	}

	return  c.JSON(user)
}

func Logout(c fiber.Ctx) error{
	cookies := fiber.Cookie{
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookies)

return c.JSON(fiber.Map{
	"message" : "Logout",
})
}