package main

import (
	"fmt"
	"log"
	"net/http"

	"gee"
)

func main() {
	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		_, err := fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
		if err != nil {
			log.Printf("发送response失败：%s", req.URL.Path)
		}
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			_, err := fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
			if err != nil {
				log.Printf("发送response失败：%s:%s", k, v)
			}
		}
	})

	r.Run(":8080")
}
