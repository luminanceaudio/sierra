package models

import (
	"github.com/luminanceaudio/sierra/src/sierra/common/sha256"
)

func NewSample(uri string, sha256 sha256.Sha256, format string, sourceUri string, duration int64) *Sample {
	return &Sample{
		Uri:       uri,
		Sha256:    sha256,
		Format:    format,
		SourceUri: sourceUri,
		Duration:  duration,
	}
}
