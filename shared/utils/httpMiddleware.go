package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReqValidator(res *gin.Context) {
	var (
		token string = res.GetHeader("token")
	)

	switch {
	case token == "":
		out := map[string]interface{}{"error": "Token can't be empty"}
		res.JSON(http.StatusUnauthorized, out)
		res.Abort()
	default:
		_, err := ValidJwt(token)
		if err != nil {
			out := map[string]interface{}{"error": "Token Not Valid :: " + err.Error()}
			res.JSON(http.StatusUnauthorized, out)
			res.Abort()
			return
		}
		res.Next()
	}

}
