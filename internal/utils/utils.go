package utils

type Data map[string]interface{}

func Sum(values []float64) float64 {
	var total float64
	for _, v := range values {
		total += v
	}
	return total
}

func Count(values []float64) float64 {
	return float64(len(values))
}

func Min(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	minVal := values[0]
	for _, v := range values {
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}

func Max(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	maxVal := values[0]
	for _, v := range values {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}
