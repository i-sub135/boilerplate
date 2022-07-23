package healthcek

import (
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
	err := h.db.Error
	ctx.JSON(200, map[string]interface{}{"status": map[string]interface{}{"db_err": err}})
}
