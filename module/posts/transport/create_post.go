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
)

func CreatePost(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(commons.CurrentUser).(commons.Requester)
		var data = postmodel.PostCreate{
			AuthorId: requester.GetUserId(),
		}
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}
		db := ctx.GetMainConnection()
		store := poststorage.NewSQLStore(db)
		redisStore := redis.NewRedisStore(ctx.GetCache())
		biz := businesss.NewCreatePostBiz(store, redisStore)
		postCreated, err := biz.CreatePost(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, commons.SimpleResponse(postCreated))
	}
}
