package router

import (
    "carParking/middleware"

    "github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

    router := mux.NewRouter()


    router.HandleFunc("/api/user/{carId}/{ownerName}", middleware.GetFreePakingSlot).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/user/{id}", middleware.DeallocateParking).Methods("PUT", "OPTIONS")


    return router
}