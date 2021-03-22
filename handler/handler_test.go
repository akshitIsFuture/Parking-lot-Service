package handler

import ("testing")

func TestGetFreePaking( t *testing.T ){
	slotStatus,_ := getFreePaking("akshit","car1")

	if slotStatus != "slot allocated" {
		t.Errorf("HELLO[\"\") failed, expected %v, got %v", "Hello Dude!",slotStatus)
	} 

}