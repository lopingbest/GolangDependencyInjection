package test

import (
	"fmt"
	"lopingbest/GolangRESTFullAPI/simple"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializeService()
	fmt.Println(err)
	fmt.Println(simpleService)
}
