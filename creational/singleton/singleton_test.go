 package singletion

 import(
	 "testing"
 )

 func TestSingletion(t *testing.T)  {
	 firstCall := GetInstance()
	 secondCall := GetInstance()

	 if firstCall == secondCall {
		 t.Log("PASS")
	 } else {
		 t.Error("NG")
	 }

 }