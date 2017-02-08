package particle

import (
	"log"
	"math/rand"
	"testing"
)

func TestNew(t *testing.T) {
	u := []float64{1, 1}
	std := []float64{.1, .1}

	pf := New(5000, []float64{100, 100}, Multinomial)
	for x := 0; x < 10; x++ {
		measurements := []float64{(rand.Float64() * 0.1) + float64(x), (rand.Float64() * 0.1) + float64(x)}
		pf.Predict(u, std)
		pf.Update(measurements, 0.1)
		pf.Resample()
		mu, est := pf.Estimate()
		log.Printf("%+v %+v", mu, est)
	}
}
