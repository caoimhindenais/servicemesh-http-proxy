package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

const PROXY_PORT = "8000"
const SERVICE_PORT = "8080"

// Serve a reverse proxy for a given url
func forwardRequest(res http.ResponseWriter, req *http.Request) {
	// parse the url
	url, _ := url.Parse("http://localhost:"+SERVICE_PORT)

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}

// Get a json decoder for a given requests body
func checkForDragons(request *http.Request)  {
	// Read body to buffer
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		panic(err)
	}

	bodyContent := string(body)
	log.Print("Request body :  ", bodyContent)
	if(strings.Contains(strings.ToLower(bodyContent), "dragon") ) {

		log.Printf("");
		log.Printf("Ahhh Ahhh ... Dragon. Wheres Bard?")

		deadDragon := []byte("{\"name\":\"Dead Dragon\"}")
		request.Body = ioutil.NopCloser(bytes.NewBuffer(deadDragon))
		request.ContentLength = int64(len(deadDragon))
	} else {

		request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	}
}

// Given a request send it to the appropriate url
func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	checkForDragons(req)
	forwardRequest(res, req)
}

/*
	Entry
*/

func main() {
	// Log setup values

	log.Printf("Proxy will run on: %s\n", PROXY_PORT)

	// start server
	http.HandleFunc("/", handleRequestAndRedirect)
	if err := http.ListenAndServe(":"+PROXY_PORT, nil); err != nil {
		panic(err)
	}
}