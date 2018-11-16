package factory_method

type FilerLoggerFactory struct{
	filePath string
}

func NewFileLoggerFactory(filePath string) *FilerLoggerFactory {
	return &FilerLoggerFactory{filePath}
}

func (self *FilerLoggerFactory) CreateLogger() Logger {
	return NewFileLogger(self.filePath)
}