package commentstorage

import (
	"MyPetProject/commons"
	commentmodel "MyPetProject/module/comment/model"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *commentmodel.CommentCreate) error {
	db := s.db
	if err := db.Create(&data).Error; err != nil {
		return commons.ErrDB(err)
	}
	return nil
}
