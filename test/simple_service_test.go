package test

import (
	"fmt"
	"lopingbest/GolangRESTFullAPI/simple"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService := simple.InitializeService()
	fmt.Println(simpleService.SimpleRepository)
}
