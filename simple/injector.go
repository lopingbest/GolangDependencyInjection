//go:build wireinject
// +build wireinject

package simple

func InitializeService() *SimpleService {
	//menggunakan build diwire,lalu mencantumkan function yang akan digunakan di provider
	wire.Build(
		NewSimpleRepository,
		NewSimpleService)

	//Data di body akan diganti oleh Google Wire
	return nil
}
