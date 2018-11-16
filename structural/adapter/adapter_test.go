package adapter

import (
	"testing"
)

func TestAdapter(t *testing.T)  {
	book := &Book{}
	book.Open()
	book.TurnPage()

	if book.GetPage() == 2 {
		t.Log("PASS")
	} else {
		t.Error("NG")
	}

	kinde := NewKindle()
	eBook := NewEBookAdapter(kinde)
	eBook.Open()
	eBook.TurnPage()
	if eBook.GetPage() == 2 {
		t.Log("PASS")
	} else {
		t.Error("NG")
	}
}
