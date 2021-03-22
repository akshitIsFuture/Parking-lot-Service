package router

import (
    "Parking-lot-Service/handler"
    "github.com/go-openapi/runtime/middleware"
    "github.com/gorilla/mux"
    "net/http" // used to access the request and response object of the api
    
)

// Router is exported and used in main.go
func Router() *mux.Router {

    router := mux.NewRouter()

    router.HandleFunc("/parking/allocate/{carId}/{ownerName}", handler.GetFreePakingSlot).Methods("PUT", "OPTIONS")
    router.HandleFunc("/parking/deallocate/{id}", handler.DeallocateParking).Methods("PUT", "OPTIONS")
    ops := middleware.RedocOpts{SpecURL:"/swagger.yaml"}
    sh := middleware.Redoc(ops, nil)
    router.Handle("/docs",sh)
    router.Handle("/swagger.yaml",http.FileServer(http.Dir("./")))


    return router
}