package main

import (
	"math"
	"math/rand"

	"github.com/barnjamin/filter/kalman"
	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
	"github.com/labstack/gommon/log"
)

func main() {
	var (
		n    = 50
		xVal = -0.37727
		x    = repeat(xVal, n)
		z    = generateRandomData(xVal-.2, xVal+.2, n)
	)

	k := kalman.New(1e-5, math.Pow(0.01, 2), 0, 1)

	var raw = []float64{}
	var vals = []float64{}
	for _, val := range z {
		raw = append(raw, val)
		vals = append(vals, k.Feed(val))
	}

	p, err := plot.New()
	if err != nil {
		log.Fatalf("Failed to create new plot: %+v", err)
	}
	p.Title.Text = "Static True Value"
	err = plotutil.AddLinePoints(p,
		"Value", plotterize(x),
		"Raw", plotterize(raw),
		"Corrected", plotterize(vals))
	if err != nil {
		log.Fatalf("Failed to add line points: %+v", err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(10*vg.Inch, 10*vg.Inch, "static_points.png"); err != nil {
		panic(err)
	}
}

func repeat(val float64, times int) []float64 {
	repeated := []float64{}
	for x := 0; x < times; x++ {
		repeated = append(repeated, val)
	}
	return repeated
}

func generateRandomData(min, max float64, length int) []float64 {
	if min > max {
		return []float64{}
	}

	delta := math.Abs(max - min)
	randoms := []float64{}
	for x := 0; x < length; x++ {
		randoms = append(randoms, (rand.Float64()*delta)+min)
	}
	return randoms

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
