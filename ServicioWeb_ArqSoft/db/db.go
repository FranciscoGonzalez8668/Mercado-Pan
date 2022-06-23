package db

import (
	"fmt"
	orderClient "pan/clients/order"
	productClient "pan/clients/product"
	userClient "pan/clients/user"
	"pan/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {

	DBName := "pan"
	DBUser := "root"
	DBPass := "root"
	DBHost := "127.0.0.1"

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3307)/"+DBName+"?charset=utf8&parseTime=True")
	/*dsn:="root:root@tcp(127.0.0.1:3307)/pan?charset=utf8mb4&parseTime=True&loc=Local"
	db, err :=gorm.Open(mysql.Open(DBUser+":"+DBPass+"@tcp("+DBHost+":3307)/"+DBName+"?charset=utf8&parseTime=True"))*/

	if err != nil {
		fmt.Println(err)
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	userClient.Db = db
	productClient.Db = db
	orderClient.Db = db

}

func main() {

}

func StartDbEngine() {

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.Order{})
	db.AutoMigrate(&model.Detalle{})
	db.AutoMigrate(&model.Adress{})

	log.Info("Finishing Migration Database Tables")
}
