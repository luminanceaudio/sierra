package samples

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sierra/common/api"
	"sierra/services/sierra/client/models"
	"sierra/services/sierra/internal/modules/sourcesample"
)

func GetSamples(w *api.Writer, r *http.Request, p httprouter.Params, j *api.JWT) (*models.GetSamplesResponse, *api.Error) {
	samples, err := sourcesample.GetAll(r.Context())
	if err != nil {
		return nil, api.NewInternalError(err, "failed getting samples")
	}

	return &models.GetSamplesResponse{
		Samples: samples,
	}, nil
}