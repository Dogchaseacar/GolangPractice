package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router[method+"-"+pattern] = handler
}

func (engine *Engine) Get(pattern string, handler HandlerFunc) {
	engine.addRoute("Get", pattern, handler)
}

func (engine *Engine) Post(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handler, ok := engine.router[req.Method+"-"+req.URL.Path]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 not found")
	}
}

func (engine *Engine) Run() {
	http.ListenAndServe("9999", engine)
}
