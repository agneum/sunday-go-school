package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanFile(t *testing.T) {

	p := Params{
		path: "test_data/data.txt",
		n:    3,
	}

	res, err := scanFile(p)

	assert.NoError(t, err)
	assert.Equal(t, 3, len(res))
}
