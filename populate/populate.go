package populate

import (
	_ "github.com/lib/pq"
	"log"
	"github.com/bayusatmoko/packform/db"
)

var baseQuery = `
DROP TABLE IF EXISTS customer_companies;
CREATE TABLE customer_companies
(company_id integer NOT NULL PRIMARY KEY,company_name varchar);
COPY customer_companies(company_id, company_name)
FROM '/sql/customer_companies.csv'
DELIMITER ','
CSV HEADER;

DROP TABLE IF EXISTS deliveries;
CREATE TABLE deliveries
(id integer NOT NULL PRIMARY KEY, order_item_id integer, delivered_quantity integer);
COPY deliveries(id, order_item_id, delivered_quantity)
FROM '/sql/deliveries.csv'
DELIMITER ','
CSV HEADER;

DROP TABLE IF EXISTS orders;
CREATE TABLE orders
(id integer NOT NULL PRIMARY KEY, created_at date, order_name varchar, customer_id varchar);
COPY orders(id, created_at, order_name, customer_id)
FROM '/sql/orders.csv'
DELIMITER ','
CSV HEADER;

DROP TABLE IF EXISTS order_items;
CREATE TABLE order_items
(id integer NOT NULL PRIMARY KEY, order_id integer, price_per_unit numeric, quantity integer, product varchar);
COPY order_items(id, order_id, price_per_unit, quantity, product)
FROM '/sql/order_items.csv'
DELIMITER ','
CSV HEADER;

DROP TABLE IF EXISTS customers;
CREATE TABLE customers
(user_id varchar NOT NULL PRIMARY KEY, login varchar, password varchar, name varchar, company_id integer, credit_cards varchar []);
COPY customers(user_id, login, password, name, company_id, credit_cards)
FROM '/sql/customers.csv'
DELIMITER ','
CSV HEADER;
`

func Seed() {
	_, error := db.GetDB().Query(baseQuery)
	if error != nil {
		log.Fatal(error)
	}
}