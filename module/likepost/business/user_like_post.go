package likebusiness

import (
	likemodel "MyPetProject/module/likepost/model"
	"context"
)

type UserLikePostStore interface {
	Create(ctx context.Context, data *likemodel.Like) error
}
type userLikePostBiz struct {
	store UserLikePostStore
	//incStore IncLikeCountResStore
}

func NewUserLikePostBiz(
	store UserLikePostStore) *userLikePostBiz {
	return &userLikePostBiz{

		store: store,
		//incStore: incStore,
	}
}

func (biz *userLikePostBiz) LikePost(ctx context.Context, data *likemodel.Like) error {
	err := biz.store.Create(ctx, data)
	if err != nil {
		return likemodel.ErrCannotDislikePost(err)
	}

	return nil
}
