package user_test

import (
	"context"

	"github.com/diazharizky/go-grpc-bootstrap/internal/app"
	"github.com/diazharizky/go-grpc-bootstrap/pb"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create handler test", func() {
	type mockCreateUser struct {
		newUser *pb.CreateParams
		err     error
	}

	type testCase struct {
		mockCreateUser

		name       string
		assertions func(*pb.CreateResponse, error)
	}

	seedCreateParams := []*pb.CreateParams{
		{
			Username: "foo",
			FullName: "Foo",
			Email:    "foo@mail.com",
		},
	}

	seedUser := []*pb.User{
		{
			Id:       1,
			Username: "foo",
			FullName: "Foo",
			Email:    "foo@mail.com",
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
					mockCreateUser: mockCreateUser{
						newUser: seedCreateParams[0],
						err:     nil,
					},
					assertions: func(gotResp *pb.CreateResponse, gotError error) {
						It("User should be match and error should be nil", func() {
							gotUser := gotResp.User

							Expect(gotUser.Id).To(BeEquivalentTo(seedUser[0].Id))
							Expect(gotUser.Username).To(BeEquivalentTo(seedUser[0].Username))
							Expect(gotUser.FullName).To(BeEquivalentTo(seedUser[0].FullName))
							Expect(gotUser.Email).To(BeEquivalentTo(seedUser[0].Email))

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
					Create(&pb.User{
						Username: tc.mockCreateUser.newUser.Username,
						FullName: tc.mockCreateUser.newUser.FullName,
						Email:    tc.mockCreateUser.newUser.Email,
					}).
					Do(func(newUser *pb.User) *pb.User {
						newUser.Id = seedUser[0].Id
						return newUser
					}).
					Return(tc.mockCreateUser.err)

				client, closer := testServer(ctx, appCtx)
				defer closer()

				tc.assertions(client.Create(ctx, tc.mockCreateUser.newUser))
			}
		})
	}
})
