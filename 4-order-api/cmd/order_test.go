package main

import (
	"bytes"
	"context"
	"encoding/json"
	"go/order-api/internal/auth"
	"go/order-api/internal/order"
	"go/order-api/internal/product"
	"go/order-api/internal/user"
	"go/order-api/pkg/middleware"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDb() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func initData(db *gorm.DB) {
	//создать пользователя
	if err := db.Create(&user.User{
		Phone:         "+79991234567",
		Code:          1234,
		SessionId:     "1234567890",
		SessionExpiry: time.Now().Add(time.Minute * 5),
	}).Error; err != nil {
		panic(err)
	}

	if err := db.Create(&product.Product{
		Name:        "Product 1",
		Price:       100,
		Description: "test1",
		Images:      pq.StringArray{},
	}).Error; err != nil {
		panic(err)
	}

	if err := db.Create(&product.Product{
		Name:        "Product 2",
		Description: "test2",
		Images:      pq.StringArray{"https://i.imgur.com/DuRQ3b1.jpeg"},
		Price:       200,
	}).Error; err != nil {
		panic(err)
	}
}

func removeData(db *gorm.DB) {
	var order order.Order
	if err := db.Preload("Products").Where("user_phone = ?", "+79991234567").First(&order).Error; err != nil {
		panic("Error finding order: " + err.Error())
	}

	// Удаляем связи между заказом и продуктами
	if err := db.Model(&order).Association("Products").Clear(); err != nil {
		panic("Error deleting product associations: " + err.Error())
	}

	// Удаляем сам заказ
	if err := db.Unscoped().
		Where("user_phone = ?", "+79991234567").
		Delete(&order).Error; err != nil {
		panic("Error deleting orders: " + err.Error())
	}

	// Удаляем продукты, если они не используются в других заказах
	productNames := []string{"Product 1", "Product 2"}
	if err := db.Unscoped().
		Where("name IN ?", productNames).
		Delete(&product.Product{}).Error; err != nil {
		panic("Error deleting products: " + err.Error())
	}

	// Удаляем пользователя
	if err := db.Unscoped().
		Where("phone = ?", "+79991234567").
		Delete(&user.User{}).Error; err != nil {
		panic("Error deleting user: " + err.Error())
	}
}

// standart test of verify user on login, choose some products to cart and make order, then check it correctly done
func TestOrderSuccess(t *testing.T) {
	//Prepare to test
	db := initDb()
	initData(db)

	//create test server
	testServer := httptest.NewServer(App())
	defer testServer.Close()

	//prepare verify user data
	data, _ := json.Marshal(&auth.VerifyRequest{
		Code:      1234,
		SessionId: "1234567890",
	})

	//send verify request
	res, err := http.Post(testServer.URL+"/auth/verify", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	//check verify response status code
	if res.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status code %d, got %d", 201, res.StatusCode)
	}

	//read verify response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	//unmarshal verify response body into struct
	var resData auth.VerifyResponse
	if err := json.Unmarshal(body, &resData); err != nil {
		t.Fatal(err)
	}

	//check token is not empty
	if resData.Token == "" {
		t.Fatal("Expected non-empty token")
	}

	//check order creation request data
	var products []product.Product
	if err := db.Find(&products).Error; err != nil {
		t.Fatalf("Failed to get products: %v", err)
	}

	//get product IDs for order creation
	var productIDs []uint
	for _, p := range products {
		productIDs = append(productIDs, p.ID)
	}

	//prepare order creation request data
	createRequest := &order.CreateRequest{
		Products:   productIDs,
		TotalPrice: 300.00,
	}

	//prepare order creation request body in JSON format
	createRequestData, err := json.Marshal(createRequest)
	if err != nil {
		t.Fatal("Failed to marshal request data:", err)
	}

	//add context to request
	authCtx := middleware.AuthContext{
		Phone: "+79991234567",
	}
	ctx := context.WithValue(context.Background(), "authData", authCtx)

	//create order request
	req, err := http.NewRequestWithContext(ctx, "POST", testServer.URL+"/order", bytes.NewReader(createRequestData))
	if err != nil {
		t.Fatal(err)
	}

	// set request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+resData.Token)

	// send order creation request
	client := &http.Client{}
	resOrder, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	// check order creation response status code
	if resOrder.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status code %d, got %d", 201, resOrder.StatusCode)
	}

	//read order creation response body
	bodyOrder, err := io.ReadAll(resOrder.Body)
	if err != nil {
		t.Fatal(err)
	}

	//unmarshal order creation response body into struct
	var resOrderData order.Order
	if err := json.Unmarshal(bodyOrder, &resOrderData); err != nil {
		t.Fatal(err)
	}

	//check order response data
	if len(resOrderData.Products) != 2 {
		t.Fatalf("Expected 2 products, got %d", len(resOrderData.Products))
	}

	//match user phone in order
	if resOrderData.UserPhone != "+79991234567" {
		t.Fatalf("Expected user ID %s, got %d", "+79991234567", resOrderData.UserID)
	}

	//remove data from db
	removeData(db)
}
