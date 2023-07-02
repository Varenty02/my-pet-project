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

func UserUpdateCommentPost(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		requester := c.MustGet(commons.CurrentUser).(commons.Requester)
		postId, err := strconv.Atoi(id)
		if err != nil {
			panic(commons.ErrInvalidRequest(err))
		}
		var data commentmodel.CommentUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}
		store := commentstorage.NewSQLStore(appCtx.GetMainConnection())
		biz := commentbusiness.NewUserUpdateCommentPostBiz(store)
		if _, err := biz.UpdateCommentPost(c.Request.Context(), requester.GetUserId(), postId, &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, commons.SimpleResponse(true))
	}
}
