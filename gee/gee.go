package gee

import (
	"fmt"
	"log"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRouter(method string, pattern string, handlerFunc HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handlerFunc
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRouter("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRouter("POST", pattern, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, r)
	} else {
		_, err := fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
		if err != nil {
			log.Printf("发送reponse失败：%s", "404")
		}
	}
}

func (engine *Engine) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, engine))
}
