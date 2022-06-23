package model

type Detalle struct {
	Id_Order      int     `gorm:"primaryKey"`
	Id_Detalle    int     `gorm:"primaryKey"`
	Id_Product    int     `gorm:"type:int"`
	Cantidad      int     `gorm:"type:int;not null"`
	Precio_Moment float32 `gorm:"type:float;not null"`
}

/*type DetalleDet struct {
	//Id_detalle    int
	Cantidad      int
	Precio_Moment float32
	Id_Product    int
	Id_Order      int
}*/
