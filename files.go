package main

import (
	"fmt"
	"os"
	"strings"
)

const inputDir = "input/"
const outputDir = "output/"
const compareDir = "compare/"

func ReadAll() (map[string][]byte, error) {
	dir, err := os.ReadDir(inputDir)
	if err != nil {
		return nil, err
	}

	data := map[string][]byte{}
	for _, file := range dir {
		if !file.IsDir() && strings.Contains(file.Name(), "svg") {
			fileData, err := os.ReadFile(fmt.Sprintf("%s%s", inputDir, file.Name()))
			if err != nil {
				return nil, err
			}
			data[file.Name()] = fileData
		}
	}
	return data, nil
}

func CopyOverComparisons() error {
	dir, err := os.ReadDir(compareDir)
	if err != nil {
		return err
	}

	for _, file := range dir {
		if !file.IsDir() && strings.Contains(file.Name(), "png") {
			fileData, err := os.ReadFile(fmt.Sprintf("%s%s", compareDir, file.Name()))
			if err != nil {
				return err
			}
			err = os.WriteFile(fmt.Sprintf("%s%s", outputDir, file.Name()), fileData, 0600)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func ClearOutput() error {
	err := os.RemoveAll(outputDir)
	if err != nil {
		return err
	}
	return os.Mkdir(outputDir, 0777)
}

func Save(testerName, testName string, data []byte) error {
	if data == nil {
		return nil
	}
	err := os.WriteFile(fmt.Sprintf("%s%s_%s.png", outputDir, testName[:len(testName)-4], testerName), data, 0600)
	if err != nil {
		return err
	}
	return nil
}

func SaveErr(testerName, testName string, err error) error {
	err2 := os.WriteFile(fmt.Sprintf("%s%s_%s_error.txt", outputDir, testName[:len(testName)-4], testerName), []byte(fmt.Sprintf("failed to run test: %v", err)), 0600)
	if err2 != nil {
		return err2
	}
	return nil
}
