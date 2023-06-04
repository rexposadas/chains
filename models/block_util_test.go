package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO: Tirso -> test all the functions in block_util.go file
func Test_GenerateBlock(t *testing.T) {

}

func Test_CalculateHash(t *testing.T) {

	testCases := []struct {
		B Block
	}{
		{B: Block{Bpm: 34}},
		{B: Block{}},
	}

	for _, item := range testCases {
		result := CalculateHash(item.B)
		assert.NotEmpty(t, result)
		t.Log(result)
	}
}
