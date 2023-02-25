package api_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/graph/api"
	"github.com/alcionai/corso/src/internal/tester"
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
	tester.Suite
}

func TestAPIUnitSuite(t *testing.T) {
	s := &APIUnitSuite{
		Suite: tester.NewUnitSuite(t),
	}
	suite.Run(t, s)
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

			suite.Run(name, func() {
				t := suite.T()

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

// TestIsLinkValid check to verify is nextLink guard check for logging
// Related to: https://github.com/alcionai/corso/issues/2520
//
//nolint:lll
func (suite *APIUnitSuite) TestIsLinkValid() {
	invalidString := `https://graph.microsoft.com/v1.0/users//mailFolders//messages/microsoft.graph.delta()?$select=id%2CisRead`
	tests := []struct {
		name        string
		inputString string
		isValid     assert.BoolAssertionFunc
	}{
		{
			name:        "Empty",
			inputString: emptyLink,
			isValid:     assert.True,
		},
		{
			name:        "Invalid",
			inputString: invalidString,
			isValid:     assert.False,
		},
		{
			name:        "Valid",
			inputString: `https://graph.microsoft.com/v1.0/users/aPerson/mailFolders/AMessage/messages/microsoft.graph.delta()?$select=id%2CisRead`,
			isValid:     assert.True,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			got := api.IsNextLinkValid(test.inputString)
			test.isValid(suite.T(), got)
		})
	}
}
