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

func FindPost(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainConnection()
		store := poststorage.NewSQLStore(db)
		redisStore := redis.NewRedisStore(ctx.GetCache())
		biz := businesss.NewFindPostBiz(store, redisStore)
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(err)
		}
		data, err := biz.FindPost(c.Request.Context(), map[string]interface{}{"id": id})
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, commons.SimpleResponse(data))
	}
}
