package builder

type Builder interface {
	CreateVehicle()
	AddWheels()
	AddEngine()
	AddDoors()
	GetVehicle() Vehicler
}