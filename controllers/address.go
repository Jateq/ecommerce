package controllers

import (
	"context"
	"fmt"
	"github.com/Jateq/ecommerce/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

func AddAddress() gin.HandlerFunc {
	return func(c *gin.Context) {

		userID := c.Query("id")
		if userID == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"err": "invalid index"})
			c.Abort()
			return
		}
		address, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			c.IndentedJSON(500, "Internal server error")
		}

		var addresses models.Address

		addresses.AddressID = primitive.NewObjectID()

		if err = c.BindJSON(&addresses); err != nil {
			c.IndentedJSON(http.StatusNotAcceptable, err.Error())
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		// Aggregation query
		matchFilter := bson.D{{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: address}}}}
		unwind := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$address"}}}}
		group := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: "$address_id"}, {Key: "count", Value: bson.D{primitive.E{Key: "$sum", Value: 1}}}}}}
		pointCursor, err := UserCollection.Aggregate(ctx, mongo.Pipeline{matchFilter, unwind, group})
		if err != nil {
			c.IndentedJSON(500, "Internal server error")
		}

		var addressInfo []bson.M
		if err := pointCursor.All(ctx, &addressInfo); err != nil {
			panic(err)
		}

		var size int32
		for _, addressNum := range addressInfo {
			count := addressNum["count"]
			size = count.(int32)
		}
		if size < 2 {
			filter := bson.D{primitive.E{Key: "_id", Value: address}}
			update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "address", Value: addresses}}}}
			_, err := UserCollection.UpdateOne(ctx, filter, update)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			c.IndentedJSON(400, "Not Allowed, only 2 addresses")
		}
		defer cancel()
		ctx.Done()

	}
}

func EditHomeAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("id")
		if userID == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"err": "invalid"})
			c.Abort()
			return
		}

		var editAddress models.Address
		if err := c.BindJSON(&editAddress); err != nil {
			c.IndentedJSON(http.StatusBadRequest, err.Error())
		}

		usertID, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			c.IndentedJSON(500, "Internal Server error")
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		filter := bson.D{primitive.E{Key: "_id", Value: usertID}}
		update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "address.0.house_name", Value: editAddress.House}, {Key: "$address.0.street_name", Value: editAddress.Street}, {Key: "address.0.city_name", Value: editAddress.City}, {Key: "address.0.pin_code", Value: editAddress.Pincode}}}}
		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(500, "Something went wrong")
			return
		}
		defer cancel()
		ctx.Done()
		c.IndentedJSON(200, "Successfully updated the home address")
	}

}

func EditWorkAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("id")
		if userID == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"err": "invalid"})
			c.Abort()
			return
		}

		var editAddress models.Address
		if err := c.BindJSON(&editAddress); err != nil {
			c.IndentedJSON(http.StatusBadRequest, err.Error())
		}

		usertID, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			c.IndentedJSON(500, "Internal Server error")
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		filter := bson.D{primitive.E{Key: "_id", Value: usertID}}
		update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "address.1.house_name", Value: editAddress.House}, {Key: "$address.1.street_name", Value: editAddress.Street}, {Key: "address.1.city_name", Value: editAddress.City}, {Key: "address.1.pin_code", Value: editAddress.Pincode}}}}
		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(500, "Something went wrong")
			return
		}
		defer cancel()
		ctx.Done()
		c.IndentedJSON(200, "Successfully updated work address")
	}
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
