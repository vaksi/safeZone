package services

import (
	"fmt"
	"gotest.tools/assert"
	"testing"
)

func TestPrimeService_Prime(t *testing.T) {
	primeSvc := &PrimeService{
		N:	int64(4),
	}

	act := primeSvc.Prime()

	assert.Equal(t, fmt.Sprintf("%v",act), fmt.Sprintf("%v",[]int64{ 2, 3, 5, 7}))
}
