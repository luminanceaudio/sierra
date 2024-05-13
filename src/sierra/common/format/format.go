package format

const (
	TypeWav Type = "wav"
)

type Type = string

var supported = map[Type]bool{
	TypeWav: true,
}

func IsSupported(t Type) bool {
	_, ok := supported[t]
	return ok
}

type Format struct {
	Type          Type
	Channels      uint16
	SampleRate    uint32
	BitsPerSample uint16
	Samples       []int16
}

func newFormat(formatType Type, channels uint16, sampleRate uint32, bitsPerSample uint16, samples []int16) *Format {
	return &Format{
		Type:          formatType,
		Channels:      channels,
		SampleRate:    sampleRate,
		BitsPerSample: bitsPerSample,
		Samples:       samples,
	}
}
