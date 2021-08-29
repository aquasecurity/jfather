package jfather

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Object(t *testing.T) {
	example := []byte(`{
	"name": "testing",
	"balance": 3.14
}`)
	target := struct {
		Name    string  `json:"name"`
		Balance float64 `json:"balance"`
	}{}
	require.NoError(t, Unmarshal(example, &target))
	assert.Equal(t, "testing", target.Name)
	assert.Equal(t, 3.14, target.Balance)
}

func Test_ObjectWithPointers(t *testing.T) {
	example := []byte(`{
	"name": "testing",
	"balance": 3.14
}`)
	target := struct {
		Name    *string  `json:"name"`
		Balance *float64 `json:"balance"`
	}{}
	require.NoError(t, Unmarshal(example, &target))
	assert.Equal(t, "testing", *target.Name)
	assert.Equal(t, 3.14, *target.Balance)
}

type nestedParent struct {
	Child *nestedChild
	Name  string
}

type nestedChild struct {
	Blah string `json:"secret"`
}

func Test_ObjectWithPointerToNestedStruct(t *testing.T) {
	example := []byte(`{
	"Child": {
		"secret": "password"
	},
	"Name": "testing"
}`)

	var parent nestedParent
	require.NoError(t, Unmarshal(example, &parent))
	assert.Equal(t, "testing", parent.Name)
	assert.Equal(t, "password", parent.Child.Blah)
}