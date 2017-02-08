package particle

import (
	"math/rand"
	"sort"
)

type Resampler func([]Particle) []Particle

func Multinomial(particles []Particle) []Particle {
	cs := cumsum(particles)

	totalWeight := 0.0
	p := make([]Particle, len(particles))
	for x := 0; x < len(particles); x++ {
		idx := sort.SearchFloat64s(cs, rand.Float64())
		p[x] = particles[idx]
		totalWeight += p[x].Weight
	}

	for idx := range p {
		p[idx].Weight /= totalWeight
	}

	return p
}

func Residual(particles []Particle) []Particle {

	totalWeight := 0.0
	p := []Particle{}

	for x := 0; x < len(particles); x++ {
		residual := int(particles[x].Weight * float64(len(particles)))
		for y := 0; y < residual; y++ {
			p = append(p, particles[x])
			totalWeight += particles[x].Weight
		}
	}

	for idx := range p {
		p[idx].Weight /= totalWeight
	}

	mSample := Multinomial(p)

	p = append(p, mSample[:(len(particles)-len(mSample))-1]...)

	return p
}

func Stratified(particles []Particle) []Particle {

	cs := cumsum(particles)

	positions := make([]float64, len(particles))
	for x := 0; x < len(particles); x++ {
		positions[x] = (float64(x) + rand.Float64()) / float64(len(particles))
	}

	i, j := 0, 0
	p := []Particle{}
	for len(p) < len(particles) {
		if positions[i] < cs[j] {
			p = append(p, particles[j])
			i++
		} else {
			j++
		}
	}

	return p
}

func Systematic(particles []Particle) []Particle {

	cs := cumsum(particles)

	positions := make([]float64, len(particles))
	offset := rand.Float64()
	for x := 0; x < len(particles); x++ {
		positions[x] = (float64(x) + offset) / float64(len(particles))
	}

	i, j := 0, 0

	p := []Particle{}
	for len(p) < len(particles) {
		if positions[i] < cs[j] {
			p = append(p, particles[j])
			i++
		} else {
			j++
		}
	}

	return p
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
