package productcontroller

import (
	"net/http"
	"project1/entities"
	"project1/models"
	"strconv"
	"text/template"
)

func Index(response http.ResponseWriter, request *http.Request) {

	var productModel models.ProductModel
	products, _ := productModel.FindAll()
	data := map[string]interface{}{
		"product": products,
	}
	tmp, _ := template.ParseFiles("view/product/index.html")
	tmp.Execute(response, data)

}

func Add(response http.ResponseWriter, request *http.Request) {

	tmp, _ := template.ParseFiles("view/product/add.html")
	tmp.Execute(response, nil)

}

func ProcessAdd(response http.ResponseWriter, request *http.Request) {

	request.ParseForm()
	var product entities.Product
	product.Name = request.Form.Get("name")
	product.Price, _ = strconv.ParseFloat(request.Form.Get("price"), 64)
	product.Quantity, _ = strconv.ParseInt(request.Form.Get("quantity"), 10, 64)
	product.Description = request.Form.Get("description")

	var productModel models.ProductModel
	productModel.Create(&product)
	http.Redirect(response, request, "/product", http.StatusSeeOther)

	tmp, _ := template.ParseFiles("view/product/add.html")
	tmp.Execute(response, nil)

}
