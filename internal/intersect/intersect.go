package intersect

import "github.com/tyz910/ray-tracer-challenge/internal/ray"

// Intersection aggregates the t value of the intersection, and the object that was intersected.
type Intersection struct {
	t   float64
	obj Intersectable
}

// Intersections is a collection of intersections.
type Intersections []*Intersection

// Intersectable is the interface implemented by objects that can be intersected.
type Intersectable interface {
	// Intersect returns the collection of intersections where the ray intersects the object.
	Intersect(r ray.Ray) Intersections
}

// New creates new intersection.
func New(t float64, obj Intersectable) *Intersection {
	return &Intersection{
		t:   t,
		obj: obj,
	}
}

// T returns the t value of the intersection.
func (i *Intersection) T() float64 {
	return i.t
}

// Object returns the object that was intersected.
func (i *Intersection) Object() Intersectable {
	return i.obj
}

// Hit returns the intersection which is actually visible from the ray’s origin.
func (xs Intersections) Hit() (h *Intersection) {
	for _, i := range xs {
		if i.t < 0 {
			continue
		}

		if h == nil || (i.t < h.t) {
			h = i
		}
	}

	return
}
