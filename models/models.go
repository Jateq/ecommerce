package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	//actually go doesn't have support of json without this ` ` symbols
	// and also _id is for mongoDB because it stores like that
	// bson - Binary Json for mongoDB actually
	// so basically third "column" used to specify how struct should be encoded and decoded when marshaling and unmarshaling
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName      *string            `json:"first_name" validate:"required,min=2,max=30"'`
	LastName       *string            `json:"last_name" validate:"required, min=2,max=30"`
	Password       *string            `json:"password" validate:"required, min=6"`
	Email          *string            `json:"email" validate:"email, required"`
	Phone          *string            `json:"phone" validate:"required"`
	Token          *string            `json:"tokengen"`
	RefreshToken   *string            `json:"refresh_token"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
	UserID         string             `json:"user_id"`
	UserCart       []ProductUser      `json:"usercart" bson:"usercart"`
	AddressDetails []Address          `json:"address" bson:"address"`
	OrderStatus    []Order            `json:"order_status" bson:"order"`
}

type Product struct {
	ProductID   primitive.ObjectID `bson:"_id"`
	ProductName *string            `json:"product_name"`
	Price       *uint64            `json:"price"`
	Rating      *uint8             `json:"rating"`
	Image       *string            `json:"image"`
}

type ProductUser struct {
	ProductID   primitive.ObjectID `bson:"_id"`
	ProductName *string            `json:"product_name" bson:"product_name"`
	Price       int                `json:"price" bson:"price"`
	Rating      *uint              `json:"rating" bson:"rating"`
	Image       *string            `json:"image" bson:"image"`
}

type Address struct {
	AddressID primitive.ObjectID `bson:"_id"`
	House     *string            `json:"house_name" bson:"house_name"`
	Street    *string            `json:"street_name" bson:"street_name"`
	City      *string            `json:"city_name" bson:"city_name"`
	Pincode   *string            `json:"pin_code" bson:"pin_code"`
}

type Order struct {
	OrderID       primitive.ObjectID `bson:"_id"`
	OrderCart     []ProductUser      `json:"order_list" bson:"order_list"`
	OrderedAt     time.Time          `json:"ordered_at" bson:"ordered_at"`
	Price         int                `json:"total_price" bson:"total_price"`
	Discount      *int               `json:"discount" bson:"discount"`
	PaymentMethod Payment            `json:"payment_method" bson:"payment_method"`
}

type Payment struct {
	Digital bool
	COD     bool
}
