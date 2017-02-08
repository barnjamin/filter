package particle

import (
	"math/rand"
	"sort"
)

type Resampler func([]Particle) []Particle

func Multinomial(particles []Particle) []Particle {
	cs := cumsum(particles)

	p := make([]Particle, len(particles))
	for x := 0; x < len(particles); x++ {
		idx := sort.SearchFloat64s(cs, rand.Float64())
		p[x] = particles[idx]
	}

	return particles
}

func Residual(particles []Particle) []Particle {
	//N = len(weights)
	//indexes = np.zeros(N, 'i')
	//
	//# take int(N*w) copies of each weight
	//num_copies = (N*np.asarray(weights)).astype(int)
	//k = 0
	//for i in range(N):
	//for _ in range(num_copies[i]): # make n copies
	//indexes[k] = i
	//k += 1
	//
	//# use multinormial resample on the residual to fill up the rest.
	//residual = w - num_copies     # get fractional part
	//residual /= sum(residual)     # normalize
	//cumulative_sum = np.cumsum(residual)
	//cumulative_sum[-1] = 1. # ensures sum is exactly one
	//indexes[k:N] = np.searchsorted(cumulative_sum, random(N-k))
	//
	//return indexes
	return particles
}

func Systematic(particles []Particle) []Particle {
	//# make N subdivisions, choose positions
	//# with a consistent random offset
	//positions = (np.arange(N) + random()) / N

	//indexes = np.zeros(N, 'i')
	//cumulative_sum = np.cumsum(weights)
	//i, j = 0, 0
	//while i < N:
	//if positions[i] < cumulative_sum[j]:
	//indexes[i] = j
	//i += 1
	//else:
	//j += 1
	//return indexes
	return particles
}

func Stratified(particles []Particle) []Particle {
	//N = len(weights)
	//# make N subdivisions, chose a random position within each one
	//positions = (random(N) + range(N)) / N
	//
	//indexes = np.zeros(N, 'i')
	//cumulative_sum = np.cumsum(weights)
	//i, j = 0, 0
	//while i < N:
	//if positions[i] < cumulative_sum[j]:
	//indexes[i] = j
	//i += 1
	//else:
	//j += 1
	//return indexes

	return particles
}

func cumsum(p []Particle) []float64 {
	cs := 0.0

	sums := make([]float64, len(p))
	for idx, particle := range p {
		cs += particle.Weight
		sums[idx] = cs
	}
	sums[len(p)-1] = 1.0

	return sums
}
