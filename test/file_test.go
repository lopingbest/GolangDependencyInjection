package test

import (
	"github.com/stretchr/testify/assert"
	"lopingbest/GolangRESTFullAPI/simple"
	"testing"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("Databases")
	assert.NotNil(t, connection)

	cleanup()
}
