package main

import (
	"embed"
	"io"
	"log"
	"path/filepath"
	"strings"
)

type SVGTester interface {
	Run(filename string, in []byte) ([]byte, error)
}

//go:embed input/*.svg
var inputFS embed.FS

//go:embed compare/*.png
var compareFS embed.FS

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

	files, err := inputFS.ReadDir("input")
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range files {
		inputfile := filepath.Join("input", fi.Name())
		log.Printf("🌠 %s", inputfile)
		f, err := inputFS.Open(inputfile)
		if err != nil {
			log.Fatal(err)
		}
		svgBytes, err := io.ReadAll(f)
		if err != nil {
			log.Fatal(err)
		}
		f.Close()

		// load comparison
		expectFilename := filepath.Join("compare", strings.TrimSuffix(fi.Name(), ".svg")+".png")
		f, err = compareFS.Open(expectFilename)
		log.Println("file", expectFilename)
		if err != nil {
			log.Fatal(err)
		}
		comparePng, err := io.ReadAll(f)
		if err != nil {
			log.Fatal(err)
		}
		f.Close()

		for testerName, tester := range svgTesters {
			log.Println(testerName)
			result, err := tester.Run(fi.Name(), svgBytes)
			if err != nil {
				log.Printf("🚧 %s %s %s", fi.Name(), testerName, err)
				err := SaveErr(testerName, fi.Name(), err)
				if err != nil {
					log.Fatal(err)
				}
				continue
			}

			similar, err := Compare(result, comparePng)
			if err != nil {
				log.Printf("⚠️🤷‍♂️ %s %s %s", fi.Name(), testerName, err)
				continue
			}
			if similar {
				log.Printf("✅ %s %s", testerName, fi.Name())
			} else {
				log.Printf("❌ %s %s", testerName, fi.Name())
			}

			err = Save(testerName, fi.Name(), result)
			if err != nil {
				log.Fatal(err)
			}
		}

	}
	log.Println("DONE")
}
