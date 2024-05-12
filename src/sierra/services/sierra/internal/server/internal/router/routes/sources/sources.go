package sources

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sierra/common/api"
	"sierra/services/sierra/client/models"
	"sierra/services/sierra/internal/modules/source"
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
