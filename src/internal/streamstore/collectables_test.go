package streamstore

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
)

type StreamStoreIntgSuite struct {
	tester.Suite
	kcloser  func()
	kwcloser func()
	ss       Streamer
}

func TestStreamStoreIntgSuite(t *testing.T) {
	suite.Run(t, &StreamStoreIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *StreamStoreIntgSuite) SetupSubTest() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	// need to initialize the repository before we can test connecting to it.
	st := storeTD.NewPrefixedS3Storage(t)

	k := kopia.NewConn(st)
	err := k.Initialize(ctx, repository.Options{}, repository.Retention{})
	require.NoError(t, err, clues.ToCore(err))

	suite.kcloser = func() { k.Close(ctx) }

	kw, err := kopia.NewWrapper(k)
	require.NoError(t, err)

	suite.kwcloser = func() { kw.Close(ctx) }

	suite.ss = NewStreamer(kw, "tenant", path.ExchangeService)
}

func (suite *StreamStoreIntgSuite) TearDownSubTest() {
	if suite.kcloser != nil {
		defer suite.kcloser()
	}

	if suite.kwcloser != nil {
		defer suite.kwcloser()
	}
}

func (suite *StreamStoreIntgSuite) TestStreamer() {
	deetsPath, err := path.FromDataLayerPath("tenant-id/exchange/user-id/email/Inbox/folder1/foo", true)
	require.NoError(suite.T(), err, clues.ToCore(err))

	locPath := path.Builder{}.Append(deetsPath.Folders()...)

	table := []struct {
		name      string
		deets     func(*testing.T) *details.Details
		errs      func(context.Context) *fault.Errors
		hasSnapID assert.ValueAssertionFunc
	}{
		{
			name:      "none",
			deets:     func(*testing.T) *details.Details { return nil },
			errs:      func(context.Context) *fault.Errors { return nil },
			hasSnapID: assert.Empty,
		},
		{
			name: "details",
			deets: func(t *testing.T) *details.Details {
				deetsBuilder := &details.Builder{}
				require.NoError(t, deetsBuilder.Add(
					deetsPath,
					locPath,
					true,
					details.ItemInfo{
						Exchange: &details.ExchangeInfo{
							ItemType: details.ExchangeMail,
							Subject:  "hello world",
						},
					}))
				return deetsBuilder.Details()
			},
			errs:      func(context.Context) *fault.Errors { return nil },
			hasSnapID: assert.NotEmpty,
		},
		{
			name:  "errors",
			deets: func(*testing.T) *details.Details { return nil },
			errs: func(ctx context.Context) *fault.Errors {
				bus := fault.New(false)
				bus.Fail(clues.New("foo"))
				bus.AddRecoverable(ctx, clues.New("bar"))
				bus.AddRecoverable(
					ctx,
					fault.FileErr(clues.New("file"), "ns", "file-id", "file-name", map[string]any{"foo": "bar"}))
				bus.AddSkip(ctx, fault.FileSkip(fault.SkipMalware, "ns", "file-id", "file-name", map[string]any{"foo": "bar"}))

				fe := bus.Errors()
				return fe
			},
			hasSnapID: assert.NotEmpty,
		},
		{
			name: "details and errors",
			deets: func(t *testing.T) *details.Details {
				deetsBuilder := &details.Builder{}
				require.NoError(t, deetsBuilder.Add(
					deetsPath,
					locPath,
					true,
					details.ItemInfo{
						Exchange: &details.ExchangeInfo{
							ItemType: details.ExchangeMail,
							Subject:  "hello world",
						},
					}))

				return deetsBuilder.Details()
			},
			errs: func(ctx context.Context) *fault.Errors {
				bus := fault.New(false)
				bus.Fail(clues.New("foo"))
				bus.AddRecoverable(ctx, clues.New("bar"))
				bus.AddRecoverable(
					ctx,
					fault.FileErr(clues.New("file"), "ns", "file-id", "file-name", map[string]any{"foo": "bar"}))
				bus.AddSkip(ctx, fault.FileSkip(fault.SkipMalware, "ns", "file-id", "file-name", map[string]any{"foo": "bar"}))

				fe := bus.Errors()
				return fe
			},
			hasSnapID: assert.NotEmpty,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			var (
				ss  = suite.ss
				err error
			)

			deets := test.deets(t)
			if deets != nil {
				err = ss.Collect(ctx, DetailsCollector(deets))
				require.NoError(t, err)
			}

			errs := test.errs(ctx)
			if errs != nil {
				err = ss.Collect(ctx, FaultErrorsCollector(errs))
				require.NoError(t, err)
			}

			snapid, err := ss.Write(ctx, fault.New(true))
			require.NoError(t, err)
			test.hasSnapID(t, snapid)

			if len(snapid) == 0 {
				return
			}

			var readDeets details.Details
			if deets != nil {
				err = ss.Read(
					ctx,
					snapid,
					DetailsReader(details.UnmarshalTo(&readDeets)),
					fault.New(true))
				require.NoError(t, err)
				require.NotEmpty(t, readDeets)

				assert.Equal(t, len(deets.Entries), len(readDeets.Entries))
				assert.Equal(t, deets.Entries[0].ParentRef, readDeets.Entries[0].ParentRef)
				assert.Equal(t, deets.Entries[0].ShortRef, readDeets.Entries[0].ShortRef)
				assert.Equal(t, deets.Entries[0].RepoRef, readDeets.Entries[0].RepoRef)
				assert.Equal(t, deets.Entries[0].LocationRef, readDeets.Entries[0].LocationRef)
				assert.Equal(t, deets.Entries[0].ItemRef, readDeets.Entries[0].ItemRef)
				assert.Equal(t, deets.Entries[0].Updated, readDeets.Entries[0].Updated)
				assert.NotNil(t, readDeets.Entries[0].Exchange)
				assert.Equal(t, *deets.Entries[0].Exchange, *readDeets.Entries[0].Exchange)
			} else {
				err := ss.Read(
					ctx,
					snapid,
					DetailsReader(details.UnmarshalTo(&readDeets)),
					fault.New(true))
				assert.ErrorIs(t, err, data.ErrNotFound)
				assert.Empty(t, readDeets)
			}

			var readErrs fault.Errors
			if errs != nil {
				err = ss.Read(
					ctx,
					snapid,
					FaultErrorsReader(fault.UnmarshalErrorsTo(&readErrs)),
					fault.New(true))
				require.NoError(t, err)
				require.NotEmpty(t, readErrs)

				assert.ElementsMatch(t, errs.Skipped, readErrs.Skipped)
				assert.ElementsMatch(t, errs.Recovered, readErrs.Recovered)
			} else {
				err := ss.Read(
					ctx,
					snapid,
					FaultErrorsReader(fault.UnmarshalErrorsTo(&readErrs)),
					fault.New(true))
				assert.ErrorIs(t, err, data.ErrNotFound)
				assert.Empty(t, readErrs)
			}
		})
	}
}
