package posttransport

import (
	"MyPetProject/commons"
	"MyPetProject/component/appctx"
	"MyPetProject/module/posts/businesss"
	poststorage "MyPetProject/module/posts/storage"
	"MyPetProject/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeletePost(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(commons.CurrentUser).(commons.Requester)
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(err)
		}
		db := ctx.GetMainConnection()
		store := poststorage.NewSQLStore(db)
		redisStore := redis.NewRedisStore(ctx.GetCache())
		biz := businesss.NewDeletePostBiz(store, redisStore)
		if err := biz.DeletePost(c.Request.Context(), requester.GetUserId(), id); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, commons.SimpleResponse(true))
	}
}
