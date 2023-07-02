package likestorage

import (
	"MyPetProject/commons"
	likemodel "MyPetProject/module/likepost/model"
	"context"
	"fmt"
	"log"
	"time"
)

const timeLayout = "2006-01-02T15:04:05.999999"

func (s *sqlStore) GetPostLikes(ctx context.Context, ids []int) (map[int]int, error) {
	result := make(map[int]int)
	type sqlData struct {
		PostId    int `gorm:"column:post_id;"`
		LikeCount int `gorm:"column:count;"`
	}
	var listLike []sqlData
	if err := s.db.Table(likemodel.Like{}.TableName()).
		Select("post_id,count(post_id) as count").
		Where("post_id in (?)", ids).
		Group("post_id").Find(&listLike).Error; err != nil {
		return nil, commons.ErrDB(err)
	}
	for _, item := range listLike {
		result[item.PostId] = item.LikeCount
	}
	return result, nil
}
func (s *sqlStore) GetUsersLikePost(ctx context.Context,
	condition map[string]interface{},
	filter *likemodel.Filter,
	paging *commons.Paging,
	moreKey ...string,
) ([]commons.SimpleUser, error) {
	var result []likemodel.Like
	db := s.db
	db = db.Table(likemodel.Like{}.TableName()).Where(condition)
	if v := filter; v != nil {
		if v.PostId > 0 {
			db = db.Where("post_id=?", v.PostId)
		}
	}
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, commons.ErrDB(err)
	}
	db = db.Preload("User")
	if v := paging.Cursor; v != "" {
		log.Println(paging.Cursor)
		timeCreated, err := time.Parse(timeLayout, v)
		if err != nil {
			return nil, commons.ErrDB(err)
		}
		db = db.Where("created_at<?", timeCreated.Format("2006-01-02 15:04:05"))
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}
	if err := db.
		Limit(paging.Limit).
		Order("created_at desc").
		Find(&result).Error; err != nil {
		return nil, commons.ErrDB(err)
	}
	users := make([]commons.SimpleUser, len(result))
	for i, item := range result {
		result[i].User.CreatedAt = item.CreatedAt
		result[i].User.UpdatedAt = nil
		users[i] = *result[i].User
		if i == len(result)-1 {
			cursorStr := fmt.Sprintf("%v", item.CreatedAt.Format(timeLayout))
			paging.NextCursor = cursorStr
		}
	}
	return users, nil
}
