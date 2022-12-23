package middleware

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"testing"
)

func Test_Middleware(t *testing.T) {
	r := chi.NewRouter()

	automata := New("D:\\develop\\Golang\\auto-test")
	r.Use(automata.Handle)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World"))
	})

	headers := make(http.Header)
	headers["Content-Type"] = []string{"application/json"}

	tss := []testCase{
		{
			uri:      "http://localhost:8080/",
			method:   "GET",
			body:     ``,
			headers:  headers,
			expected: `Hello World`,
			status:   200,
		},
	}

	for _, ts := range tss {
		t.Run(ts.uri, func(t *testing.T) {
			req := newRequest(ts.method, ts.uri, ts.body, ts.headers)
			res := executeRequest(r, req)
			assert(t, res, ts)
		})
	}

	//TODO: check if the test file was created and if it contains the expected content
}
