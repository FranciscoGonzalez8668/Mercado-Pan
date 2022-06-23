package dto

type CartDto struct {
	Products_cart ProductsDto `json:"products_cart"`
	Cant_cart     []int       `json:"cant_cart"`
	Price_cart    float32     `json:"price_cart"`
}
