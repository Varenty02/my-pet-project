package postmodel

import (
	"MyPetProject/commons"
	"time"
)

var EntityName = "POST"

type Post struct {
	Id         int                 `json:"id" gorm:"id"`
	Name       string              `json:"name" gorm:"name"`
	Content    string              `json:"content,omitempty" gorm:"content"`
	CategoryId int                 `json:"category_id,omitempty" gorm:"category_id"`
	AuthorId   int                 `json:"author_id,omitempty" gorm:"author_id"`
	User       *commons.SimpleUser `json:"author,omitempty" gorm:"foreignKey:AuthorId"`
	Status     int                 `json:"status,omitempty" gorm:"status;default:1"`
	CreatedAt  *time.Time          `json:"created_at,omitempty" gorm:"created_at"`
	UpdatedAt  *time.Time          `json:"updated_at,omitempty" gorm:"updated_at"`
	LikeCount  int                 `json:"like_count,omitempty" gorm:"-"`
}
type PostCreate struct {
	Id         int    `json:"id" gorm:"id"`
	Name       string `json:"name" gorm:"name"`
	Content    string `json:"content" gorm:"content"`
	CategoryId int    `json:"category_id" gorm:"category_id"`
	AuthorId   int    `json:"author_id" gorm:"author_id"`
	Status     int    `json:"status" gorm:"status;default:1"`
}
type PostUpdate struct {
	Name       *string `json:"name" gorm:"name"`
	Content    *string `json:"content" gorm:"content"`
	CategoryId *int    `json:"category_id" gorm:"category_id"`
	AuthorId   *int    `json:"author_id" gorm:"-"`
}

func (Post) TableName() string       { return "posts" }
func (PostUpdate) TableName() string { return "posts" }
func (PostCreate) TableName() string { return "posts" }
