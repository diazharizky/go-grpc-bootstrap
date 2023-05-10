package users

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type listTestSuite struct {
	suite.Suite
}

func TestListTestSuite(t *testing.T) {
	suite.Run(t, &listTestSuite{})
}

func (ts listTestSuite) TestList() {
	testCases := []struct {
		name string
	}{}

	for idx, tc := range testCases {
		testNo := idx + 1

		ts.T().Run(
			fmt.Sprintf("(%d) %s", testNo, tc.name),
			func(t *testing.T) {
				ctx := context.Background()

				client, closer := testServer(ctx)
				defer closer()

				client.List(ctx, nil)
			},
		)
	}
}
