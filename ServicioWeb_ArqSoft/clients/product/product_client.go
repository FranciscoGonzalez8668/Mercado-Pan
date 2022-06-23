package product

import (
	"pan/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

//FUNCION LISTA
//Params: array product.product
//Proces: Busca en la base de datos todos los productos que contiene ese nombre
// o parte de el y los devuelve.

func GetProductByName(key string) []model.Product {
	var products []model.Product
	Db.Where("name_product LIKE ?", key+"%").Find(&products)

	return products
}

func GetProductCart(id_products []int) []model.Product {
	var products []model.Product
	Db.Where("id_product IN (?)", id_products).Find(&products)
	if products == nil {
		log.Debug("Fue nulo")
		products[0].Id_Product = 0
	}
	return products

}
func GetProductById(id_product int) model.Product {
	var product model.Product

	Db.Where("id_product = ?", id_product).First(&product)

	return product
}

//FUNCION LISTA
//
//Params: 	string
//Return:	array model.producto
//Proces:	Busca en BD productos con categoria @params@
//
//
func GetProductByCat(cat string) []model.Product { //

	var products []model.Product

	Db.Where("category LIKE ?", cat+"%").Find(&products)

	log.Debug("product: ", products)

	return products

	/*var productAux model.Product
	for j := 0; j < 2; j++ {
		fmt.Println("se entra ford client i= ", j)
		productAux.Id_Product = j
		productAux.Category = "electronico"
		productAux.Desciption = "computadora lenovo"
		productAux.Name_product = "Lenovo Computador"
		productAux.Price = 3.25
		productAux.Stock = 1
		products = append(products, productAux)
	}
	fmt.Println("array created")*/

}

func GetProductAll() []model.Product {
	var products []model.Product

	Db.Find(&products)

	log.Debug("products: ", products)

	return products
}
