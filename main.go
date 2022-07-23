package main

import (
	"api-majoo/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
func main() {
	mode := os.Getenv("MODE")
	gin.SetMode(mode)

	route := gin.Default()
	service.Routes(route)

	_ = route.Run()
}
