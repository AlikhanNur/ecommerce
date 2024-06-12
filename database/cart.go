package database

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var (
	ErrCantFindProduct    = errors.New("can't find product")
	ErrCantDecodeProducts = errors.New("can't decode products")
	ErrUserIdIsNotValid   = errors.New("this user is not valid")
	ErrCantUpdateUser     = errors.New("can't add this product to the cart")
	ErrCantRemoveItemCart = errors.New("can't remove item cart")
	ErrCantGetItem        = errors.New("was unable to get item")
	ErrCantBuyCartItem    = errors.New("cannot update the purchase")
)

func AddProductToCart() gin.HandlerFunc {

}

func RemoveCartItem() gin.HandlerFunc {

}

func BuyItemFromCart() gin.HandlerFunc {

}

func InstantBuyer() gin.HandlerFunc {

}
