package commentbusiness

import (
	commentmodel "MyPetProject/module/comment/model"
	"context"
)

type UserDeleteCommentPostStore interface {
	Delete(ctx context.Context, userId, postId int) error
}
type userDeleteCommentPostBiz struct {
	store UserDeleteCommentPostStore
}

func NewUserDeleteCommentPostBiz(
	store UserDeleteCommentPostStore) *userDeleteCommentPostBiz {
	return &userDeleteCommentPostBiz{

		store: store,
	}
}

func (biz *userDeleteCommentPostBiz) DeleteCommentPost(ctx context.Context, userId, postId int) error {
	err := biz.store.Delete(ctx, userId, postId)
	if err != nil {
		return commentmodel.ErrCannotCommentPost(err)
	}

	return nil
}
