package main

import (
	"bytes"
	"fmt"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
	"image"
	"image/png"
)

type OkSVG struct {
	width int
}

func (v OkSVG) Run(inputFilename string, data []byte) ([]byte, error) {
	buf := bytes.NewBuffer(data)
	icon, err := oksvg.ReadIconStream(buf, oksvg.IgnoreErrorMode)
	if err != nil {
		return nil, err
	}
	w := icon.ViewBox.W
	h := icon.ViewBox.H
	if w <= 0 || h <= 0 {
		return nil, fmt.Errorf("parsed svg resulted in a viewBox with 0 width or height")
	}
	ratio := h / w

	width := v.width
	height := int(float64(width) * ratio)

	icon.SetTarget(0, 0, float64(width), float64(height))
	rgba := image.NewRGBA(image.Rect(0, 0, width, height))
	icon.Draw(rasterx.NewDasher(width, height, rasterx.NewScannerGV(width, height, rgba, rgba.Bounds())), 1)

	wBuf := bytes.NewBuffer(nil)
	err = png.Encode(wBuf, rgba)
	if err != nil {
		return nil, err
	}

	return wBuf.Bytes(), nil
}
