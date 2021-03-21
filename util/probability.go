package util

import (
	"fmt"
	"gonum.org/v1/gonum/stat/distuv"
)

const (
	significantThreshold = 0.95
)

func IsSignificant(attempts, successes int) bool {
	binomial := distuv.Binomial{
		N: float64(attempts),
		P: 0.5,
	}

	fmt.Println(binomial.CDF(float64(successes)))
	return binomial.CDF(float64(successes)) > significantThreshold
}
