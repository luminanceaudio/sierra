package sources

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sierra/common/api"
	"sierra/common/uri"
	"sierra/services/sierra/client/models"
	"sierra/services/sierra/internal/modules/source"
	"sierra/services/sierra/internal/modules/sourcesample"
)

func DeleteSource(w *api.Writer, r *http.Request, p httprouter.Params, j *api.JWT, request *models.DeleteSourceRequest) (*models.DeleteSourceResponse, *api.Error) {
	uri, err := uri.New(request.Uri)
	if err != nil {
		return nil, api.NewBadRequestError(err, "Invalid source URI")
	}

	err = sourcesample.DeleteBySource(r.Context(), uri)
	if err != nil {
		return nil, api.NewInternalError(err, "failed deleting sources")
	}

	err = source.Delete(r.Context(), uri)
	if err != nil {
		return nil, api.NewInternalError(err, "failed getting sources")
	}

	return &models.DeleteSourceResponse{}, nil
}
