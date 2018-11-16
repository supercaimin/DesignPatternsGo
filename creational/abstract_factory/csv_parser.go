package abstract_factory


const OPTION_CONTAINS_HEADER = true
const OPTION_CONTAINS_NO_HEADER = false

type CsvParser struct {
	skipHeaderLine bool
}

func NewCsvParser(skipHeaderLine bool)(*CsvParser)  {
	return &CsvParser{skipHeaderLine}
}

func (self *CsvParser) Parse(input string) []interface{}{
	return nil
}