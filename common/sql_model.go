package common

import (
	"fmt"
	"time"
)

type SQLModel struct {
	Id        int        `json:"-" gorm:"column:id"`
	FakeId    *UID       `json:"id" gorm:"-"`
	Status    int        `json:"status" gorm:"column:status";default:1,"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdateAt  *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

const (
	DbTypeRestaurant = 1
	DbTypeUser       = 1
)

func (m *SQLModel) GenUID(dbType int) {
	uid := NewUID(uint32(m.Id), dbType, DbTypeUser)
	m.FakeId = &uid
	fmt.Println(m.FakeId)
}
