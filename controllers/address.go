package controllers

import (
	"context"
	"github.com/Jateq/ecommerce/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func AddAddress() gin.HandlerFunc {

}

func EditAddress() gin.HandlerFunc {

}

func EditWorkAddress() gin.HandlerFunc {

}
func DeleteAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Query("id")
		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"err": "invalid search index"})
			c.Abort()
			return
		}
		addresses := make([]models.Address, 0)
		usert_id, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			c.IndentedJSON(500, "Internal Server error")
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		filter := bson.D{primitive.E{Key: "_id", Value: usert_id}}
		update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "address", Value: addresses}}}}
		// setting it empty rather than deleting from db it's the most efficient way
		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(404, "Invalid command")
			return
		}
		defer cancel()
		ctx.Done()
		c.IndentedJSON(200, "Successfully deleted")
	}
}

// Interview question about golang: Why golang is so fast and light?
// Golang can't even print it on its own, you always need to use packages, completely modular
// GO expect you not to import packages if you are not using it

// Another one: Why you need a timeout?
// Whenever server works with database, you can't have it to endlessly waiting the result
