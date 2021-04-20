package models

// Seller represents seller
type Seller struct {
	ID            uint32 `bson:"id_seller" json:"id_seller"`
	Email         string `bson:"email" json:"email"`
	Name          string `bson:"name" json:"name"`
	Password      string `bson:"password" json:"password"`
	PickupAddress string `bson:"pickup_address" json:"pickup_address"`
}
