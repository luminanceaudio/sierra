package sample

type Sample struct {
	Name string
}

func New(name string) *Sample {
	return &Sample{
		Name: name,
	}
}
