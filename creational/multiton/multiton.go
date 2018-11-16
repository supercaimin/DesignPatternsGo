package multiton

const (
	INSTANCE_MYSQL = "mysql"
	INSTANCE_SQLLITE = "sqllite"
) 

type INSTANCE_NAME string

type SQLConnection struct {
	Name INSTANCE_NAME
}

type Multiton struct {
	instances map[INSTANCE_NAME] *SQLConnection
}

func NewMultiton() *Multiton {
	return &Multiton{make(map[INSTANCE_NAME] *SQLConnection, 2)}
}

func (self *Multiton)GetInstance(instanceName INSTANCE_NAME) *SQLConnection {
	if instance, ok := self.instances[instanceName]; ok {
		return instance
	}
	instance := &SQLConnection{instanceName}
	self.instances[instanceName] = instance
	return instance
}