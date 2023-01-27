package api_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/graph/api"
)

type mockNextLink struct {
	nextLink *string
}

func (l mockNextLink) GetOdataNextLink() *string {
	return l.nextLink
}

type mockDeltaNextLink struct {
	mockNextLink
	deltaLink *string
}

func (l mockDeltaNextLink) GetOdataDeltaLink() *string {
	return l.deltaLink
}

type testInput struct {
	name         string
	inputLink    *string
	expectedLink string
}

// Needs to be var not const so we can take the address of it.
var (
	emptyLink = ""
	link      = "foo"
	link2     = "bar"

	nextLinkInputs = []testInput{
		{
			name:         "empty",
			inputLink:    &emptyLink,
			expectedLink: "",
		},
		{
			name:         "nil",
			inputLink:    nil,
			expectedLink: "",
		},
		{
			name:         "non_empty",
			inputLink:    &link,
			expectedLink: link,
		},
	}
)

type APIUnitSuite struct {
	suite.Suite
}

func TestAPIUnitSuite(t *testing.T) {
	suite.Run(t, new(APIUnitSuite))
}

func (suite *APIUnitSuite) TestNextLink() {
	for _, test := range nextLinkInputs {
		suite.T().Run(test.name, func(t *testing.T) {
			l := mockNextLink{nextLink: test.inputLink}
			assert.Equal(t, test.expectedLink, api.NextLink(l))
		})
	}
}

func (suite *APIUnitSuite) TestNextAndDeltaLink() {
	deltaTable := []testInput{
		{
			name:         "empty",
			inputLink:    &emptyLink,
			expectedLink: "",
		},
		{
			name:         "nil",
			inputLink:    nil,
			expectedLink: "",
		},
		{
			name: "non_empty",
			// Use a different link so we can see if the results get swapped or something.
			inputLink:    &link2,
			expectedLink: link2,
		},
	}

	for _, next := range nextLinkInputs {
		for _, delta := range deltaTable {
			name := strings.Join([]string{next.name, "next", delta.name, "delta"}, "_")

			suite.T().Run(name, func(t *testing.T) {
				l := mockDeltaNextLink{
					mockNextLink: mockNextLink{nextLink: next.inputLink},
					deltaLink:    delta.inputLink,
				}
				gotNext, gotDelta := api.NextAndDeltaLink(l)

				assert.Equal(t, next.expectedLink, gotNext)
				assert.Equal(t, delta.expectedLink, gotDelta)
			})
		}
	}
}
