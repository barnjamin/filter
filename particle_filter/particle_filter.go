package particle

import (
	"math"
	"math/rand"
	"time"

	gaussian "github.com/chobie/go-gaussian"
)

/*
*Randomly generate a bunch of particles*
Particles can have position, heading, and/or whatever other state variable you need to estimate. Each has a weight (probability) indicating how likely it matches the actual state of the system. Initialize each with the same weight.

*Predict next state of the particles*
Move the particles based on how you predict the real system is behaving.

*Update*
Update the weighting of the particles based on the measurement. Particles that closely match the measurements are weighted higher than particles which don't match the measurements very well.

*Resample*
Discard highly improbable particle and replace them with copies of the more probable particles.

*Compute Estimate*
Optionally, compute weighted mean and covariance of the set of particles to get a state estimate.
*/

type ParticleFilter struct {
	N          int        // Number of particles to sample
	Dimensions []float64  // Slice of dimension max values
	Particles  []Particle // Slice of particles sampled
	resampler  Resampler
}

type Particle struct {
	Weight     float64
	Dimensions []float64
}

func init() {
	rand.Seed(time.Now().Unix())
}

func New(particleCount int, dimensions []float64, resampler Resampler) *ParticleFilter {

	//Initialize to 1/N
	weight := 1.0 / float64(particleCount)

	//Initialize Particles
	particles := make([]Particle, particleCount)
	for x := 0; x < len(particles); x++ {
		dims := make([]float64, len(dimensions))
		for idx, dim := range dimensions {
			dims[idx] = rand.Float64() * dim
		}
		particles[x] = Particle{
			Weight:     weight,
			Dimensions: dims,
		}
	}

	return &ParticleFilter{
		N:          particleCount,
		Dimensions: dimensions,
		Particles:  particles,
		resampler:  resampler,
	}
}

// Predict the next movement of the object we're tracking
// u is the process model, and it contains the predicted dimensional advancment
// std expresses the uncertainty in the process model prediction by dimension
func (p *ParticleFilter) Predict(u, std []float64) {

	if len(u) != len(std) {
		return //TODO:: Return Error
	}

	//Advance each particle
	for _, particle := range p.Particles {
		for idx := range u {
			particle.Dimensions[idx] += (rand.NormFloat64()*std[idx] + u[idx])
		}
	}

}

func (p *ParticleFilter) Update(measurements []float64, variance float64) {

	totalWeight := 0.0

	g := gaussian.NewGaussian(0.0, math.Sqrt(variance))

	//Reweight
	for idx, particle := range p.Particles {
		sum := 0.0
		for idx, m := range measurements {
			sum += math.Pow(particle.Dimensions[idx]-m, 2)
		}
		distance := math.Sqrt(sum)

		p.Particles[idx].Weight += math.Max(g.Pdf(distance), 1e-12)
		totalWeight += p.Particles[idx].Weight
	}

	//Normalize
	for idx := range p.Particles {
		p.Particles[idx].Weight /= totalWeight
	}
}

func (p *ParticleFilter) Resample() {
	if p.performResample() {
		p.Particles = p.resampler(p.Particles)
	}
}
func (p *ParticleFilter) performResample() bool {
	//TODO:: add other thresholding funcs
	sumOfSquares := 0.0
	for _, particle := range p.Particles {
		sumOfSquares += math.Pow(particle.Weight, 2)
	}

	return (1.0 / sumOfSquares) > (float64(len(p.Particles)) / 2.0)
}

func (p *ParticleFilter) Estimate() ([]float64, []float64) {

	avg := make([]float64, len(p.Dimensions))
	variance := make([]float64, len(p.Dimensions))

	for _, particle := range p.Particles {
		for idx, val := range particle.Dimensions {
			avg[idx] += (val * particle.Weight)
		}
	}

	for _, particle := range p.Particles {
		for idx, val := range particle.Dimensions {
			variance[idx] += (math.Pow(val-avg[idx], 2) * particle.Weight)
		}
	}

	return avg, variance
}
