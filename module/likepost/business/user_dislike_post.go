package likebusiness

import (
	likemodel "MyPetProject/module/likepost/model"
	"context"
)

type UserDislikePostStore interface {
	Delete(ctx context.Context, userId, postId int) error
}
type userDislikePostBiz struct {
	store UserDislikePostStore
}

func NewUserDislikePostBiz(
	store UserDislikePostStore) *userDislikePostBiz {
	return &userDislikePostBiz{

		store: store,
	}
}

func (biz *userDislikePostBiz) DislikePost(ctx context.Context, userId, postId int) error {
	err := biz.store.Delete(ctx, userId, postId)
	if err != nil {
		return likemodel.ErrCannotDislikePost(err)
	}

	return nil
}
