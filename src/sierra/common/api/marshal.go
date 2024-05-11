package api

import (
	"encoding/json"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func Marshal(v any) ([]byte, error) {
	if msg, isProtoMessage := interface{}(v).(proto.Message); isProtoMessage {
		return protojson.MarshalOptions{EmitUnpopulated: false}.Marshal(msg)
	}

	return json.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	if msg, isProtoMessage := interface{}(v).(proto.Message); isProtoMessage {
		return protojson.Unmarshal(data, msg)
	}

	return json.Unmarshal(data, v)
}
