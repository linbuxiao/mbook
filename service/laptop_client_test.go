package service

import (
	"context"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"mbook/pb"
	"mbook/sample"
	"mbook/serializer"
	"net"
	"testing"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	testLaptopServer, serverAddr := StartTestLaptopServer(t)

	testLaptopClient := newTestLaptopClient(t, serverAddr)

	laptop := sample.NewLaptop()

	expectedID := laptop.Id

	req := &pb.CreateLaptopRequest{Laptop: laptop}

	res, err := testLaptopClient.CreateLaptop(context.Background(), req)

	require.NoError(t, err)

	require.NotNil(t, res)

	require.Equal(t, expectedID, res.Id)

	other, err := testLaptopServer.Store.Find(expectedID)

	require.NoError(t, err)

	require.NotNil(t, other)

	requireSameLaptop(t, laptop, other)
}

func StartTestLaptopServer(t *testing.T) (*LaptopServer, string) {
	testLaptopServer := NewLapTopServer(NewInMemoryLaptopStore())

	grpcServer := grpc.NewServer()

	pb.RegisterLaptopServiceServer(grpcServer, testLaptopServer)

	listener, err := net.Listen("tcp", ":0") // random port
	require.NoError(t, err)

	go grpcServer.Serve(listener)

	return testLaptopServer, listener.Addr().String()
}

func newTestLaptopClient(t *testing.T, serverAddr string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())

	require.NoError(t, err)

	return pb.NewLaptopServiceClient(conn)
}

func requireSameLaptop(t *testing.T, laptop1 *pb.Laptop, laptop2 *pb.Laptop) {
	json1, err := serializer.ProtobufToJSON(laptop1)
	require.NoError(t, err)

	json2, err := serializer.ProtobufToJSON(laptop2)
	require.NoError(t, err)

	require.Equal(t, json1, json2)
}
