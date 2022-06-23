package dto

type LoginDto struct {
	Email    string `jsno:"email"`
	Password string `json:"password"`
}

type LoginOkDto struct {
	Token    string `json:"token"`
	Id_login string `json:"id_login"`
}
