package commentbusiness

import (
	commentmodel "MyPetProject/module/comment/model"
	"context"
)

type UserUpdateCommentPostStore interface {
	Update(ctx context.Context, userId, postId int, data *commentmodel.CommentUpdate) (*commentmodel.CommentUpdate, error)
}
type userUpdateCommentPostBiz struct {
	store UserUpdateCommentPostStore
}

func NewUserUpdateCommentPostBiz(
	store UserUpdateCommentPostStore) *userUpdateCommentPostBiz {
	return &userUpdateCommentPostBiz{

		store: store,
	}
}

func (biz *userUpdateCommentPostBiz) UpdateCommentPost(ctx context.Context, userId, postId int, data *commentmodel.CommentUpdate) (*commentmodel.CommentUpdate, error) {
	dataUpdated, err := biz.store.Update(ctx, userId, postId, data)
	if err != nil {
		return nil, commentmodel.ErrCannotCommentPost(err)
	}

	return dataUpdated, nil
}
