package main

import (
	"fmt"

	"github.com/tyz910/ray-tracer-challenge/internal/canvas"
	"github.com/tyz910/ray-tracer-challenge/internal/canvas/color"
	"github.com/tyz910/ray-tracer-challenge/internal/canvas/image"
	"github.com/tyz910/ray-tracer-challenge/internal/math/tuple"
	"github.com/tyz910/ray-tracer-challenge/internal/render"
	"github.com/tyz910/ray-tracer-challenge/internal/render/light"
	"github.com/tyz910/ray-tracer-challenge/internal/render/material"
	"github.com/tyz910/ray-tracer-challenge/internal/render/ray"
	"github.com/tyz910/ray-tracer-challenge/internal/render/shape/sphere"
)

func main() {
	rayOrigin := tuple.Point(0.0, 0.0, -5.0)

	wallZ := 10.0
	wallSize := 7.0
	wallHalf := wallSize / 2.0

	cnvPixels := 300
	pixelSize := wallSize / float64(cnvPixels)

	cnv := canvas.New(cnvPixels, cnvPixels)
	shape := sphere.New()

	m := material.New()
	m.SetColor(color.Magenta())
	shape.SetMaterial(m)

	lightPos := tuple.Point(-10.0, 10.0, -10.0)
	lightColor := color.White()
	l := light.New(lightPos, lightColor)

	for y := 0; y < cnvPixels; y++ {
		worldY := -(pixelSize*float64(y) - wallHalf)

		for x := 0; x < cnvPixels; x++ {
			worldX := pixelSize*float64(x) - wallHalf

			pos := tuple.Point(worldX, worldY, wallZ)
			r := ray.New(rayOrigin, pos.Sub(rayOrigin).Normalize())
			xs := shape.Intersect(r)

			if h := xs.Hit(); h != nil {
				m := h.Object().Material()
				p := r.Position(h.T())
				n := h.Object().NormalAt(p)
				eye := r.Direction().Negate()

				pixelColor := render.Lighting(m, l, p, eye, n)
				cnv.SetPixel(x, y, pixelColor)
			}
		}
	}

	if err := image.NewPPM(cnv).Save("images/ppm/shadow.ppm"); err != nil {
		fmt.Printf("failed to save image: %v", err)
	}
}
