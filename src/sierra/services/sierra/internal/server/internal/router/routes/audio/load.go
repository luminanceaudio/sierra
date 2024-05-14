package audio

import (
	"github.com/julienschmidt/httprouter"
	"github.com/luminanceaudio/sierra/src/sierra/common/api"
	"github.com/luminanceaudio/sierra/src/sierra/common/uri"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/modules/source"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/modules/sourcesample"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/sierradb"
	"io"
	"net/http"
	"strings"
)

func Load(w *api.Writer, r *http.Request, p httprouter.Params, j *api.JWT) (*any, *api.Error) {
	uriStr := p.ByName("uri")
	uriStr = strings.TrimPrefix(uriStr, "/")

	fileUri, err := uri.New(uriStr)
	if err != nil {
		return nil, api.NewBadRequestError(err, "Invalid sample URI")
	}

	sample, err := sourcesample.Get(r.Context(), fileUri)
	if err != nil {
		if sierradb.IsNotFound(err) {
			return nil, api.NewNotFoundError(err, "Sample not found")
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
			return nil, api.NewNotFoundError(err, "Source not found")
		}

		return nil, api.NewInternalError(err, "failed getting source from db")
	}

	audioFile, err := src.Open(fileUri.Path())
	if err != nil {
		return nil, api.NewBadRequestError(err, "Unable to load sample")
	}
	defer audioFile.Close()

	_, err = io.Copy(w, audioFile)
	if err != nil {
		return nil, api.NewInternalError(err, "Unable to load sample")
	}

	return nil, nil
}
