package main

import (
	"kalman"
	"math"
	"math/rand"

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

	k := kalman.New()

	var raw = make(plotter.XYs, n)
	var vals = make(plotter.XYs, n)
	for idx, val := range z {
		k.Feed(val)
		raw = append(raw, struct {
			X float64
			Y float64
		}{float64(idx), val})
		vals = append(vals, struct {
			X float64
			Y float64
		}{float64(idx), k.LastValue})
	}

	p, err := plot.New()
	if err != nil {
		log.Fatalf("Failed to create new plot: %+v", err)
	}
	err = plotutil.AddLinePoints(p,
		"Value", x,
		"Raw", raw,
		"Corrected", vals)
	if err != nil {
		log.Fatalf("Failed to add line points: %+v", err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

func repeat(val float64, times int) plotter.XYs {
	repeated := make(plotter.XYs, times)
	for x := 0; x < times; x++ {
		repeated[x] = struct {
			X float64
			Y float64
		}{float64(x), val}
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
