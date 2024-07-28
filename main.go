package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func rootApi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome to back-end\n")

	body, _ := io.ReadAll(r.Body)

	fmt.Println(r.Method, r.URL, r.Proto, r.Host,
		"<--"+r.RemoteAddr, "( "+string(body)+" )")

	for name, values := range r.Header {
		if strings.HasPrefix(name, "X-") {
			fmt.Print(name, values, "\n")
		}
	}

}

func handleReq() {
	http.HandleFunc("/", rootApi)
	log.Fatal(http.ListenAndServe(":1080", nil))
}

func main() {
	handleReq()
}
