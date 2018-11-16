package prototype

type Sweater struct {
	tag string
}

func NewSweater(tag string) *Sweater {
	return &Sweater{tag}
}

func (self *Sweater)Clone() Cloth {
	sweater := &Sweater{}
	sweater.tag = self.tag
	return sweater
}
func (self *Sweater)GetTag() string {
	return self.tag
}