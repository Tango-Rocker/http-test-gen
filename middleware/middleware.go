package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type generator interface {
	Gen(tss []testCase, path string) error
}

type Automata struct {
	tss       []testCase
	generator generator
	path      string
}

func New(path string) *Automata {
	return &Automata{
		tss:       make([]testCase, 0),
		generator: &jenniferGenerator{},
		path:      path,
	}
}

func (auto *Automata) Handle(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ts := testCase{}

		ts.uri = r.URL.RequestURI()
		if ts.uri == "/auto-generate-test" {
			err := auto.generator.Gen(auto.tss, auto.path)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Test Generated Successfully"))
			return
		}

		ts.method = r.Method
		ts.headers = r.Header
		ts.body = getBody(r)

		obs := &observer{Writer: w, Buffer: make([]byte, 0)}
		handler.ServeHTTP(obs, r)

		ts.status = obs.Status
		ts.expected = string(obs.Buffer)

		auto.tss = append(auto.tss, ts)
	})
}

type observer struct {
	Writer http.ResponseWriter
	Status int
	Buffer []byte
}

func (o *observer) Header() http.Header {
	return o.Writer.Header()
}

func (o *observer) Write(buffer []byte) (int, error) {
	o.Buffer = append(o.Buffer, buffer...)
	return o.Writer.Write(buffer)
}

func (o *observer) WriteHeader(statusCode int) {
	o.Status = statusCode
	o.Writer.WriteHeader(statusCode)
}

func getBody(r *http.Request) string {
	body, _ := ioutil.ReadAll(r.Body)
	buf := bytes.NewBuffer(body)
	r.Body = ioutil.NopCloser(buf)

	return string(body)
}
