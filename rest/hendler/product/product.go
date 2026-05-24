package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/saurav11sarkar/resapi/internal/model"
	"github.com/saurav11sarkar/resapi/utils"
)

var products []model.Product
var nextProductId = 1

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var body model.Product

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.SendJson(w, http.StatusBadRequest, model.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	body.Id = nextProductId
	nextProductId++
	products = append(products, body)
	utils.SendJson(w, http.StatusCreated, model.Response{
		Success: true,
		Message: "Product create successfully!",
		Data:    body,
	})
	return
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	utils.SendJson(w, http.StatusOK, model.Response{
		Success: true,
		Message: "Product get successfully!",
		Data:    products,
	})
	return
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.SendJson(w, http.StatusBadRequest, model.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	var product model.Product
	for _, product = range products {
		if product.Id == id {
			utils.SendJson(w, http.StatusOK, model.Response{
				Success: true,
				Message: "Product get successfully!",
				Data:    product,
			})
			return
		}
	}
	utils.SendJson(w, http.StatusNotFound, model.Response{
		Success: false,
		Message: "Product not found!",
		Data:    nil,
	})

}
