package dto

type OrderDto struct {
	Id_order  int          `json:"id_order"`
	Total     float32      `json:"total"`
	Det_order []DetalleDto `json:"det_orden"`
	Adress    AdressDto    `json:"adress"`
}
type NewOrderDto struct {
	Token    string          `json:"id_user"`
	Detalles []DetalleOrdDto `json:"detalles"`
}

type OrdersHistoryDto struct {
	Id_User int        `json:"id_user"`
	Order   []OrderDto `json:"order"`
}

type OrderDetDto struct {
	Id_order  int            `json:"id_user"`
	Total     float32        `json:"total"`
	Det_order DetallesDetDto `json:"det_orden"`
}
