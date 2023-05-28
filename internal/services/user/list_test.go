package user_test

import (
	"context"

	"github.com/diazharizky/go-grpc-bootstrap/internal/app"
	"github.com/diazharizky/go-grpc-bootstrap/pb"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ = Describe("List handler test", func() {
	type mockListUser struct {
		users []*pb.User
		err   error
	}

	type testCase struct {
		mockListUser

		name       string
		assertions func(*pb.ListResponse, error)
	}

	seedUsers := []*pb.User{
		{
			Id:       1,
			Username: "foo",
			FullName: "Foo",
			Email:    "foo@mail.com",
		},
		{
			Id:       2,
			Username: "bar",
			FullName: "Bar",
			Email:    "bar@mail.com",
		},
	}

	contexts := []struct {
		name      string
		testCases []testCase
	}{
		{
			name: "Succeed",
			testCases: []testCase{
				{
					name: "Succeed",
					mockListUser: mockListUser{
						users: seedUsers,
						err:   nil,
					},
					assertions: func(gotResp *pb.ListResponse, gotError error) {
						It("User should be match and error should be nil", func() {
							gotUsers := gotResp.Users

							for idx, gotUser := range gotUsers {
								Expect(gotUser.Id).To(BeEquivalentTo(seedUsers[idx].Id))
								Expect(gotUser.Username).To(BeEquivalentTo(seedUsers[idx].Username))
								Expect(gotUser.FullName).To(BeEquivalentTo(seedUsers[idx].FullName))
								Expect(gotUser.Email).To(BeEquivalentTo(seedUsers[idx].Email))
							}

							Expect(gotError).To(BeNil())
						})
					},
				},
			},
		},
	}

	for _, ctx := range contexts {
		Context(ctx.name, func() {
			for _, tc := range ctx.testCases {
				ctx := context.Background()

				mockCtrl := gomock.NewController(GinkgoT())
				defer mockCtrl.Finish()

				mockUserRepository := app.NewMockIUserRepository(mockCtrl)

				appCtx := &app.Context{
					UserRepository: mockUserRepository,
				}

				mockUserRepository.
					EXPECT().
					List().
					Return(tc.mockListUser.users, tc.mockListUser.err)

				client, closer := testServer(ctx, appCtx)
				defer closer()

				tc.assertions(client.List(ctx, &emptypb.Empty{}))
			}
		})
	}
})
