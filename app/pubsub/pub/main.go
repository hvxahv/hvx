package main

import (
	"github.com/gin-gonic/gin"
	"hvxahv/pkg/middleware"
)

func main()  {
	r := gin.Default()
	r.Use(middleware.CORS())
	r.GET("/pub/new", PubHandler)
	r.Run(":8960")
}
