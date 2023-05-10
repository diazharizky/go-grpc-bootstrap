package users

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/diazharizky/go-grpc-bootstrap/internal/app"
	"github.com/diazharizky/go-grpc-bootstrap/internal/app/mocks"
	"github.com/diazharizky/go-grpc-bootstrap/pb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/types/known/emptypb"
)

type listTestSuite struct {
	suite.Suite
	mockUserRepository *mocks.IUserRepository
	appCtx             *app.Context
}

func TestListTestSuite(t *testing.T) {
	suite.Run(t, &listTestSuite{})
}

func (ts *listTestSuite) resetMocks() {
	ts.mockUserRepository = &mocks.IUserRepository{}

	ts.appCtx = &app.Context{
		UserRepository: ts.mockUserRepository,
	}
}

func (ts listTestSuite) TestList() {
	type want struct {
		err    error
		result *pb.ListResponse
	}

	type mockListUser struct {
		users []*pb.User
		err   error
	}

	testCases := []struct {
		mockListUser
		want
		name string
	}{
		{
			name: "Succeed test",
			mockListUser: mockListUser{
				users: []*pb.User{},
			},
			want: want{
				result: &pb.ListResponse{
					Ok:    true,
					Users: &pb.UserList{},
				},
			},
		},
		{
			name: "Failed test",
			mockListUser: mockListUser{
				users: []*pb.User{},
				err:   errors.New("Internal server error"),
			},
			want: want{
				result: &pb.ListResponse{},
			},
		},
	}

	for idx, tc := range testCases {
		testNo := idx + 1

		ts.T().Run(
			fmt.Sprintf("(%d) %s", testNo, tc.name),
			func(t *testing.T) {
				ts.resetMocks()

				ctx := context.Background()

				client, closer := testServer(ctx, ts.appCtx)
				defer closer()

				assertions := assert.New(t)

				ts.mockUserRepository.
					On("List").
					Return(tc.mockListUser.users, tc.mockListUser.err)

				gotResult, gotErr := client.List(ctx, &emptypb.Empty{})
				if gotErr != nil {
					fmt.Println("error", gotErr.Error())
					assertions.Equal(tc.want.err, gotErr)
				}

				assertions.Equal(tc.want.result, gotResult)
			},
		)
	}
}
