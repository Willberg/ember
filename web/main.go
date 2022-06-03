package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	port = flag.String("p", "20000", "监控的端口")
)

func main() {
	flag.Parse()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%s", *port), nil))
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
	sb.WriteString("\n\n\n")
	fmt.Fprintf(os.Stdout, "%s", sb.String())
	fmt.Fprintf(w, "%s", sb.String())
}
