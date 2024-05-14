package sources

import (
	"github.com/julienschmidt/httprouter"
	"github.com/luminanceaudio/sierra/src/sierra/common/api"
	"github.com/luminanceaudio/sierra/src/sierra/common/uri"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/client/models"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/modules/source"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/modules/sourcesample"
	"net/http"
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
