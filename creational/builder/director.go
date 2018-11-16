package builder

type Director struct{}

func (self *Director)build(builder Builder)  Vehicler{
	builder.CreateVehicle()
	builder.AddDoors()
	builder.AddEngine()
	builder.AddWheels()
	
	return builder.GetVehicle()
}