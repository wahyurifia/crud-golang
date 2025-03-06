package categorycontroller

import (
	"net/http"
	"project-crud/models/categorymodel"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
 // sama seperti di expressjs -> trycatch -> panggil model/servicesnya yang akan berinteraksi dengan database
	categories := categorymodel.GetAll()
	data := map[string]any {
		"categories": categories,
	}

 // sama seperti di expressjs -> res.json({})
	temp, err := template.ParseFiles("views/category/index.html")
	if err != nil  {
		panic(err)
	}
	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	
}