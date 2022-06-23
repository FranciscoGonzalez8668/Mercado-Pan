package product

import (
	//"fmt"

	"pan/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetDetalles(id_order int) []model.Detalle {
	var detalles []model.Detalle
	Db.Where("id_order = ? ", id_order).Find(&detalles)
	return detalles

}
func GetOrders(id_user int) []model.Order {
	var orders []model.Order
	Db.Where("id_user = ?", id_user).Find(&orders)
	return orders
}

func NewOrder(order model.Order, detalles []model.Detalle) model.Order {
	var product model.Product
	var LastDetalleDb model.Detalle
	var LastOrdenDb model.Order
	//var detalleAUX model.Detalle

	//Se verifica que haya stock de los productos
	for j := 0; j < len(detalles); j++ {
		Db.Where("id_product = ?", detalles[j].Id_Product).First(&product)
		if product.Stock < detalles[j].Cantidad {
			return order
		}
	}
	Db.Order("id_order DESC").First(&LastOrdenDb)
	order.Id_Order = 1 + LastOrdenDb.Id_Order
	//se seta el id_order en los detalles
	for i := 0; i < len(detalles); i++ {
		Db.Last(&LastDetalleDb)
		detalles[i].Id_Order = order.Id_Order
		detalles[i].Id_Detalle = 2 + i + LastDetalleDb.Id_Detalle
	}
	result := Db.Create(&order)
	if result.Error != nil {
		log.Error("no se pudo crear la orden")
	}
	log.Debug(detalles)
	for k := 0; k < len(detalles); k++ {
		result = Db.Create(&detalles[k])
		if result.Error != nil {
			Db.Where("id_order = ?", order.Id_Order).Delete(&order)
			return order
		}
	}

	for k := 0; k < len(detalles); k++ {
		Db.Model(&model.Product{}).Where("id_product = ?", detalles[k].Id_Product).Update("stock", gorm.Expr("stock - ?", detalles[k].Cantidad))
	}

	log.Debug("order created: ", order.Id_Order)
	return order
}
