package poststorage

import (
	"MyPetProject/commons"
	postmodel "MyPetProject/module/posts/model"
	"context"
)

func (s *sqlStore) Update(ctx context.Context, id int, data *postmodel.PostUpdate) (*postmodel.PostUpdate, error) {
	if err := s.db.Where("id=?", id).Updates(&data).Error; err != nil {
		return nil, commons.ErrDB(err)
	}
	return data, nil
}
