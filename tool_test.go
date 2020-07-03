package ginicoef

import "testing"

func TestTrapZ(t *testing.T) {
	xArray := []float64{0, 0.25, 0.5, 0.75, 1}
	yArray := []float64{0, 0.1, 0.3, 0.6, 1}
	x := TrapZ(xArray, yArray)
	expect := 0.375
	if x == expect {
		t.Logf("success, %v", x)
	} else {
		t.Errorf("error, expect:%v, got:%v\n", expect, x)
	}
}
