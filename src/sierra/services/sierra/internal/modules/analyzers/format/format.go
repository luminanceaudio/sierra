package format

type Type = string

const (
	Wav Type = "wav"
)

type Format struct {
	Type Type
}

func NewFormat(formatType Type) *Format {
	return &Format{
		Type: formatType,
	}
}
