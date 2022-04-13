package config

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, Message interface{}, Data interface{}) {
	defer ResponseStruct.Set(9999)
	c.JSON(200, gin.H{
		"code":    1000,
		"status":  "SUCCESS",
		"message": fmt.Sprintf("%v", Message),
		"value":   Data,
	})
}
func ErrorResponse(c *gin.Context, Message interface{}) {
	defer ResponseStruct.Set(9999)
	c.AbortWithStatusJSON(400, gin.H{
		"code":    ResponseStruct.Get().ResponseCode,
		"status":  "FAILED",
		"message": fmt.Sprintf("%v", Message),
	})

}
func ResponseGlobal(c *gin.Context, codeHttp int, Data map[string]interface{}) {
	c.JSON(codeHttp, Data)
}

var ResponseStruct = responseStruct{9999}

type responseStruct struct {
	ResponseCode int
}

func (r *responseStruct) Get() *responseStruct {
	return r
}

func (r *responseStruct) Set(code int) *responseStruct {
	r.ResponseCode = code
	return r
}
