package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	fmt.Println("webhook example")

	postToFoo()

	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		displayAndLogHookCall(r)
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func postToFoo() {
	go func() {
		for {
			resp := call("/foo", "gerbau")
			defer resp.Body.Close()

			body := getBody(resp)
			fmt.Printf("received: %s\n", body)
			time.Sleep(time.Second)
		}
	}()
}
func call(endpoint string, authz string) *http.Response {
	//resp, err := http.Post("http://127.0.0.1:8080/foo", "text/plain", strings.NewReader("Brown Fox"))

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://127.0.0.1:8080/foo", strings.NewReader("Brown Fox"))
	req.Header.Add("Authorization", "gerbau")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	return resp
}
func getBody(resp *http.Response) []byte {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	return body
}

func displayAndLogHookCall(r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	log.Printf("hook called, body: %s, headers: %v", body, r.Header)
}
