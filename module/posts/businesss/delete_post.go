package businesss

import (
	"MyPetProject/commons"
	postmodel "MyPetProject/module/posts/model"
	"context"
	"errors"
)

type DeletePostStore interface {
	Find(ctx context.Context, conditions map[string]interface{}, moreKey ...string) (*postmodel.Post, error)
	Delete(ctx context.Context, id int) error
	//FindUserPost(ctx context.Context, conditions map[string]interface{}, moreKey ...string) (*int, error)
}
type deletePostBiz struct {
	store      DeletePostStore
	redisStore RefreshRedisStore
}

func NewDeletePostBiz(store DeletePostStore, redisStore RefreshRedisStore) *deletePostBiz {
	return &deletePostBiz{
		store:      store,
		redisStore: redisStore,
	}
}
func (biz *deletePostBiz) DeletePost(ctx context.Context, userId int, id int) error {
	data, err := biz.store.Find(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return commons.ErrCannotNotFound(postmodel.EntityName, err)
	}
	if data.Status == 0 {
		return errors.New("id empty")
	}

	if data.AuthorId != userId {
		return commons.ErrNoPermission(err)
	}
	if err := biz.store.Delete(ctx, id); err != nil {
		return commons.ErrCannotDeleteEntity(postmodel.EntityName, err)
	}
	biz.redisStore.RefreshRedis(data.Id)
	return nil
}
