package samples

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/luminanceaudio/sierra/src/sierra/common/api"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/client/models"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/modules/sourcesample"
	"net/http"
	"strconv"
)

const (
	maxPageSize = 500
)

func SearchSamples(w *api.Writer, r *http.Request, p httprouter.Params, j *api.JWT) (*models.GetSamplesResponse, *api.Error) {
	pageStr := p.ByName("page")
	if pageStr == "" {
		return nil, api.NewBadRequestError(nil, "Must specify page parameter")
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return nil, api.NewBadRequestError(err, "Page must be a number")
	}

	if page <= 0 {
		return nil, api.NewBadRequestError(nil, "Page must be a positive number")
	}

	sizeStr := r.URL.Query().Get("size")
	if sizeStr == "" {
		return nil, api.NewBadRequestError(nil, "Must specify size parameter")
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return nil, api.NewBadRequestError(err, "Page must be a number")
	}

	if size <= 0 {
		return nil, api.NewBadRequestError(nil, "Size must be a positive number")
	}

	if size > maxPageSize {
		return nil, api.NewBadRequestError(nil, fmt.Sprintf("Size must be less than or equal to %d", maxPageSize))
	}

	query := p.ByName("query")

	samples, err := sourcesample.Query(r.Context(), query, page, size)
	if err != nil {
		return nil, api.NewInternalError(err, "failed getting samples")
	}

	return &models.GetSamplesResponse{
		Samples: samples,
	}, nil
}

func GetSamplesCount(w *api.Writer, r *http.Request, p httprouter.Params, j *api.JWT) (*models.GetSamplesCountResponse, *api.Error) {
	query := p.ByName("query")

	count, err := sourcesample.QueryCount(r.Context(), query)
	if err != nil {
		return nil, api.NewInternalError(err, "failed getting samples")
	}

	return &models.GetSamplesCountResponse{
		Count: int64(count),
	}, nil
}
