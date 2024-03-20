package usermodel

import "go-food-delivery/common"

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"-" gorm:"column:email"`
	Salt            string        `json:"-" gorm:"column:salt"`
	LastName        string        `json:"last_name" gorm:"column:last_name"`
	Password        string        `json:"-" gorm:"column:password"`
	FirstName       string        `json:"first_name" gorm:"column:first_name"`
	Phone           string        `json:"phone" gorm:"column:phone"`
	Role            string        `json:"role" gorm:"column:role"`
	Avatar          *common.Image `json:"avatar, omitempty" gorm:"column:avatar;type:json"`
}

// func(u *User) GetUserId() int {
// 	return u.Id
// }

// func (u *User) GetEmail() string {
// 	return u.Email
// }

// func (u *User) GetRole() string {
// 	return u.Role
// }

func (User) TableName() string {
	return "users"
}

func (u *User) Mask(isadmin bool) {
	u.GenUID(common.DbTypeUser)
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email"`
	Password        string `json:"password" gorm:"column:password"`
}
