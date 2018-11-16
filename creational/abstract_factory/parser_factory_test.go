package abstract_factory

import (
	"fmt"
	"testing"
)

var factory *ParserFactory = new(ParserFactory)

func TestCreateCsvParser(t *testing.T)  {
	var parser Parser = factory.CreateCsvParser(OPTION_CONTAINS_HEADER)
	if v,ok := parser.(*CsvParser); ok {
		t.Log("PASS")
		fmt.Printf("%T \n", v)
	} else {
		t.Error("NG")
	}
}

func TestCreateJsonParser(t *testing.T)  {
	var parser Parser = factory.CreateJsonParser()
	if v,ok := parser.(*JsonParser); ok {
		t.Log("PASS")
		fmt.Printf("%T \n", v)
	} else {
		t.Error("NG")
	}
}