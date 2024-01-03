package main

import (
	"database/sql"
	
)

type product struct {
	ID int `json: "id"`
	Name string `json: "name"`
	Quantity int `json: "quantity"`
	Price float64 `json: "price"`
}

func getProducts(db *sql.DB)([]product, error) {
	query := "SELECT id, name, quantity, price from products"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	products := []product{}
	for rows.Next() {
		var p product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, p.Quantity)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}