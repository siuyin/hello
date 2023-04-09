package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/open-policy-agent/opa/sdk"
)

var opa *sdk.OPA

func main() {
	webAndBundleInit()

	opa = opaEvalRun()

	select {}
}

func webAndBundleInit() {
	go func() {
		http.Handle("/bundles/", http.StripPrefix("/bundles/", http.FileServer(http.Dir("./bundles"))))
		http.HandleFunc("/bundles/bundle.tar.gz", bundleTarGz)

		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
	log.Println("web and OPA bundle server initializing")
}

func bundleTarGz(w http.ResponseWriter, r *http.Request) {
	log.Println("OPA downloading bundle")
	w.Header().Add("Content-Type", "application/gzip")

	zw := gzip.NewWriter(w)
	defer zw.Close()
	tw := tar.NewWriter(zw)
	defer tw.Close()

	filesys := os.DirFS("./bundles")
	fs.WalkDir(filesys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		info, err := d.Info()
		if err != nil {
			return err
		}
		hdr, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		hdr.Name = path
		if d.IsDir() {
			hdr.Name = hdr.Name + "/"
		}
		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		body, err := os.ReadFile("./bundles/" + path)
		if err != nil {
			return err
		}
		if _, err := tw.Write(body); err != nil {
			return err
		}

		return nil
	})

}

func opaEvalRun() *sdk.OPA {
	opa := opaInit()

	go func() {
		ctx := context.Background()
		defer opa.Stop(ctx)

		for {
			result, err := opa.Decision(ctx, sdk.DecisionOptions{
				Path:  "/example/authz/award_value",
				Input: map[string]interface{}{"open": "sesame"},
			})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(result.Result)
			time.Sleep(time.Second)
		}
	}()

	log.Println("OPA ready")
	return opa
}

func opaInit() *sdk.OPA {
	config := []byte(`
services: 
  acme:
    url: http://localhost:8080
bundles:
  authz:
    service: acme
    resource: bundles/bundle.tar.gz
    polling:
      min_delay_seconds: 10
      max_delay_seconds: 30
decision_logs:
  console: true
`)
	ctx := context.Background()
	opa, err := sdk.New(ctx, sdk.Options{
		Config: bytes.NewReader(config),
	})
	if err != nil {
		log.Fatal(err)
	}
	return opa
}
