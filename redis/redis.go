package redis

import (
	"MyPetProject/commons"
	postmodel "MyPetProject/module/posts/model"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"strings"
	"time"
)

type RedisStore struct {
	redis *redis.Client
}

func NewRedisStore(redis *redis.Client) *RedisStore {
	return &RedisStore{
		redis: redis,
	}
}
func (r *RedisStore) GetPosts(paging *commons.Paging) ([]postmodel.Post, error) {
	client := r.redis
	key := fmt.Sprintf("posts_page:%d,limit:%d", paging.Page, paging.Limit)
	cacheData, err := client.Get(key).Result()
	if err == redis.Nil {
		//k có dữ liệu
		return nil, nil
	} else if err != nil {
		// có lỗi
		return nil, err
	} else {
		// Đã có dữ liệu
		var posts []postmodel.Post
		err := json.Unmarshal([]byte(cacheData), &posts)
		if err != nil {
			return nil, err
		}

		return posts, nil
	}
}
func (r *RedisStore) PushPostsToRedis(posts []postmodel.Post, paging *commons.Paging) error {
	jsonData, err := json.Marshal(posts)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("posts_page:%d,limit:%d", paging.Page, paging.Limit)
	err = r.redis.Set(key, jsonData, time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}
func (r *RedisStore) GetPostById(id int) (*postmodel.Post, error) {
	client := r.redis
	key := fmt.Sprintf("post_id:%d", id)
	cacheData, err := client.Get(key).Result()
	if err == redis.Nil {
		//k có dữ liệu
		return nil, nil
	} else if err != nil {
		// có lỗi
		return nil, err
	} else {
		// Đã có dữ liệu
		var post *postmodel.Post
		err := json.Unmarshal([]byte(cacheData), &post)
		if err != nil {
			return nil, err
		}

		return post, nil
	}
}
func (r *RedisStore) PushPostToRedis(post *postmodel.Post) error {
	jsonData, err := json.Marshal(post)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("post_id:%d", post.Id)
	err = r.redis.Set(key, jsonData, time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}
func (r *RedisStore) RefreshRedis(id int) {
	client := r.redis
	keys, err := client.Keys("*").Result()
	if err != nil {
		log.Println(err)
	}
	for _, key := range keys {
		if strings.HasPrefix(key, "posts") {
			err := client.Del(key).Err()
			if err != nil {
				log.Println(err)
			}
		}
	}
	if id != 0 {
		key := fmt.Sprintf("post_id:%d", id)
		if err := client.Del(key); err != nil {
			log.Println(err)
		}
	}

}
