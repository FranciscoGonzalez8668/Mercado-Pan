package model

type Product struct {
	Id_Product   int     `gorm:"primaryKey"`
	Name_product string  `gorm:"type:varchar(350);not null;unique"`
	Price        float64 `gorm:"type:float;not null"`
	Stock        int     `gorm:"type:int;not null"`
	Category     string  `gorm:"type:varchar(350);not null"`
	Description  string  `gorm:"type:varchar(3000); not null"`

	/*Name string `json:"name"`
	UniversalCode string `json:"universal_code"`
	Picture string `json:"picture_url"`
	Price float32 `json:"base_price"`

	*/

}
