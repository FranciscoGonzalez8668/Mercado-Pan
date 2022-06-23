package services

import (
	"fmt"
	orderCliente "pan/clients/order"
	productClient "pan/clients/product"
	userClient "pan/clients/user"
	"pan/dto"
	"pan/model"
	e "pan/utils/errors"

	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

type orderService struct{}

type orderServiceInterface interface {
	OrderHistory(string) (dto.OrdersHistoryDto, e.ApiError)
	NewOrder(NewOrderDto dto.NewOrderDto) (dto.OrderDto, e.ApiError)
}

var (
	OrderService orderServiceInterface
)

func init() {
	OrderService = &orderService{}
}
func getUserFromToken(receivedToken string) (int, error) {
	log.Debug("AK")
	token, err := jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Debug("1")
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		log.Debug("2")
		return JwtKey, nil
	})
	log.Debug("TOKENOTKEONT", token)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usId := (claims["Id_User"].(float64))
		usIntId := int(usId)
		return usIntId, nil
	} else {
		log.Debug("4")
		return -1, err
	}

}

func (s *orderService) OrderHistory(Token string) (dto.OrdersHistoryDto, e.ApiError) {
	var orders []model.Order
	var detalles []model.Detalle
	var addres model.Adress
	idUser, err := getUserFromToken(Token)
	//Funcion de verificar token

	var orderAux dto.OrderDto
	var detalleAUX dto.DetalleDto
	var productAUX dto.ProductDto
	var product model.Product
	var order dto.OrdersHistoryDto
	var addresAux dto.AdressDto
	orders = orderCliente.GetOrders(idUser)
	addres = userClient.GetAdress(idUser)

	addresAux.City = addres.City
	addresAux.Id_adress = addres.Id_adress
	addresAux.Number = addres.Number
	addresAux.Street_Name = addres.Street_Name
	addresAux.Neighborhood = addres.Neighborhood
	if err != nil {
		return order, e.NewBadRequestApiError("Token Invalida")
	}

	// seteo de []order.orderhistorydto lugar por lugar
	for j := 0; j < len(orders); j++ {
		detalles = orderCliente.GetDetalles(orders[j].Id_Order)
		orderAux.Det_order = nil
		for k := 0; k < len(detalles); k++ {
			// seteo del producto dentro del detalleAUX a insertar
			product = productClient.GetProductById(detalles[k].Id_Product)
			productAUX.Id_Product = product.Id_Product
			productAUX.Stock = product.Stock
			productAUX.Category = product.Category
			productAUX.Price = product.Price
			productAUX.Name_product = product.Name_product
			detalleAUX.Product = productAUX

			//setup detalle a insertar
			detalleAUX.Cant_det = detalles[k].Cantidad
			detalleAUX.Price_det = detalles[k].Precio_Moment
			orderAux.Det_order = append(orderAux.Det_order, detalleAUX)

		}
		orderAux.Id_order = orders[j].Id_Order
		orderAux.Total = orders[j].Total
		orderAux.Adress = addresAux
		// setup historial con auxiliares
		order.Order = append(order.Order, orderAux)

	}
	order.Id_User = idUser
	return order, nil

}
func (s *orderService) NewOrder(NewOrderDto dto.NewOrderDto) (dto.OrderDto, e.ApiError) {
	var order model.Order
	var detalles []model.Detalle
	var detalleAUX model.Detalle
	id_user, err := getUserFromToken(NewOrderDto.Token)
	var total float32 = 0

	//falta escribir todos los datos en el producto para mostrar bien la order
	for j := 0; j < len(NewOrderDto.Detalles); j++ {
		detalleAUX.Cantidad = NewOrderDto.Detalles[j].Cant_det
		detalleAUX.Id_Product = NewOrderDto.Detalles[j].Id_product
		detalleAUX.Precio_Moment = float32(NewOrderDto.Detalles[j].Precio_Moment)
		detalles = append(detalles, detalleAUX)
		total = total + float32(NewOrderDto.Detalles[j].Precio_Moment)*float32(NewOrderDto.Detalles[j].Cant_det)
	}
	order.Total = total
	order.Id_User = id_user
	order = orderCliente.NewOrder(order, detalles)
	log.Debug("newOrderClient Complete")
	var orderDto dto.OrderDto
	if err != nil {
		return orderDto, e.NewBadRequestApiError("Token Invalida")
	}
	if order.Id_Order == 0 {
		return orderDto, e.NewBadRequestApiError("No stock enought")
	}
	//hay una diferencia entre los dto y los modelos ya que los dto tienen datos de otras tablas de la db

	orderDto.Id_order = order.Id_Order
	orderDto.Total = order.Total

	return orderDto, nil
}

/*
 ACA VA MI INTENTO
func (s *orderService) OrderHistoy(idUser int) (dto.OrderDetDto, e.ApiError){


	var orders []model.OrderDet
	orders = orderCliente.GetHistory (idUser)
	var orders dto.OrderDetDto

	for(i=0, i</* cantidad de orderdenes , i++){

		ordersDto[i] = [i]model.OrderDet



	}

	return ordersDto



}


*/
