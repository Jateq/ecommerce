package database

import (
	"errors"
)

var (
	ErrCantFindProduct    = errors.New("can't find the product")
	ErrCantDecodeProducts = errors.New("can't find the product")
	ErrUserIdIsNotValid   = errors.New("this user is not valid")
	ErrCantUpdateUser     = errors.New("can't add this product to cart")
	ErrCantRemoveItemCart = errors.New("can't remove this item from cart")
	ErrCantGetItem        = errors.New("unable to get item from the cart")
	ErrCantBuyItem        = errors.New("can't update the purchase")
)

func AddProductToCart() error {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}
