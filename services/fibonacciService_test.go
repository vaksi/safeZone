package services

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFibonacciService_Fibonacci(t *testing.T) {
	fibSvc := &FibonacciService{
		N: int64(4),
	}

	act := fibSvc.Fibonacci()

	assert.Equal(t, fmt.Sprintf("%v", act), fmt.Sprintf("%v", []int64{0, 1, 1, 2}))
}
