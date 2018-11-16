package builder

type Vehicler interface{
	SetPart(key string, value interface{})
}
type Vehicle struct {
	data	map[string] interface{}
}

func (self *Vehicle)SetPart(key string, value interface{})  {
	self.data[key] = value
}

type Door struct {}

type Engine struct{}

type Wheel struct{}

type Car struct {
	Vehicle
}

type Truck struct {
	Vehicle
}



