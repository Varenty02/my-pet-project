package likestorage

import (
	"MyPetProject/commons"
	likemodel "MyPetProject/module/likepost/model"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *likemodel.Like) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return commons.ErrDB(err)
	}
	return nil
}
