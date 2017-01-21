package kalman

import (
	"kalman"
	"log"
	"math"
	"math/rand"
)

func Example() {

	var (
		q            = 1e-5
		actualValue  = -0.377
		r            = math.Pow(0.01, 2)
		initialError = 0
		initialValue = 0
	)

	k := kalman.New(q, r, initialError, initialValue)

	measuredValue := func() float64 {
		return (rand.NormFloat64()*0.3 + actualValue)
	}

	for x := 0; x < 50; x++ {
		measured := measuedValue()
		estimate := k.Feed(measured)
		log.Printf("Actual: %.3f Measured: %.3f New Estimate: %.3f", actualValue, measured, estimate)
	}

}
