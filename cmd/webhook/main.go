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

	callWebhookWorker()

	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		displayAndLogHookCall(r)
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

const baseURL = "http://127.0.0.1:8080"

func callWebhookWorker() {
	go func() {
		client := &http.Client{}
		for {
			resp := call(client, "/foo", "gerbau") // make client into a type if more parameters needed.
			defer resp.Body.Close()

			body := getBody(resp)
			fmt.Printf("received: %s\n", body)

			time.Sleep(time.Second)
		}
	}()
}
func call(client *http.Client, endpoint string, authz string) *http.Response {
	//resp, err := http.Post("http://127.0.0.1:8080/foo", "text/plain", strings.NewReader("Brown Fox"))

	msg := fmt.Sprintf("Brown Fox: %s", time.Now().Format("15:04:05.000"))
	req, err := http.NewRequest("POST", baseURL+endpoint, strings.NewReader(msg))
	req.Header.Add("Authorization", authz)
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
