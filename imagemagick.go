package main

import (
	"fmt"
	"os/exec"
)

// https://imagemagick.org/index.php

type ImageMagick struct {
}

func (v ImageMagick) Run(inputFilename string, data []byte) ([]byte, error) {
	cmd := exec.Command("convert", fmt.Sprintf("input/%s", inputFilename), fmt.Sprintf("output/%s_imagemagick.png", inputFilename[:len(inputFilename)-4]))
	err := cmd.Run()

	if err != nil {
		return nil, err
	}
	return nil, nil
}
