package prototype

import(
	"testing"
)

func TestPrototype(t *testing.T)  {
	sweater := NewSweater("s1")
	newSweater := sweater.Clone()

	if _,ok := newSweater.(*Sweater); ok {
		t.Log("PASS")
	} else {
		t.Error("NG")
	}

	overcoat := NewOvercoat("c1")
	newOvercoat := overcoat.Clone()

	if _,ok := newOvercoat.(*Overcoat); ok {
		t.Log("PASS")
	} else {
		t.Error("NG")
	}
}