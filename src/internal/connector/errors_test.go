package connector_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"

	connect "github.com/alcionai/corso/internal/connector"
)

type GraphConnectorErrorSuite struct {
	suite.Suite
}

func TestGraphConnectorErrorSuite(t *testing.T) {
	suite.Run(t, new(GraphConnectorErrorSuite))
}

func (suite *GraphConnectorErrorSuite) TestCreateErrorList() {
	var list connect.ErrorList
	list = make([]error, 0)
	suite.Equal(len(list), 0)
	list = append(list, errors.New("This one"))
	suite.Equal(len(list), 1)
	err2 := errors.New("I have two")
	list = append(list, err2)
	suite.Equal(len(list), 2)

}

func (suite *GraphConnectorErrorSuite) TestErrorListFormatCheck() {
	listing := connect.NewErrorList()
	emptyReturn := listing.GetErrors()
	suite.Equal(len(emptyReturn), 0)
	err1 := errors.New("cauliflower")
	listing = append(listing, err1)
	suite.Equal(len(listing.GetErrors()), 20)
}

func (suite *GraphConnectorErrorSuite) TestErrorWrapAndAppend() {
	listing := connect.NewErrorList()
	err := errors.New("Network failure")
	listing = connect.WrapErrorAndAppend("arch789", err, listing)
	suite.Equal(len(listing), 1)
	suite.T().Logf("%s", listing.GetErrors())
}
