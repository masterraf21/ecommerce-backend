package models

// Counter represents counter for
type Counter struct {
	BuyerID   uint32 `bson:"id_buyer" json:"id_buyer"`
	ProductID uint32 `bson:"id_product" json:"id_product"`
	SellerID  uint32 `bson:"id_seller" json:"id_seller"`
	OrderID   uint32 `bson:"id_order" json:"id_order"`
}

// CounterRepository repo for counter
type CounterRepository interface {
	Get(collectionName string, identifier string) (uint32, error)
}
