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

type MailBackupHandlerUnitSuite struct {
	tester.Suite
}

func TestMailBackupHandlerUnitSuite(t *testing.T) {
	suite.Run(t, &MailBackupHandlerUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *MailBackupHandlerUnitSuite) TestHandler_CanSkipItemFailure() {
	t := suite.T()

	h := newMailBackupHandler(api.Client{})
	cause, result := h.CanSkipItemFailure(nil, "", "", control.Options{})

	assert.False(t, result)
	assert.Equal(t, fault.SkipCause(""), cause)
}
