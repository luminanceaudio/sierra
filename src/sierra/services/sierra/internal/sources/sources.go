package sources

import "sierra/services/sierra/internal/source"

type Sources struct {
	List []*source.OneofSource `json:"list"`
}

func New() *Sources {
	return &Sources{}
}

func (s *Sources) Add(source *source.OneofSource) {
	s.List = append(s.List, source)
}

func (s *Sources) Purge() {
	s.List = []*source.OneofSource{}
}
