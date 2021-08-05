package models

import (
	"fmt"
	"project1/config"
	"project1/entities"
)

type ProductModel struct {
}

func (*ProductModel) FindAll() ([]entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		
		fmt.Println(err)
		return nil, err
	} else {
		rows, err2 := db.Query("select * from product")
		if err2 != nil {
			return nil, err2
		} else {
			var products []entities.Product
			for rows.Next() {
				var product entities.Product
				rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Description)
				products = append(products, product)
			}
			return products, nil
		}
	}
}

func (*ProductModel) Create(product *entities.Product)bool{
	fmt.Println("hello")
	db, err := config.GetDB()

	if err != nil {
		fmt.Println(err)
		return false
	} else {
		result, err2  :=db.Exec("insert into product(name, price, quantity, description) values(?,?,?,?)",product.Name,product.Price,product.Quantity,product.Description)
		if err2 != nil {
			return false
		} else {
			rowsAffected,_ := result.RowsAffected()
			return rowsAffected >0
	}
}
}
