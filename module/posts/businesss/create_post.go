package businesss

import (
	"MyPetProject/commons"
	postmodel "MyPetProject/module/posts/model"
	"context"
)

type CreatePostStore interface {
	Create(ctx context.Context, data *postmodel.PostCreate) (*postmodel.PostCreate, error)
}
type RefreshRedisStore interface {
	RefreshRedis(id int)
}
type createPostBiz struct {
	store      CreatePostStore
	redisStore RefreshRedisStore
}

func NewCreatePostBiz(store CreatePostStore, redisStore RefreshRedisStore) *createPostBiz {
	return &createPostBiz{
		store:      store,
		redisStore: redisStore,
	}
}
func (biz *createPostBiz) CreatePost(ctx context.Context, data *postmodel.PostCreate) (*postmodel.PostCreate, error) {

	data, err := biz.store.Create(ctx, data)
	if err != nil {
		return nil, commons.ErrCannotCreateEntity(postmodel.EntityName, err)
	}
	biz.redisStore.RefreshRedis(data.Id)
	return data, nil
}
