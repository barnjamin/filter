package kalman

type Kalman struct {
	Q                float64
	Gain             float64 // Gain
	MeasurementError float64 //estimate of measurement variance
	EstimateValue    float64
	EstimateError    float64
}

func New(q, r, lastErr, lastVal float64) *Kalman {
	return &Kalman{
		Q:                q,
		MeasurementError: r,
		EstimateValue:    lastVal,
		EstimateError:    lastErr,
	}
}

func (k *Kalman) Feed(measurement float64) float64 {
	k.Gain = k.EstimateError / (k.EstimateError + k.MeasurementError)
	k.EstimateValue = k.EstimateValue + k.Gain*(measurement-k.EstimateValue)
	k.EstimateError = ((1 - k.Gain) * k.EstimateError) + k.Q
	return k.EstimateValue
}
