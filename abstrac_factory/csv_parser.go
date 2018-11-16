package abstract_factory

import(
	"fmt"
)

const OPTION_CONTAINS_HEADER = true
const OPTION_CONTAINS_NO_HEADER = false

type CsvParser struct {
	skipHeaderLine bool
}

func NewCsvParser(skipHeaderLine bool)(*CsvParser)  {
	return &CsvParser{skipHeaderLine}
}

func (self *CsvParser) parse(input string) []interface{}{
	headerWasParsed := false
	parsedLines := make([]string)
	return nil
}