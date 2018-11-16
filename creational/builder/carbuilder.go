package builder

type CarBuilder struct {
	car *Car
}


func (self *CarBuilder)CreateVehicle(){
	self.car = &Car{}
	self.car.data = make(map[string]interface{}, 20)
}
func (self *CarBuilder)AddWheels(){
	self.car.SetPart("wheel1", &Wheel{})
	self.car.SetPart("wheel2", &Wheel{})
	self.car.SetPart("wheel3", &Wheel{})
	self.car.SetPart("wheel4", &Wheel{})
}

func (self *CarBuilder)AddEngine(){
	self.car.SetPart("engine", &Engine{})
}

func (self *CarBuilder)AddDoors(){
	self.car.SetPart("door1", &Door{})
	self.car.SetPart("door2", &Door{})
}

func (self *CarBuilder)GetVehicle() Vehicler{
	return self.car
}