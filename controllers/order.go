package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/bayusatmoko/packform/models"
)

var orderModel = new(models.OrderModel)

func GetOrders(c *gin.Context) {
	// parameters
	keyword := c.Query("keyword")
	offset := c.Query(("offset"))
	startDate := c.Query(("startDate"))
	endDate := c.Query(("endDate"))
	startDateString := "'"
	startDateString += startDate
	startDateString += "'"
	endDateString := "'"
	endDateString += endDate
	endDateString += "'"
	fmt.Println(keyword)
	fmt.Println(offset)

	orders, total, err := orderModel.All()

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.JSON(200, gin.H{
		"orders": orders,
		"rows": total,
	})
	panic((err))
}