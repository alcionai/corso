package backup

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/utils/testdata"
	"github.com/alcionai/corso/src/internal/tester"
	dtd "github.com/alcionai/corso/src/pkg/backup/details/testdata"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type BackupUnitSuite struct {
	tester.Suite
}

func TestBackupUnitSuite(t *testing.T) {
	suite.Run(t, &BackupUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *BackupUnitSuite) TestGenericDetailsCore() {
	t := suite.T()

	expected := append(
		append(
			dtd.GetItemsForVersion(
				t,
				path.ExchangeService,
				path.EmailCategory,
				0,
				-1),
			dtd.GetItemsForVersion(
				t,
				path.ExchangeService,
				path.EventsCategory,
				0,
				-1)...),
		dtd.GetItemsForVersion(
			t,
			path.ExchangeService,
			path.ContactsCategory,
			0,
			-1)...)

	ctx, flush := tester.NewContext(t)
	defer flush()

	bg := testdata.VersionedBackupGetter{
		Details: dtd.GetDetailsSetForVersion(t, 0),
	}

	sel := selectors.NewExchangeBackup([]string{"user-id"})
	sel.Include(sel.AllData())

	output, err := genericDetailsCore(
		ctx,
		bg,
		"backup-ID",
		sel.Selector,
		control.DefaultOptions())
	assert.NoError(t, err, clues.ToCore(err))
	assert.ElementsMatch(t, expected, output.Entries)
}
