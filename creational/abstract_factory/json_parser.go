package abstract_factory

type JsonParser struct {}

func NewJsonParser() (*JsonParser) {
	return &JsonParser{}
}

func (self *JsonParser) Parse(input string) []interface{} {
	return nil
}