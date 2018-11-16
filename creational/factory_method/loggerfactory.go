package factory_method

type FactoryMethod interface {
	CreateLogger() Logger
}