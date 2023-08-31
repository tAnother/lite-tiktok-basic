package middleware

import (
	"context"
	"net/http"

	"github.com/RaymondCode/simple-demo/controller"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/tools"

	// "github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			token = c.PostForm("token")
		}
		userP, err := checkLogin(token)
		if err != nil || userP == nil {
			c.JSON(http.StatusOK, controller.UserResponse{
				Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
			})
			c.Abort() // This prevents the controller from being called
			return
		}

		// If needed, you can add the user to the context for subsequent use in your application
		c.Set("user", *userP)

		c.Next()
	}
}

func checkLogin(token string) (*model.User, error) {
	ctx := context.Background()
	serUser, err := tools.GetClient().Get(ctx, token).Result() // todo: i don't think we should store a serialized user in redis...
	if err != nil {
		return nil, err
	}
	user, err := model.DeserializeUser(serUser)
	if err != nil {
		return nil, err
	}
	return user, nil
}
