package triemux_test

import (
    "fmt"
    "net/http"
    "testing"

    "github.com/josestg/dsa/triemux"
)

func TestMux_Handle(t *testing.T) {
    mux := triemux.New()
    mux.Use(triemux.Logger())
    mux.SetErrorHandler(func(w http.ResponseWriter, r *http.Request, err error) bool {
        t.Log("this is the error", err)
        w.WriteHeader(http.StatusOK)
        return true
    })
    mux.Handle("GET", "/api/v1/users", func(w http.ResponseWriter, r *http.Request) {
        _, _ = fmt.Fprintf(w, "method=%q path=%q get all users", r.Method, r.URL.Path)
    })
    mux.Handle("GET", "/api/v1/users/:id/:age", func(w http.ResponseWriter, r *http.Request) {
        ctx := r.Context()
        vars := triemux.Params(ctx)
        _, _ = fmt.Fprintf(w, "method=%q path=%q get user by id %q age %q", r.Method, r.URL.Path, vars["id"], vars["age"])
    })
    mux.Handle("POST", "/api/v1/users", func(w http.ResponseWriter, r *http.Request) {
        _, _ = fmt.Fprintf(w, "method=%q path=%q create users", r.Method, r.URL.Path)
    })

    srv := http.Server{
        Addr:    "localhost:8080",
        Handler: mux,
    }

    if err := srv.ListenAndServe(); err != nil {
        panic(err)
    }
}
