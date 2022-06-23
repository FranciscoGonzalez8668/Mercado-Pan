package model

type Adress struct {
	Id_adress    int    `gorm:"primaryKey"`
	Id_user      int    `gorm:"primaryKey"`
	Street_Name  string `gorm:"type:varchar(350);not null"`
	Number       int    `gorm:"type:int"`
	Neighborhood string `gorm:"type:varchar(350)"`
	City         string `gorm:"type:varchar(350);not null"`
	// hace falta poner el id usuaurio ya que se relacionan?
}

type Adresss []Adress
