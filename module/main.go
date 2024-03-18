package main

import (
	"encoding/json"
	"fmt"
	"go-food-delivery/component/appctx"
	"go-food-delivery/module/restaurant/transport/ginrestaurant"
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
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	restaurants := r.Group("/v1")
	restaurants.POST("/restaurants", ginrestaurant.CreateRestaurantdb(db))
	restaurants.GET("/:id", ginrestaurant.FindDataWithCondition(db))
	// restaurants.POST("/restaurants", func(c *gin.Context) {
	// 	var data Restaurant

	// 	if err := c.ShouldBind(&data); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 		return
	// 	}
	// 	db.Create(&data)
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"data": data,
	// 	})

	// 	//   r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// })

	// restaurants.GET("/restaurants/:id", func(c *gin.Context) {
	// 	var data Restaurant
	// 	id, err := strconv.Atoi(c.Param("id"))
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 	}
	// 	db.Where("id = ?", id).First(&data)
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"data": data,
	// 	})
	// })

	// restaurants.GET("/list_res", func(c *gin.Context) {
	// 	var data []Restaurant
	// 	type Paging struct {
	// 		Page  int `json:"page" form:"page"`
	// 		Limit int `json:"limit" form:"limit"`
	// 	}

	// 	var pagingData Paging

	// 	if err := c.ShouldBind(&pagingData); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": err.Error(),
	// 		})

	// 		return
	// 	}

	// 	if pagingData.Page <= 0 {
	// 		pagingData.Page = 1
	// 	}

	// 	if pagingData.Limit <= 0 {
	// 		pagingData.Limit = 5
	// 	}
	// 	db.Offset((pagingData.Page - 1) * pagingData.Limit).Order("id desc").Limit(pagingData.Limit).Find(&data)
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"data": data,
	// 	})
	// })
	restaurants.PATCH("/:id", ginrestaurant.UpdateData(db))
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(db))

	// restaurants.PATCH("/:id", func(c *gin.Context) {
	// 	id, err := strconv.Atoi(c.Param("id"))
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 		return
	// 	}
	// 	var data RestaurantUpdate

	// 	if err := db.Where("id = ?", id).Error; err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 		return
	// 	}

	// 	db.Table(Restaurant{}.TableName()).Where("id = ?", id).Delete(nil)
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"data": 1,
	// 	})

	// 	if err := c.ShouldBind(&data); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 		return
	// 	}

	// 	db.Where("id = ?", id).Updates(&data)
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"data": data,
	// 	})
	// })

	r.Run()
}
