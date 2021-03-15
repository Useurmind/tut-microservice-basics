package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestIncrementor(t *testing.T) {
	inc := NewIncrementor(0)

	value1, err := inc.Increment()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, value1)
}
