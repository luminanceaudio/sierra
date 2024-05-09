package format

import (
	"github.com/sirupsen/logrus"
	"os"
)

// TODO: Make it so when a new analyzer version is out it will invalidate the existing data

const (
	FormatWav = "WAV"
)

func AnalyzeFormat(sampleFile *os.File) (*Format, bool, error) {
	format, valid, err := TryParseWav(sampleFile)
	if err != nil {
		logrus.WithError(err).Error("failed to analyze wav format")
	}
	if valid {
		return format, true, nil
	}

	return nil, false, nil
}
