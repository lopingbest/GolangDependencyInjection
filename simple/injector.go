//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitializeService() (*SimpleService, error) {
	//menggunakan build diwire,lalu mencantumkan function yang akan digunakan di provider
	wire.Build(
		NewSimpleRepository, NewSimpleService)

	//Data di body akan diganti oleh Google Wire
	return nil, nil
}
