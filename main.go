package main

import (
	"net/http"
	"io"
	"os"
	"io/ioutil"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://hello:8080")
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}

		io.WriteString(w, "name: "+os.Getenv("HOSTNAME")+" ---> "+string(body))
	})

	http.HandleFunc("/name", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "this HOSTNAME is: " + os.Getenv("HOSTNAME"))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
