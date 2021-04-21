package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Order represents order
type Order struct {
	ID              uint32     `bson:"id_order" json:"id_order"`
	BuyerID         uint32     `bson:"id_buyer" json:"id_buyer"`
	Buyer           *Buyer     `bson:"buyer" json:"buyer"`
	SellerID        uint32     `bson:"id_seller" json:"id_seller"`
	Seller          *Seller    `bson:"seller" json:"seller"`
	SourceAddress   string     `bson:"source_address" json:"source_address"`
	DeliveryAddress string     `bson:"delivery_address" json:"delivery_address"`
	Items           []*Product `bson:"items" json:"items"`
	Quantity        uint32     `bson:"quantity" json:"quantity"`
	Price           float32    `bson:"price" json:"price"`
	TotalPrice      float32    `bson:"total_price" json:"total_price"`
	Status          string     `bson:"status" json:"status"`
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
