package posttransport

import (
	"MyPetProject/commons"
	"MyPetProject/component/appctx"
	likestorage "MyPetProject/module/likepost/storage"
	"MyPetProject/module/posts/businesss"
	postmodel "MyPetProject/module/posts/model"
	poststorage "MyPetProject/module/posts/storage"
	"MyPetProject/redis"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListPost(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainConnection()
		store := poststorage.NewSQLStore(db)
		likeStore := likestorage.NewSQLStore(db)
		redisStore := redis.NewRedisStore(ctx.GetCache())
		biz := businesss.NewListPostBiz(store, likeStore, redisStore)
		var filter = postmodel.Filter{}
		filter.Status = []int{1}
		var paging = commons.Paging{}
		paging.Fulfil()
		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}
		data, err := biz.ListPost(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, commons.SimpleResponse(data))
	}
}
