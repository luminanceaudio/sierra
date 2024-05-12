package models

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func NewSource(uri string, createTime time.Time) *Source {
	return &Source{
		Uri:        uri,
		CreateTime: timestamppb.New(createTime),
	}
}
