package main

import (
	"github.com/bayusatmoko/packform/controllers"
	"github.com/bayusatmoko/packform/db"
	"github.com/bayusatmoko/packform/populate"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	populate.Seed()
	r := gin.Default()
	r.GET("/orders", controllers.GetOrders)
	r.SetTrustedProxies(nil)
	r.Run()
}
