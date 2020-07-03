package ginicoef

import (
	"sort"
)

func GiniCoefficient(itemSlice []float64) float64 {
	itemSlice = append(itemSlice, 0)
	sort.Sort(sort.Float64Slice(itemSlice))
	yArray := CumSum(itemSlice)
	xArray := make([]float64, len(yArray))
	// Normalized
	for i := 0; i < len(yArray); i++ {
		yArray[i] = yArray[i] / yArray[len(yArray)-1]
	}

	for i := 0; i < len(xArray); i++ {
		xArray[i] = float64(i) / float64(len(xArray)-1)
	}

	B := TrapZ(xArray, yArray)
	A := 0.5 - B
	return A / (A + B)
}
