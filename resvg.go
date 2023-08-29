package main

import (
	"fmt"
	"os/exec"
)

// https://github.com/RazrFalcon/resvg

type ReSVG struct {
}

func (v ReSVG) Run(inputFilename string, data []byte) ([]byte, error) {
	cmd := exec.Command("resvg", fmt.Sprintf("input/%s", inputFilename), fmt.Sprintf("output/%s_resvg.png", inputFilename[:len(inputFilename)-4]))
	err := cmd.Run()

	if err != nil {
		return nil, err
	}
	return nil, nil
}
