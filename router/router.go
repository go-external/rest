// Minimalistic wrapper over default http router,
// to make it easier to create API using default http/net

package main

import "net/http"

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type MuxWrapper interface {
	Get(path string, handler HandlerFunc)
	Post(path string, handler HandlerFunc)
	Put(path string, handler HandlerFunc)
	Patch(path string, handler HandlerFunc)
	Delete(path string, handler HandlerFunc)
}

type muxWrapper struct {
	*http.ServeMux
}

func NewMuxWrapper(mux *http.ServeMux) MuxWrapper {
	return &muxWrapper{mux}
}

func (m *muxWrapper) Get(path string, handler HandlerFunc) {
	m.HandleFunc(path, registerHandler(http.MethodGet, handler))
}

func (m *muxWrapper) Post(path string, handler HandlerFunc) {
	m.HandleFunc(path, registerHandler(http.MethodPost, handler))
}

func (m *muxWrapper) Put(path string, handler HandlerFunc) {
	m.HandleFunc(path, registerHandler(http.MethodPut, handler))
}

func (m *muxWrapper) Patch(path string, handler HandlerFunc) {
	m.HandleFunc(path, registerHandler(http.MethodPatch, handler))
}

func (m *muxWrapper) Delete(path string, handler HandlerFunc) {
	m.HandleFunc(path, registerHandler(http.MethodDelete, handler))
}

func registerHandler(method string, handler HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if method == r.Method {
			handler(w, r)
		}
	}
}
