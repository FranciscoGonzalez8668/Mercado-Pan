package dto

type ProductDto struct {
	Id_Product   int     `json:"id_product"`
	Name_product string  `json:"name_product"`
	Price        float64 `json:"cost"`
	Stock        int     `json:"stock"`
	Category     string  `json:"category"`
	Descripcion  string  `json:"descripcion"`

	/*Name string `json:"name"`
	UniversalCode string `json:"universal_code"`
	Picture string `json:"picture_url"`
	Price float32 `json:"base_price"`

	*/
}
type ProductCart struct {
	Id_Product   int     `json:"id_product"`
	Name_product string  `json:"name_product"`
	Price        float64 `json:"cost"`
}
type IdProCart struct {
	Id_products []string `json:"id_products"`
}

type ProductsDto []ProductDto
