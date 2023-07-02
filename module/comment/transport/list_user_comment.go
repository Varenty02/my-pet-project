package commenttransport

import (
	"MyPetProject/commons"
	"MyPetProject/component/appctx"
	commentbusiness "MyPetProject/module/comment/business"
	commentmodel "MyPetProject/module/comment/model"
	commentstorage "MyPetProject/module/comment/storage"
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
		filter := commentmodel.Filter{
			PostId: postId,
		}
		var paging commons.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(commons.ErrInvalidRequest(err))
		}
		paging.Fulfil()
		store := commentstorage.NewSQLStore(appCtx.GetMainConnection())
		biz := commentbusiness.NewListUserCommentPostBiz(store)
		result, err := biz.ListUsers(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, commons.NewSuccessResponse(result, paging, filter))
	}
}
