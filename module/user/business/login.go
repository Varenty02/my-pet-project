package userbusiness

import (
	"MyPetProject/commons"
	"MyPetProject/component/tokenprovider"
	usermodel "MyPetProject/module/user/model"
	"context"
)

type FindStore interface {
	Find(ctx context.Context, conditions map[string]interface{}, moreKey ...string) (*usermodel.User, error)
}
type loginBiz struct {
	store         FindStore
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBiz(store FindStore, tokenProvider tokenprovider.Provider, hasher Hasher, expiry int) *loginBiz {
	return &loginBiz{
		store:         store,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}
func (biz *loginBiz) Login(ctx context.Context, data *usermodel.UserLogin, moreKey ...string) (*tokenprovider.Token, error) {
	user, err := biz.store.Find(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}
	passHash := biz.hasher.Hash(data.Password + user.Salt)
	if user.Password != passHash {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}
	payload := tokenprovider.TokenPayload{UserId: user.Id, Role: user.Role}
	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, commons.ErrInternal(err)
	}
	return accessToken, nil
}
