package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
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

type pageSet struct {
	cfg         *brow.Cfg
	recs        []brow.Rec
	page        brow.Page
	recsPerPage int
}

func createPage(cfg *brow.Cfg, recs []brow.Rec, page brow.Page, wg *sync.WaitGroup) {
	const numPerPage = 100
	go func() {
		defer wg.Done()

		ps := newPageSet(cfg, recs, page, numPerPage)
		recSet, linkSet := ps.paginate()
		for i := 0; i < len(linkSet); i++ {
			f := ps.createFile(linkSet[i])
			defer f.Close()

			ps.writeOutput(f, recSet[i])
		}
	}()
}
func newPageSet(cfg *brow.Cfg, recs []brow.Rec, page brow.Page, numPerPage int) *pageSet {
	return &pageSet{cfg, recs, page, numPerPage}
}
func (ps *pageSet) paginate() ([][]brow.Rec, []string) {
	recSet := [][]brow.Rec{}
	recs := ps.filteredRecs()
	pls := ps.pageLinks()
	for pgNum := range pls {
		recSet = append(recSet, ps.recRange(pgNum, recs))
	}
	return recSet, pls
}
func (ps *pageSet) recRange(pgNum int, recs []brow.Rec) []brow.Rec {
	startIndex := pgNum * ps.recsPerPage
	endIndex := min((pgNum+1)*ps.recsPerPage, len(recs))
	return recs[startIndex:endIndex]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func (ps *pageSet) createFile(fn string) *os.File {
	f, err := os.Create(filepath.Join(ps.cfg.OutputDir, fn))
	if err != nil {
		log.Fatalf("createFile: %v: %v", fn, err)
	}
	return f
}
func (ps *pageSet) writeOutput(f *os.File, recs []brow.Rec) {
	const master = `<!DOCTYPE html>
<html>
<head>
<title>Browse</title>
<style>
  body { font-size: 1em; font-family: Arial, Helvetica, sans-serif; }
  nav { margin-bottom: 1em; }
  .entry {margin-left: 0.5em; margin-bottom: 2em; }
  .attr { background-color: LightBlue; margin:0.1em; padding:0.2em; border-radius: 10px; }
  .page-links { margin: 1em; }
  .page-link { margin: 0.5em; }
</style>
</head>

<body>
<h1>Listing</h1>
{{define "nav"}}
  <nav>
    {{ range .Cfg.Pages -}}
      {{if eq .Name $.CurrentPage.Name}}{{.Name}}
      {{else}}<a href="{{.Filename}}">{{.Name}}</a>{{end}}
    {{ end }}
  </nav>
{{end}}
{{template "nav" .}}

{{define "pageLinks"}}
  <div class="page-links">
    page:
    {{range $index, $element := .PageLinks}}
      {{if eq $.PageNum $index}}<span class="page-link">{{$index}}</span>
      {{else}}<a class="page-link" href="{{.}}">{{$index}}</a>{{end}}
    {{end}}
  </div>
{{end}}
{{template "pageLinks" .}}

{{range $index, $element := .Recs}}
  <div class="entry">
    {{$index}}.
    <a href="{{.Thumbnail}}" target="_blank"><img src="{{.Thumbnail}}" height="150px"/></a><br>
    {{.Link}}<br>
    {{range .Attr -}}<span class="attr">{{.}}</span>{{- end -}}
  </div>
{{end}}

{{template "pageLinks" .}}
{{template "nav" .}}

</body>
</html>`
	t := template.Must(template.New("master").Parse(master))
	err := t.Execute(f, struct {
		Cfg         *brow.Cfg
		Recs        []brow.Rec
		CurrentPage brow.Page
		PageNum     int
		PageLinks   []string
	}{
		ps.cfg,
		recs,
		ps.page,
		pageNum(f),
		ps.pageLinks(),
	})
	if err != nil {
		log.Println(err)
	}
}
func pageNum(f *os.File) int {
	re := regexp.MustCompile(`-(\d{1,}).html$`)
	matches := re.FindStringSubmatch(f.Name())
	if matches == nil {
		return 0 // page zero has no -n suffix.
	}
	n, err := strconv.Atoi(matches[1]) // matches[0] is the whole string.
	if err != nil {
		panic(fmt.Sprintf("bad page number in file %s", f.Name()))
	}
	return n
}
func (ps *pageSet) pageLinks() []string {
	recs := ps.filteredRecs()
	ret := []string{ps.page.Filename}
	if len(recs) < ps.recsPerPage {
		return ret
	}
	var i int
	for i = 1; i < len(recs)/ps.recsPerPage; i++ { // i := 1 as zero case handled above.
		ret = append(ret, linkFilename(ps.page.Filename, i))
	}
	if len(recs)%ps.recsPerPage > 0 {
		ret = append(ret, linkFilename(ps.page.Filename, i)) // take care of last page
	}
	return ret
}
func (ps *pageSet) filteredRecs() []brow.Rec {
	return brow.Filter(ps.recs, ps.page)
}
func linkFilename(fn string, index int) string {
	return fmt.Sprintf("%s-%d.html", strings.Split(filepath.Base(fn), ".")[0], index)
}

func writeRatings(cfg *brow.Cfg, recs []brow.Rec, mainWG *sync.WaitGroup) {
	go func() {
		defer mainWG.Done()

		createOutputDir(cfg)

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
