package main

import (
	"os"
	"sync"
	"testing"

	"github.com/siuyin/hello/brow"
)

func TestCreatePages(t *testing.T) {
	cfg := readConfig("testdata/sample.yaml")
	os.RemoveAll(cfg.OutputDir)
	recs := brow.ReadData(cfg)
	var wg sync.WaitGroup
	wg.Add(1)
	createPages(cfg, recs, &wg)

	t.Run("chkIndex", func(t *testing.T) {
		d, err := os.Open(cfg.OutputDir)
		if err != nil {
			t.Error(err)
		}
		defer d.Close()

		fns, err := d.Readdirnames(0)
		if err != nil {
			t.Error(err)
		}

		if v := len(fns); v != len(cfg.Pages) {
			t.Errorf("unexpected value: %v: expected %d", v, len(cfg.Pages))
		}
	})
}

func TestWriteRatings(t *testing.T) {
	cfg := readConfig("testdata/sample.yaml")
	os.RemoveAll(cfg.OutputDir)
	recs := brow.ReadData(cfg)

	var wg sync.WaitGroup
	wg.Add(1)
	writeRatings(cfg, recs, &wg)
	wg.Wait()

	t.Run("checkOutput", func(t *testing.T) {
		_, err := os.Stat("bar/ratings.csv")
		if err != nil {
			t.Error(err)
		}
	})
}

func TestPages(t *testing.T) {
	cfg := readConfig("testdata/sample.yaml")
	recs := brow.ReadData(cfg)
	page := cfg.Pages[0]

	ps := newPageSet(cfg, recs, page, 1) // 1 rec per page.
	pls := ps.pageLinks()
	if v := len(pls); v < 2 {
		t.Errorf("there should be at least two pages, unexpected value: %v", v)
	}
}
