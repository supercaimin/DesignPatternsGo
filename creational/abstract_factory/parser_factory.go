package abstract_factory

type ParserFactory struct {}

func (self *ParserFactory) CreateCsvParser(skipHeaderLine bool)( *CsvParser) {
	return NewCsvParser(skipHeaderLine)
}

func (self *ParserFactory) CreateJsonParser()( *JsonParser) {
	return NewJsonParser()
}