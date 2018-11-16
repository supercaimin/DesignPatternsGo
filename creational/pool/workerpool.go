package pool

import(
	"fmt"
)
type WorkerPool struct {
	occupiedWorkers  map[string] *StringReverseWorker
	freeWorkers 	 map[string] *StringReverseWorker

}

func NewWorkerPool() *WorkerPool {
	workerPool := new(WorkerPool)
	workerPool.occupiedWorkers = make(map[string] *StringReverseWorker)
	workerPool.freeWorkers = make(map[string] *StringReverseWorker)
	return workerPool
}

func (self *WorkerPool) Get() *StringReverseWorker  {
	var worker *StringReverseWorker
	if len(self.freeWorkers) == 0 {
		worker = NewStringReverseWorker()
		fmt.Println(worker)
	} else {
		for _,v := range self.freeWorkers {
			worker = v
			break
		}
	}
	delete(self.freeWorkers, worker.Key)
	self.occupiedWorkers[worker.Key] = worker
	self.Log()
	return worker
}

func (self *WorkerPool) Dispose(worker *StringReverseWorker) {
	if _, ok := self.occupiedWorkers[worker.Key]; ok {
		self.freeWorkers[worker.Key] = worker
		delete(self.occupiedWorkers, worker.Key)
	}
	self.Log()
}

func (self *WorkerPool) Count() int {
	return len(self.occupiedWorkers) + len(self.freeWorkers)
}

func (self *WorkerPool) Log() {
	fmt.Println("free : ", len(self.freeWorkers))
	fmt.Println("occupied : ", len(self.occupiedWorkers))
}