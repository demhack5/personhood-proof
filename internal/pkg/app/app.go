package app

import (
	"context"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/go-chi/chi"
	"github.com/oklog/run"
	"github.com/pkg/errors"
)

type httpServer struct {
	listener net.Listener
	router   chi.Router
	port     int
	name     string
}

func newHTTPServer(port int) (*httpServer, error) {
	router := chi.NewMux()
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return nil, errors.Errorf("cannot create listener: %v", err)
	}
	return &httpServer{
		listener: listener,
		router:   router,
		port:     port,
	}, nil
}

func (hs *httpServer) runner() (func() error, func(error)) {
	s := &http.Server{Handler: hs.router}
	return func() error {
			return errors.Wrap(s.Serve(hs.listener), "http server")
		}, func(err error) {
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()

			s.SetKeepAlivesEnabled(false)
			err = errors.Wrap(s.Shutdown(ctx), "shutdown")
			_ = err
		}
}

type App struct {
	runGroup run.Group

	http *httpServer

	mu    sync.RWMutex
	ready bool
}

func New() (*App, error) {
	httpServerRunner, err := newHTTPServer(8080)
	if err != nil {
		return nil, err
	}
	app := &App{
		http: httpServerRunner,
	}
	app.AddRunner(httpServerRunner.runner())
	return app, nil
}

func (a *App) AddRunner(exec func() error, close func(error)) {
	a.runGroup.Add(exec, close)
}

func (a *App) Run() error {
	a.setIsReady(true)
	return nil
}

func (a *App) setIsReady(v bool) {
	a.mu.Lock()
	a.ready = v
	a.mu.Unlock()
}

func (a *App) checkIsReady() error {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if !a.ready {
		return errors.New("app is not ready")
	}
	return nil
}
