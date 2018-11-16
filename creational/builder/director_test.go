package builder

import(
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T)  {
	var director *Director = &Director{}
	var vehicle Vehicler 
	vehicle = director.build(&CarBuilder{})
	if v, ok := vehicle.(*Car); ok {
		t.Log("PASS")
		fmt.Println(v)
	} else {
		t.Error("NG")
	}
	vehicle = director.build(&TruckBuilder{})
	if v, ok := vehicle.(*Truck); ok {
		t.Log("PASS")
		fmt.Println(v)
	} else {
		t.Error("NG")
	}
}