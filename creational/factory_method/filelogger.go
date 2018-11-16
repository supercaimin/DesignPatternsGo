package factory_method

import (
	"io/ioutil"
)

type FileLogger struct {
	filePath string
}

func NewFileLogger(filePath string)( *FileLogger)  {
	return &FileLogger{filePath}
}

func (self *FileLogger)Log(message string) {
	ioutil.WriteFile(self.filePath, []byte(message), 0644)
}