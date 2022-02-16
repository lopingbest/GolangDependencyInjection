//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
	"io"
	"os"
)

func InitializeService(isError bool) (*SimpleService, error) {
	//menggunakan build diwire,lalu mencantumkan function yang akan digunakan di provider
	wire.Build(
		NewSimpleRepository, NewSimpleService)

	//Data di body akan diganti oleh Google Wire
	return nil, nil
}

func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabaseMongoDB,
		NewDatabasePostgreSQL,
		NewDatabaseRepository,
	)
	return nil
}

var fooset = wire.NewSet(
	NewFooRepository,
	NewFooService,
)

var barset = wire.NewSet(
	NewBarRepository,
	NewBarService,
)

//Injector
func InitializedFooBarService() *FooBarService {
	wire.Build(fooset, barset, NewFooBarService)
	return nil
}

/*func InitializedHelloService() *HelloService {
	wire.Build(NewHelloService, NewSayHelloImpl)
	return nil
}*/

//Binding Interface
var helloset = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)), //yang butuh SayHello akan di balikan pointer
)

func InitializedHelloService() *HelloService {
	wire.Build(helloset, NewHelloService)
	return nil
}

//Bagian mana yang mau di inject
func InitializedFoobar() *FooBar {
	wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "Foo", "Bar"))
	return nil
}

var fooValue = &Foo{}
var barValue = &Bar{}

func InitializedFooBarUsingValue() *FooBar {
	wire.Build(wire.Value(fooValue), wire.Value(barValue), wire.Struct(new(FooBar), "*"))
	return nil
}

func InitializedReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

func InitializedConfiguration() *Configuration {
	wire.Build(
		NewApplication,
		wire.FieldsOf(new(*Application), "Configuration"),
	)
	return nil
}

func InitializedConnection(name string) (*Connection, func()) {
	wire.Build(
		NewConnection,
		NewFile)
	return nil, nil
}
