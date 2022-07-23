package login

import (
	"api-majoo/shared/model"
	"api-majoo/shared/utils"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Handler struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Handler {
	return &Handler{db}
}

func (h *Handler) Implement(ctx *gin.Context) {

	var loginParam model.UserLoginRequest

	_ = ctx.BindJSON(&loginParam)

	if loginParam.UserName == "" || loginParam.Password == "" {
		ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "user unauthorized please input user & password",
		})
		return
	}

	algorithm := md5.New()
	algorithm.Write([]byte(loginParam.Password))
	password := hex.EncodeToString(algorithm.Sum(nil))

	data, err := h.getUserByUsernamePassword(loginParam.UserName, password)

	switch {
	case err != nil:
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	case data.ID == 0:
		ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "user unauthorized",
		})
		return

	default:
		jwtClm := map[string]interface{}{
			"ID":       data.ID,
			"Name":     data.Name,
			"UserName": data.UserName,
		}
		jwtCreated, _ := utils.CreateJwt(jwtClm)
		ctx.JSON(http.StatusOK, map[string]interface{}{"Token": jwtCreated})
	}

}
