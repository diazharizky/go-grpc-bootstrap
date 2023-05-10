package repositories

import (
	"testing"

	"github.com/diazharizky/go-grpc-bootstrap/internal/app"
	"github.com/diazharizky/go-grpc-bootstrap/pb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type userRepositoryTestSuite struct {
	suite.Suite
	repo app.IUserRepository
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, &userRepositoryTestSuite{})
}

func (ts *userRepositoryTestSuite) SetupTest() {
	ts.repo = NewUserRepository()
}

func (ts userRepositoryTestSuite) TestList() {
	type want struct {
		err    error
		result []*pb.User
	}

	dummyUsers := []*pb.User{
		{
			Username:  "aominedaiki",
			FirstName: "Aomine",
			LastName:  "Daiki",
			Email:     "daiki.aomine@knb.com",
			Age:       18,
		},
		{
			Username:  "kiseryota",
			FirstName: "Kise",
			LastName:  "Ryota",
			Email:     "ryota.kise@knb.com",
			Age:       19,
		},
	}

	testCases := []struct {
		want
		name string
	}{
		{
			name: "Succeed test",
			want: want{
				result: dummyUsers,
			},
		},
	}

	for _, tc := range testCases {
		ts.T().Run(tc.name, func(t *testing.T) {
			assertions := assert.New(t)

			gotResult, gotErr := ts.repo.List()
			if gotErr != nil {
				assertions.Equal(tc.want.err, gotErr)
				return
			}

			assertions.Equal(tc.want.result, gotResult)
		})
	}
}

func (ts userRepositoryTestSuite) TestGet() {
	type want struct {
		err    error
		result []*pb.User
	}

	testCases := []struct {
		want
		name string
	}{
		{
			name: "Succeed test",
			want: want{
				result: []*pb.User{},
			},
		},
	}

	for _, tc := range testCases {
		ts.T().Run(tc.name, func(t *testing.T) {
			assertions := assert.New(t)

			gotResult, gotErr := ts.repo.Get()
			if gotErr != nil {
				assertions.Equal(tc.want.err, gotErr)
				return
			}

			assertions.Equal(tc.want.result, gotResult)
		})
	}
}

func (ts userRepositoryTestSuite) TestCreate() {
	testCases := []struct {
		name    string
		params  pb.User
		wantErr error
	}{
		{
			name:   "Succeed test",
			params: pb.User{},
		},
	}

	for _, tc := range testCases {
		ts.T().Run(tc.name, func(t *testing.T) {
			assertions := assert.New(t)

			gotErr := ts.repo.Create(tc.params)
			if gotErr != nil {
				assertions.Equal(tc.wantErr, gotErr)
			}
		})
	}
}

func (ts userRepositoryTestSuite) TestUpdate() {
	testCases := []struct {
		name    string
		params  pb.User
		wantErr error
	}{
		{
			name:   "Succeed test",
			params: pb.User{},
		},
	}

	for _, tc := range testCases {
		ts.T().Run(tc.name, func(t *testing.T) {
			assertions := assert.New(t)

			gotErr := ts.repo.Update(tc.params)
			if gotErr != nil {
				assertions.Equal(tc.wantErr, gotErr)
			}
		})
	}
}

func (ts userRepositoryTestSuite) TestDelete() {
	testCases := []struct {
		name    string
		param   string
		wantErr error
	}{
		{
			name:  "Succeed test",
			param: "some_username",
		},
	}

	for _, tc := range testCases {
		ts.T().Run(tc.name, func(t *testing.T) {
			assertions := assert.New(t)

			gotErr := ts.repo.Delete(tc.param)
			if gotErr != nil {
				assertions.Equal(tc.wantErr, gotErr)
			}
		})
	}
}
