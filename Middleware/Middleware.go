package Middleware

import (
	"github.com/RaymondCode/simple-demo/Entry"
	"github.com/RaymondCode/simple-demo/controller"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserResponse struct {
	Entry.Response
	User Entry.User `json:"user"`
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			token = c.PostForm("token")
		}
		userP, err := service.CheckLogin(token)
		if err != nil || userP == nil {
			c.JSON(http.StatusOK, controller.UserResponse{
				Response: Entry.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
			})
			c.Abort() // This prevents the handler from being called
			return
		}

		// If needed, you can add the user to the context for subsequent use in your application
		c.Set("user", *userP)

		c.Next()
	}
}
