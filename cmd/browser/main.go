package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/siuyin/hello/brow"
)

func main() {
	if !checkUsage() {
		return
	}
	cfg := readConfig(os.Args[1])
	recs := brow.ReadData(cfg)
	createPages(cfg, recs)
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
		log.Fatalf("createPages: %s: %v", c.OutputDir, err)
	}
	for _, p := range c.Pages {
		createPage(c, recs, p)
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

func createPage(c *brow.Cfg, recs []brow.Rec, p brow.Page) {
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
	}{c, filter(recs, p), p},
	); err != nil {
		log.Println(err)
	}
}
func filter(recs []brow.Rec, p brow.Page) []brow.Rec {
	recs = mediaFilter(recs, p)
	recs = attrFilter(recs, p)
	return recs
}
func mediaFilter(recs []brow.Rec, p brow.Page) []brow.Rec {
	op := []brow.Rec{}
	if (p.Filter == "" || p.Filter == "ALL") && (p.Ext == "" || p.Ext == "ALL") {
		return recs
	}
	for _, r := range recs {
		if strings.ToLower(filepath.Ext(r.Link)) != p.Ext {
			continue
		}
		op = append(op, r)
	}
	return op
}
func attrFilter(recs []brow.Rec, p brow.Page) []brow.Rec {
	if (p.Filter == "" || p.Filter == "ALL") && p.Ext != "" {
		return recs
	}

	op := []brow.Rec{}
	cond := p.Name
	if p.Ext != "" {
		cond = p.Filter
	}
	for _, r := range recs {
		for _, a := range r.Attr {
			if strings.Contains(a, cond) {
				op = append(op, r)
				continue
			}
		}
	}
	return op
}
