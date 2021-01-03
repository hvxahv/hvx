package main

import (
	"github.com/gin-gonic/gin"
	"hvxahv/pkg/http"
)

func main()  {
	r := gin.Default()
	r.Use(http.CORS())
	r.GET("/pub/new", PubHandler)
	r.Run(":8960")
}
