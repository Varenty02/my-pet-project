package usertransport

import (
	"MyPetProject/commons"
	"MyPetProject/component/appctx"
	"MyPetProject/component/hasher"
	"MyPetProject/component/tokenprovider/jwt"
	userbusiness "MyPetProject/module/user/business"
	usermodel "MyPetProject/module/user/model"
	userstorage "MyPetProject/module/user/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUser usermodel.UserLogin
		if err := c.ShouldBind(&loginUser); err != nil {
			panic(commons.ErrInvalidRequest(err))
		}
		db := ctx.GetMainConnection()
		store := userstorage.NewSQLStore(db)
		tokenprovider := jwt.NewTokenJWTProvider(ctx.GetSecretKey())
		md5 := hasher.NewHasher()
		biz := userbusiness.NewLoginBiz(store, tokenprovider, md5, 60*60*24*30)
		account, err := biz.Login(c.Request.Context(), &loginUser)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, commons.SimpleResponse(account))
	}
}
