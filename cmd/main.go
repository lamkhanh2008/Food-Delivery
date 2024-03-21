package main

import (
	"encoding/json"
	"fmt"
	"go-food-delivery/component/appctx"
	"go-food-delivery/middleware"
	"go-food-delivery/module/restaurant/transport/ginrestaurant"

	"go-food-delivery/module/user/servicesupport/upload/transport/ginupload"
	"go-food-delivery/module/user/transport/ginuser"

	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Configuration struct {
	MYSQL_CONNECT_STRING string `json:"MYSQL_CONNECT_STRING"`
	SECRET_KEY           string `json:"SECRET_KEY"`
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

func main() {
	file, _ := os.Open("./config.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	fmt.Println(configuration)
	if err != nil {
		fmt.Println("error: ", err)
	}

	database, err := gorm.Open(mysql.Open(configuration.MYSQL_CONNECT_STRING), &gorm.Config{})
	db := appctx.NewAppContext(database, configuration.SECRET_KEY)
	if err != nil {
		log.Fatalln(err)
	}
	r := gin.Default()
	r.Static("/static", "./")
	r.Use(middleware.Recover(db))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	restaurants := r.Group("/res")
	restaurants.POST("/upload", ginupload.UploadImage(db))
	restaurants.POST("/up", ginrestaurant.CreateRestaurantdb(db))
	restaurants.GET("/:id", ginrestaurant.FindDataWithCondition(db))
	restaurants.PATCH("/:id", ginrestaurant.UpdateData(db))
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(db))
	restaurants.GET("/list", ginrestaurant.ListDataWithCondition(db))

	users := r.Group("/user")
	users.POST("/up", ginuser.Register(db))
	users.POST("/authenticate", ginuser.Login(db))
	users.GET("/profile", middleware.RequiredAuth(db), ginuser.Profile(db))
	r.Run("localhost:8020")
}
