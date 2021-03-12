package http

import (
	"net/http"

	"github.com/phungvandat/source-template/utils/logger"
)

type mux struct {
	trees   map[string]*mux
	method  string
	hdlFunc HandleFunc
}

func newMux() *mux {
	return &mux{
		trees: make(map[string]*mux),
	}
}

// NewRouter is init router interface
func NewRouter() Router {
	return newMux()
}

// Router interface
type Router interface {
	Group(pattern string, fn func(r Router))
	POST(pattern string, hdl HandleFunc)
	Build() map[string]HandleFunc
}

func (m *mux) Group(pattern string, fn func(r Router)) {
	r := newMux()
	fn(r)
	m.trees[pattern] = r
}

func (m *mux) POST(pattern string, hdl HandleFunc) {
	if hdl, ok := m.trees[pattern]; ok && hdl.hdlFunc != nil {
		logger.Warning("duplicate router at pattern %v", pattern)
	}
	m.trees[pattern] = &mux{
		method:  http.MethodPost,
		hdlFunc: hdl,
	}
}

func (m *mux) Build() map[string]HandleFunc {
	var mRoute = make(map[string]HandleFunc)
	for pattern := range m.trees {
		m.trees[pattern].reassembleRoute(pattern, mRoute)
	}

	return mRoute
}

func (m *mux) reassembleRoute(parent string, mRoute map[string]HandleFunc) {
	if len(m.trees) == 0 &&
		m.method != "" &&
		m.hdlFunc != nil {
		mRoute[parent] = m.hdlFunc
		return
	}
	for pattern := range m.trees {
		parent += pattern
		m.trees[pattern].reassembleRoute(parent, mRoute)
	}
}
