package prototype

type Overcoat struct {
	tag string
}

func NewOvercoat(tag string) *Overcoat {
	return &Overcoat{tag}
}

func (self *Overcoat)Clone() Cloth {
	overcoat := &Overcoat{}
	overcoat.tag = self.tag
	return overcoat
}
func (self *Overcoat)GetTag() string {
	return self.tag
}