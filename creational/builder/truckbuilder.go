package builder

type TruckBuilder struct {
	truck *Truck
}

func (self *TruckBuilder)CreateVehicle(){
	self.truck = &Truck{}
	self.truck.data = make(map[string]interface{}, 30)
}
func (self *TruckBuilder)AddWheels(){
	self.truck.SetPart("wheel1", &Wheel{})
	self.truck.SetPart("wheel2", &Wheel{})
	self.truck.SetPart("wheel3", &Wheel{})
	self.truck.SetPart("wheel4", &Wheel{})
	self.truck.SetPart("wheel5", &Wheel{})
	self.truck.SetPart("wheel6", &Wheel{})
}

func (self *TruckBuilder)AddEngine(){
	self.truck.SetPart("engine", &Engine{})
}

func (self *TruckBuilder)AddDoors(){
	self.truck.SetPart("door1", &Door{})
	self.truck.SetPart("door2", &Door{})
}

func (self *TruckBuilder)GetVehicle() Vehicler{
	return self.truck
}