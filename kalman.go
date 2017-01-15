package kalman

import "math"

type Kalman struct {
	R         float64 //estimate of measurement variance
	K         float64 // Gain
	LastValue float64
	LastError float64
}

func New() *Kalman {
	return &Kalman{
		R:         math.Pow(0.1, 2),
		K:         0.0,
		LastValue: 0.0,
		LastError: 1.0,
	}

}

func (k *Kalman) Feed(measurement float64) float64 {
	k.K = k.LastError / (k.LastError + k.R)
	k.LastValue = k.LastValue + k.K*(measurement-k.LastValue)
	k.LastError = (1 - k.K) * k.LastError
	return k.LastValue
}
