package main

import (
	"math/rand"
	"time"

	"github.com/barnjamin/filter/kalman"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
	"github.com/labstack/gommon/log"
)

func main() {
	rand.Seed(time.Now().Unix())

	var (
		n       = 50
		xVector = vectorize(10, -0.3, n)
		z       = generateFromVector(xVector, 2, n)
	)

	k := kalman.New(1e-3, 0.01, 100, 11)

	var raw = []float64{}
	var vals = []float64{}
	for _, val := range z {
		raw = append(raw, val)

		k.Feed(val)

		vals = append(vals, k.EstimateValue)
	}

	p, err := plot.New()
	if err != nil {
		log.Fatalf("Failed to create new plot: %+v", err)
	}
	p.Title.Text = "Constant Rate Value"
	err = plotutil.AddLinePoints(p,
		"Value", plotterize(xVector),
		"Raw", plotterize(raw),
		"Corrected", plotterize(vals))

	if err != nil {
		log.Fatalf("Failed to add line points: %+v", err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(10*vg.Inch, 10*vg.Inch, "constant_rate.png"); err != nil {
		panic(err)
	}
}

func vectorize(start, velocity float64, length int) []float64 {
	vectorized := []float64{}
	for x := 0; x < length; x++ {
		vectorized = append(vectorized, start+(float64(x)*velocity))
	}
	return vectorized
}

func generateFromVector(actuals []float64, mErr float64, length int) []float64 {
	measured := []float64{}
	for _, val := range actuals {
		max, min := val+mErr, val-mErr
		delta := max - min
		measured = append(measured, (rand.Float64()*delta)+min)
	}

	return measured
}

func plotterize(vals []float64) plotter.XYs {
	plotski := make(plotter.XYs, len(vals))
	for idx, val := range vals {
		plotski[idx] = struct {
			X float64
			Y float64
		}{float64(idx), val}
	}
	return plotski
}
