package commenttransport

import (
	"MyPetProject/commons"
	"MyPetProject/component/appctx"
	commentbusiness "MyPetProject/module/comment/business"
	commentstorage "MyPetProject/module/comment/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserDeleteCommentPost(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		requester := c.MustGet(commons.CurrentUser).(commons.Requester)
		postId, err := strconv.Atoi(id)
		if err != nil {
			panic(commons.ErrInvalidRequest(err))
		}
		store := commentstorage.NewSQLStore(appCtx.GetMainConnection())
		biz := commentbusiness.NewUserDeleteCommentPostBiz(store)
		if err := biz.DeleteCommentPost(c.Request.Context(), requester.GetUserId(), postId); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, commons.SimpleResponse(true))
	}
}
