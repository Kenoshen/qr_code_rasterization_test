package main

import (
	"log"
)

type SVGTester interface {
	Run(inputFilename string, data []byte) ([]byte, error)
}

func main() {
	err := ClearOutput()
	if err != nil {
		log.Fatal(err)
	}

	svgTesters := map[string]SVGTester{
		"oksvg": OkSVG{},
		//"v8":          V8{},
		//"webview":     Webview{},
		"imagemagick": ImageMagick{},
		"resvg":       ReSVG{},
		"cons2p":      ConS2P{},
	}

	input, err := ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for fileName, data := range input {
		for testerName, tester := range svgTesters {
			log.Println("RUN:", fileName, "for", testerName)
			result, err := tester.Run(fileName, data)
			if err != nil {
				err := SaveErr(testerName, fileName, err)
				if err != nil {
					log.Fatal(err)
				}
			}

			err = Save(testerName, fileName, result)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	err = CopyOverComparisons()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DONE")
}
