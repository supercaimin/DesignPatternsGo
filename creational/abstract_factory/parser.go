package abstract_factory


type Parser interface {
	Parse(input string) []interface{}
}