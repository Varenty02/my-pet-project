package businesss

import (
	"MyPetProject/commons"
	postmodel "MyPetProject/module/posts/model"
	"context"
	"log"
)

type ListPostStore interface {
	List(ctx context.Context, filter *postmodel.Filter, paging *commons.Paging, moreKey ...string) ([]postmodel.Post, error)
}
type LikePostStore interface {
	GetPostLikes(ctx context.Context, ids []int) (map[int]int, error)
}
type GetPostsRedisStore interface {
	GetPosts(paging *commons.Paging) ([]postmodel.Post, error)
	PushPostsToRedis(posts []postmodel.Post, paging *commons.Paging) error
}
type listPostBiz struct {
	store      ListPostStore
	likeStore  LikePostStore
	redisStore GetPostsRedisStore
}

func NewListPostBiz(store ListPostStore, likeStore LikePostStore, redisStore GetPostsRedisStore) *listPostBiz {
	return &listPostBiz{
		store:      store,
		likeStore:  likeStore,
		redisStore: redisStore,
	}
}
func (biz *listPostBiz) ListPost(ctx context.Context, filter *postmodel.Filter, paging *commons.Paging, moreKey ...string) ([]postmodel.Post, error) {
	redisData, redisErr := biz.redisStore.GetPosts(paging)

	if redisErr == nil && redisData == nil {
		data, err := biz.fetchPostData(ctx, filter, paging)
		if err != nil {
			return nil, commons.ErrCannotListEntity(postmodel.EntityName, err)
		}

		// Update LikeCount for each post
		err = biz.updatePostLikeCount(ctx, data)
		if err != nil {
			log.Println(err)
			return data, err
		}

		// Push data to Redis
		biz.redisStore.PushPostsToRedis(data, paging)

		return data, nil
	}

	if redisErr != nil && redisData == nil {
		data, err := biz.fetchPostData(ctx, filter, paging)
		if err != nil {
			return nil, commons.ErrCannotListEntity(postmodel.EntityName, err)
		}

		// Update LikeCount for each post
		err = biz.updatePostLikeCount(ctx, data)
		if err != nil {
			log.Println(err)
			return data, err
		}

		return data, nil
	}

	return redisData, nil
}

func (biz *listPostBiz) fetchPostData(ctx context.Context, filter *postmodel.Filter, paging *commons.Paging) ([]postmodel.Post, error) {
	data, err := biz.store.List(ctx, filter, paging, "User")
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (biz *listPostBiz) updatePostLikeCount(ctx context.Context, data []postmodel.Post) error {
	ids := make([]int, len(data))
	for i := range data {
		ids[i] = data[i].Id
	}

	likeMap, err := biz.likeStore.GetPostLikes(ctx, ids)
	if err != nil {
		return err
	}

	for i, item := range data {
		data[i].LikeCount = likeMap[item.Id]
	}

	return nil
}
