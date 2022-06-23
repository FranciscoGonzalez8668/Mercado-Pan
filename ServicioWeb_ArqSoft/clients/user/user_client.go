package product

import (
	"pan/dto"
	"pan/model"

	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

//Params: 	dto.LoginDto
//Return:	array model.User
//Proces:	Busca en BD user con el mismo email
//
func GetUserByEmail(LoginDto dto.LoginDto) model.User {
	var user model.User

	/*if LoginDto.Email != "W" || LoginDto.Password != "abc" {
		user.Id_user = 0
		user.Email = "notfound"
		user.Lname = "nil"
		user.Name = "nil"
		user.Password = ""
		return user
	}
	user.Id_user = 1
	user.Email = LoginDto.Email
	user.Lname = "Gonzalez"
	user.Name = "Francisco"
	user.Password = "abc"*/
	Db.First(&user, "email = ? AND password = ?", LoginDto.Email, LoginDto.Password)
	//Db.Where("email= ? AND password= ?", LoginDto.Email, LoginDto.Password).First(&user) //No se creo la base de datos todavia
	return user
}
func GetAdress(idUser int) model.Adress {
	var adress model.Adress

	Db.Where("id_user = ?", idUser).First(&adress)

	return adress

}
