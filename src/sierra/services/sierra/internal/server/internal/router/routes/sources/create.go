package sources

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sierra/common/api"
	"sierra/common/uri"
	"sierra/services/sierra/client/models"
	"sierra/services/sierra/internal/indexer"
	"sierra/services/sierra/internal/modules/source"
	"sierra/services/sierra/internal/sierradb/sierraent"
)

func CreateSource(w *api.Writer, r *http.Request, p httprouter.Params, j *api.JWT, request *models.CreateSourceRequest) (*models.CreateSourceResponse, *api.Error) {
	uri, err := uri.New(request.Uri)
	if err != nil {
		return nil, api.NewBadRequestError(err, "Invalid source URI")
	}

	err = source.Create(r.Context(), uri)
	if err != nil {
		if sierraent.IsConstraintError(err) {
			return nil, api.NewBadRequestError(err, "Source already exists")
		}
		return nil, api.NewInternalError(err, "failed creating source")
	}

	src, err := source.Get(r.Context(), uri)
	if err != nil {
		return nil, api.NewInternalError(err, "failed getting source")
	}

	err = indexer.Singleton().Index(context.WithoutCancel(r.Context()), src, false, true)
	if err != nil {
		return nil, api.NewInternalError(err, "failed indexing sources")
	}

	return &models.CreateSourceResponse{}, nil
}
