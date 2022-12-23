# http-test-gen
dynamically generate test for your rest api


## Configuration

```go 
//any router solution that implements http.HandleFunc should suffice
r := chi.NewRouter()

automata := New("C:\\your\\favortie\\path\\")
r.Use(automata.Handle)
```

we register our midleware with the router
to inspect all request/response and populate a test case repository


Once you are done with your live testing session
do the following request to generate your new test file

GET localhost:8080/auto-generate-test

enjoy your free test file!


