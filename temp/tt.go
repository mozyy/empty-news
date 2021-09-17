package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ello, world"))
	})
	go func() {
		err := http.ListenAndServe(":9099", nil)
		if err != nil {
			fmt.Println("error: ", err)
		} else {
			fmt.Println("success ")
		}
	}()
	time.Sleep(2 * time.Second)
	resp, err := http.Get("http://localhost:9099")
	if err != nil {
		fmt.Println("error2: ", err)
		return
	}
	body, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Println("error2: ", err)
		return
	}
	fmt.Println(string(body))
	time.Sleep(2 * time.Hour)
}
