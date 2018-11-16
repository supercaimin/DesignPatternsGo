package factory_method

import (
	"testing"
)

func TestStdoutLoggerFactoryMethod(t *testing.T)  {
	loggerFactory := &StdoutLoggerFactory{}
	logger := loggerFactory.CreateLogger()
	if _, ok := logger.(*StdoutLogger); ok {
		t.Log("PASS")
	} else {
		t.Error("NG")
	}
}

func TestFileLoggerFactoryMethod(t *testing.T)  {
	loggerFactory := NewFileLoggerFactory("./")
	logger := loggerFactory.CreateLogger()
	if _, ok := logger.(*FileLogger); ok {
		t.Log("PASS")
	} else {
		t.Error("NG")
	}
}