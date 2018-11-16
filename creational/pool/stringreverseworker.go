package pool

import (
	"fmt"
	"time"
	"math/rand"
)
type StringReverseWorker struct {
	createdAt 	time.Time
	Key			string
}

func NewStringReverseWorker() *StringReverseWorker  {
	instance := &StringReverseWorker{}
	instance.generateKey()
	return instance
}

func(self *StringReverseWorker)generateKey(){
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
	   result = append(result, bytes[r.Intn(len(bytes))])
	}
	// 休眠1秒钟，否则连续调用产生的key是一样的
	time.Sleep(time.Duration(1) * time.Second)
	self.Key = string(result)
 }
func (self *StringReverseWorker) Run(text string) string {
	r := []rune(text)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func (self *StringReverseWorker) String() string {
	return fmt.Sprintf("Worker %s, Created at %s", self.Key, self.createdAt.UTC())
}