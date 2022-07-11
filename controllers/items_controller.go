package controllers

import (
	"encoding/json"
	"github.com/rmortale/bookstore_items-api/domain/items"
	"github.com/rmortale/bookstore_items-api/services"
	"github.com/rmortale/bookstore_items-api/utils/http_utils"
	"github.com/rmortale/bookstore_oauth-go/oauth"
	"github.com/rmortale/bookstore_utils-go/rest_errors"
	"io/ioutil"
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
		//http_utils.RespondErrorJson(w, *err)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondErrorJson(w, *respErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondErrorJson(w, *respErr)
		return
	}

	itemRequest.Seller = oauth.GetCallerId(r)

	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondErrorJson(w, *createErr)
		return
	}
	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
