package samples

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/luminanceaudio/sierra/src/sierra/common/api"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/client/models"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/modules/sourcesample"
	"net/http"
)

const (
	maxPageSize = 500
)

func QuerySamples(w *api.Writer, r *http.Request, p httprouter.Params, j *api.JWT, req models.QuerySamplesRequest) (*models.QuerySamplesResponse, *api.Error) {
	if req.Page == 0 {
		return nil, api.NewBadRequestError(nil, "Must specify page parameter")
	}

	if req.Size <= 0 {
		return nil, api.NewBadRequestError(nil, "Size must be a positive number")
	}

	if req.Size > maxPageSize {
		return nil, api.NewBadRequestError(nil, fmt.Sprintf("Size must be less than or equal to %d", maxPageSize))
	}

	samples, err := sourcesample.Query(r.Context(), req.Query, int(req.Page), int(req.Size), req.SortDirection, req.SortColumn)
	if err != nil {
		return nil, api.NewInternalError(err, "failed getting samples")
	}

	return &models.QuerySamplesResponse{
		Samples: samples,
	}, nil
}

func QuerySamplesCount(w *api.Writer, r *http.Request, p httprouter.Params, j *api.JWT, req models.QuerySamplesCountRequest) (*models.QuerySamplesCountResponse, *api.Error) {
	count, err := sourcesample.QueryCount(r.Context(), req.Query)
	if err != nil {
		return nil, api.NewInternalError(err, "failed getting samples")
	}

	return &models.QuerySamplesCountResponse{
		Count: int64(count),
	}, nil
}
