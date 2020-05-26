// Package brow implements browser functions.
package brow

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

// Cfg is a browser configuration.
type Cfg struct {
	InputFile string `yaml:"InputFile"`
	OutputDir string `yaml:"OutputDir"`
	Pages     []Page `yaml:"Pages"`
}

// Page is a browser output page.
type Page struct {
	Name     string `yaml:"Name"`
	Filename string `yaml:"Filename"`
	Filter   string `yaml:"Filter"`
	Ext      string `yaml:"Ext"` // "" means all extensions
}

// Rec is a browser input record.
type Rec struct {
	Link      string
	Thumbnail string
	Attr      []string
}

// ReadConfig reads a browser configuration from io.Reader r.
func ReadConfig(r io.Reader) (*Cfg, error) {
	c := new(Cfg)

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return c, fmt.Errorf("readConfig read: %v", err)
	}

	if err := yaml.Unmarshal(b, c); err != nil {
		return c, fmt.Errorf("readConfig %v", err)
	}
	return c, nil
}

// ReadData reads from the input data file given in configuration c,
// and returns a slice of browser records.
func ReadData(c *Cfg) []Rec {
	f, err := os.Open(c.InputFile)
	if err != nil {
		log.Fatalf("openInputFile: %s: %v", c.InputFile, err)
	}
	defer f.Close()

	cr := csv.NewReader(f)
	return parseData(cr)
}
func parseData(cr *csv.Reader) []Rec {
	rs, err := cr.ReadAll()
	if err != nil {
		log.Fatalf("parseData: %v", err)
	}
	op := []Rec{}
	for _, r := range rs[1:] { // skip header
		rec := Rec{Link: r[0], Thumbnail: r[1]}
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

// Filter filters recs to return only requested records.
func Filter(recs []Rec, p Page) []Rec {
	recs = mediaFilter(recs, p)
	recs = attrFilter(recs, p)
	return recs
}
func mediaFilter(recs []Rec, p Page) []Rec {
	op := []Rec{}
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
func attrFilter(recs []Rec, p Page) []Rec {
	if (p.Filter == "" || p.Filter == "ALL") && p.Ext != "" {
		return recs
	}

	op := []Rec{}
	cond := p.Name
	if p.Ext != "" {
		cond = p.Filter
	}
	for _, r := range recs {
		for _, a := range r.Attr {
			if strings.Contains(strings.ToLower(a), cond) {
				op = append(op, r)
				continue
			}
		}
	}
	return op
}

// Rating rates an image.
type Rating struct {
	Link string
	Val  string
}

// ImageRating extracts image ratings from recs.
func ImageRating(recs []Rec) []Rating {
	rat := []Rating{}
	page := Page{Ext: ".jpg", Filter: "ALL"}
	recs = Filter(recs, page)
	for _, r := range recs {
		rtg := rating(r)
		if rtg != "" {
			rat = append(rat, Rating{Link: r.Link, Val: rtg})
		}
	}
	return rat
}
func rating(rec Rec) string {
	for _, a := range rec.Attr {
		matched, err := regexp.MatchString(`f\d`, a)
		if err != nil {
			log.Printf("rating: %v: %v", rec, err)
		}
		if matched {
			return a
		}
	}
	return ""
}
