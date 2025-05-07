package utils

import (
	"fmt"
	"go.uber.org/zap"
	"math"
	"server-aggregation/pkg/log"
	"strconv"
)

var mathLogName = "utils.math"

// float64=>int
func Float64ToInt(f float64) (dst int) {
	b := fmt.Sprintf("%0.0f", f)
	dst, err := strconv.Atoi(b)
	if err != nil {
		log.New().Named(mathLogName).Error("Float64ToInt err:%s", zap.Error(err))
	}
	return
}

// int -> float64
func IntegerChangeToIntegerFloat64(value int) float64 {
	return float64(value) / math.Pow10(2)
}

func RoundToDecimal(value float64, decimalPlaces int) float64 {
	factor := math.Pow(10, float64(decimalPlaces))
	return math.Round(value*factor) / factor
}
func TruncateToDecimal(value float64, decimalPlaces int) float64 {
	factor := math.Pow(10, float64(decimalPlaces))
	return math.Trunc(value*factor) / factor
}
