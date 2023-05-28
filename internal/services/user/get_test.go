package user_test

import (
	"context"

	"github.com/diazharizky/go-grpc-bootstrap/internal/app"
	"github.com/diazharizky/go-grpc-bootstrap/pb"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get handler test", func() {
	type mockGetUser struct {
		username string
		user     *pb.User
		err      error
	}

	type testCase struct {
		mockGetUser

		name       string
		assertions func(*pb.GetResponse, error)
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
					mockGetUser: mockGetUser{
						username: seedUser[0].Username,
						user:     seedUser[0],
						err:      nil,
					},
					assertions: func(gotResp *pb.GetResponse, gotError error) {
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
					Get(tc.mockGetUser.username).
					Return(tc.mockGetUser.user, tc.mockGetUser.err)

				client, closer := testServer(ctx, appCtx)
				defer closer()

				param := &pb.GetUsernameParam{
					Username: tc.mockGetUser.username,
				}
				tc.assertions(client.Get(ctx, param))
			}
		})
	}
})
