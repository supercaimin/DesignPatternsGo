package simple_factory

type SimpleFactory struct {}

func (self *SimpleFactory)CreateRobot() *Robot  {
	return &Robot{}
}