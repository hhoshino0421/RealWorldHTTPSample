package TestEcho

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {

	dump, err := httputil.DumpRequest(r, true)

	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	fmt.Println(string(dump))

	_, err = fmt.Fprintf(w, "<html><body>hello</body></html>\n")
	if err != nil {
		log.Println(err)
	}

}

func TestEcoMain() {

	var httpServer http.Server

	http.HandleFunc("/", handler)

	log.Println("start http listening : 18888")

	httpServer.Addr = ":18888"

	log.Println(httpServer.ListenAndServe())

}
