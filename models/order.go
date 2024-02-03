package models

import (
	"github.com/bayusatmoko/packform/db"
	"log"
)

// Order ...
type Order struct {
	Order_name       string
	Product          string
	Company_name     string
	Name             string
	Created_at       string
	Total_amount     *float64
	Delivered_amount *float64
}

type OrderModel struct{}

func (m OrderModel) All() (orders []Order, total int64, err error) {
	query := `
	SELECT o.order_name, oi.product,
	cc.company_name, cu.name,
	o.created_at,
	oi.price_per_unit * oi.quantity AS total_amount, d.delivered_quantity * oi.price_per_unit AS delivered_amount,
	COUNT(*) OVER()
	FROM orders o LEFT JOIN customers cu ON o.customer_id = cu.user_id
	LEFT JOIN customer_companies cc ON cu.company_id = cc.company_id
	LEFT JOIN order_items oi ON o.id = oi.order_id
	LEFT JOIN deliveries d ON d.order_item_id = oi.id
	`
	orders = make([]Order, 0)
	rows, err := db.GetDB().Query(query)
	total = 0
	if rows != nil {
		for rows.Next() {
			var order_name string
			var product string
			var company_name string
			var name string
			var created_at string
			var total_amount *float64
			var delivered_amount *float64
			var count int64

			if err := rows.Scan(&order_name, &product, &company_name, &name, &created_at, &total_amount, &delivered_amount, &count); err != nil {
				log.Fatal(err)
			}

			orders = append(orders, Order{Order_name: order_name,
				Product:          product,
				Company_name:     company_name,
				Name:             name,
				Created_at:       created_at,
				Total_amount:     total_amount,
				Delivered_amount: delivered_amount})

			total = count
		}
	}

	return orders, total, err
}
