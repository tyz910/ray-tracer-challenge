package main

import (
	"fmt"

	"github.com/tyz910/ray-tracer-challenge/internal/canvas"
	"github.com/tyz910/ray-tracer-challenge/internal/canvas/image"
	"github.com/tyz910/ray-tracer-challenge/internal/color"
	"github.com/tyz910/ray-tracer-challenge/internal/ray"
	"github.com/tyz910/ray-tracer-challenge/internal/sphere"
	"github.com/tyz910/ray-tracer-challenge/internal/tuple"
)

func main() {
	rayOrigin := tuple.Point(0.0, 0.0, -5.0)

	wallZ := 10.0
	wallSize := 7.0
	wallHalf := wallSize / 2.0

	cnvPixels := 300
	pixelSize := wallSize / float64(cnvPixels)
	pixelColor := color.Red()

	cnv := canvas.New(cnvPixels, cnvPixels)
	shape := sphere.New()

	for y := 0; y < cnvPixels; y++ {
		worldY := -(pixelSize*float64(y) - wallHalf)

		for x := 0; x < cnvPixels; x++ {
			worldX := pixelSize*float64(x) - wallHalf

			pos := tuple.Point(worldX, worldY, wallZ)
			r := ray.New(rayOrigin, pos.Sub(rayOrigin).Normalize())
			xs := shape.Intersect(r)

			if h := xs.Hit(); h != nil {
				cnv.SetPixel(x, y, pixelColor)
			}
		}
	}

	if err := image.NewPPM(cnv).Save("images/ppm/shadow.ppm"); err != nil {
		fmt.Printf("failed to save image: %v", err)
	}
}
