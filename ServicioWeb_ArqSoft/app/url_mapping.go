package app

import (
	orderController "pan/controllers/order"
	productController "pan/controllers/product"
	userController "pan/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	router.GET("/OK", userController.OK)

	router.POST("/login", userController.LoginUser) //identificar el usuario ^ devolver una toquen

	router.GET("/product/product/:KeyPro", productController.GetProductByName) // Busca productos por Palabra clave

	router.GET("/product/all", productController.GetProductAll)

	router.PUT("/compra/add", orderController.NewOrder) //Debe crear la nueva orden verificar cantidades

	router.POST("/cart/ProLo", productController.GetProductCart) //CartController.CartMananger) //Sumar Restar sacar y agregar productos en carrito

	router.GET("/User/OrH/:IU", orderController.GetHistory) //Devolver compras por user

	router.GET("/product/cat/:KeyCat", productController.GetProductByCat) //product.Controller.GetProductByCat)//Traer Productos por Categoria

	router.GET("/order/order:id_user", orderController.GetHistory)

	router.POST("/order", orderController.NewOrder)

	log.Info("Finishing mappings configurations")

}
