package homecontroller

import (
	"net/http"
	"text/template"
)

func Welcome(writer http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("views/home/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(writer, nil)
}