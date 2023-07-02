package liketransport

import (
	"MyPetProject/commons"
	"MyPetProject/component/appctx"
	likebusiness "MyPetProject/module/likepost/business"
	likestorage "MyPetProject/module/likepost/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserDislikePost(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		requester := c.MustGet(commons.CurrentUser).(commons.Requester)
		postId, err := strconv.Atoi(id)
		if err != nil {
			panic(commons.ErrInvalidRequest(err))
		}
		store := likestorage.NewSQLStore(appCtx.GetMainConnection())
		biz := likebusiness.NewUserDislikePostBiz(store)
		if err := biz.DislikePost(c.Request.Context(), requester.GetUserId(), postId); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, commons.SimpleResponse(true))
	}
}
