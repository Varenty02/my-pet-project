package businesss

import (
	"MyPetProject/commons"
	postmodel "MyPetProject/module/posts/model"
	"context"
	"errors"
)

type UpdatePostStore interface {
	Find(ctx context.Context, conditions map[string]interface{}, moreKey ...string) (*postmodel.Post, error)
	Update(ctx context.Context, id int, data *postmodel.PostUpdate) (*postmodel.PostUpdate, error)
	//FindUserPost(ctx context.Context, conditions map[string]interface{}, moreKey ...string) (*int, error)
}
type updatePostBiz struct {
	store      UpdatePostStore
	redisStore RefreshRedisStore
}

func NewUpdatePostBiz(store UpdatePostStore, redisStore RefreshRedisStore) *updatePostBiz {
	return &updatePostBiz{
		store:      store,
		redisStore: redisStore,
	}
}
func (biz *updatePostBiz) UpdatePost(ctx context.Context, userId, id int, data *postmodel.PostUpdate) (*postmodel.PostUpdate, error) {
	dataExist, err := biz.store.Find(ctx, map[string]interface{}{"id": 1})
	if err != nil {
		return nil, commons.ErrCannotNotFound(postmodel.EntityName, err)
	}
	if dataExist.Status == 0 {
		return nil, errors.New("id empty")
	}
	if dataExist.AuthorId != userId {
		return nil, commons.ErrNoPermission(err)
	}
	newData, err := biz.store.Update(ctx, id, data)
	if err != nil {
		return nil, commons.ErrCannotUpdateEntity(postmodel.EntityName, err)
	}
	biz.redisStore.RefreshRedis(id)
	return newData, nil
}
