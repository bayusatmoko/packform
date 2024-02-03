package main

import (
	"github.com/bayusatmoko/packform/populate"
	"github.com/gin-gonic/gin"
	"github.com/bayusatmoko/packform/controllers"
)

func main() {
	populate.Seed()
	r := gin.Default()
	r.GET("/orders", controllers.GetOrders)
	r.SetTrustedProxies(nil)
	r.Run()
}
