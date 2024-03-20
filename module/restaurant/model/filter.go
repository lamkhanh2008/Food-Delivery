package restaurantmodel

type Filter struct {
	OwnerId int `json:"owner_id" gorm:"column:owner_id"`
}
