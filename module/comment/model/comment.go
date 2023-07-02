package commentmodel

import (
	"MyPetProject/commons"
	"fmt"
	"time"
)

type CommentCreate struct {
	PostId    int        `json:"post_id" gorm:"column:post_id;"`
	UserId    int        `json:"user_id" gorm:"column:user_id;"`
	Content   string     `json:"content" gorm:"content;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}
type Comment struct {
	PostId    int                 `json:"post_id" gorm:"column:post_id;"`
	UserId    int                 `json:"user_id" gorm:"column:user_id;"`
	User      *commons.SimpleUser `json:"user" gorm:"column:user"`
	Content   string              `json:"content" gorm:"content;"`
	CreatedAt *time.Time          `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time          `json:"updated_at" gorm:"column:updated_at;"`
}
type CommentUpdate struct {
	Content *string `json:"content" gorm:"content;"`
}

func (Comment) TableName() string       { return "comments" }
func (CommentCreate) TableName() string { return "comments" }
func (CommentUpdate) TableName() string { return "comments" }
func ErrCannotCommentPost(err error) *commons.AppError {
	return commons.NewCustomError(err, fmt.Sprintf("Cannot comment this post"), fmt.Sprintf("ErrCannotCommentPost"))
}
func ErrCannotDeleteCommentPost(err error) *commons.AppError {
	return commons.NewCustomError(err, fmt.Sprintf("Cannot delete comment this post"), fmt.Sprintf("ErrCannotDeleteCommentPost"))
}
