package model

type User struct {
	Id_user  int    `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(350);not null"`
	Lname    string `gorm:"type:varchar(350);not null"`
	Email    string `gorm:"type:varchar(500);not null;unique"`
	Password string `gorm:"type:varchar(350);not null"`
	//Cart     string `json:"cart"` nose si hacen falta en el modelo ya que on irian en la base de datos
	//Token    string `json:"token"`
}
