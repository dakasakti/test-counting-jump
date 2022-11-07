package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// [0, 0, 0, 0, 1, 0, 0]
func jump(slice []int) int {
	var count int

	fmt.Println("slice awal", slice)
	for i := 0; i < len(slice); i++ {
		if slice[2] == 0 {
			slice = slice[2:]
			fmt.Println("Slice ganti 2x :", slice)
			count++
		} else {
			slice = slice[1:]
			fmt.Println("Slice ganti else :", slice)
			count++
		}
	}

	return count + 1
}

func TestJump(t *testing.T) {
	assert := assert.New(t)

	testcases := []struct {
		Name   string
		Input  []int
		Output int
	}{
		{
			Name:   "test-01",
			Input:  []int{0, 0, 0, 0, 1, 0, 0},
			Output: 4,
		},
		{
			Name:   "test-02",
			Input:  []int{0, 0, 1, 0, 1, 0, 0},
			Output: 4,
		},
		{
			Name:   "test-03",
			Input:  []int{0, 0, 1, 0, 1, 0, 1, 0, 0, 0},
			Output: 5,
		},
	}

	for _, val := range testcases {
		result := jump(val.Input)
		assert.Equal(val.Output, result)
	}
}
