package userstorage

import (
	"MyPetProject/commons"
	usermodel "MyPetProject/module/user/model"
	"context"
)

func (s *sqlStore) Find(ctx context.Context, conditions map[string]interface{}, moreKey ...string) (*usermodel.User, error) {
	var user usermodel.User
	if err := s.db.Where(conditions).First(&user).Error; err != nil {
		return nil, commons.ErrDB(err)
	}
	return &user, nil
}
