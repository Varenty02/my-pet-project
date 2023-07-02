package usertransport

import (
	"MyPetProject/commons"
	"MyPetProject/component/appctx"
	"MyPetProject/component/hasher"
	"MyPetProject/module/user/business"
	usermodel "MyPetProject/module/user/model"
	userstorage "MyPetProject/module/user/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainConnection()
		store := userstorage.NewSQLStore(db)
		hasher := hasher.NewHasher()
		biz := userbusiness.NewRegisterBiz(store, hasher)
		var user = usermodel.UserCreate{}
		if err := c.ShouldBind(&user); err != nil {
			panic(err)
		}
		result, err := biz.Register(c.Request.Context(), &user)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, commons.SimpleResponse(result))
	}
}
