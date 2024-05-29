package utils

import (
	"math"
)

func RoundToNDecimals(value float64, decimals int) float64 {
	factor := math.Pow(10, float64(decimals))
	return math.Round(value*factor) / factor
}

func RoundTo2Decimals(value float64) float64 {
	return RoundToNDecimals(value, 2)
}
