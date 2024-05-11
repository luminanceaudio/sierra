package api

import (
	"fmt"
	"io"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

type Router struct {
	Router *httprouter.Router
}

func NewRouter() *Router {
	return &Router{
		Router: httprouter.New(),
	}
}

func (r *Router) ListenAndServe(addr string) error {
	s := http.Server{
		Handler:           r.Router,
		Addr:              addr,
		ReadHeaderTimeout: 5 * time.Second,
	}

	return s.ListenAndServe()
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

		writer := NewWriter(w)

		jwt, aErr := authorization.Protect(r)
		if aErr != nil {
			writer.writeError(aErr)
			return
		}

		response, aErr := h(writer, r, p, jwt)
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
	})
}

func RoutePOST[T any, R any](router *Router, authorization Authorization, endpoint string, h func(*Writer, *http.Request, httprouter.Params, *JWT, T) (*R, *Error)) {
	router.Router.POST(endpoint, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer panicHandler(w)
		logrus.Infof("Handling request - POST %s", endpoint)

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

		response, aErr := h(writer, r, p, jwt, request)
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
		return NewBadRequestError(nil, "missing request body")
	}

	err = Unmarshal(b, v)
	if err != nil {
		return NewBadRequestError(err, "failed unmarshalling body")
	}

	return nil
}

func SuccessEndpoint(w *Writer, r *http.Request, p httprouter.Params, j *JWT) (*any, *Error) {
	return nil, nil
}
