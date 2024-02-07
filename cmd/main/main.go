package main

import (
	"Inventory-APIs/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
)

func main()  {
	r := mux.NewRouter()
	routes.InventoryRouters(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9090", r))

}
