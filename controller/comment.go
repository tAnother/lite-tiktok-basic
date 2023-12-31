package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tAnother/lite-tiktok-basic/model"
	"github.com/tAnother/lite-tiktok-basic/proto"
)

type CommentListResponse struct {
	proto.Response
	CommentList []model.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	proto.Response
	Comment model.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")

	if user, exist := usersLoginInfo[token]; exist {
		if actionType == "1" {
			text := c.Query("comment_text")
			c.JSON(http.StatusOK, CommentActionResponse{Response: proto.Response{StatusCode: 0},
				Comment: model.Comment{
					ID:         1,
					UserID:     user.ID,
					Content:    text,
					CreateDate: "05-01",
				}})
			return
		}
		c.JSON(http.StatusOK, proto.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, proto.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    proto.Response{StatusCode: 0},
		CommentList: DemoComments,
	})
}
