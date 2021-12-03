package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"mbook/pb"
)

type LaptopServer struct {
	Store LaptopStore
}

func NewLapTopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{
		Store: store,
	}
}

func (server *LaptopServer) CreateLaptop(context context.Context, request *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	laptop := request.GetLaptop()
	log.Printf("receive a create-laptop request request with id: %s", laptop.Id)

	if len(laptop.Id) > 0 {
		_, err := uuid.Parse(laptop.Id)

		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop id is not a valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()

		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new laptop ID: %v", err)
		}

		laptop.Id = id.String()
	}

	err := server.Store.Save(laptop)

	if err != nil {
		code := codes.Internal

		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}

		return nil, status.Errorf(code, "cannot save laptop to the store: %s", err)
	}

	log.Printf("saved laptop with id: %s", laptop.Id)

	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}

	return res, nil
}
