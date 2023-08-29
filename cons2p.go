package main

import (
	"bytes"
	"os/exec"
)

// https://www.npmjs.com/package/convert-svg-to-png

type ConS2P struct {
}

func (v ConS2P) Run(inputFilename string, data []byte) ([]byte, error) {
	cmd := exec.Command("convert-svg-to-png")

	cmd.Stdin = bytes.NewBuffer(data)
	buf := bytes.NewBuffer(nil)
	cmd.Stdout = buf
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
