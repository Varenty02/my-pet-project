package commons

import (
	"time"
)

type SimpleUser struct {
	Id        int        `json:"id" gorm:"column:id;"`
	LastName  string     `json:"last_name" gorm:"column:last_name;"`
	FirstName string     `json:"first_name" gorm:"column:first_name;"`
	Role      string     `json:"role,omitempty" gorm:"column:role;"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func (SimpleUser) TableName() string {
	return "users"
}
