// Server for debugging HTTP requests;

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:8081", "Listen address")
	flag.Parse()

	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			var b strings.Builder

			fmt.Fprintf(&b, "%v\t%v\t%v\tHost: %v\n", r.RemoteAddr, r.Method, r.URL, r.Host)
			for name, headers := range r.Header {
				for _, h := range headers {
					fmt.Fprintf(&b, "%v: %v\n", name, h)
				}
			}
			log.Println(b.String())

			fmt.Fprintf(w, "hello %s\n", r.URL)
		})

	log.Println("Starting server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
