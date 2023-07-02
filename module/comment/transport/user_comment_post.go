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

func UserCommentPost(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		requester := c.MustGet(commons.CurrentUser).(commons.Requester)
		postId, err := strconv.Atoi(id)
		if err != nil {
			panic(commons.ErrInvalidRequest(err))
		}
		data := commentmodel.CommentCreate{
			PostId: postId,
			UserId: requester.GetUserId(),
		}
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}
		store := commentstorage.NewSQLStore(appCtx.GetMainConnection())
		biz := commentbusiness.NewUserCommentPostBiz(store)
		if err := biz.LikePost(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, commons.SimpleResponse(true))
	}
}
