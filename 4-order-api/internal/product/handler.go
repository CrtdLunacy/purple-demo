package product

import (
	"go/order-api/pkg/request"
	"go/order-api/pkg/response"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type ProoductHandlerDeps struct {
	ProductRepository *ProductRepository
}

type ProductHandler struct {
	ProductRepository *ProductRepository
}

func NewProductHandler(router *http.ServeMux, deps ProoductHandlerDeps) {
	handler := &ProductHandler{
		ProductRepository: deps.ProductRepository,
	}
	router.HandleFunc("POST /product", handler.Create())
	router.HandleFunc("GET /products", handler.GetProducts())
	router.HandleFunc("GET /product/{id}", handler.GetSingleProduct())
	router.HandleFunc("PATCH /product/{id}", handler.Update())
	router.HandleFunc("DELETE /product/{id}", handler.Delete())
}

func (productHandler *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[CreateRequest](w, req)
		if err != nil {
			response.ResponseJSON(w, err.Error(), http.StatusBadRequest)
			return
		}

		createdProduct, err := productHandler.ProductRepository.Create(&Product{
			Name:        body.Name,
			Description: body.Description,
			Images:      body.Images,
			Price:       body.Price,
		})
		if err != nil {
			response.ResponseJSON(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.ResponseJSON(w, createdProduct, http.StatusAccepted)
	}
}

func (productHandler *ProductHandler) GetProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		queryParams := req.URL.Query()
		var limit, offset *int

		if l := queryParams.Get("limit"); l != "" {
			parsedLimit, err := strconv.Atoi(l)
			if err == nil {
				limit = &parsedLimit
			}
		}

		if o := queryParams.Get("offset"); o != "" {
			parsedOffset, err := strconv.Atoi(o)
			if err == nil {
				offset = &parsedOffset
			}
		}

		products, err := productHandler.ProductRepository.GetProducts(limit, offset)
		if err != nil {
			response.ResponseJSON(w, err.Error(), http.StatusNotFound)
			return
		}

		response.ResponseJSON(w, products, http.StatusOK)
	}
}

func (productHandler *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[UpdateRequest](w, req)
		if err != nil {
			response.ResponseJSON(w, err.Error(), http.StatusBadRequest)
			return
		}

		idString := req.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			response.ResponseJSON(w, err.Error(), http.StatusBadRequest)
			return
		}

		product, err := productHandler.ProductRepository.Update(&Product{
			Model:       gorm.Model{ID: uint(id)},
			Name:        body.Name,
			Description: body.Description,
			Images:      body.Images,
			Price:       body.Price,
		})
		if err != nil {
			response.ResponseJSON(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.ResponseJSON(w, product, http.StatusOK)
	}
}

func (productHandler *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		idString := req.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			response.ResponseJSON(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = productHandler.ProductRepository.GetById(uint(id))
		if err != nil {
			response.ResponseJSON(w, err.Error(), http.StatusNotFound)
			return
		}

		err = productHandler.ProductRepository.Delete(uint(id))
		if err != nil {
			response.ResponseJSON(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.ResponseJSON(w, nil, http.StatusOK)
	}
}

func (productHandler *ProductHandler) GetSingleProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		idString := req.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			response.ResponseJSON(w, err.Error(), http.StatusBadRequest)
			return
		}

		product, err := productHandler.ProductRepository.GetById(uint(id))
		if err != nil {
			response.ResponseJSON(w, err.Error(), http.StatusNotFound)
			return
		}

		response.ResponseJSON(w, product, http.StatusOK)
	}
}
