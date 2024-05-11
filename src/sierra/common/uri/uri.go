package uri

import (
	"fmt"
	"strings"
)

const (
	uriSeparator = "://"
)

type URI struct {
	scheme string
	path   string
}

func New(uriStr string) (*URI, error) {
	split := strings.SplitN(uriStr, uriSeparator, 2)
	if len(split) != 2 {
		return nil, fmt.Errorf("invalid source uri: %s", uriStr)
	}

	return &URI{
		scheme: split[0],
		path:   split[1],
	}, nil
}

func NewFromParts(schema, path string) *URI {
	return &URI{
		scheme: schema,
		path:   path,
	}
}

func (u *URI) String() string {
	return u.scheme + uriSeparator + u.path
}

func (u *URI) Scheme() string {
	return u.scheme
}

func (u *URI) Path() string {
	return u.path
}
