// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package simple

import (
	"github.com/google/wire"
)

// Injectors from injector.go:

func InitializeService(isError bool) (*SimpleService, error) {
	simpleRepository := NewSimpleRepository(isError)
	simpleService, err := NewSimpleService(simpleRepository)
	if err != nil {
		return nil, err
	}
	return simpleService, nil
}

func InitializedDatabaseRepository() *DatabaseRepository {
	databasePostgreSQL := NewDatabasePostgreSQL()
	databaseMongoDB := NewDatabaseMongoDB()
	databaseRepository := NewDatabaseRepository(databasePostgreSQL, databaseMongoDB)
	return databaseRepository
}

//Injector
func InitializedFooBarService() *FooBarService {
	fooRepository := NewFooRepository()
	fooService := NewFooService(fooRepository)
	barRepository := NewBarRepository()
	barService := NewBarService(barRepository)
	fooBarService := NewFooBarService(fooService, barService)
	return fooBarService
}

func InitializedHelloService() *HelloService {
	sayHelloImpl := NewSayHelloImpl()
	helloService := NewHelloService(sayHelloImpl)
	return helloService
}

//Bagian mana yang mau di inject
func InitializedFoobar() *FooBar {
	foo := NewFoo()
	bar := NewBar()
	fooBar := &FooBar{
		Foo: foo,
		Bar: bar,
	}
	return fooBar
}

// injector.go:

var fooset = wire.NewSet(
	NewFooRepository,
	NewFooService,
)

var barset = wire.NewSet(
	NewBarRepository,
	NewBarService,
)

//Binding Interface
var helloset = wire.NewSet(
	NewSayHelloImpl, wire.Bind(new(SayHello), new(*SayHelloImpl)),
)
