package format

import (
	"sierra/services/sierra/internal/format"
)

type Format struct {
	Type          format.Type
	Channels      uint16
	SampleRate    uint32
	BitsPerSample uint16
	Samples       []int16
}

func newFormat(formatType format.Type, channels uint16, sampleRate uint32, bitsPerSample uint16, samples []int16) *Format {
	return &Format{
		Type:          formatType,
		Channels:      channels,
		SampleRate:    sampleRate,
		BitsPerSample: bitsPerSample,
		Samples:       samples,
	}
}
