package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yudiagonal/go-crud-form-validation/controller"
)

func main() {
	// config.ConnectDB()
	http.HandleFunc("/index", controller.Index)
	http.HandleFunc("/dosen/create", controller.Create)
	http.HandleFunc("/dosen/edit", controller.Edit)
	http.HandleFunc("/dosen/delete", controller.Delete)

	fmt.Println("server started at localhost:9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
