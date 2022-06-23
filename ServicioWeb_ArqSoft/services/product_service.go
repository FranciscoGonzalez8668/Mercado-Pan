package services

import (
	productCliente "pan/clients/product"
	"pan/dto"
	"pan/model"
	e "pan/utils/errors"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type productService struct{}

//EL SERVICE SE COMUNICA CON EL MODEL (BUSINESS CLASSES) Y CONTROLLER(DTOS)
type productServiceInterface interface {
	GetProductByName(Key string) (dto.ProductsDto, e.ApiError)
	GetProductByCat(key string) (dto.ProductsDto, e.ApiError)
	GetProductAll() (dto.ProductsDto, e.ApiError)
	GetProductsCart(dto.IdProCart) ([]dto.ProductCart, e.ApiError)
}

var (
	ProductService productServiceInterface
)

func init() {
	ProductService = &productService{}

}

func (s *productService) GetProductByName(key string) (dto.ProductsDto, e.ApiError) {

	var product []model.Product = productCliente.GetProductByName(key)
	var productDto []dto.ProductDto

	var productAuxDto dto.ProductDto

	if product == nil { // si no se encuentran productos NOT FOUND
		productAuxDto.Id_Product = 0
		productDto = append(productDto, productAuxDto)
		log.Debug("f", productDto)
		return productDto, e.NewBadRequestApiError("Product Not Found")
	}

	for i := 0; i < len(product); i++ { // guardar datos en productDto[] para devolver al cliente

		productAuxDto.Id_Product = product[i].Id_Product
		productAuxDto.Category = product[i].Category
		productAuxDto.Descripcion = product[i].Description
		productAuxDto.Name_product = product[i].Name_product
		productAuxDto.Price = product[i].Price
		productAuxDto.Stock = product[i].Stock
		productDto = append(productDto, productAuxDto)
	}
	return productDto, nil

}

func (s *productService) GetProductByCat(cat string) (dto.ProductsDto, e.ApiError) {

	var product []model.Product = productCliente.GetProductByCat(cat)
	var productDto []dto.ProductDto

	var productAuxDto dto.ProductDto
	if product == nil { // si no se encuentran productos NOT FOUND
		productAuxDto.Id_Product = 0
		productDto = append(productDto, productAuxDto)
		log.Debug("f", productDto)

		return productDto, e.NewBadRequestApiError("Product null")
	}

	for i := 0; i < len(product); i++ { // guardar datos en productDto[] para devolver al cliente

		productAuxDto.Id_Product = product[i].Id_Product
		productAuxDto.Category = product[i].Category
		productAuxDto.Descripcion = product[i].Description
		productAuxDto.Name_product = product[i].Name_product
		productAuxDto.Price = product[i].Price
		productAuxDto.Stock = product[i].Stock
		productDto = append(productDto, productAuxDto)
	}

	return productDto, nil

}
func (s *productService) GetProductAll() (dto.ProductsDto, e.ApiError) {
	var productsModel []model.Product = productCliente.GetProductAll()
	var productsDto dto.ProductsDto

	var productsAux dto.ProductDto

	if productsModel[0].Id_Product == 0 {
		productsDto[0].Id_Product = 0
		return productsDto, e.NewBadRequestApiError("No products in DataBase")
	}

	for k := 0; k < len(productsModel); k++ {
		productsAux.Id_Product = productsModel[k].Id_Product
		productsAux.Category = productsModel[k].Category
		productsAux.Descripcion = productsModel[k].Description
		productsAux.Name_product = productsModel[k].Name_product
		productsAux.Price = productsModel[k].Price
		productsAux.Stock = productsModel[k].Stock
		productsDto = append(productsDto, productsAux)
	}
	return productsDto, nil
}
func (s *productService) GetProductsCart(products dto.IdProCart) ([]dto.ProductCart, e.ApiError) {
	var id_products []int
	var intAux int
	for k := 0; k < len(products.Id_products); k++ {
		intAux, _ = strconv.Atoi(products.Id_products[k])
		id_products = append(id_products, intAux)
	}

	var productsModel []model.Product = productCliente.GetProductCart(id_products)
	var productsDto []dto.ProductCart
	var productAux dto.ProductCart
	log.Debug("Productos Carrito", productsModel)
	if productsModel[0].Id_Product == 0 {
		productsDto = nil
		return productsDto, e.NewBadRequestApiError("Bad products request")
	}
	for k := 0; k < len(productsModel); k++ {
		productAux.Id_Product = productsModel[k].Id_Product
		productAux.Name_product = productsModel[k].Name_product
		productAux.Price = productsModel[k].Price
		productsDto = append(productsDto, productAux)
	}
	log.Debug("DTO DETALLE: ", productsDto)
	return productsDto, nil
}
