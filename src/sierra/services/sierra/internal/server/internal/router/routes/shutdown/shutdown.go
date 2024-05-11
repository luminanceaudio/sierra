package shutdown

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sierra/common/api"
)

func Shutdown(w *api.Writer, r *http.Request, p httprouter.Params, j *api.JWT, request any) (*any, *api.Error) {
	return nil, api.NewShutdownError()
}
