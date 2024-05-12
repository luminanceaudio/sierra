package sources

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sierra/common/api"
	"sierra/common/uri"
	"sierra/services/sierra/client/models"
	"sierra/services/sierra/internal/modules/source"
)

func CreateSource(w *api.Writer, r *http.Request, p httprouter.Params, j *api.JWT, request *models.CreateSourceRequest) (*models.CreateSourceResponse, *api.Error) {
	uri, err := uri.New(request.Uri)
	if err != nil {
		return nil, api.NewBadRequestError(err, "Invalid source URI")
	}

	err = source.Create(r.Context(), uri)
	if err != nil {
		return nil, api.NewInternalError(err, "failed getting sources")
	}

	return &models.CreateSourceResponse{}, nil
}
