package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/siuyin/hello/brow"
)

func main() {
	if !checkUsage() {
		return
	}
	cfg := readConfig(os.Args[1])
	recs := brow.ReadData(cfg)

	var wg sync.WaitGroup
	wg.Add(2)
	createPages(cfg, recs, &wg)
	writeRatings(cfg, recs, &wg)
	wg.Wait()
}

func checkUsage() bool {
	if len(os.Args) != 2 {
		fmt.Printf(`Usage: browser <config.yaml> 
eg.
  browser config.yaml

`)
		return false
	}
	return true
}

func readConfig(path string) *brow.Cfg {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("readConfig open: %s: %v", path, err)
	}
	defer f.Close()

	cfg, err := brow.ReadConfig(f)
	if err != nil {
		panic(err)
	}
	return cfg
}

func createPages(c *brow.Cfg, recs []brow.Rec, mainWG *sync.WaitGroup) {
	defer mainWG.Done()

	createOutputDir(c)

	var wg sync.WaitGroup
	for _, p := range c.Pages {
		wg.Add(1)
		createPage(c, recs, p, &wg)
	}
	wg.Wait()
}
func createOutputDir(c *brow.Cfg) {
	if err := os.MkdirAll(c.OutputDir, 0700); err != nil {
		panic(fmt.Sprintf("createOutputDir: %s: %v", c.OutputDir, err))
	}
}

func createPage(cfg *brow.Cfg, recs []brow.Rec, page brow.Page, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()

		f := createFile(cfg, page)
		defer f.Close()

		writeOutput(f, cfg, recs, page)
	}()
}
func createFile(c *brow.Cfg, p brow.Page) *os.File {
	f, err := os.Create(filepath.Join(c.OutputDir, p.Filename))
	if err != nil {
		log.Fatalf("createPage: %v: %v", p, err)
	}
	return f
}
func writeOutput(f *os.File, cfg *brow.Cfg, recs []brow.Rec, page brow.Page) {
	t := template.Must(template.New("master").Parse(master))
	err := t.Execute(f, struct {
		Cfg         *brow.Cfg
		Recs        []brow.Rec
		CurrentPage brow.Page
	}{
		cfg,
		brow.Filter(recs, page),
		page,
	})
	if err != nil {
		log.Println(err)
	}
}

const master = `<!DOCTYPE html>
<html>
<head>
<title>Browse</title>
<style>
  body { font-size: 1em; font-family: Arial, Helvetica, sans-serif; }
  nav { margin-bottom: 1em; }
  .entry {margin-left: 0.5em; margin-bottom: 2em; }
  .attr { background-color: LightBlue; margin:0.1em; padding:0.2em; border-radius: 10px; }
</style>
</head>

<body>
<h1>Listing</h1>
<nav>
  {{ range .Cfg.Pages -}}
    {{if eq .Name $.CurrentPage.Name}} {{.Name}} {{else}} <a href="{{.Filename}}">{{.Name}}</a> {{end}}
  {{ end }}
</nav>

{{range $index, $element := .Recs}}
  <div class="entry">
    {{$index}}.
    <a href="{{.Thumbnail}}" target="_blank"><img src="{{.Thumbnail}}" height="150px"/></a><br>
    {{.Link}}<br>
    {{range .Attr -}}<span class="attr">{{.}}</span>{{- end -}}
  </div>
{{end}}

</body>

</html>`

func writeRatings(cfg *brow.Cfg, recs []brow.Rec, mainWG *sync.WaitGroup) {
	go func() {
		defer mainWG.Done()

		rats := brow.ImageRating(recs)
		write(cfg, rats)
	}()
}
func write(cfg *brow.Cfg, rats []brow.Rating) {
	f, err := os.Create(filepath.Join(cfg.OutputDir, "ratings.csv"))
	if err != nil {
		log.Fatalf("writeRatings: %v", err)
	}
	defer f.Close()

	cw := csv.NewWriter(f)
	defer cw.Flush()

	cw.Write([]string{"Link", "Rating"})
	for _, r := range rats {
		cw.Write([]string{r.Link, r.Val})
	}
}
