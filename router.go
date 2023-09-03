package main

import (
	"net/http"

	"github.com/RaymondCode/simple-demo/controller"
	"github.com/RaymondCode/simple-demo/middleware"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/service"

	"github.com/RaymondCode/simple-demo/config"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")
	r.LoadHTMLGlob("templates/*")

	// home page
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	userRepo := repository.NewUserRepository(config.DbCon())
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService) // TODO: move out from router.go

	publicRouter := r.Group("/douyin")
	{
		publicRouter.POST("/user/register/", userController.Register)
		publicRouter.POST("/user/login/", userController.Login)
		publicRouter.GET("/feed/", controller.Feed)
	}

	// need authentication
	authRouter := r.Group("/douyin")
	authRouter.Use(middleware.TokenAuth())
	{
		// basic apis
		authRouter.GET("/user/", userController.UserInfo)
		authRouter.POST("/publish/action/", controller.Publish)
		authRouter.GET("/publish/list/", controller.PublishList)
		// extra apis - I
		authRouter.POST("/favorite/action/", controller.FavoriteAction)
		authRouter.GET("/favorite/list/", controller.FavoriteList)
		authRouter.POST("/comment/action/", controller.CommentAction)
		authRouter.GET("/comment/list/", controller.CommentList)
		// extra apis - II
		// authRouter.POST("/relation/action/", controller.RelationAction)
		// authRouter.GET("/relation/follow/list/", controller.FollowList)
		// authRouter.GET("/relation/follower/list/", controller.FollowerList)
		// authRouter.GET("/relation/friend/list/", controller.FriendList)
		// authRouter.GET("/message/chat/", controller.MessageChat)
		// authRouter.POST("/message/action/", controller.MessageAction)
	}
}
