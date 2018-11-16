// 需要借助静态变量来实现
package singletion

type Singletion struct{}

var instance *Singletion

func GetInstance() *Singletion {
	if instance != nil {
		return instance
	}
	return new(Singletion)
}
