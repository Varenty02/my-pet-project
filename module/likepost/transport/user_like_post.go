package liketransport

import (
	"MyPetProject/commons"
	"MyPetProject/component/appctx"
	likebusiness "MyPetProject/module/likepost/business"
	likemodel "MyPetProject/module/likepost/model"
	likestorage "MyPetProject/module/likepost/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserLikePost(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		requester := c.MustGet(commons.CurrentUser).(commons.Requester)
		postId, err := strconv.Atoi(id)
		if err != nil {
			panic(commons.ErrInvalidRequest(err))
		}
		data := likemodel.Like{
			PostId: postId,
			UserId: requester.GetUserId(),
		}
		store := likestorage.NewSQLStore(appCtx.GetMainConnection())
		biz := likebusiness.NewUserLikePostBiz(store)
		if err := biz.LikePost(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, commons.SimpleResponse(true))
	}
}
