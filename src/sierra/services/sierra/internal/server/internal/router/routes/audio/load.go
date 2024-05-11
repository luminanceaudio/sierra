package audio

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"sierra/common/api"
	"sierra/common/uri"
	"sierra/services/sierra/internal/modules/source"
	"sierra/services/sierra/internal/modules/sourcesample"
	"sierra/services/sierra/internal/sierradb"
	"strings"
)

func Load(w *api.Writer, r *http.Request, p httprouter.Params, j *api.JWT) (*any, *api.Error) {
	uriStr := p.ByName("uri")
	uriStr = strings.TrimPrefix(uriStr, "/")

	fileUri, err := uri.New(uriStr)
	if err != nil {
		return nil, api.NewBadRequestError(err, "got bad uri")
	}

	sample, err := sourcesample.Get(r.Context(), fileUri)
	if err != nil {
		if sierradb.IsNotFound(err) {
			return nil, api.NewNotFoundError(err, "source sample not found")
		}

		return nil, api.NewInternalError(err, "failed getting source sample from db")
	}

	sourceUri, err := uri.New(sample.SourceUri)
	if err != nil {
		return nil, api.NewInternalError(err, "got bad source uri from db")
	}

	src, err := source.Get(r.Context(), sourceUri)
	if err != nil {
		if sierradb.IsNotFound(err) {
			return nil, api.NewNotFoundError(err, "source not found")
		}

		return nil, api.NewInternalError(err, "failed getting source from db")
	}

	audioFile, err := src.Open(fileUri.Path())
	if err != nil {
		return nil, api.NewBadRequestError(err, "got bad source")
	}
	defer audioFile.Close()

	_, err = io.Copy(w, audioFile)
	if err != nil {
		return nil, api.NewInternalError(err, "failed to read audio")
	}

	return nil, nil
}