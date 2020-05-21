package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
)

type rec struct {
	Link      string
	Thumbnail string
	Attr      []string
}

func main() {
	fmt.Println("media browser")
	if !checkUsage() {
		return
	}
	cr, f := openInputFile()
	defer f.Close()

	recs := parseData(cr)
	writeHTML(recs)
}

func checkUsage() bool {
	if len(os.Args) != 3 {
		fmt.Printf(`Usage: browser <data.csv> <output.html>
eg.
  browser data.csv output.html

`)
		return false
	}
	return true
}

func openInputFile() (*csv.Reader, *os.File) {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("openInput: %v", err)
	}
	return csv.NewReader(f), f
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

func writeHTML(recs []rec) {
	writePage(recs, filter, "fav")
}
func writePage(recs []rec, filter func([]rec, string) []rec, s string) {
	recs = filter(recs, s)
	w, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalf("writeHTML: %v", err)
	}
	defer w.Close()

	tpl := `<!DOCTYPE html>
<html>
<head>
<title>Browse</title>
<style>
  body { font-size: 1em; font-family: Arial, Helvetica, sans-serif; }
  .entry {margin-left: 0.5em; margin-bottom: 2em; }
  .attr { background-color: LightBlue; margin:0.1em; padding:0.2em; border-radius: 10px; }
</style>
</head>

<body>
<h1>Listing</h1>
{{range $index, $element := .}}
  <div class="entry">
    {{$index}}.
    <a href="{{.Thumbnail}}"><img src="{{.Thumbnail}}" height="150px"/></a><br>
    {{.Link}}<br>
    {{range .Attr -}}<span class="attr">{{.}}</span>{{- end -}}
  </div>
{{end}}
<table>
</table>
</body>

</html>
`
	t := template.Must(template.New("fav").Parse(tpl))
	t.Execute(w, recs)
	fmt.Printf("Output written to %s.\n\n", os.Args[2])
}
func filter(recs []rec, s string) []rec {
	op := []rec{}
	for _, r := range recs {
		for _, a := range r.Attr {
			if strings.Contains(a, s) {
				op = append(op, r)
				continue
			}
		}
	}
	return op
}
