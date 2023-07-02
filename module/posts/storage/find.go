package poststorage

import (
	"MyPetProject/commons"
	postmodel "MyPetProject/module/posts/model"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) Find(ctx context.Context, conditions map[string]interface{}, moreKey ...string) (*postmodel.Post, error) {
	var data postmodel.Post
	db := s.db.Table(postmodel.Post{}.TableName()).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,last_name, first_name")
	})
	if err := db.Select("id,name,content,category_id,author_id,status").Where(conditions).First(&data).Error; err != nil {
		return nil, commons.ErrDB(err)
	}
	return &data, nil
}

//func (s *sqlStore) FindUserPost(ctx context.Context, conditions map[string]interface{}, moreKey ...string) (*int, error) {
//	var data postmodel.Post
//	if err := s.db.Where(conditions).First(&data).Error; err != nil {
//		return nil, commons.ErrDB(err)
//	}
//	return &data.AuthorId, nil
//}
