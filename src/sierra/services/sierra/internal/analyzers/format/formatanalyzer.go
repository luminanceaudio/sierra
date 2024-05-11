package format

import (
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"slices"
)

// TODO: Make it so when a new analyzer version is out it will invalidate the existing data

var blacklistedFilenames = []string{
	".DS_Store",
}

// AnalyzeFormat returns a parsed format of an audio file.
// Returns false if format is not supported.
func AnalyzeFormat(sampleFile *os.File) (*Format, bool, error) {
	if slices.Contains(blacklistedFilenames, filepath.Base(sampleFile.Name())) {
		return nil, false, nil
	}

	format, valid, err := TryParseWav(sampleFile)
	if err != nil {
		logrus.WithField("path", sampleFile.Name()).WithError(err).Error("failed to analyze wav format")
	}
	if valid {
		return format, true, nil
	}

	return nil, false, nil
}
