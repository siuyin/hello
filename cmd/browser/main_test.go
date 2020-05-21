package main

import (
	"os"
	"testing"
)

func TestReadConfig(t *testing.T) {
	cfg := readConfig("testdata/sample.yaml")
	if v := len(cfg.Pages); v != 3 {
		t.Errorf("unexpected value: %v", v)
	}
	if v := cfg.InputFile; v != "/h/Downloads/data.csv" {
		t.Errorf("unexpected value: %v", v)
	}
}

func TestCreatePages(t *testing.T) {
	cfg := readConfig("testdata/sample.yaml")
	os.RemoveAll(cfg.OutputDir)
	recs := readData(cfg)
	createPages(cfg, recs)

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
