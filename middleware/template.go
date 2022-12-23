package middleware

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

//This file is meant to be used by the tojen tool to generate the required jennifer constructs
//to generate the test file dynamically
//tojen gen template.go test_gen.go
//https://github.com/aloder/tojen

import (
	"github.com/go-chi/chi/v5"
)

type testCase struct {
	uri      string
	method   string
	body     string
	headers  http.Header
	expected string
	status   int
}

func Test_HttpServer(t *testing.T) {

	r := setupRoutes()
	headers := getHeaders()
	tss := getTestCases(headers)

	for _, ts := range tss {
		t.Run(ts.uri, func(t *testing.T) {
			req := newRequest(ts.method, ts.uri, ts.body, ts.headers)
			res := executeRequest(r, req)
			assert(t, res, ts)
		})
	}
}

func setupRoutes() http.Handler {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World"))
	})

	return r
}

func getHeaders() http.Header {
	headers := make(http.Header)
	headers["Content-Type"] = []string{"application/json"}
	return headers
}

func getTestCases(headers http.Header) []testCase {
	return []testCase{
		{
			uri:      "/more/tests?id=5",
			method:   "POST",
			body:     `{"value":"text","payload":{"items":[{"id":25}],"origin":"ARG"}}`,
			headers:  headers,
			expected: "More tests",
			status:   200,
		},
		{
			uri:      "/more/tests?id=5",
			method:   "POST",
			body:     `{"value":"text","payload":{"items":[{"id":25}],"origin":"ARG"}}`,
			headers:  headers,
			expected: "More tests",
			status:   200,
		},
	}
}
func executeRequest(h http.Handler, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	return rr
}
func newRequest(method string, uri string, body string, headers http.Header) *http.Request {
	req, _ := http.NewRequest(method, uri, bytes.NewBufferString(body))
	req.Header = headers
	return req
}
func assert(t *testing.T, res *httptest.ResponseRecorder, ts testCase) {
	if res.Code != ts.status {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", res.Code, ts.status)
	}
	if res.Body.String() != ts.expected {
		t.Errorf("Response body is wrong. Have: %s, want: %s.", res.Body.String(), ts.expected)
	}
}
