package main

import (
	"fmt"
	"math"

	"github.com/tyz910/ray-tracer-challenge/internal/canvas"
	"github.com/tyz910/ray-tracer-challenge/internal/canvas/image"
	"github.com/tyz910/ray-tracer-challenge/internal/color"
	"github.com/tyz910/ray-tracer-challenge/internal/matrix"
	"github.com/tyz910/ray-tracer-challenge/internal/tuple"
)

func main() {
	size := 300
	points := 12

	cnv := canvas.New(size, size)
	center := tuple.Point(0.0, 0.0, 0.0)
	radius := float64(size) * 0.4
	angle := (2.0 * math.Pi) / float64(points)

	for h := 0; h < points; h++ {
		p := matrix.Transform(
			matrix.Translation(0.0, radius, 0.0),
			matrix.RotationZ(-float64(h)*angle),
			matrix.Translation(float64(cnv.Width())/2.0, float64(cnv.Height())/2.0, 0.0),
		).TupMul(center)

		fmt.Println(h, p)

		cnv.WritePixel(p.X(), float64(cnv.Width())-p.Y(), color.White())
	}

	if err := image.NewPPM(cnv).Save("images/ppm/clock.ppm"); err != nil {
		fmt.Printf("failed to save image: %v", err)
	}
}
