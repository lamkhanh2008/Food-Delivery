package restaurantmodel

import (
	"errors"
	"go-food-delivery/common"
	"strings"
)

const RestaurantTableName string = "restaurants"
const EntityName string = "restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	// Id              int    `json:"id" gorm:"column:id"`
	Name  string         `json:"name" gorm:"column:name"`
	Addr  string         `json:"addr" gorm:"column:addr"`
	Type  string         `json:"type" gorm:"column:type"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo"`
	Cover *common.Images `json:"cover" gorm:"column:cover"`
	// Status          int    `json"status" gorm"column:status"`
	// OwnerId int    `json:"owner_id" gorm"column:owner_id"`
}

func (Restaurant) TableName() string { return RestaurantTableName }

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)
}

type RestaurantCreate struct {
	// Id   *int    `json:"id" gorm:"column:id"`
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name"`
	Addr            string         `json:"addr" gorm:"column:addr"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo"`
	Cover           *common.Images `json:"cover" gorm:"column:cover"`
}

func (data *RestaurantCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

func (data RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)
	if data.Name == "" {
		return ErrNameIsEmpty
	}
	return nil
}

// type RestaurantUpdate struct {
// 	Name *string `json:"name" gorm:"column:name"`
// 	Addr *string `json:"addr" gorm:"column:addr"`
// }

// func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

type RestaurantUpdate struct {
	Name *string       `json:"name" gorm:"column:name"`
	Addr *string       `json:"addr" gorm:"column:addr"`
	Logo *common.Image `json:"logo" gorm:"column:logo"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

var (
	ErrNameIsEmpty = errors.New("Name cant  be empy")
)
