package format

type Type = string

type Format interface {
}

var supported = map[Type]Format{
	TypeWav: &Wav{},
}

func IsSupported(t Type) bool {
	_, ok := supported[t]
	return ok
}
