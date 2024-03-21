package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/developer2connect/sample-go-mongo/models"
	"github.com/developer2connect/sample-go-mongo/repository"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ProductHandlers contains handlers for product related requests
type ProductHandlers struct {
	repo *repository.ProductRepository
}

// NewProductHandlers creates a new instance of ProductHandlers
func NewProductHandlers(repo *repository.ProductRepository) *ProductHandlers {
	return &ProductHandlers{repo: repo}
}

// CreateProductEndpoint creates a new product
func (ph *ProductHandlers) CreateProductEndpoint(w http.ResponseWriter, req *http.Request) {
	var product models.Product
	err := json.NewDecoder(req.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	err = ph.repo.CreateProduct(ctx, &product)
	if err != nil {
		http.Error(w, "Error creating product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// GetProductEndpoint retrieves a product by its ID
func (ph *ProductHandlers) GetProductEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	product, err := ph.repo.GetProductByID(ctx, id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(product)
}

// UpdateProductEndpoint updates an existing product
func (ph *ProductHandlers) UpdateProductEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var product models.Product
	err = json.NewDecoder(req.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	err = ph.repo.UpdateProduct(ctx, id, &product)
	if err != nil {
		http.Error(w, "Error updating product", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}
