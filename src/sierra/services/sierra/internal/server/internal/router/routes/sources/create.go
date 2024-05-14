package sources

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"github.com/luminanceaudio/sierra/src/sierra/common/api"
	"github.com/luminanceaudio/sierra/src/sierra/common/uri"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/client/models"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/indexer"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/modules/source"
	"net/http"
)

func CreateSource(w *api.Writer, r *http.Request, p httprouter.Params, j *api.JWT, request *models.CreateSourceRequest) (*models.CreateSourceResponse, *api.Error) {
	uri, err := uri.New(request.Uri)
	if err != nil {
		return nil, api.NewBadRequestError(err, "Invalid source URI")
	}

	err = source.Create(r.Context(), uri)
	if err != nil {
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
