package main

import (
	// "cart/db"
	"cart/db"
	"cart/helper"
	"cart/rabbitmq"
	"cart/router"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	_ "google.golang.org/grpc"
)

func main() {
	config := helper.GetConfiguration()
	go db.ConnectToDatabse()
	go rabbitmq.ConnectToRabbitMQ()
	r := mux.NewRouter()
	router.GenerateRoutes(r)

	log.Println("Serving at port = " + strconv.Itoa(config.Port) + "...")
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Port), r))
}
