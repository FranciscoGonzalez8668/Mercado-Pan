package model

type Order struct {
	Id_Order int     `gomr:"primaryKey"`
	Id_User  int     `gorm:"type:int"`
	Total    float32 `gorm:"type:float"`
}

/*type OrderDet struct {
	Id_Orden      int
	Total         float32
	Product       string
	Cantidad      int
	Precio_Moment float32
}*/
