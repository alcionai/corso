package exchange

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type ExchangeRestoreUnitSuite struct {
	tester.Suite
}

func TestExchangeRestoreUnitSuite(t *testing.T) {
	suite.Run(t, &ExchangeRestoreUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ExchangeRestoreUnitSuite) TestGetRestoreResource() {
	var (
		id        = "id"
		name      = "name"
		cfgWithPR = control.DefaultRestoreConfig(dttm.HumanReadable)
	)

	cfgWithPR.ProtectedResource = id

	table := []struct {
		name           string
		cfg            control.RestoreConfig
		orig           idname.Provider
		cache          map[string]string
		expectErr      assert.ErrorAssertionFunc
		expectProvider assert.ValueAssertionFunc
		expectID       string
		expectName     string
	}{
		{
			name:       "use original",
			cfg:        control.DefaultRestoreConfig(dttm.HumanReadable),
			orig:       idname.NewProvider("oid", "oname"),
			expectErr:  assert.NoError,
			expectID:   "oid",
			expectName: "oname",
		},
		{
			name:       "use new",
			cfg:        cfgWithPR,
			orig:       idname.NewProvider("oid", "oname"),
			cache:      map[string]string{id: name},
			expectErr:  assert.NoError,
			expectID:   id,
			expectName: name,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			svc, result, err := GetRestoreResource(
				ctx,
				api.Client{},
				test.cfg,
				idname.NewCache(test.cache),
				test.orig)
			test.expectErr(t, err, clues.ToCore(err))
			require.NotNil(t, result)
			assert.Equal(t, path.ExchangeService, svc)
			assert.Equal(t, test.expectID, result.ID())
			assert.Equal(t, test.expectName, result.Name())
		})
	}
}
