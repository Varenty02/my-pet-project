package middleware

import (
	"MyPetProject/commons"
	"MyPetProject/component/appctx"
	"MyPetProject/component/tokenprovider/jwt"
	userstorage "MyPetProject/module/user/storage"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

func ErrWrongAuthHeader(err error) *commons.AppError {
	return commons.NewCustomError(
		err,
		"wrong authen header",
		"ErrWrongAuthHeader",
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}
	return parts[1], nil
}
func RequireAuth(ctx appctx.AppContext) func(c *gin.Context) {
	tokenProvider := jwt.NewTokenJWTProvider(ctx.GetSecretKey())
	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}
		db := ctx.GetMainConnection()
		store := userstorage.NewSQLStore(db)
		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}
		user, err := store.Find(c.Request.Context(), map[string]interface{}{"id": payload.UserId})
		if err != nil {
			panic(err)
		}
		if user.Status == 0 {
			panic(commons.ErrNoPermission(errors.New("user has been deleted or banned")))
		}
		c.Set(commons.CurrentUser, user)
		c.Next()
	}
}
