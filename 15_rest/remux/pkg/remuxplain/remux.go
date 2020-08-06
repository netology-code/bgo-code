package remuxplain

import (
	"errors"
	"net/http"
	"strings"
	"sync"
)

type ReMux struct {
	mu    sync.RWMutex
	plain map[Method]map[string]http.Handler
	notFoundHandler http.Handler
}

var defaultNotFoundHandler = func(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNotFound)
}

func NewReMux() *ReMux {
	return &ReMux{
		notFoundHandler: http.HandlerFunc(defaultNotFoundHandler),
	}
}

var (
	ErrInvalidPath      = errors.New("invalid path")
	ErrInvalidMethod    = errors.New("invalid http method")
	ErrNilHandler       = errors.New("nil handler")
	ErrAmbiguousMapping = errors.New("ambiguous mapping")
)

type Method string

const (
	GET     Method = "GET"
	POST    Method = "POST"
	PUT     Method = "PUT"
	PATCH   Method = "PATCH"
	DELETE  Method = "DELETE"
	OPTIONS Method = "OPTIONS"
	HEAD    Method = "HEAD"
)

func (r *ReMux) RegisterPlain(method Method, path string, handler http.Handler) error {
	if !isValidMethod(method) {
		return ErrInvalidMethod
	}

	if !strings.HasPrefix(path, "/") {
		return ErrInvalidPath
	}

	if handler == nil {
		return ErrNilHandler
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	// запрещаем добавлять дубликаты
	if _, exists := r.plain[method][path]; exists {
		return ErrAmbiguousMapping
	}

	if r.plain == nil {
		r.plain = make(map[Method]map[string]http.Handler)
	}

	if r.plain[method] == nil {
		r.plain[method] = make(map[string]http.Handler)
	}

	r.plain[method][path] = handler
	return nil
}

func (r *ReMux) NotFound(handler http.Handler) error {
	if handler == nil {
		return ErrNilHandler
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	r.notFoundHandler = handler
	return nil
}

func (r *ReMux) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	r.mu.RLock()
	handlers, exists := r.plain[Method(request.Method)]
	if !exists {
		if r.notFoundHandler != nil {
			r.notFoundHandler.ServeHTTP(writer, request)
		}
		return
	}

	handler, ok := handlers[request.URL.Path]
	if !ok {
		if r.notFoundHandler != nil {
			r.notFoundHandler.ServeHTTP(writer, request)
		}
		return
	}
	r.mu.RUnlock()

	handler.ServeHTTP(writer, request)
}

func isValidMethod(method Method) bool {
	for _, m := range []Method{GET, POST, PUT, PATCH, DELETE} {
		if m == method {
			return true
		}
	}
	return false
}
