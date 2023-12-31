package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tAnother/lite-tiktok-basic/config"
	"github.com/tAnother/lite-tiktok-basic/proto"
)

func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			token = c.PostForm("token") // TODO: change it so that we choose how to get token base on request method
		}
		fmt.Println(token)

		userID, err := checkLogin(token)
		if err != nil || userID == 0 {
			c.JSON(http.StatusUnauthorized, proto.Response{
				StatusCode: proto.BadCredentials,
				StatusMsg:  "Invalid token. Please login again.",
			})
			c.Abort() // This prevents the controller from being called
			return
		}

		// If needed, you can add the user to the context for subsequent use in your application
		// c.Set("user", *userP)

		c.Next()
	}
}

func checkLogin(token string) (userID int64, err error) {
	ctx := context.Background()
	result, err := config.RedisClient().Get(ctx, token).Result()
	if err != nil || result == "" {
		return 0, err
	}
	userID, err = strconv.ParseInt(result, 10, 64)
	return userID, err
}
