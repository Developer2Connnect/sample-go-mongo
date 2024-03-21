package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ProductRepository defines methods for interacting with products
type ProductRepository struct {
	collection *mongo.Collection
}

// NewProductRepository creates a new instance of ProductRepository
func NewProductRepository(collection *mongo.Collection) *ProductRepository {
	return &ProductRepository{
		collection: collection,
	}
}

// CreateProduct inserts a new product into the database
func (pr *ProductRepository) CreateProduct(ctx context.Context, product *Product) error {
	_, err := pr.collection.InsertOne(ctx, product)
	if err != nil {
		log.Printf("Error inserting product: %v\n", err)
		return err
	}
	return nil
}

// GetProductByID retrieves a product by its ID from the database
func (pr *ProductRepository) GetProductByID(ctx context.Context, id primitive.ObjectID) (*Product, error) {
	var product Product
	err := pr.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		log.Printf("Error retrieving product: %v\n", err)
		return nil, err
	}
	return &product, nil
}

// UpdateProduct updates an existing product in the database
func (pr *ProductRepository) UpdateProduct(ctx context.Context, id primitive.ObjectID, product *Product) error {
	_, err := pr.collection.ReplaceOne(ctx, bson.M{"_id": id}, product)
	if err != nil {
		log.Printf("Error updating product: %v\n", err)
		return err
	}
	return nil
}
