package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var bots Bots
var tr *http.Transport
var config Config
var err error

func main() {
	fmt.Println("Telegram Webhook Gateway initializing...")
	jsonInit()
	config, err = getConfig()
	if err != nil {
		log.Fatal(err)
	}
	bots, err = getBots()
	if err != nil {
		log.Fatal(err)
	}
	tr = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port 8443.")
	err := http.ListenAndServeTLS(":"+config.ListenPort, config.CertFile, config.KeyFile, nil)
	if err != nil {
		log.Fatal("could not start listener:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	for _, a := range bots.Bot {
		if a.URI == r.RequestURI {
			req, err := http.NewRequest("POST", a.Host+a.URI, r.Body)
			if err != nil {
				log.Println(err)
				return
			}
			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{Transport: tr}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("> "+a.URI+" FAIL:", err)
				http.Error(w, http.StatusText(http.StatusGatewayTimeout), http.StatusGatewayTimeout)
				return
			}
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			w.Write(body)
			fmt.Println("> " + a.URI + " OK")
			return
		}
	}
	fmt.Println("Invalid URI '" + r.RequestURI + "', ignoring...")
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}
