package ginicoef

import (
	"testing"
)

func TestGiniCoefficient(t *testing.T) {
	x := GiniCoefficient([]float64{1, 2, 3, 4})
	expect := 0.25
	if x == expect {
		t.Logf("success, %v", x)
	} else {
		t.Errorf("error, expect:%v, got:%v\n", expect, x)
	}
}
