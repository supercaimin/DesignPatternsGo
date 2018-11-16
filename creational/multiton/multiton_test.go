package multiton

import(
	"testing"
)

func TestMultiton(t *testing.T)  {
	multiton := NewMultiton()
	mysqlInstance := multiton.GetInstance(INSTANCE_MYSQL)
	if mysqlInstance.Name == INSTANCE_MYSQL {
		t.Log("PASS")
	} else {
		t.Error("NG")
	}

	sqlliteInstance := multiton.GetInstance(INSTANCE_SQLLITE)
	if sqlliteInstance.Name == INSTANCE_SQLLITE {
		t.Log("PASS")
	} else {
		t.Error("NG")
	}
}