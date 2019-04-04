package services

import (
	"gotest.tools/assert"
	"testing"
)

func TestSumService_Sum(t *testing.T) {
	sumSvc := SumService{
		X: 1,
		Y: 2,
	}
	act := sumSvc.Sum()
	assert.Equal(t, act, int64(3))
}
