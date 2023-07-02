package main

import (
	"MyPetProject/component/appctx"
	"MyPetProject/middleware"
	commenttransport "MyPetProject/module/comment/transport"
	liketransport "MyPetProject/module/likepost/transport"
	posttransport "MyPetProject/module/posts/transport"
	usertransport "MyPetProject/module/user/transport"
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRoute(ctx appctx.AppContext, v1 *gin.RouterGroup) {
	v1 = v1.Group("/", middleware.Recover(ctx))
	//test
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
	//posts
	post := v1.Group("/post", middleware.RequireAuth(ctx), middleware.Authorize(ctx, "author"))
	post.POST("/", posttransport.CreatePost(ctx))
	post.GET("/:id", posttransport.FindPost(ctx))
	post.GET("/", posttransport.ListPost(ctx))
	post.DELETE("/:id", posttransport.DeletePost(ctx))
	post.PUT("/:id", posttransport.UpdatePost(ctx))
	//like
	post.POST("/:id/like", liketransport.UserLikePost(ctx))
	post.DELETE("/:id/dislike", liketransport.UserDislikePost(ctx))
	post.GET("/:id/liked-users", liketransport.ListUser(ctx))
	//comment
	post.POST("/:id/comment", commenttransport.UserCommentPost(ctx))
	post.DELETE("/:id/delete-comment", commenttransport.UserDeleteCommentPost(ctx))
	post.GET("/:id/comment-users", commenttransport.ListUser(ctx))
	post.PUT("/:id/update-comment", commenttransport.UserUpdateCommentPost(ctx))
	//user
	v1.POST("/register", usertransport.Register(ctx))
	v1.POST("/login", usertransport.Login(ctx))
}
