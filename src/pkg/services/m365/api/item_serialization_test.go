package api

import (
	"context"
	"fmt"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
)

type ItemSerializationUnitSuite struct {
	tester.Suite
}

func TestItemSerializationUnitSuite(t *testing.T) {
	suite.Run(t, &ItemSerializationUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ItemSerializationUnitSuite) TestConcurrentItemSerialization() {
	var (
		user      = "a-user"
		instances = 100
	)

	table := []struct {
		name                   string
		serializer             func(t *testing.T, ctx context.Context, idx int) []byte
		deserializeAndGetField func(t *testing.T, bs []byte) string
	}{
		{
			name: "Exchange Mail",
			serializer: func(t *testing.T, ctx context.Context, idx int) []byte {
				subject := fmt.Sprintf("%d", idx)

				item := models.NewMessage()
				item.SetSubject(&subject)

				bs, err := Mail{}.Serialize(ctx, item, user, subject)
				require.NoError(t, err, clues.ToCore(err))

				return bs
			},
			deserializeAndGetField: func(t *testing.T, bs []byte) string {
				item, err := BytesToMessageable(bs)
				require.NoError(
					t,
					err,
					"deserializing message of %q: %v",
					string(bs),
					clues.ToCore(err))

				return ptr.Val(item.GetSubject())
			},
		},
		{
			name: "Exchange Event",
			serializer: func(t *testing.T, ctx context.Context, idx int) []byte {
				subject := fmt.Sprintf("%d", idx)

				item := models.NewEvent()
				item.SetSubject(&subject)

				bs, err := Events{}.Serialize(ctx, item, user, subject)
				require.NoError(t, err, clues.ToCore(err))

				return bs
			},
			deserializeAndGetField: func(t *testing.T, bs []byte) string {
				item, err := BytesToEventable(bs)
				require.NoError(
					t,
					err,
					"deserializing event of %q: %v",
					string(bs),
					clues.ToCore(err))

				return ptr.Val(item.GetSubject())
			},
		},
		{
			name: "Exchange Contact",
			serializer: func(t *testing.T, ctx context.Context, idx int) []byte {
				name := fmt.Sprintf("%d", idx)

				item := models.NewContact()
				item.SetGivenName(&name)

				bs, err := Contacts{}.Serialize(ctx, item, user, name)
				require.NoError(t, err, clues.ToCore(err))

				return bs
			},
			deserializeAndGetField: func(t *testing.T, bs []byte) string {
				item, err := BytesToContactable(bs)
				require.NoError(
					t,
					err,
					"deserializing contact of %q: %v",
					string(bs),
					clues.ToCore(err))

				return ptr.Val(item.GetGivenName())
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flusher := tester.NewContext(t)
			defer flusher()

			output := make([][]byte, instances)

			for i := 0; i < instances; i++ {
				output[i] = test.serializer(t, ctx, i)
			}

			for i := 0; i < instances; i++ {
				got := test.deserializeAndGetField(t, output[i])
				// I'm lazy and don't want to deal with the error from atoi functions.
				assert.Equal(t, fmt.Sprintf("%d", i), got, "item output")
			}
		})
	}
}
