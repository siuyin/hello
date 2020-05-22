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
	createPages(cfg, recs)
	writeRatings(cfg, recs)
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

func createPages(c *brow.Cfg, recs []brow.Rec) {
	if err := os.MkdirAll(c.OutputDir, 0700); err != nil {
		panic(fmt.Sprintf("createPages: %s: %v", c.OutputDir, err))
	}
	var wg sync.WaitGroup
	for _, p := range c.Pages {
		wg.Add(1)
		createPage(c, recs, p, &wg)
	}
	wg.Wait()
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
  {{range .Cfg.Pages}} {{if eq .Name $.CurrentPage.Name}} {{.Name}} {{else}} <a href="{{.Filename}}">{{.Name}}</a> {{end}}{{end}}
</nav>

{{range $index, $element := .Recs}}
  <div class="entry">
    {{$index}}.
    <a href="{{.Thumbnail}}"><img src="{{.Thumbnail}}" height="150px"/></a><br>
    {{.Link}}<br>
    {{range .Attr -}}<span class="attr">{{.}}</span>{{- end -}}
  </div>
{{end}}

</body>

</html>
`

func createPage(c *brow.Cfg, recs []brow.Rec, p brow.Page, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()

		f, err := os.Create(filepath.Join(c.OutputDir, p.Filename))
		if err != nil {
			log.Fatalf("createPage: %v: %v", p, err)
		}
		defer f.Close()

		t := template.Must(template.New("master").Parse(master))
		if err := t.Execute(f, struct {
			Cfg         *brow.Cfg
			Recs        []brow.Rec
			CurrentPage brow.Page
		}{c, brow.Filter(recs, p), p},
		); err != nil {
			log.Println(err)
		}
	}()
}

func writeRatings(cfg *brow.Cfg, recs []brow.Rec) {
	rats := brow.ImageRating(recs)
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
