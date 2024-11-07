package handler

import (
	"context"
	"encoding/json"
	"example_product/pkg/dto"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (h *Handler) CreateProduct(response http.ResponseWriter, request *http.Request) {
	var product dto.Product
	if err := json.NewDecoder(request.Body).Decode(&product); err != nil {
		responseWithJson(response, http.StatusInternalServerError, []byte(`{ "message": "`+err.Error()+`" }`))
		return
	}

	ctx, cancel := newContext()
	defer cancel()

	result, _ := h.ProductCollection.InsertOne(ctx, product)
	id := result.InsertedID
	product.ID = id.(primitive.ObjectID)

	responseWithJson(response, http.StatusCreated, product)

	// Phát thông báo cho các client WebSocket
	productJSON, _ := json.Marshal(product)
	h.Hub.Broadcast(productJSON)
}

func (h *Handler) GetProductByID(response http.ResponseWriter, request *http.Request) {
	var product dto.Product
	id, _ := primitive.ObjectIDFromHex(mux.Vars(request)["id"])

	ctx, cancel := newContext()
	defer cancel()

	err := h.ProductCollection.FindOne(ctx, dto.Product{ID: id}).Decode(&product)
	if err != nil {
		responseWithJson(response, http.StatusInternalServerError, []byte(`{ "message": "`+err.Error()+`" }`))
		return
	}

	responseWithJson(response, http.StatusCreated, product)
}

func (h *Handler) GetAllProduct(response http.ResponseWriter, request *http.Request) {
	var products []dto.Product

	ctx, cancel := newContext()
	defer cancel()

	cursor, err := h.ProductCollection.Find(ctx, bson.M{})
	if err != nil {
		responseWithJson(response, http.StatusInternalServerError, []byte(`{ "message": "`+err.Error()+`" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var product dto.Product
		cursor.Decode(&product)
		products = append(products, product)
	}
	if err := cursor.Err(); err != nil {
		responseWithJson(response, http.StatusInternalServerError, []byte(`{ "message": "`+err.Error()+`" }`))
		return
	}

	responseWithJson(response, http.StatusCreated, products)
}

func (h *Handler) UpdateProduct(response http.ResponseWriter, request *http.Request) {
	id, _ := primitive.ObjectIDFromHex(mux.Vars(request)["id"])

	var product dto.Product
	if err := json.NewDecoder(request.Body).Decode(&product); err != nil {
		responseWithJson(response, http.StatusInternalServerError, []byte(`{ "message": "`+err.Error()+`" }`))
		return
	}

	ctx, cancel := newContext()
	defer cancel()

	result, err := h.ProductCollection.UpdateOne(ctx, dto.Product{ID: id}, bson.M{"$set": product})
	if err != nil {
		responseWithJson(response, http.StatusInternalServerError, []byte(`{ "message": "`+err.Error()+`" }`))
		return
	}

	responseWithJson(response, http.StatusCreated, result)

	// Phát thông báo cho các client WebSocket
	productJSON, _ := json.Marshal(product)
	h.Hub.Broadcast(productJSON)
}

func (h *Handler) DeleteProduct(response http.ResponseWriter, request *http.Request) {
	id, _ := primitive.ObjectIDFromHex(mux.Vars(request)["id"])

	ctx, cancel := newContext()
	defer cancel()

	result, err := h.ProductCollection.DeleteOne(ctx, dto.Product{ID: id})
	if err != nil {
		responseWithJson(response, http.StatusInternalServerError, []byte(`{ "message": "`+err.Error()+`" }`))
		return
	}

	responseWithJson(response, http.StatusCreated, result)
}

func (h *Handler) QueryProduct(response http.ResponseWriter, request *http.Request) {
	var products []dto.Product

	// Parse the limit from query parameters (default to 10 if not provided)
	limitParam := request.URL.Query().Get("limit")
	limit := int64(10) // default limit
	if limitParam != "" {
		if l, err := strconv.ParseInt(limitParam, 10, 64); err == nil {
			limit = l
		}
	}

	// Create the aggregation pipeline
	pipeline := mongo.Pipeline{
		{{
			Key:   "$match",
			Value: bson.M{"age": bson.M{"$gt": 20}},
		}},
		{{
			Key:   "$limit",
			Value: limit,
		}},
	}

	ctx, cancel := newContext()
	defer cancel()

	cursor, err := h.ProductCollection.Aggregate(ctx, pipeline)
	if err != nil {
		responseWithJson(response, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product dto.Product
		if err := cursor.Decode(&product); err != nil {
			responseWithJson(response, http.StatusInternalServerError, map[string]string{"message": err.Error()})
			return
		}
		products = append(products, product)
	}
	if err := cursor.Err(); err != nil {
		responseWithJson(response, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	responseWithJson(response, http.StatusOK, products)
}

func responseWithJson(writer http.ResponseWriter, status int, object interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(object)
}

func newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
