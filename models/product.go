package models

// Product represents product
type Product struct {
	ID          uint32  `bson:"id_product" json:"id_product"`
	ProductName string  `bson:"product_name" json:"product_name"`
	Description string  `bson:"description" json:"description"`
	Price       float32 `bson:"price" json:"price"`
	Seller      Seller  `bson:"seller" json:"seller"`
}
