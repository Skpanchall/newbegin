package middleware

import (
	"fmt"
	"net/http"
	"time"

	utils "github.com/Skpanchall/newbegin/utils"
)

type middlewareWriter struct {
	http.ResponseWriter
	code int
}

func (w *middlewareWriter) WriteHeader(code int) {
	w.code = code
	w.ResponseWriter.WriteHeader(code)
}

type AppHandler func(http.ResponseWriter, *http.Request) error
type MiddlewareFunc func(AppHandler) AppHandler

func WrapperFunc(w http.ResponseWriter, r *http.Request, h AppHandler) error { // calling first middleware and then call a h handler
	wrappedHandler := middleWareWrapper(h, AuthMiddleware, LoggingMiddleware)
	return wrappedHandler(w, r)
}

func RegisterRoute(path string, h AppHandler) { // registere routes for wrapper func
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		err := WrapperFunc(w, r, h)
		if err != nil {
			utils.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

func middleWareWrapper(h AppHandler, middlewares ...MiddlewareFunc) AppHandler { // this function will wrap the handler with the middlewares
	for i := 0; i < len(middlewares); i++ {
		h = middlewares[i](h)
	}
	return h
}

func AuthMiddleware(next AppHandler) AppHandler {

	return func(w http.ResponseWriter, r *http.Request) error {
		token := r.URL.Query().Get("token")

		if token != "123" {
			return fmt.Errorf("unauthorized")
		}

		return next(w, r)
	}
}

func LoggingMiddleware(next AppHandler) AppHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		start := time.Now()
		mw := &middlewareWriter{ResponseWriter: w}
		err := next(mw, r)
		duration := time.Since(start)
		fmt.Println("[", r.Method, "]", r.URL.Path, "time : ", duration, "status :", mw.code)
		return err
	}
}
