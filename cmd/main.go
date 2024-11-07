package main

import (
	"context"
	"example_product/pkg/handler"
	"example_product/pkg/middleware"
	"example_product/pkg/websocket"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	productCollection := client.Database("example_product").Collection("product")
	userCollection := client.Database("example_product").Collection("users")
	hub := websocket.NewWebSocketHub()

	h := &handler.Handler{
		ProductCollection: productCollection,
		UserCollection:    userCollection,
		Hub:               hub,
	}

	go h.Hub.Run()

	router := mux.NewRouter().StrictSlash(true)

	// Endpoint WebSocket
	router.HandleFunc("/ws", h.Hub.ConnectWebsocket).Methods(http.MethodGet)

	// Các route REST API
	router.HandleFunc("/login", h.Login).Methods(http.MethodPost)
	router.HandleFunc("/register", h.Register).Methods(http.MethodPost)

	// Các route sản phẩm có bảo vệ bằng JWT
	productRouter := router.PathPrefix("/product").Subrouter()
	productRouter.Use(middleware.JWTAuthMiddleware)
	productRouter.HandleFunc("", h.GetAllProduct).Methods(http.MethodGet)
	productRouter.HandleFunc("/{id}", h.GetProductByID).Methods(http.MethodGet)
	productRouter.HandleFunc("", h.CreateProduct).Methods(http.MethodPost)
	productRouter.HandleFunc("/{id}", h.UpdateProduct).Methods(http.MethodPut)
	productRouter.HandleFunc("/{id}", h.DeleteProduct).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8080", router))
}
