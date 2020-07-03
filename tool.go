package ginicoef

func CumSum(itemSlice []float64) []float64 {
	sum := 0.0
	newSlice := make([]float64, len(itemSlice))
	for i, item := range itemSlice {
		sum += item
		newSlice[i] = sum
	}
	return newSlice
}

// https://en.wikipedia.org/wiki/Trapezoidal_rule
func TrapZ(xArray []float64, yArray []float64) float64 {
	var area float64
	for i := 0; i < len(xArray); i++ {
		if i > 0 {
			area += 0.5 * (xArray[i] - xArray[i-1]) *  (yArray[i-1] + yArray[i])
		}
	}
	return area
}
