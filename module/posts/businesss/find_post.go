package businesss

import (
	"MyPetProject/commons"
	postmodel "MyPetProject/module/posts/model"
	"context"
	"errors"
)

type FindPostStore interface {
	Find(ctx context.Context, conditions map[string]interface{}, moreKey ...string) (*postmodel.Post, error)
}
type GetPostRedisStore interface {
	GetPostById(id int) (*postmodel.Post, error)
	PushPostToRedis(post *postmodel.Post) error
}
type findPostBiz struct {
	store      FindPostStore
	redisStore GetPostRedisStore
}

func NewFindPostBiz(store FindPostStore, redisStore GetPostRedisStore) *findPostBiz {
	return &findPostBiz{
		store:      store,
		redisStore: redisStore,
	}
}
func (biz *findPostBiz) FindPost(ctx context.Context, conditions map[string]interface{}, moreKey ...string) (*postmodel.Post, error) {
	id := conditions["id"].(int)
	redisData, redisErr := biz.redisStore.GetPostById(id)
	if redisErr == nil && redisData == nil {
		data, err := biz.fetchPostData(ctx, conditions)
		if err != nil {
			return nil, commons.ErrCannotNotFound(postmodel.EntityName, err)
		}
		biz.redisStore.PushPostToRedis(data)
		return data, nil
	}
	if redisErr != nil && redisData == nil {
		data, err := biz.fetchPostData(ctx, conditions)
		if err != nil {
			return nil, commons.ErrCannotNotFound(postmodel.EntityName, err)
		}
		return data, nil
	}
	return redisData, nil
}
func (biz *findPostBiz) fetchPostData(ctx context.Context, conditions map[string]interface{}) (*postmodel.Post, error) {
	data, err := biz.store.Find(ctx, conditions)
	if err != nil {
		return nil, commons.ErrCannotNotFound(postmodel.EntityName, err)
	}
	if data.Status == 0 {
		return nil, errors.New("id empty")
	}
	return data, nil
}
