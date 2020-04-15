package main

import (
	"log"
	"net/http"
	"os"
	"cart/routes"
	"github.com/gorilla/mux"
)

func main ()  {
	
	addr := determineListenAddress()

	r := mux.NewRouter()
	routes.GenerateRoutes(r)
	
	log.Println("Serving at port = " + addr + "...")
	log.Fatal(http.ListenAndServe(addr, r))
}


func determineListenAddress() string {
	port := os.Getenv("PORT")
	if port == "" {
		return ":7000"
	}
	return ":" + port
}