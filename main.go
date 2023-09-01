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
	var results []*Result
	for _, fi := range files {
		inputName := fi.Name()
		r := &Result{inputName: inputName}
		results = append(results, r)
		inputfile := filepath.Join("input", inputName)
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
			t := &Tester{
				testerName: testerName,
				took:       took,
			}
			r.testers = append(r.testers, t)
			if err != nil {
				t.errorMessage = err.Error()
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
	for _, r := range results {
		fmt.Printf("\n\n### %s\n", r.inputName)
		baseFile := strings.TrimSuffix(r.inputName, ".svg")
		fmt.Printf("`input`\n![img](%s)\n\n", filepath.Join("input", r.inputName))
		fmt.Printf("`expected output`\n![img](%s)\n\n", filepath.Join("compare", baseFile+".png"))
		for _, t := range r.testers {
			fmt.Printf("`%s - %.2fs`\n", t.testerName, t.took.Seconds())
			fmt.Printf("![%s](%s)\n\n", t.testerName, filepath.Join("output", baseFile+"_"+t.testerName+".png"))
		}
	}

	err = CopyOverComparisons()
	if err != nil {
		log.Fatal(err)
	}
}

type Result struct {
	inputName string
	testers   []*Tester
}

type Tester struct {
	testerName   string
	took         time.Duration
	errorMessage string
}
