package main

import (
	"Parking-lot-Service/router"
	"fmt"
	"log"
	"net/http"
	"go.elastic.co/apm/module/apmhttp"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on the port 8080 akshit...")
	log.Fatal(http.ListenAndServe(":8080", apmhttp.Wrap(r)))
}
