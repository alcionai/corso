package streamstore

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
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
			[][]string{tester.AWSStorageCredEnvs}),
	})
}

func (suite *StreamStoreIntgSuite) SetupSubTest() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	// need to initialize the repository before we can test connecting to it.
	st := tester.NewPrefixedS3Storage(t)

	k := kopia.NewConn(st)
	require.NoError(t, k.Initialize(ctx))

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
	table := []struct {
		name  string
		deets func() *details.Details
		errs  func() *fault.Errors
	}{
		{
			name:  "none",
			deets: func() *details.Details { return nil },
			errs:  func() *fault.Errors { return nil },
		},
		{
			name: "details",
			deets: func() *details.Details {
				deetsBuilder := &details.Builder{}
				deetsBuilder.Add(
					"rr", "sr", "pr", "lr",
					true,
					details.ItemInfo{
						Exchange: &details.ExchangeInfo{Subject: "hello world"},
					})

				return deetsBuilder.Details()
			},
			errs: func() *fault.Errors { return nil },
		},
		{
			name:  "errors",
			deets: func() *details.Details { return nil },
			errs: func() *fault.Errors {
				bus := fault.New(false)
				bus.Fail(clues.New("foo"))
				bus.AddRecoverable(clues.New("bar"))
				bus.AddRecoverable(fault.FileErr(clues.New("file"), "file-id", "file-name", map[string]any{"foo": "bar"}))
				bus.AddSkip(fault.FileSkip(fault.SkipMalware, "file-id", "file-name", map[string]any{"foo": "bar"}))

				fe := bus.Errors()
				return &fe
			},
		},
		{
			name: "details and errors",
			deets: func() *details.Details {
				deetsBuilder := &details.Builder{}
				deetsBuilder.Add(
					"rr", "sr", "pr", "lr",
					true,
					details.ItemInfo{
						Exchange: &details.ExchangeInfo{Subject: "hello world"},
					})

				return deetsBuilder.Details()
			},
			errs: func() *fault.Errors {
				bus := fault.New(false)
				bus.Fail(clues.New("foo"))
				bus.AddRecoverable(clues.New("bar"))
				bus.AddRecoverable(fault.FileErr(clues.New("file"), "file-id", "file-name", map[string]any{"foo": "bar"}))
				bus.AddSkip(fault.FileSkip(fault.SkipMalware, "file-id", "file-name", map[string]any{"foo": "bar"}))

				fe := bus.Errors()
				return &fe
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			ctx, flush := tester.NewContext()
			defer flush()

			var (
				t   = suite.T()
				ss  = suite.ss
				err error
			)

			deets := test.deets()
			if deets != nil {
				err = ss.Collect(ctx, DetailsCollector(deets))
				require.NoError(t, err)
			}

			errs := test.errs()
			if errs != nil {
				err = ss.Collect(ctx, FaultErrorsCollector(errs))
				require.NoError(t, err)
			}

			snapid, err := ss.Write(ctx, fault.New(true))
			require.NoError(t, err)
			require.NotEmpty(t, snapid)

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
