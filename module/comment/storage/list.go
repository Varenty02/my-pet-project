package commentstorage

import (
	"MyPetProject/commons"
	commentmodel "MyPetProject/module/comment/model"
	"context"
	"fmt"
	"log"
	"time"
)

const timeLayout = "2006-01-02T15:04:05.999999"

func (s *sqlStore) GetUsersCommentPost(ctx context.Context,
	condition map[string]interface{},
	filter *commentmodel.Filter,
	paging *commons.Paging,
	moreKey ...string,
) ([]commentmodel.Comment, error) {
	var result []commentmodel.Comment
	db := s.db
	db = db.Table(commentmodel.Comment{}.TableName()).Where(condition)
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
	for i, item := range result {
		result[i].User.CreatedAt = item.CreatedAt
		result[i].User.UpdatedAt = item.UpdatedAt
		if i == len(result)-1 {
			cursorStr := fmt.Sprintf("%v", item.CreatedAt.Format(timeLayout))
			paging.NextCursor = cursorStr
		}
	}
	return result, nil
}
