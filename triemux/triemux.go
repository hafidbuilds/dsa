package triemux

import (
    "context"
    "errors"
    "fmt"
    "net/http"
    "strings"
    "time"
)

type ctxKey struct{}

func Params(ctx context.Context) Vars {
    vars, ok := ctx.Value(ctxKey{}).(Vars)
    if !ok {
        return make(Vars)
    }
    return vars
}

type Mux struct {
    // routes map[key]http.HandlerFunc
    routes     *trie
    errHandler func(w http.ResponseWriter, r *http.Request, err error) bool
    mids       []Middleware
}

func New() *Mux {
    return &Mux{
        routes:     newTrie(),
        errHandler: nil,
    }
}

type Middleware = func(http.HandlerFunc) http.HandlerFunc

func (m *Mux) Use(middlewares ...Middleware) {
    m.mids = middlewares
}

func (m *Mux) SetErrorHandler(h func(w http.ResponseWriter, r *http.Request, err error) bool) {
    m.errHandler = h
}

func (m *Mux) Handle(method string, path string, handler http.HandlerFunc) {
    m.routes.insert(method, path, handler)
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    h, vars, err := m.routes.find(r.Method, r.URL.Path)
    if err != nil {
        if m.errHandler != nil {
            if m.errHandler(w, r, err) {
                return
            }
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        // if errors.Is(err, errMethodFound) {
        //     w.WriteHeader(http.StatusMethodNotAllowed)
        //     return
        // }
        // if errors.Is(err, errPathNotFound) {
        //     w.WriteHeader(http.StatusNotFound)
        //     return
        // }
    }

    ctx := r.Context()
    ctx = context.WithValue(ctx, ctxKey{}, vars)

    for _, mid := range m.mids {
        h = mid(h) // f o g
    }

    h(w, r.WithContext(ctx))
}

func Logger() Middleware {
    return func(f http.HandlerFunc) http.HandlerFunc {
        g := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // custom code
            now := time.Now()
            fmt.Println("before handler", r.Method, r.URL)
            f(w, r)
            fmt.Println("after handler", time.Since(now))
            // custom code
        })
        return g
    }
}

type Segment = string
type Method = string

type node struct {
    segment    Segment
    isTerminal bool
    isVariable bool
    children   map[Segment]*node
    handlers   map[Method]http.HandlerFunc
}

type trie struct {
    root *node
}

func newTrie() *trie {
    return &trie{
        root: &node{
            segment:    "*",
            isTerminal: false,
            children:   make(map[Segment]*node),
            handlers:   make(map[Method]http.HandlerFunc),
        },
    }
}

var (
    errPathNotFound = errors.New("path not found")
    errMethodFound  = errors.New("method not found")
)

func (t *trie) insert(method string, path string, handler http.HandlerFunc) {
    // `/`
    segments := strings.Split(path, "/")
    p := t.root
    for i, s := range segments {
        n, ok := p.children[s]
        if !ok {
            n = &node{
                segment:    strings.TrimPrefix(s, ":"),
                isTerminal: false,
                isVariable: strings.HasPrefix(s, ":"),
                children:   make(map[Segment]*node),
                handlers:   make(map[Method]http.HandlerFunc),
            }
            if !n.isVariable {
                p.children[s] = n
            } else {
                p.children["<vars>"] = n
            }

        }
        p = n
        if i == len(segments)-1 {
            n.isTerminal = true
            if n.handlers == nil {
                n.handlers = make(map[Method]http.HandlerFunc)
            }
            if _, ok := n.handlers[method]; ok {
                panic(fmt.Sprintf("duplicate method handler: method=%q path=%q", method, path))
            }
            n.handlers[method] = handler
        }
    }
}

type Vars map[string]string

func (t *trie) find(method Method, path string) (http.HandlerFunc, Vars, error) {
    // `/`
    segments := strings.Split(path, "/")
    p := t.root
    vars := make(Vars)
    for _, s := range segments {
        n, ok := p.children[s]
        if !ok {
            n, ok = p.children["<vars>"]
            if !ok {
                return nil, nil, fmt.Errorf("path segment not found: %q from %q: %w", s, path, errPathNotFound)
            }
            vars[n.segment] = s
        }
        p = n
    }
    if !p.isTerminal || p.handlers == nil {
        return nil, nil, fmt.Errorf("method not found: %q: %w", method, errMethodFound)
    }
    h, ok := p.handlers[method]
    if !ok {
        return nil, nil, fmt.Errorf("method not found: %q: %w", method, errMethodFound)
    }
    return h, vars, nil
}
