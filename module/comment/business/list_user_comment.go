package commentbusiness

import (
	"MyPetProject/commons"
	commentmodel "MyPetProject/module/comment/model"
	"context"
)

type ListUserCommentPostStore interface {
	GetUsersCommentPost(ctx context.Context,
		condition map[string]interface{},
		filter *commentmodel.Filter,
		paging *commons.Paging,
		moreKey ...string,
	) ([]commentmodel.Comment, error)
}
type listUserCommentPostBiz struct {
	store ListUserCommentPostStore
}

func NewListUserCommentPostBiz(store ListUserCommentPostStore) *listUserCommentPostBiz {
	return &listUserCommentPostBiz{store: store}
}
func (biz *listUserCommentPostBiz) ListUsers(
	ctx context.Context,
	filter *commentmodel.Filter,
	paging *commons.Paging,
) ([]commentmodel.Comment, error) {
	comments, err := biz.store.GetUsersCommentPost(ctx, nil, filter, paging)
	if err != nil {
		return nil, commons.ErrCannotListEntity("COMMENT", err)

	}
	return comments, nil
}
