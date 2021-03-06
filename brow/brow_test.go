package brow

import (
	"os"
	"testing"
)

func TestReadConfig(t *testing.T) {
	cfg := testCfg(t)
	if v := cfg.OutputDir; v != "bar" {
		t.Errorf("unexpected value: %v", v)
	}
	if v := cfg.Pages[1]; v.Name != "foo" {
		t.Errorf("unexpected value: %v", v.Name)
	}
}
func testCfg(t *testing.T) *Cfg {
	f, err := os.Open("testdata/sample.yaml")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	cfg, err := ReadConfig(f)
	if err != nil {
		t.Error(err)
	}
	return cfg
}

func TestReadData(t *testing.T) {
	recs := testData(t)
	if len(recs) == 0 {
		t.Errorf("records should be returned")
	}
	if v := recs[3].Link; v != "cat.avi" {
		t.Errorf("unexpected value: %v", v)
	}
}
func testData(t *testing.T) []Rec {
	cfg := testCfg(t)
	recs := ReadData(cfg)
	return recs
}

func TestFilter(t *testing.T) {
	cfg := testCfg(t)
	p := cfg.Pages[1] // cfg: Name: foo, Filename: foo.html, dat: foo.jpg,foos.jpg,fav,foo
	recs := ReadData(cfg)
	if v := len(recs); v < 4 {
		t.Errorf("unexpected value: %v", v)
	}

	orecs := Filter(recs, p)
	if v := len(orecs); v != 1 {
		t.Errorf("unexpected value: %v", v)
	}
}

func TestImageRating(t *testing.T) {
	exp := []Rating{
		{"foo.jpg", "f3"},
	}
	recs := testData(t)
	rat := ImageRating(recs)
	if v := len(rat); v != len(exp) {
		t.Errorf("unexpected value: %v", v)
	}
	if v := rat[0].Val; v != "f3" {
		t.Errorf("unexpected value: %q", v)
	}
}
