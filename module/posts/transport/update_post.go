package posttransport

import (
	"MyPetProject/commons"
	"MyPetProject/component/appctx"
	"MyPetProject/module/posts/businesss"
	postmodel "MyPetProject/module/posts/model"
	poststorage "MyPetProject/module/posts/storage"
	"MyPetProject/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdatePost(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(commons.CurrentUser).(commons.Requester)
		userId := requester.GetUserId()
		var data = postmodel.PostUpdate{
			AuthorId: &userId,
		}
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(err)
		}
		db := ctx.GetMainConnection()
		store := poststorage.NewSQLStore(db)
		redisStore := redis.NewRedisStore(ctx.GetCache())
		biz := businesss.NewUpdatePostBiz(store, redisStore)
		newData, err := biz.UpdatePost(c.Request.Context(), userId, id, &data)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, commons.SimpleResponse(newData))
	}
}
