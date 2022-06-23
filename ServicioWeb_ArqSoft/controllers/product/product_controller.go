package productController

import (
	"fmt"
	"net/http"
	"pan/dto"
	service "pan/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetProductByName(c *gin.Context) {

	name := c.Param("KeyPro")
	log.Debug(name)
	var productDto dto.ProductsDto

	productDto, err := service.ProductService.GetProductByName(name)

	if err != nil {

		c.JSON(err.Status(), err)
		return

	}

	c.JSON(http.StatusOK, productDto)

}

func GetProductByCat(c *gin.Context) {

	cat := c.Param("KeyCat")
	log.Debug(cat)

	productsDto, err := service.ProductService.GetProductByCat(cat)
	fmt.Println("todas las consultas ok")
	if err != nil {
		c.JSON(err.Status(), productsDto)
		return
	}
	c.JSON(http.StatusOK, productsDto)
}

func GetProductAll(c *gin.Context) {
	productsDto, err := service.ProductService.GetProductAll()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productsDto)
}

func GetProductCart(c *gin.Context) {
	var productsCart dto.IdProCart
	var products []dto.ProductCart
	err := c.BindJSON(&productsCart)
	log.Debug("productos", productsCart)
	if err != nil {
		log.Debug(err)
		c.JSON(http.StatusBadRequest, products)
		return
	}
	products, err = service.ProductService.GetProductsCart(productsCart)
	if err != nil {
		log.Debug(err)
		c.JSON(http.StatusBadRequest, products)
		return
	}
	c.JSON(http.StatusOK, products)
}
