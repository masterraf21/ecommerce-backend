package models

// Buyer represents buyer
type Buyer struct {
	ID              uint32 `bson:"id_buyer" json:"id_buyer"`
	Email           string `bson:"email" json:"email"`
	Name            string `bson:"name" json:"name"`
	Password        string `bson:"password" json:"password"`
	DeliveryAddress string `bson:"delivery_address" json:"delivery_address"`
}
