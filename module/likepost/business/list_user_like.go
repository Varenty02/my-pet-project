package likebusiness

import (
	"MyPetProject/commons"
	likemodel "MyPetProject/module/likepost/model"
	"context"
)

type ListUserLikePostStore interface {
	GetUsersLikePost(ctx context.Context,
		condition map[string]interface{},
		filter *likemodel.Filter,
		paging *commons.Paging,
		moreKey ...string,
	) ([]commons.SimpleUser, error)
}
type listUserLikePostBiz struct {
	store ListUserLikePostStore
}

func NewListUserLikePostBiz(store ListUserLikePostStore) *listUserLikePostBiz {
	return &listUserLikePostBiz{store: store}
}
func (biz *listUserLikePostBiz) ListUsers(
	ctx context.Context,
	filter *likemodel.Filter,
	paging *commons.Paging,
) ([]commons.SimpleUser, error) {
	users, err := biz.store.GetUsersLikePost(ctx, nil, filter, paging)
	if err != nil {
		return nil, commons.ErrCannotListEntity("LIKE", err)

	}
	return users, nil
}
