package remuxregex

import (
	"errors"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

type ReMux struct {
	mu    sync.RWMutex
	plain map[Method]map[string]http.Handler
	regex map[Method]map[*regexp.Regexp]http.Handler
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

func (r *ReMux) RegisterRegex(method Method, path *regexp.Regexp, handler http.Handler) error {
	if !isValidMethod(method) {
		return ErrInvalidMethod
	}

	if !strings.HasPrefix(path.String(), `^/`) {
		return ErrInvalidPath
	}

	if !strings.HasSuffix(path.String(), `$`) {
		return ErrInvalidPath
	}

	if handler == nil {
		return ErrNilHandler
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.regex[method][path]; exists {
		return ErrAmbiguousMapping
	}

	if r.regex == nil {
		r.regex = make(map[Method]map[*regexp.Regexp]http.Handler)
	}

	if r.regex[method] == nil {
		r.regex[method] = make(map[*regexp.Regexp]http.Handler)
	}

	r.regex[method][path] = handler
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
	var resultHandler http.Handler
	if handlers, exists := r.plain[Method(request.Method)]; exists {
		if handler, ok := handlers[request.URL.Path]; ok {
			resultHandler = handler
		}
	}
	if resultHandler == nil {
		if handlers, exists := r.regex[Method(request.Method)]; exists {
			for path, handler := range handlers {
				if path.MatchString(request.URL.Path) {
					resultHandler = handler
					break
				}
			}
		}
	}
	if resultHandler == nil {
		resultHandler = r.notFoundHandler
	}
	r.mu.RUnlock()

	resultHandler.ServeHTTP(writer, request)
}

func isValidMethod(method Method) bool {
	for _, m := range []Method{GET, POST, PUT, PATCH, DELETE} {
		if m == method {
			return true
		}
	}
	return false
}
