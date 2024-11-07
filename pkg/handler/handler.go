package handler

import (
	"example_product/pkg/websocket"

	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	ProductCollection *mongo.Collection
	UserCollection    *mongo.Collection
	Hub               *websocket.WebSocketHub
}
