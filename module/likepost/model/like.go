package likemodel

import (
	"MyPetProject/commons"
	"fmt"
	"time"
)

type Like struct {
	PostId    int                 `json:"post_id" gorm:"column:post_id;"`
	UserId    int                 `json:"user_id" gorm:"column:user_id;"`
	CreatedAt *time.Time          `json:"created_at" gorm:"column:created_at;"`
	User      *commons.SimpleUser `json:"user"`
}

func (Like) TableName() string { return "likes" }
func ErrCannotLikePost(err error) *commons.AppError {
	return commons.NewCustomError(err, fmt.Sprintf("Cannot like this post"), fmt.Sprintf("ErrCannotLikePost"))
}
func ErrCannotDislikePost(err error) *commons.AppError {
	return commons.NewCustomError(err, fmt.Sprintf("Cannot dislike this post"), fmt.Sprintf("ErrCannotDislikePost"))
}
