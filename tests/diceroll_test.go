package tests

import (
	"testing"

	. "example.com/gottrpg/v2/simulation"
)

func TestDiceRoll(t *testing.T) {
	die, i := Icosahedron, 30
	for i > 0 {
		roll := DiceRoll(len(die), 1)

		// ensure roll is not higher than the max value of the provided die
		if roll > die[len(die)-1] {
			t.Fail()
		}
		// ensure roll is not lower than the min value of the provided die
		if roll < die[0] {
			t.Fail()
		}
		i--
	}
}

func TestDistributions(t *testing.T) {
	i := 3000
	die := Icosahedron
	sig := float64(0)
	distribution := []int64{}

	mean, sd, variance := PolyhedralBaseStats(int64(len(die)))

	for i > 0 {
		distribution = append(distribution, DiceRoll(len(die), 1))
		i--
	}

	sampleMean := Mean(distribution)
	sampleSd, sampleVariance := StandardDeviation(die)

	if sampleMean > mean {
		sig = 1 - (mean / sampleMean)
	} else {
		sig = 1 - (sampleMean / mean)
	}

	// ensure the means are within .05 of each other
	if sig > .05 {
		t.Fail()
	}

	// with a large enough sample size the standard deviations should be equivalent
	if sd != sampleSd {
		t.Fail()
	}

	// results may ... vary
	if variance != sampleVariance {
		t.Fail()
	}
}

/** Some quick function checks against Hexahedrons specifically */

func TestMean(t *testing.T) {
	if Mean(Hexahedron) != 3.5 {
		t.Fail()
	}
}

func TestSigma(t *testing.T) {
	if Sigma(Hexahedron, nil) != 21 {
		t.Fail()
	}
}

func TestSigmaModifier(t *testing.T) {
	mean := float64(Mean(Hexahedron))
	numerator := Sigma(Hexahedron, func(n int64) float64 {
		val := float64(n) - mean
		return val * val
	})
	if numerator != 17.5 {
		t.Fail()
	}
}

func TestStandardDeviation(t *testing.T) {
	sd, _ := StandardDeviation(Hexahedron)
	if sd != 1.707825127659933 {
		t.Fail()
	}
}
