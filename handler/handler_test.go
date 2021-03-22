package handler

import ("testing")

func TestGetFreePaking( t *testing.T ){
	slotStatus,_ := getFreePaking("akshit","car1")

	if slotStatus != "slot allocated" {
		t.Errorf("getFreePaking[\"\") failed, expected %v, got %v", "slot allocated",slotStatus)
	} 

}

func TestDeallocateCarSpace( t *testing.T ){
	slotStatus,_ := DeallocateCarSpace("car1")

	if slotStatus != "carId car 1 left the parking slot" {
		t.Errorf("DeallocateCarSpace[\"\") failed, expected %v, got %v", "carId car 1 left the parking slot",slotStatus)
	} 

}