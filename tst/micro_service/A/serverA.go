package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	// resp, err := http.Get("http://esigo.dev.b:8090/test")
	// url := "http://esigo.dev.b:8090/test"
	url := "http://ingress-nginx-controller.ingress-nginx.svc.cluster.local:80/serviceb/test"
	// url := "http://microapp-service-b.myapp.svc.cluster.local:80/test"
	http.NewRequest("GET", url, nil)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}
	req.Header = r.Header
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}
		bodyString := string(bodyBytes)
		log.Println(bodyString)
		w.Write(bodyBytes)
	} else {
		w.Write([]byte("Gorilla!\n"))
	}
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", HandleRoot)

	// Bind to a port and pass our router in
	log.Println(http.ListenAndServe(":80", r))
}
