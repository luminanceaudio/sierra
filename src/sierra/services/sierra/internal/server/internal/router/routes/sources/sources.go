package sources

import (
	"github.com/julienschmidt/httprouter"
	"github.com/luminanceaudio/sierra/src/sierra/common/api"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/client/models"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/modules/source"
	"net/http"
)

func GetSources(w *api.Writer, r *http.Request, p httprouter.Params, j *api.JWT) (*models.GetSourcesResponse, *api.Error) {
	sources, err := source.GetAllAsModel(r.Context())
	if err != nil {
		return nil, api.NewInternalError(err, "failed getting sources")
	}

	return &models.GetSourcesResponse{
		Sources: sources,
	}, nil
}
