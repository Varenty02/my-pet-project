package userstorage

import (
	"MyPetProject/commons"
	usermodel "MyPetProject/module/user/model"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *usermodel.UserCreate) (*usermodel.UserCreate, error) {
	if err := s.db.Create(&data).Error; err != nil {
		return nil, commons.ErrDB(err)
	}
	return data, nil
}
