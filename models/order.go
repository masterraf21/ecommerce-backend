package models

// Order represents order
type Order struct {
	ID              uint32    `bson:"id_order" json:"id_order"`
	Buyer           Buyer     `bson:"buyer" json:"buyer"`
	Seller          Seller    `bson:"seller" json:"seller"`
	SourceAddress   string    `bson:"source_address" json:"source_address"`
	DeliveryAddress string    `bson:"delivery_address" json:"delivery_address"`
	Items           []Product `bson:"items" json:"items"`
	Quantity        uint32    `bson:"quantity" json:"quantity"`
	Price           float32   `bson:"price" json:"price"`
	TotalPrice      float32   `bson:"total_price" json:"total_price"`
}
