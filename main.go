package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"os"
)

var version string = "0.0.4";
var hostname string = "undef";
var testenv string = "unset";


func main() {
	var err error;
	var err_bool bool;
	println("Starting ver:", version)

	hostname, err = os.Hostname()

	if err != nil {
		fmt.Println(err)
	}


	testenv, err_bool = os.LookupEnv("GO_CONTAINER_ENV")
	if !err_bool {
		fmt.Println("GO_CONTAINER_ENV is not present")
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var dataset string;

	content, err := ioutil.ReadFile("data/data.txt")
	if err != nil {
		fmt.Println(err)
		dataset = "error reading data/data.txt"
	} else {
		dataset = string(content)
	}


	log.Printf("Request received from %s\nstorage data: %s\nGO_CONTAINER_ENV: %s", r.RemoteAddr, dataset, testenv)
	fmt.Fprintf(w, "from %s, v%s data:%s env:%s", hostname, version, dataset, testenv)
}
