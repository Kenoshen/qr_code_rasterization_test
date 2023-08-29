package main

import (
	"bytes"
	"github.com/srwiley/rasterx"
	"image"
	"image/png"
)
import "github.com/srwiley/oksvg"

type OkSVG struct {
}

func (v OkSVG) Run(inputFilename string, data []byte) ([]byte, error) {
	buf := bytes.NewBuffer(data)
	icon, err := oksvg.ReadIconStream(buf, oksvg.IgnoreErrorMode)
	if err != nil {
		return nil, err
	}
	w := int(icon.ViewBox.W)
	h := int(icon.ViewBox.H)
	icon.SetTarget(0, 0, float64(w), float64(h))
	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	icon.Draw(rasterx.NewDasher(w, h, rasterx.NewScannerGV(w, h, rgba, rgba.Bounds())), 1)

	wBuf := bytes.NewBuffer(nil)
	err = png.Encode(wBuf, rgba)
	if err != nil {
		return nil, err
	}

	return wBuf.Bytes(), nil
}
