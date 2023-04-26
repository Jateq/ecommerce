package models

import (
	_ "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID
	FirstName
	LastName
	Password
	Email
	Phone
	Token
	RefreshToken
	CreatedAt
	UpdatedAt
	UserID
	UserCart
	AddresDetails
	OrderStatus
}
