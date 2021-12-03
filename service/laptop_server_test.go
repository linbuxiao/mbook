package service

import (
	"context"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mbook/pb"
	"mbook/sample"
	"testing"
)

func TestLaptopServer_CreateLaptop(t *testing.T) {
	t.Parallel()

	laptopNoID := sample.NewLaptop()
	laptopNoID.Id = ""

	laptopInValidID := sample.NewLaptop()
	laptopInValidID.Id = "invalid-uuid"

	duplicateStore := NewInMemoryLaptopStore()
	laptopDuplicateID := sample.NewLaptop()

	err := duplicateStore.Save(laptopDuplicateID)
	require.NoError(t, err)
	testCases := []struct {
		name   string
		laptop *pb.Laptop
		store  LaptopStore
		code   codes.Code
	}{
		{
			name:   "success_with_id",
			laptop: sample.NewLaptop(),
			store:  NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "success_no_id",
			laptop: laptopNoID,
			store:  NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "failure_invalid_id",
			laptop: laptopInValidID,
			store:  NewInMemoryLaptopStore(),
			code:   codes.InvalidArgument,
		},
		{
			name:   "failure_duplicate_id",
			laptop: laptopDuplicateID,
			store:  duplicateStore,
			code:   codes.AlreadyExists,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			req := &pb.CreateLaptopRequest{
				Laptop: tc.laptop,
			}

			server := NewLapTopServer(tc.store)
			res, err := server.CreateLaptop(context.TODO(), req)
			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Id)
				if len(tc.laptop.Id) > 0 {
					require.Equal(t, tc.laptop.Id, res.Id)
				}
			} else {
				require.Error(t, err)
				require.Nil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tc.code, st.Code())
			}
		})
	}
}
