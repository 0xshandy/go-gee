package gee

import (
	"log"
	"net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.router.addRouter("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.router.addRouter("POST", pattern, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContext(w, r)
	engine.router.handle(c)
}

func (engine *Engine) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, engine))
}
