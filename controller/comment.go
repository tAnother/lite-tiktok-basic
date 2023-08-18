package controller

import (
	"github.com/RaymondCode/simple-demo/Entry"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommentListResponse struct {
	Entry.Response
	CommentList []Entry.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Entry.Response
	Comment Entry.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")

	if user, exist := usersLoginInfo[token]; exist {
		if actionType == "1" {
			text := c.Query("comment_text")
			c.JSON(http.StatusOK, CommentActionResponse{Response: Entry.Response{StatusCode: 0},
				Comment: Entry.Comment{
					Id:         1,
					User:       user,
					Content:    text,
					CreateDate: "05-01",
				}})
			return
		}
		c.JSON(http.StatusOK, Entry.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Entry.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Entry.Response{StatusCode: 0},
		CommentList: DemoComments,
	})
}
