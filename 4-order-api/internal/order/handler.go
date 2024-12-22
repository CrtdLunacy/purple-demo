package order

import (
	"fmt"
	"go/order-api/configs"
	"go/order-api/pkg/middleware"
	"go/order-api/pkg/request"
	"go/order-api/pkg/response"
	"net/http"
	"strconv"
)

type OrderHandlerDeps struct {
	*OrderRepository
	*OrderService
	*configs.Config
}

type OrderHandler struct {
	*OrderRepository
	*OrderService
}

func NewOrderHandler(router *http.ServeMux, deps OrderHandlerDeps) {
	handler := &OrderHandler{
		OrderRepository: deps.OrderRepository,
		OrderService:    deps.OrderService,
	}

	router.Handle("GET /orders", middleware.AuthCheck(handler.GetAll(), deps.Config))
	router.Handle("GET /order/{id}", middleware.AuthCheck(handler.GetById(), deps.Config))
	router.Handle("POST /order", middleware.AuthCheck(handler.Create(), deps.Config))
}

func (orderHandler *OrderHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[CreateRequest](w, req)
		if err != nil {
			response.ResponseJSON(w, ErrBadRequest, http.StatusBadRequest)
			return
		}

		ctx := req.Context()
		authCtx, ok := ctx.Value("authData").(middleware.AuthContext)
		if !ok {
			response.ResponseJSON(w, ErrContextTypeChecking, http.StatusAccepted)
			return
		}

		fmt.Println(authCtx)
		order, err := orderHandler.OrderService.CreateOrder(&OrderData{
			UserId:     authCtx.UserID,
			Products:   body.Products,
			TotalPrice: body.TotalPrice,
			UserPhone:  authCtx.Phone,
		})
		if err != nil {
			response.ResponseJSON(w, ErrOrderCreate, http.StatusBadRequest)
			return
		}

		response.ResponseJSON(w, order, http.StatusCreated)
	}
}

func (orderHandler *OrderHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		authCtx, ok := ctx.Value("authData").(middleware.AuthContext)
		if !ok {
			response.ResponseJSON(w, ErrContextTypeChecking, http.StatusAccepted)
			return
		}

		userOrders, err := orderHandler.OrderRepository.GetAll(authCtx.UserID)
		if err != nil {
			response.ResponseJSON(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.ResponseJSON(w, userOrders, http.StatusAccepted)
	}
}

func (orderHandler *OrderHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id, err := strconv.Atoi(req.PathValue("id"))
		if err != nil {
			response.ResponseJSON(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := req.Context()
		authCtx, ok := ctx.Value("authData").(middleware.AuthContext)
		if !ok {
			response.ResponseJSON(w, ErrContextTypeChecking, http.StatusAccepted)
			return
		}

		userOrder, err := orderHandler.OrderRepository.GetByID(authCtx.UserID, uint(id))
		if err != nil {
			response.ResponseJSON(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.ResponseJSON(w, userOrder, http.StatusAccepted)
	}
}
