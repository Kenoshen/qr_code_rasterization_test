package main

import (
	"embed"
	"fmt"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"
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
		"chrome":      &Chrome{},
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
			start := time.Now()
			result, err := tester.Run(fi.Name(), svgBytes)
			took := time.Since(start)
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
				log.Printf("✅ %s %s took:%s", testerName, fi.Name(), took)
			} else {
				log.Printf("❌ %s %s took:%s", testerName, fi.Name(), took)
			}

			err = Save(testerName, fi.Name(), result)
			if err != nil {
				log.Fatal(err)
			}
		}

	}
	log.Println("DONE")

	// print out markdown
	headers := []string{"name", "input", "expected"}
	for k, _ := range svgTesters {
		headers = append(headers, k)
	}
	fmt.Printf("| %s |\n", strings.Join(headers, " | "))
	col := len(headers)
	headers = []string{}
	for i := 0; i < col; i++ {
		headers = append(headers, "---")
	}
	fmt.Printf("| %s |\n", strings.Join(headers, " | "))

	for _, fi := range files {
		var row []string
		row = append(row, fi.Name())
		baseFile := strings.TrimSuffix(fi.Name(), ".svg")
		row = append(row, fmt.Sprintf("![img](%s)", filepath.Join("input", fi.Name())))
		row = append(row, fmt.Sprintf("![img](%s)", filepath.Join("compare", baseFile+".png")))
		for testerName, _ := range svgTesters {
			row = append(row, fmt.Sprintf("![%s](%s)", testerName, filepath.Join("output", baseFile+"_"+testerName+".png")))
		}
		fmt.Printf("| %s |\n", strings.Join(row, " | "))
	}
}
