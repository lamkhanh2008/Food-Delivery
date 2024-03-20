package main

import (
	"encoding/json"
	"fmt"
	"go-food-delivery/component/appctx"
	"go-food-delivery/middleware"
	"go-food-delivery/module/restaurant/transport/ginrestaurant"
	"go-food-delivery/module/upload/transport/ginupload"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Configuration struct {
	MYSQL_CONNECT_STRING string
}

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
	Addr string `json:"addr" gorm:"column:addr"`
}

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name"`
	Addr *string `json:"addr" gotm:"column:addr"`
}

func (Restaurant) TableName() string { return "restaurants" }

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

// func main() {
// 	file, _ := os.Open("./cmd/config.json")
// 	decoder := json.NewDecoder(file)
// 	configuration := Configuration{}
// 	err := decoder.Decode(&configuration)
// 	if err != nil {
// 		fmt.Println("error: ", err)
// 	}
// 	// dsn := os.Getenv("MYSQL_CONNECT_STRING")
// 	// dsn := "root:2008@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"

// 	db, err := gorm.Open(mysql.Open(configuration.MYSQL_CONNECT_STRING), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	newRestaurant := Restaurant{Name: "Tani", Addr: "9 Pham van hai"}
// 	if err := db.Create(&newRestaurant).Error; err != nil {
// 		log.Println(err)
// 	}
// 	log.Println("id", newRestaurant.Id)

// 	var myRestaurant Restaurant
// 	if err := db.Where("id = ?", 3).First(&myRestaurant).Error; err != nil {
// 		log.Println(err)
// 	}

// 	log.Println(myRestaurant)
// 	newname := "lam"
// 	updateRes := RestaurantUpdate{
// 		Name: &newname,
// 	}

// 	if err := db.Where("id = ?", 3).Updates(&updateRes).Error; err != nil {
// 		log.Println(err)
// 	}

// 	log.Println(myRestaurant)
// }

func main() {
	file, _ := os.Open("./../cmd/config.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error: ", err)
	}
	// dsn := os.Getenv("MYSQL_CONNECT_STRING")
	// dsn := "root:2008@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(configuration.MYSQL_CONNECT_STRING), &gorm.Config{})
	db := appctx.NewAppContext(database)
	if err != nil {
		log.Fatalln(err)
	}
	r := gin.Default()
	r.Static("/static", "./static")
	r.Use(middleware.Recover(db))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	restaurants := r.Group("/v1")
	restaurants.POST("/upload", ginupload.UploadImage(db))
	restaurants.POST("/restaurants", ginrestaurant.CreateRestaurantdb(db))
	restaurants.GET("/:id", ginrestaurant.FindDataWithCondition(db))
	restaurants.PATCH("/:id", ginrestaurant.UpdateData(db))
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(db))
	restaurants.GET("/list", ginrestaurant.ListDataWithCondition(db))
	r.Run()
}
