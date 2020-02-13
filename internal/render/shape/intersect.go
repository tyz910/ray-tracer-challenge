package shape

// Intersection aggregates the t value of the intersection, and the object that was intersected.
type Intersection struct {
	t   float64
	obj Shape
}

// Intersections is a collection of intersections.
type Intersections []*Intersection

// NewIntersection creates new intersection.
func NewIntersection(t float64, obj Shape) *Intersection {
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
func (i *Intersection) Object() Shape {
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
