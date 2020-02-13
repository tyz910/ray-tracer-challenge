package main

import (
	"fmt"

	"github.com/tyz910/ray-tracer-challenge/internal/canvas"
	"github.com/tyz910/ray-tracer-challenge/internal/canvas/color"
	"github.com/tyz910/ray-tracer-challenge/internal/canvas/image"
	"github.com/tyz910/ray-tracer-challenge/internal/math/tuple"
)

type projectile struct {
	Position tuple.Tuple
	Velocity tuple.Tuple
}

type environment struct {
	Gravity tuple.Tuple
	Wind    tuple.Tuple
}

func tick(env environment, proj projectile) projectile {
	pos := proj.Position.Add(proj.Velocity)
	vel := proj.Velocity.Add(env.Gravity).Add(env.Wind)

	return projectile{pos, vel}
}

func main() {
	e := environment{
		Gravity: tuple.Vector(0.0, -0.1, 0.0),
		Wind:    tuple.Vector(-0.01, 0.0, 0.0),
	}

	p := projectile{
		Position: tuple.Point(0.0, 1.0, 0.0),
		Velocity: tuple.Vector(1.0, 1.8, 0).Normalize().Mul(11.25),
	}

	c := canvas.New(900, 550)

	i := 0
	for {
		if p.Position.Y() <= 0 {
			break
		}

		fmt.Printf("#%d %v\n", i, p.Position)
		p = tick(e, p)
		i++

		c.WritePixel(p.Position.X(), float64(c.Height())-p.Position.Y(), color.White())
	}

	fmt.Printf("\nNumber of ticks: %d\n", i)

	if err := image.NewPPM(c).Save("images/ppm/projectile.ppm"); err != nil {
		fmt.Printf("failed to save image: %v", err)
	}
}
