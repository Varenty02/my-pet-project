package poststorage

import (
	"MyPetProject/commons"
	postmodel "MyPetProject/module/posts/model"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *postmodel.PostCreate) (*postmodel.PostCreate, error) {
	if err := s.db.Create(&data).Error; err != nil {
		return nil, commons.ErrDB(err)
	}
	return data, nil
}
