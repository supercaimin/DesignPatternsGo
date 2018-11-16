package simple_factory

import "testing"

func TestSimpleFactory(t *testing.T)  {
	factory := &SimpleFactory{}
	robot := factory.CreateRobot()
	robot.Walk("Home")
}