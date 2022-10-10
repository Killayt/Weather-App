package gl

import (
	"image"
	"image/color"

	"fyne.io/fyne"
)

type captureImage struct {
	pix           []uint8
	width, height int
}

func (c *captureImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (c *captureImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, c.width, c.height)
}

func (c *captureImage) At(x, y int) color.Color {
	start := ((c.height-y-1)*c.width + x) * 4
	return color.RGBA{R: c.pix[start], G: c.pix[start+1], B: c.pix[start+2], A: c.pix[start+3]}
}

func (p *glPainter) Capture(c fyne.Canvas) image.Image {
	pos := fyne.NewPos(c.Size().Width, c.Size().Height)
	width, height := c.PixelCoordinateForPosition(pos)
	pixels := make([]uint8, width*height*4)

	p.context.RunWithContext(func() {
		p.glCapture(int32(width), int32(height), &pixels)
	})

	return &captureImage{
		pix:    pixels,
		width:  width,
		height: height,
	}
}
