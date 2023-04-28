package controllers

import (
	"context"
	"github.com/Jateq/ecommerce/database"
	"github.com/Jateq/ecommerce/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"time"
)

func HashPassword(password string) string {

}

func VerifyPassword(userPassword string, givenPassword string) (bool, string) {

}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			//c.JSON{http.StatusBadRequest, gin.H{"error": err.Error()}}
			//return
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var Validate = validator.New()
		validationErr := Validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			return
		}
		count, err := database.UserData(database.Client, "UserCollection").CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
		}

		//count, err = UserCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		count, err = database.UserData(database.Client, "UserCollection").CountDocuments(ctx, bson.M{"email": user.Phone})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "this phone number is already used"})
			return
		}
	}
}
func Login() gin.HandlerFunc {

}

func AddProductAdmin() gin.HandlerFunc {

}

func AllProducts() gin.HandlerFunc {

}

func SearchProduct() gin.HandlerFunc {

}
