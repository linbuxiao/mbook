package serializer

import (
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"mbook/pb"
	"mbook/sample"
	"testing"
)

func TestSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop() // 由 pb 生成

	err := WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}

	err = ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)
}
