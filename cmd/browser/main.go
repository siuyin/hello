package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type cfg struct {
	InputFile string `yaml:"InputFile"`
	OutputDir string `yaml:"OutputDir"`
	Pages     []page `yaml:"Pages"`
}
type page struct {
	Name     string `yaml:"Name"`
	Filename string `yaml:"Filename"`
	Filter   string `yaml:"Filter"`
	Ext      string `yaml:"Ext"` // "" means all extensions
}
type rec struct {
	Link      string
	Thumbnail string
	Attr      []string
}

func main() {
	if !checkUsage() {
		return
	}
	cfg := readConfig(os.Args[1])
	recs := readData(cfg)
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

func readConfig(path string) *cfg {
	c := new(cfg)
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("readConfig open: %s: %v", path, err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("readConfig read: %s: %v", path, err)
	}

	if err := yaml.Unmarshal(b, c); err != nil {
		log.Fatalf("readConfig %s: %v", path, err)
	}
	return c
}

func readData(c *cfg) []rec {
	f, err := os.Open(c.InputFile)
	if err != nil {
		log.Fatalf("openInputFile: %s: %v", c.InputFile, err)
	}
	defer f.Close()

	cr := csv.NewReader(f)
	return parseData(cr)
}
func parseData(cr *csv.Reader) []rec {
	rs, err := cr.ReadAll()
	if err != nil {
		log.Fatalf("parseData: %v", err)
	}
	op := []rec{}
	for _, r := range rs[1:] { // skip header
		rec := rec{Link: r[0], Thumbnail: r[1]}
		rec.Attr = attrs(r)
		op = append(op, rec)
	}
	return op
}
func attrs(r []string) []string {
	op := []string{}
	for _, e := range r[2:] {
		if len(e) == 0 {
			continue
		}
		op = append(op, e)
	}
	return op
}

func createPages(c *cfg, recs []rec) {
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

func createPage(c *cfg, recs []rec, p page) {
	f, err := os.Create(filepath.Join(c.OutputDir, p.Filename))
	if err != nil {
		log.Fatalf("createPage: %v: %v", p, err)
	}
	defer f.Close()

	t := template.Must(template.New("master").Parse(master))
	if err := t.Execute(f, struct {
		Cfg         *cfg
		Recs        []rec
		CurrentPage page
	}{c, filter(recs, p), p},
	); err != nil {
		log.Println(err)
	}
}
func filter(recs []rec, p page) []rec {
	recs = mediaFilter(recs, p)
	recs = attrFilter(recs, p)
	return recs
}
func mediaFilter(recs []rec, p page) []rec {
	op := []rec{}
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
func attrFilter(recs []rec, p page) []rec {
	if (p.Filter == "" || p.Filter == "ALL") && p.Ext != "" {
		return recs
	}

	op := []rec{}
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
