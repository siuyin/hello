package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"text/template"
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
		rec.Attr = append(rec.Attr, r[2:]...)
		op = append(op, rec)
	}
	return op
}
func writeHTML(recs []rec) {
	w, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalf("writeHTML: %v", err)
	}
	t := template.Must(template.New("output").Parse(`
<!DOCTYPE html>
<html>
<head>
<title>Browse</title>
<style>
body { font-size: 1em; font-family: Arial, Helvetica, sans-serif; }
.entry {margin-left: 0.5em; margin-bottom: 2em; }
</style>
</head>

<body>
<h1>Listing</h1>
{{range .}}
  <div class="entry">
    <img src="{{.Thumbnail}}" height="150px"/><br>
    {{.Link}}<br>
    {{range .Attr}} {{.}} {{end}}
  </div>
{{end}}
<table>
</table>
</body>

</html>
`))
	t.Execute(w, recs)
}
