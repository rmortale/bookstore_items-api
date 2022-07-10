package controllers

import (
	"fmt"
	"github.com/rmortale/bookstore_items-api/domain/items"
	"github.com/rmortale/bookstore_items-api/services"
	"github.com/rmortale/bookstore_oauth-go/oauth"
	"net/http"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct {
}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		//TODO
		return
	}
	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}
	result, err := services.ItemsService.Create(item)
	if err != nil {
		return
	}
	fmt.Println(result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
