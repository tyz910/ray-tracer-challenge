package image

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"

	"github.com/tyz910/ray-tracer-challenge/internal/canvas"
	"github.com/tyz910/ray-tracer-challenge/internal/color"
)

const rowMaxLen = 70

// PPM represents Portable Pixmap image format.
type PPM struct {
	data string
}

// NewPPM creates new PPM image.
func NewPPM(cnv canvas.Canvas) *PPM {
	builder := new(ppmBuilder)
	builder.WriteHeader(cnv.Width(), cnv.Height())

	for y := 0; y < cnv.Height(); y++ {
		for x := 0; x < cnv.Width(); x++ {
			builder.WriteColor(cnv.Pixel(x, y))
		}

		builder.WriteNewRow()
	}

	return &PPM{
		data: builder.String(),
	}
}

func (ppm *PPM) String() string {
	return ppm.data
}

// Save saves the .ppm image to disk.
func (ppm *PPM) Save(filename string) error {
	return ioutil.WriteFile(filename, []byte(ppm.data), 0644)
}

type ppmBuilder struct {
	data   strings.Builder
	rowLen int
}

// WriteHeader appends the image header.
func (pb *ppmBuilder) WriteHeader(width, height int) {
	pb.data.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", width, height))
}

func (pb *ppmBuilder) WriteNewRow() {
	pb.data.WriteString("\n")
	pb.rowLen = 0
}

func (pb *ppmBuilder) WriteColor(c color.Color) {
	pb.writeColorChannel(c.Red())
	pb.writeColorChannel(c.Green())
	pb.writeColorChannel(c.Blue())
}

func (pb *ppmBuilder) writeColorChannel(c float64) {
	cStr := strconv.Itoa(convertColorChannel(c))
	cLen := len(cStr)

	if pb.rowLen+cLen+1 > rowMaxLen {
		pb.WriteNewRow()
	}

	if pb.rowLen > 0 {
		pb.data.WriteString(" ")
		pb.rowLen++
	}

	pb.data.WriteString(cStr)
	pb.rowLen += cLen
}

func (pb *ppmBuilder) String() string {
	return pb.data.String()
}

func convertColorChannel(c float64) int {
	i := int(math.Ceil(c * 255))

	if i < 0 {
		i = 0
	}

	if i > 255 {
		i = 255
	}

	return i
}
