package factory_method

type StdoutLoggerFactory struct{}

func (self *StdoutLoggerFactory) CreateLogger() Logger {
	return &StdoutLogger{}
}