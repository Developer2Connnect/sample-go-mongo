package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SKU represents a SKU in the e-commerce portal
type SKU struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	SkuID          string             `bson:"sku_id,omitempty"`
	Size           string             `bson:"size,omitempty"`
	Stock          int                `bson:"stock,omitempty"`
	OnHold         int                `bson:"onhold,omitempty"`
	Dimension      string             `bson:"dimension,omitempty"`
	EAN            string             `bson:"ean,omitempty"`
	StockAvailable int                `bson:"stock_available,omitempty"`
}

// Color represents a Color in the e-commerce portal
type Color struct {
	ID            primitive.ObjectID   `bson:"_id,omitempty"`
	Slug          string               `bson:"slug,omitempty"`
	Name          string               `bson:"name,omitempty"`
	Title         string               `bson:"title,omitempty"`
	Description   string               `bson:"description,omitempty"`
	Images        []string             `bson:"images,omitempty"`
	Sku           []primitive.ObjectID `bson:"sku,omitempty"`
	Fabric        string               `bson:"fabric,omitempty"`
	MasterColor   string               `bson:"master_color,omitempty"`
	OriginCountry string               `bson:"origin_country,omitempty"`
	TotalQuantity int                  `bson:"total_quantity,omitempty"`
}

// Product represents a Product in the e-commerce portal
type Product struct {
	ID                   primitive.ObjectID   `bson:"_id,omitempty"`
	ProductID            string               `bson:"product_id,omitempty"`
	ArticleCode          string               `bson:"article_code,omitempty"`
	Price                int                  `bson:"price,omitempty"`
	OfferPrice           int                  `bson:"offer_price,omitempty"`
	Gender               string               `bson:"gender,omitempty"`
	Category             string               `bson:"category,omitempty"`
	Color                []primitive.ObjectID `bson:"color,omitempty"`
	Brand                string               `bson:"brand,omitempty"`
	Seo                  interface{}          `bson:"seo,omitempty"`
	MarketedBy           string               `bson:"marketed_by,omitempty"`
	ManufacturingDetails string               `bson:"manufacturing_details,omitempty"`
	HsnCode              string               `bson:"hsn_code,omitempty"`
	GstRate              float64              `bson:"gst_rate,omitempty"`
	Gst                  float64              `bson:"gst,omitempty"`
	Discount             float64              `bson:"discount,omitempty"`
	Quantity             int                  `bson:"quantity,omitempty"`
}

// CalculateStockAvailable calculates stockAvailable before saving SKU
func (sku *SKU) CalculateStockAvailable() {
	sku.StockAvailable = max(0, sku.Stock-sku.OnHold)
}

// CalculateTotalQuantity calculates totalQuantity before saving Color
func (color *Color) CalculateTotalQuantity() {
	totalQuantity := 0
	for _, skuID := range color.Sku {
		// Fetch SKU from database and calculate totalQuantity
	}
	color.TotalQuantity = totalQuantity
}

// CalculateQuantityGstDiscount calculates quantity, gst, and discount before saving Product
func (product *Product) CalculateQuantityGstDiscount() {
	totalQuantity := 0
	for _, colorID := range product.Color {
		// Fetch Color from database and sum up totalQuantity
	}
	product.Quantity = totalQuantity
	product.Gst = (product.GstRate * float64(product.OfferPrice)) / 100
	product.Discount = (float64(product.Price) - float64(product.OfferPrice)) / 100
}

// Helper function to calculate max value
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// SaveSKU saves a SKU to the database
func SaveSKU(ctx context.Context, sku *SKU) error {
	// Assuming you have a MongoDB client named "client"
	collection := client.Database("test").Collection("skus")
	_, err := collection.InsertOne(ctx, sku)
	if err != nil {
		log.Printf("Error inserting SKU: %v\n", err)
		return err
	}
	return nil
}

// SaveColor saves a Color to the database
func SaveColor(ctx context.Context, color *Color) error {
	// Assuming you have a MongoDB client named "client"
	collection := client.Database("test").Collection("colors")
	_, err := collection.InsertOne(ctx, color)
	if err != nil {
		log.Printf("Error inserting Color: %v\n", err)
		return err
	}
	return nil
}

// SaveProduct saves a Product to the database
func SaveProduct(ctx context.Context, product *Product) error {
	// Assuming you have a MongoDB client named "client"
	collection := client.Database("test").Collection("products")
	_, err := collection.InsertOne(ctx, product)
	if err != nil {
		log.Printf("Error inserting Product: %v\n", err)
		return err
	}
	return nil
}
