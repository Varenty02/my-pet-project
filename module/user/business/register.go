package userbusiness

import (
	"MyPetProject/commons"
	usermodel "MyPetProject/module/user/model"
	"context"
	"errors"
)

type RegisterStore interface {
	Find(ctx context.Context, conditions map[string]interface{}, moreKey ...string) (*usermodel.User, error)
	Create(ctx context.Context, data *usermodel.UserCreate) (*usermodel.UserCreate, error)
}
type Hasher interface {
	Hash(data string) string
}
type registerBiz struct {
	store  RegisterStore
	hasher Hasher
}

func NewRegisterBiz(store RegisterStore, hasher Hasher) *registerBiz {
	return &registerBiz{
		store:  store,
		hasher: hasher,
	}
}
func (biz *registerBiz) Register(ctx context.Context, data *usermodel.UserCreate) (*usermodel.UserCreate, error) {
	user, _ := biz.store.Find(ctx, map[string]interface{}{"email": data.Email})
	if user != nil {
		return nil, errors.New("This user has been exist")
	}
	salt := commons.GenSalt(30)
	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "author"
	data, err := biz.store.Create(ctx, data)
	if err != nil {
		return nil, commons.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	return data, nil
}
