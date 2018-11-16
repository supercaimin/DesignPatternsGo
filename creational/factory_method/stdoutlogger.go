package factory_method

import(
	"fmt"
)

type StdoutLogger struct {
}

func (self *StdoutLogger)Log(messge string)  {
	fmt.Println(messge)
}