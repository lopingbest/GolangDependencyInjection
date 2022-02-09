package test

import (
	"github.com/stretchr/testify/assert"
	"lopingbest/GolangRESTFullAPI/simple"
	"testing"
)

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.InitializeService(true)
	assert.NotNil(t, err)
	assert.Nil(t, simpleService)
}

func TestSimpleServiceSuccess(t *testing.T) {
	simpleservice, err := simple.InitializeService(false)
	assert.Nil(t, err)
	assert.NotNil(t, simpleservice)
}
