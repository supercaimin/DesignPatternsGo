package pool

import(
	"fmt"
	"testing"
)

func TestPool(t *testing.T)  {
	pool := NewWorkerPool()
	worker1 := pool.Get()
	worker2 := pool.Get()

	if pool.Count() == 2 {
		t.Log("PASS")
	} else {
		t.Error("NG")
	}
	str := worker1.Run("abdcdfdfdfdfd")
	fmt.Println(str)
	pool.Dispose(worker1)
	pool.Dispose(worker2)

}