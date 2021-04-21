package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Order represents order
type Order struct {
	ID              uint32        `bson:"id_order" json:"id_order"`
	BuyerID         uint32        `bson:"id_buyer" json:"id_buyer"`
	Buyer           *Buyer        `bson:"buyer" json:"buyer"`
	SellerID        uint32        `bson:"id_seller" json:"id_seller"`
	Seller          *Seller       `bson:"seller" json:"seller"`
	SourceAddress   string        `bson:"source_address" json:"source_address"`
	DeliveryAddress string        `bson:"delivery_address" json:"delivery_address"`
	Products        []OrderDetail `bson:"products" json:"products"`
	TotalPrice      float32       `bson:"total_price" json:"total_price"`
	Status          string        `bson:"status" json:"status"`
}

// OrderDetail will detail the order
type OrderDetail struct {
	ProductID  uint32   `json:"id_product" bson:"id_product"`
	Product    *Product `json:"product" bson:"product"`
	Quantity   uint32   `json:"quantity" bson:"quantity"`
	TotalPrice float32  `json:"total_price" bson:"total_price"`
}

// ProductDetail for body
type ProductDetail struct {
	ProductID uint32 `json:"id_product"`
	Quantity  uint32 `json:"quantity"`
}

// OrderBody body from json
type OrderBody struct {
	BuyerID         uint32          `json:"id_buyer"`
	SellerID        uint32          `json:"id_seller"`
	SourceAddress   string          `json:"source_address"`
	DeliveryAddress string          `json:"delivery_address"`
	Products        []ProductDetail `json:"products"`
	TotalPrice      float32         `json:"total_price"`
}

// OrderRepository reprresents repo functions for order
type OrderRepository interface {
	Store(order *Order) (primitive.ObjectID, error)
	GetAll() ([]Order, error)
	GetByID(id uint32) (*Order, error)
	GetByOID(oid primitive.ObjectID) (*Order, error)
	UpdateArbitrary(id uint32, key string, value interface{}) error
	GetBySellerID(sellerID uint32) ([]Order, error)
	GetByBuyerID(buyerID uint32) ([]Order, error)
	GetByBuyerIDAndStatus(buyerID uint32, status string) ([]Order, error)
	GetBySellerIDAndStatus(sellerID uint32, status string) ([]Order, error)
	GetByStatus(status string) ([]Order, error)
}

// OrderUsecase usecase for order
type OrderUsecase interface {
	CreateOrder(order OrderBody) (uint32, error)
	AcceptOrder(id uint32) error
	GetAll() ([]Order, error)
	GetByID(id uint32) (*Order, error)
	GetBySellerID(sellerID uint32) ([]Order, error)
	GetByBuyerID(buyerID uint32) ([]Order, error)
	GetByBuyerIDAndStatus(buyerID uint32, status string) ([]Order, error)
	GetBySellerIDAndStatus(sellerID uint32, status string) ([]Order, error)
	GetByStatus(status string) ([]Order, error)
}
