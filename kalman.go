package kalman

type Kalman struct {
	Q         float64
	R         float64 //estimate of measurement variance
	K         float64 // Gain
	LastValue float64
	LastError float64
}

func New(q, r, lastErr, lastVal float64) *Kalman {
	return &Kalman{
		Q:         q,
		R:         r,
		LastValue: lastVal,
		LastError: lastErr,
	}

}

func (k *Kalman) Feed(measurement float64) float64 {
	k.K = k.LastError / (k.LastError + k.R)
	k.LastValue = k.LastValue + k.K*(measurement-k.LastValue)
	k.LastError = ((1 - k.K) * k.LastError) + k.Q
	return k.LastValue
}
