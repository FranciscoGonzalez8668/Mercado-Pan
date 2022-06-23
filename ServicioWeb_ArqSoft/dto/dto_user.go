package dto

type UserDto struct {
	Id_user  int    `json:"id_user"`
	Name     string `json:"name"`
	Lname    string `json:"lname"`
	Email    string `jsno:"email"`
	Password string `json:"password"`
	//Cart     string `json:"cart"`
	//Token    string `json:"token"`
}

type UsersDto []UserDto
