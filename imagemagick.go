package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

// https://imagemagick.org/index.php

type ImageMagick struct {
}

func (v ImageMagick) Run(inputFilename string, data []byte) ([]byte, error) {
	outputFilename := fmt.Sprintf("output/%s_imagemagick.png", inputFilename[:len(inputFilename)-4])
	cmd := exec.Command("convert", fmt.Sprintf("input/%s", inputFilename), outputFilename)
	err := cmd.Run()

	if err != nil {
		return nil, err
	}
	f, err := os.Open(outputFilename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(f)
}
