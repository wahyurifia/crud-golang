package productcontroller

import (
	"net/http"
	"project-crud/entities"
	"project-crud/models/categorymodel"
	"project-crud/models/productmodel"
	"strconv"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	products := productmodel.GetAll()
	data := map[string]any {
		"products": products,
	}

	temp, err := template.ParseFiles("views/product/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/product/create.html")
		if err != nil {panic(err)}

		categories := categorymodel.GetAll()
		data := map[string]any {
			"categories": categories,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var product entities.Product

		categoryId, err := strconv.Atoi( r.FormValue("category_id"))
		if err != nil {panic(err)}

		stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil {panic(err)}
		
		product.Name = r.FormValue("name")
		product.Category.Id = uint(categoryId)
		product.Stock = stock
		product.Description = r.FormValue("description")
		product.CreatedAt = time.Now()
		product.UpdatedAt = time.Now()

		berhasil := productmodel.Create(product)
		if !berhasil {
			temp, _ := template.ParseFiles("views/product/create.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/product/edit.html")
		if err != nil {panic(err)}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {panic(err)}

		product := productmodel.GetById(id)
		categories := categorymodel.GetAll()
		data := map[string]any {
			"product": product,
			"categories": categories,
		}
		
		temp.Execute(w, data)
	}	

	if r.Method == "POST" {
		var product entities.Product
		stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil {panic(err)}

		categoryId, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {panic(err)}

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {panic(err)}

		product.Name = r.FormValue("name")
		product.Stock = stock
		product.Category.Id = uint(categoryId)
		product.Description = r.FormValue("description")
		product.UpdatedAt = time.Now()

		berhasil := productmodel.Update(id, product)
		if !berhasil {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
		}
		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {panic(err)}

	categories := categorymodel.GetAll()
	product := productmodel.GetById(id)
	data := map[string]any {
		"product": product,
		"categories": categories,
	}

	temp, err := template.ParseFiles("views/product/detail.html")
	if err != nil {panic(err)}

	temp.Execute(w, data)
}
func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {panic(err)}

	if err := productmodel.Delete(id); err != nil{
		panic(err)
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}