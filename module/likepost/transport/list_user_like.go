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

func ListUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		postId, err := strconv.Atoi(id)
		if err != nil {
			panic(commons.ErrInvalidRequest(err))
		}
		filter := likemodel.Filter{
			PostId: postId,
		}
		var paging commons.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(commons.ErrInvalidRequest(err))
		}
		paging.Fulfil()
		store := likestorage.NewSQLStore(appCtx.GetMainConnection())
		biz := likebusiness.NewListUserLikePostBiz(store)
		result, err := biz.ListUsers(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, commons.NewSuccessResponse(result, paging, filter))
	}
}
