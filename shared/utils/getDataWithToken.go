package utils

import (
	"github.com/gin-gonic/gin"
)

func GetDataWithToken(res *gin.Context) map[string]interface{} {
	token := res.GetHeader("token")
	data, _ := ValidJwt(token)
	return data
}
