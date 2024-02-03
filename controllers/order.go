package controllers

import (
	"github.com/bayusatmoko/packform/models"
	"github.com/gin-gonic/gin"
)

var orderModel = new(models.OrderModel)

func GetOrders(c *gin.Context) {
	// parameters
	keyword := c.Query("keyword")
	offset := c.Query(("offset"))
	startDate := c.Query(("startDate"))
	endDate := c.Query(("endDate"))

	orders, total, err := orderModel.All(startDate, endDate, keyword, offset)
	if err != nil {
		panic(err)
	}

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.JSON(200, gin.H{
		"orders": orders,
		"rows":   total,
	})
}
