package resp2

import (
	"net/http"

	"manger/pkg/misc/error2"
	logger "manger/pkg/misc/log"

	"github.com/gin-gonic/gin"
)

type resp interface{}

type r struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data resp   `json:"data"`
}

// Format 统一返回值格式
func Format(c *gin.Context, resp resp, err error) {
	if err != nil {
		if err1, ok := err.(error2.Error); ok {
			c.JSON(http.StatusOK, &r{
				Code: err1.Code,
				Msg:  err1.Error(),
				Data: resp,
			})
			return
		}
		logger.Logger.Error(err)
		c.JSON(http.StatusOK, &r{
			Code: error2.Unknown,
			Msg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &r{
		Code: error2.Success,
		Data: resp,
	})
	return
}
