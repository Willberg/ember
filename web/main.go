package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:20000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	sb := &strings.Builder{}
	sb.WriteString(fmt.Sprintf("%s %s %s\n\n", r.Method, r.URL.RequestURI(), r.Proto))
	hasHost := false
	for k, s := range r.Header {
		if strings.EqualFold(k, "host") {
			hasHost = true
		}
		sb.WriteString(fmt.Sprintf("%s: %s\n", k, strings.Join(s, ",")))
	}
	if !hasHost {
		sb.WriteString(fmt.Sprintf("Host: %s\n", r.Host))
	}
	sb.WriteString("\n")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
	}
	bs := strings.Split(string(body), "&")
	sb.WriteString(strings.Join(bs, "\n"))
	fmt.Printf("%s", sb.String())
	fmt.Fprintf(w, "%s", sb.String())
}
