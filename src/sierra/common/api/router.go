package api

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

type MiddlewareFunc func(w *Writer, r *http.Request, p httprouter.Params) *Error

type Router struct {
	Router *httprouter.Router
	Server *http.Server
}

func NewRouter() *Router {
	router := httprouter.New()

	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			setCorsHeaders(w)
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})

	return &Router{
		Router: router,
	}
}

func (r *Router) ListenAndServe(addr string) error {
	r.Server = &http.Server{
		Handler:           r.Router,
		Addr:              addr,
		ReadHeaderTimeout: 5 * time.Second,
	}

	return r.Server.ListenAndServe()
}

func (r *Router) Shutdown() {
	logrus.Info("Shutting down HTTP server")
	err := r.Server.Shutdown(context.Background())
	if err != nil {
		logrus.WithError(err).Error("failed to shutdown http server")
	}
}

func panicHandler(w http.ResponseWriter) {
	panicErrInterface := recover()
	if panicErrInterface == nil {
		return
	}

	debug.PrintStack()

	var panicErr error
	var ok bool
	if panicErr, ok = panicErrInterface.(error); !ok {
		panicErr = fmt.Errorf("panic with value: %v", panicErrInterface)
	}

	aErr := NewInternalError(panicErr, "Panicked during handling of API route")
	if aErr != nil {
		NewWriter(w).writeError(aErr)
		return
	}
}

func RouteGET[T any](router *Router, authorization Authorization, endpoint string, h func(*Writer, *http.Request, httprouter.Params, *JWT) (*T, *Error)) {
	router.Router.GET(endpoint, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer panicHandler(w)
		logrus.WithField("params", p).Infof("Handling request - GET %s", endpoint)

		setCorsHeaders(w)

		writer := NewWriter(w)

		jwt, aErr := authorization.Protect(r)
		if aErr != nil {
			writer.writeError(aErr)
			return
		}

		var shouldShutdown bool
		response, aErr := h(writer, r, p, jwt)
		if aErr != nil && errors.Is(aErr.Err, shutdownError) {
			aErr = nil
			shouldShutdown = true
		}

		if aErr != nil {
			writer.writeError(aErr)
			return
		}

		if response != nil {
			aErr = writer.writeJSON(http.StatusOK, response)
			if aErr != nil {
				writer.writeError(aErr)
				return
			}
		}

		logrus.Infof("Handled request - GET %s", endpoint)

		if shouldShutdown {
			router.Shutdown()
		}
	})
}

func RoutePOST[T any, R any](router *Router, authorization Authorization, endpoint string, h func(*Writer, *http.Request, httprouter.Params, *JWT, T) (*R, *Error)) {
	router.Router.POST(endpoint, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer panicHandler(w)
		logrus.Infof("Handling request - POST %s", endpoint)

		setCorsHeaders(w)

		writer := NewWriter(w)

		jwt, aErr := authorization.Protect(r)
		if aErr != nil {
			writer.writeError(aErr)
			return
		}

		var request T
		aErr = UnmarshalBody(r, &request)
		if aErr != nil {
			writer.writeError(aErr)
			return
		}

		var shouldShutdown bool
		response, aErr := h(writer, r, p, jwt, request)
		if aErr != nil && errors.Is(aErr.Err, shutdownError) {
			aErr = nil
			shouldShutdown = true
		}

		if aErr != nil {
			writer.writeError(aErr)
			return
		}

		if response != nil {
			aErr = writer.writeJSON(http.StatusOK, response)
			if aErr != nil {
				writer.writeError(aErr)
				return
			}
		}

		logrus.Infof("Handled request POST %s", endpoint)

		if shouldShutdown {
			go router.Shutdown()
		}
	})
}

func UnmarshalBody(r *http.Request, v any) *Error {
	if r.Header.Get("Content-Type") != "application/json" {
		return nil
	}

	if r.ContentLength == 0 {
		return nil
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return NewInternalError(err, "failed reading request body")
	}

	if len(b) == 0 {
		return NewBadRequestError(nil, "Missing request body")
	}

	err = Unmarshal(b, v)
	if err != nil {
		return NewBadRequestError(err, "Bad request body")
	}

	return nil
}

func SuccessEndpoint(w *Writer, r *http.Request, p httprouter.Params, j *JWT) (*any, *Error) {
	return nil, nil
}

func setCorsHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Referer, User-Agent, Accept-Language, Content-Type")
}
