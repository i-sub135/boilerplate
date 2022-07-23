package service

import (
	"api-majoo/feature/getmerchant"
	"api-majoo/feature/healthcek"
	"api-majoo/feature/login"
	"api-majoo/shared/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func Routes(engine *gin.Engine) {
	//endpoint health
	var DBR = DBConnect{
		DSN: GetDSN(),
	}
	conn, err := DBR.Conn()
	if err != nil {
		log.Fatal(err)
	}

	engine.GET("/health", healthcek.New(conn).Implement)
	engine.POST("/login", login.New(conn).Implement)

	merchan := engine.Group("/merchant")

	//register JWT validations
	merchan.Use(utils.ReqValidator)
	merchan.GET("", getmerchant.New(conn).Implement)

}
