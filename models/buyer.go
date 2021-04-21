package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Buyer represents buyer
type Buyer struct {
	ID              uint32 `bson:"id_buyer" json:"id_buyer"`
	Email           string `bson:"email" json:"email"`
	Name            string `bson:"name" json:"name"`
	Password        string `bson:"password" json:"password"`
	DeliveryAddress string `bson:"delivery_address" json:"delivery_address"`
}

// BuyerBody body for buyer
type BuyerBody struct {
	Email           string `json:"email"`
	Name            string `json:"name"`
	Password        string `json:"password"`
	DeliveryAddress string `json:"delivery_address"`
}

// BuyerRepository represents repo functions for Buyer
type BuyerRepository interface {
	Store(buyer *Buyer) (primitive.ObjectID, error)
	GetAll() ([]Buyer, error)
	GetByID(id uint32) (*Buyer, error)
	GetByOID(oid primitive.ObjectID) (*Buyer, error)
	UpdateArbitrary(id uint32, key string, value interface{}) error
}

// BuyerUsecase will create usecase for buyer
type BuyerUsecase interface {
	CreateBuyer(buyer BuyerBody) (uint32, error)
	GetAll() ([]Buyer, error)
	GetByID(id uint32) (*Buyer, error)
}
