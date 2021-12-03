package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtobufToJSON(message proto.Message) ([]byte, error) {
	marshaller := protojson.MarshalOptions{
		Multiline: true,
	}
	data, err := marshaller.Marshal(message)

	return data, err
}
