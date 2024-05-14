package shutdown

import (
	"github.com/julienschmidt/httprouter"
	"github.com/luminanceaudio/sierra/src/sierra/common/api"
	"net/http"
)

func Shutdown(w *api.Writer, r *http.Request, p httprouter.Params, j *api.JWT, request any) (*any, *api.Error) {
	return nil, api.NewShutdownError()
}
