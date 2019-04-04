package services

import (
	"gotest.tools/assert"
	"testing"
)

func TestMultiplyService_Multiply(t *testing.T) {
	mulSvc := MultiplyService{
		X: 1,
		Y: 2,
	}

	act := mulSvc.Multiply()
	assert.Equal(t, act, int64(2))
}
