package abstract_factory

import (
	"encoding/json"
	"strings"
)
type JsonParser struct {}

func NewJsonParser() (*JsonParser) {
	return &JsonParser{}
}

func (self *JsonParser) Parser(input string) []interface{} {
	return nil
}