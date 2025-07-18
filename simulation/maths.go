package simulation

import (
	"fmt"
	"math"

	"github.com/mazen160/go-random"
)

/** C'mon, take a chance and roll the dice */
func DiceRoll(polyhedron, quantity int) int64 {
	var total int64 = 0
	for quantity > 0 {
		result, _ := random.IntRange(1, polyhedron+1)
		total += int64(result)
		quantity--
	}
	return total
}

/** Test Logging to show mean, standard deviation, and variance */
func LogBasicStats(mean, standardDeviation, variance float64) {
	fmt.Printf("\tMean: %v\n\tSD: %v\n\tVariance: %v\n", mean, standardDeviation, variance)
}

/** Just an average function */
func Mean(nums []int64) float64 {
	return Sigma(nums, nil) / float64(len(nums))
}

/** Return the mean, standard deviation, and variation of a given polyhedral */
func PolyhedralBaseStats(sides int64) (float64, float64, float64) {
	nums := []int64{}
	sd, variance := float64(0), float64(0)

	for sides > 0 {
		nums = append(nums, sides)
		sides--
	}

	sd, variance = StandardDeviation(nums)

	return Mean(nums), sd, variance
}

/** Sum a slice of int64 values, with an optional function to control additional operations that occur during summation */
func Sigma(nums []int64, modifier func(int64) float64) float64 {
	i := float64(0)
	for _, num := range nums {
		if modifier != nil {
			i += modifier(num)
		} else {
			i += float64(num)
		}
	}
	return i
}

/** Return the standard deviation of a slice of int64s */
func StandardDeviation(nums []int64) (float64, float64) {
	mean := float64(Mean(nums))
	numerator := Sigma(nums, func(n int64) float64 {
		val := float64(n) - mean
		return val * val
	})
	variance := numerator / float64(len(nums))
	return math.Sqrt(variance), variance
}
