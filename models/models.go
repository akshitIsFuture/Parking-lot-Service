package models

// User schema of the parking table
type ParkingDetails  struct {
	ParkingId int64 `json:"parkingId"`
	CarId string `json:"carId"`
	OwnerName string `json:"ownerName"`
}

