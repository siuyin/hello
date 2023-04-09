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
		http.HandleFunc("/rocket/launch", chkRocketLaunch)
		http.HandleFunc("/example/authz/allow", chkAuthzAllow)
		http.HandleFunc("/example/authz/award_value", chkAuthzValue)
		http.Handle("/bundles/", http.StripPrefix("/bundles/", http.FileServer(http.Dir("./bundles"))))
		http.HandleFunc("/bundles/bundle.tar.gz", bundleTarGz)

		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
	log.Println("web and OPA bundle server initializing")
}

type rocketLaunchParams struct {
	FuelKg             float32
	O2Kg               float32
	AvionicsGo         bool
	FlightDirectorNoGo bool
}

func chkRocketLaunch(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	params := rocketLaunchParams{}
	params.FuelKg = 115
	params.O2Kg = 212
	params.AvionicsGo = true
	params.FlightDirectorNoGo = true
	fmt.Fprintf(w, "rocket launch parameters: %#v\n", params)
	result, err := opa.Decision(ctx, sdk.DecisionOptions{
		Path:  "/rocket/launch",
		Input: map[string]interface{}{"params": params},
	})
	if err != nil {
		http.Error(w, "could not evaluate policy", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%#v\n", result.Result)
}

func chkAuthzAllow(w http.ResponseWriter, r *http.Request) {
	opn := r.FormValue("open")
	ctx := context.Background()
	result, err := opa.Decision(ctx, sdk.DecisionOptions{
		Path:  "/example/authz/allow",
		Input: map[string]interface{}{"open": opn},
	})
	if err != nil {
		http.Error(w, "could not evaluate policy", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%#v\n", result.Result)
}
func chkAuthzValue(w http.ResponseWriter, r *http.Request) {
	opn := r.FormValue("open")
	ctx := context.Background()
	result, err := opa.Decision(ctx, sdk.DecisionOptions{
		Path:  "/example/authz/award_value",
		Input: map[string]interface{}{"open": opn},
	})
	if err != nil {
		http.Error(w, "could not evaluate policy", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%#v\n", result.Result)
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
		select {}

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
