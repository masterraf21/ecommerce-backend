package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Seller represents seller
type Seller struct {
	ID            uint32 `bson:"id_seller" json:"id_seller"`
	Email         string `bson:"email" json:"email"`
	Name          string `bson:"name" json:"name"`
	Password      string `bson:"password" json:"password"`
	PickupAddress string `bson:"pickup_address" json:"pickup_address"`
}

// SellerRepository represents repo functions for seller
type SellerRepository interface {
	Store(seller *Seller) (primitive.ObjectID, error)
	GetAll() ([]Seller, error)
	GetByID(id uint32) (*Seller, error)
	GetByOID(oid primitive.ObjectID) (*Seller, error)
	UpdateArbitrary(id uint32, key string, value interface{}) error
}

// SellerBody body for buyer
type SellerBody struct {
	Email         string `json:"email"`
	Name          string `json:"name"`
	Password      string `json:"password"`
	PickupAddress string `json:"pickup_address"`
}

// SellerUsecase for seller usecase
type SellerUsecase interface {
	CreateSeller(seller SellerBody) (uint32, error)
	GetAll() ([]Seller, error)
	GetByID(id uint32) (*Seller, error)
}
