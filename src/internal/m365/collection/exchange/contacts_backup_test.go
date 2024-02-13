package exchange

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type ContactsBackupHandlerUnitSuite struct {
	tester.Suite
}

func TestContactsBackupHandlerUnitSuite(t *testing.T) {
	suite.Run(t, &ContactsBackupHandlerUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ContactsBackupHandlerUnitSuite) TestHandler_CanSkipItemFailure() {
	t := suite.T()

	h := newContactBackupHandler(api.Client{})
	cause, result := h.CanSkipItemFailure(nil, "", "", control.Options{})

	assert.False(t, result)
	assert.Equal(t, fault.SkipCause(""), cause)
}
