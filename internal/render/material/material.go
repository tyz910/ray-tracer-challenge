package material

import "github.com/tyz910/ray-tracer-challenge/internal/canvas/color"

// Material encapsulates surface color and four attributes from the Phong reflection model.
type Material struct {
	color color.Color

	ambient   float64
	diffuse   float64
	specular  float64
	shininess float64
}

// New creates new material.
func New() Material {
	return Material{
		color: color.White(),

		ambient:   0.1,
		diffuse:   0.9,
		specular:  0.9,
		shininess: 200.0,
	}
}

// Ambient returns the ambient reflection. It is background lighting,
// or light reflected from other objects in the environment.
func (m Material) Ambient() float64 {
	return m.ambient
}

// Diffuse returns the diffuse reflection. It is light reflected from a matte surface.
func (m Material) Diffuse() float64 {
	return m.diffuse
}

// Specular returns the specular reflection. It is the reflection of the light source itself
// and results in what is called a specular highlight — the bright spot on a curved surface.
func (m Material) Specular() float64 {
	return m.specular
}

// Shininess returns the shininess of the material.
// The higher the shininess, the smaller and tighter the specular highlight.
func (m Material) Shininess() float64 {
	return m.shininess
}

// Color returns the surface color.
func (m Material) Color() color.Color {
	return m.color
}

// SetAmbient changes the ambient reflection of the material.
func (m *Material) SetAmbient(ambient float64) {
	m.ambient = ambient
}

// SetDiffuse changes the diffuse reflection of the material.
func (m *Material) SetDiffuse(diffuse float64) {
	m.diffuse = diffuse
}

// SetSpecular changes the specular reflection of the material.
func (m *Material) SetSpecular(specular float64) {
	m.specular = specular
}

// SetShininess changes the shininess of the material.
func (m *Material) SetShininess(shininess float64) {
	m.shininess = shininess
}

// SetColor changes the surface color of the material.
func (m *Material) SetColor(c color.Color) {
	m.color = c
}
