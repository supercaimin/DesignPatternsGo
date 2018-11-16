package prototype

type Cloth interface {
	Clone() Cloth
	GetTag() string
}