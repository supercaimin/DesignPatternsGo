package static_factory

import (
	"testing"
)

func TestStaticFactory(t *testing.T) {
	formatNumber := Factory(FORMAT_NUMBER_TYPE)
	if _,ok := formatNumber.(*FormatNumber); ok {
		t.Log("PASS")
	} else {
		t.Error("NG")
	}
	formatString := Factory(FORMAT_STRING_TYPE)
	if _,ok := formatString.(*FormatString); ok {
		t.Log("PASS")
	} else {
		t.Error("NG")
	}
}