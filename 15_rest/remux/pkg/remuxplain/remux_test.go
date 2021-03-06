package remuxplain

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReMux_Plain(t *testing.T) {
	mux := NewReMux()
	if err := mux.RegisterPlain(GET, "/get", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(GET))
	})); err != nil {
		t.Fatal(err)
	}
	if err := mux.RegisterPlain(POST, "/post", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(POST))
	})); err != nil {
		t.Fatal(err)
	}

	type args struct {
		method Method
		path   string
	}

	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "GET", args: args{method: GET, path: "/get"}, want: []byte(GET)},
		{name: "POST", args: args{method: POST, path: "/post"}, want: []byte(POST)},
		// TODO: write for other methods
	}

	for _, tt := range tests {
		request := httptest.NewRequest(string(tt.args.method), tt.args.path, nil)
		response := httptest.NewRecorder()
		mux.ServeHTTP(response, request)
		got := response.Body.Bytes()
		if !bytes.Equal(tt.want, got) {
			t.Errorf("got %s, want %s", got, tt.want)
		}
	}
}

func TestReMux_NotFound(t *testing.T) {
	mux := NewReMux()

	type args struct {
		method Method
		path   string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "GET", args: args{method: GET, path: "/get"}, want: http.StatusNotFound},
		{name: "POST", args: args{method: POST, path: "/post"}, want: http.StatusNotFound},
		// TODO: write for other methods
	}

	for _, tt := range tests {
		request := httptest.NewRequest(string(tt.args.method), tt.args.path, nil)
		response := httptest.NewRecorder()
		mux.ServeHTTP(response, request)
		got := response.Result().StatusCode
		if tt.want != got {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}
