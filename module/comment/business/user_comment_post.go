package commentbusiness

import (
	commentmodel "MyPetProject/module/comment/model"
	"context"
)

type UserCommentPostStore interface {
	Create(ctx context.Context, data *commentmodel.CommentCreate) error
}
type userCommentPostBiz struct {
	store UserCommentPostStore
}

func NewUserCommentPostBiz(
	store UserCommentPostStore) *userCommentPostBiz {
	return &userCommentPostBiz{

		store: store,
	}
}

func (biz *userCommentPostBiz) LikePost(ctx context.Context, data *commentmodel.CommentCreate) error {
	err := biz.store.Create(ctx, data)
	if err != nil {
		return commentmodel.ErrCannotCommentPost(err)
	}

	return nil
}
