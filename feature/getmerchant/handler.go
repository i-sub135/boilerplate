package getmerchant

import (
	"api-majoo/shared/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Handler {
	return &Handler{db}
}

func (h *Handler) Implement(ctx *gin.Context) {

	dataWithToken := utils.GetDataWithToken(ctx)
	userID := dataWithToken["ID"].(float64)

	merchant, _ := h.getMerchantByUserID(int(userID))
	outlets, _ := h.getOutletByMerchantID(merchant.ID)
	response := h.toResponse(merchant, outlets)

	ctx.JSON(200, response)
}
