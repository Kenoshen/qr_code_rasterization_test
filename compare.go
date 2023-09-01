package main

import (
	"bytes"
	"fmt"
	"image/png"

	"github.com/vitali-fedulov/images4"
)

func Compare(a, b []byte) (bool, error) {
	aa, err := png.Decode(bytes.NewBuffer(a))
	if err != nil {
		return false, fmt.Errorf("a %w", err)
	}
	bb, err := png.Decode(bytes.NewBuffer(b))
	if err != nil {
		return false, fmt.Errorf("b %w", err)
	}
	ai, bi := images4.Icon(aa), images4.Icon(bb)
	return images4.Similar(ai, bi), nil
}
