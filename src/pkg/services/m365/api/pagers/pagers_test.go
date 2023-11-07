package pagers

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/tester"
)

// ---------------------------------------------------------------------------
// mock impls & stubs
// ---------------------------------------------------------------------------

// next and delta links

type nextLink struct {
	nextLink *string
}

func (l nextLink) GetOdataNextLink() *string {
	return l.nextLink
}

type deltaNextLink struct {
	nextLink
	deltaLink *string
}

func (l deltaNextLink) GetOdataDeltaLink() *string {
	return l.deltaLink
}

// mock item

var _ getIDModAndAddtler = &testItem{}

func removedItem(id string) testItem {
	return testItem{
		id: id,
		additionalData: map[string]any{
			graph.AddtlDataRemoved: struct{}{},
		},
	}
}

func addedItem(id string, modTime time.Time) testItem {
	return testItem{
		id:      id,
		modTime: modTime,
	}
}

type testItem struct {
	id             string
	modTime        time.Time
	additionalData map[string]any
}

//nolint:revive
func (ti testItem) GetId() *string {
	return &ti.id
}

func (ti testItem) GetLastModifiedDateTime() *time.Time {
	return &ti.modTime
}

func (ti testItem) GetAdditionalData() map[string]any {
	return ti.additionalData
}

// mock page

type testPage struct {
	values    []testItem
	nextLink  string
	deltaLink string
}

func (p testPage) GetOdataNextLink() *string {
	return ptr.To(p.nextLink)
}

func (p testPage) GetOdataDeltaLink() *string {
	// delta is not tested here
	return ptr.To(p.deltaLink)
}

func (p testPage) GetValue() []testItem {
	return p.values
}

// mock item pagers

type pageResult struct {
	items      []testItem
	err        error
	errCode    string
	needsReset bool
}

var (
	_ NonDeltaHandler[testItem] = &testIDsNonDeltaMultiPager{}
	_ DeltaHandler[testItem]    = &testIDsDeltaMultiPager{}
)

type testIDsNonDeltaMultiPager struct {
	t             *testing.T
	pageIdx       int
	pages         []pageResult
	validModTimes bool
	needsReset    bool
}

func (p *testIDsNonDeltaMultiPager) GetPage(
	ctx context.Context,
) (NextLinkValuer[testItem], error) {
	if p.pageIdx >= len(p.pages) {
		return testPage{}, clues.New("result out of expected range")
	}

	res := p.pages[p.pageIdx]
	p.needsReset = res.needsReset
	p.pageIdx++

	if res.err != nil {
		return testPage{}, res.err
	}

	if len(res.errCode) > 0 {
		ierr := odataerrors.NewMainError()
		ierr.SetCode(ptr.To(res.errCode))

		err := odataerrors.NewODataError()
		err.SetErrorEscaped(ierr)

		return testPage{}, err
	}

	var (
		nextLink  string
		deltaLink string
	)

	if p.pageIdx < len(p.pages) {
		// Value doesn't matter as long as it's not empty.
		nextLink = "next"
	} else {
		deltaLink = "delta"
	}

	return testPage{
		values:    res.items,
		nextLink:  nextLink,
		deltaLink: deltaLink,
	}, nil
}

func (p *testIDsNonDeltaMultiPager) SetNextLink(string) {}

func (p *testIDsNonDeltaMultiPager) Reset(context.Context) {
	if !p.needsReset {
		require.Fail(p.t, "reset should not be called")
	}

	p.needsReset = false
}

func (p testIDsNonDeltaMultiPager) ValidModTimes() bool {
	return p.validModTimes
}

func newDeltaPager(p *testIDsNonDeltaMultiPager) *testIDsDeltaMultiPager {
	return &testIDsDeltaMultiPager{
		testIDsNonDeltaMultiPager: p,
	}
}

type testIDsDeltaMultiPager struct {
	*testIDsNonDeltaMultiPager
}

func (p *testIDsDeltaMultiPager) GetPage(
	ctx context.Context,
) (DeltaLinkValuer[testItem], error) {
	linker, err := p.testIDsNonDeltaMultiPager.GetPage(ctx)
	deltaLinker := linker.(DeltaLinkValuer[testItem])

	return deltaLinker, err
}

// ---------------------------------------------------------------------------
// Tests
// ---------------------------------------------------------------------------

type PagerUnitSuite struct {
	tester.Suite
}

func TestPagerUnitSuite(t *testing.T) {
	suite.Run(t, &PagerUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *PagerUnitSuite) TestBatchEnumerateItems() {
	item1 := addedItem("foo", time.Now())
	item2 := addedItem("bar", time.Now())

	tests := []struct {
		name      string
		getPager  func(*testing.T) NonDeltaHandler[testItem]
		expect    []testItem
		expectErr require.ErrorAssertionFunc
	}{
		{
			name: "OnePage",
			getPager: func(t *testing.T) NonDeltaHandler[testItem] {
				return &testIDsNonDeltaMultiPager{
					t: t,
					pages: []pageResult{
						{
							items: []testItem{
								item1,
								item2,
							},
						},
					},
				}
			},
			expect: []testItem{
				item1,
				item2,
			},
			expectErr: require.NoError,
		},
		{
			name: "TwoPages",
			getPager: func(t *testing.T) NonDeltaHandler[testItem] {
				return &testIDsNonDeltaMultiPager{
					t: t,
					pages: []pageResult{
						{
							items: []testItem{
								item1,
							},
						},
						{
							items: []testItem{
								item2,
							},
						},
					},
				}
			},
			expect: []testItem{
				item1,
				item2,
			},
			expectErr: require.NoError,
		},
		{
			name: "TwoPages ErrorAfterFirst",
			getPager: func(t *testing.T) NonDeltaHandler[testItem] {
				return &testIDsNonDeltaMultiPager{
					t: t,
					pages: []pageResult{
						{
							items: []testItem{
								item1,
							},
						},
						{
							err: assert.AnError,
						},
						{
							items: []testItem{
								item2,
							},
						},
					},
				}
			},
			expect: []testItem{
				item1,
			},
			expectErr: require.Error,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			result, err := BatchEnumerateItems(ctx, test.getPager(t))
			test.expectErr(t, err, clues.ToCore(err))

			require.ElementsMatch(t, test.expect, result)
		})
	}
}

func assertSliceEmptyOr[S ~[]E, E any](
	t *testing.T,
	expect S,
	got S,
	assertionFunc assert.ComparisonAssertionFunc,
	msgAndArgs ...any,
) {
	if len(expect) == 0 {
		assert.Empty(t, got, msgAndArgs)
		return
	}

	assertionFunc(t, expect, got, msgAndArgs)
}

func assertMapEmptyOr[M ~map[K]V, K comparable, V any](
	t *testing.T,
	expect M,
	got M,
	assertionFunc assert.ComparisonAssertionFunc,
	msgAndArgs ...any,
) {
	if len(expect) == 0 {
		assert.Empty(t, got, msgAndArgs)
		return
	}

	assertionFunc(t, expect, got, msgAndArgs)
}

func assertAddedAndRemoved(
	t *testing.T,
	validModTimes bool,
	wantAdded []testItem,
	gotAdded map[string]time.Time,
	wantRemoved []testItem,
	gotRemoved []string,
) {
	epoch, err := time.Parse(time.DateOnly, "1970-01-01")
	require.NoError(t, err, clues.ToCore(err))

	expectAdded := map[string]time.Time{}
	for _, item := range wantAdded {
		expectAdded[item.id] = item.modTime
	}

	if validModTimes {
		assertMapEmptyOr(
			t,
			expectAdded,
			gotAdded,
			assert.Equal,
			"added item IDs and mod times")
	} else {
		assertSliceEmptyOr(
			t,
			maps.Keys(expectAdded),
			maps.Keys(gotAdded),
			assert.ElementsMatch,
			"added item IDs")

		for _, modtime := range gotAdded {
			assert.True(t, modtime.After(epoch), "mod time after epoch")
			assert.False(t, modtime.IsZero(), "non-zero mod time")
		}
	}

	expectRemoved := []string{}
	for _, item := range wantRemoved {
		expectRemoved = append(expectRemoved, item.id)
	}

	assertSliceEmptyOr(
		t,
		expectRemoved,
		gotRemoved,
		assert.ElementsMatch,
		"removed item IDs")
}

type modTimeTest struct {
	name          string
	validModTimes bool
}

var (
	addedItem1 = addedItem("a_uno", time.Now())
	addedItem2 = addedItem("a_dos", time.Now())

	removedItem1 = removedItem("r_uno")
	removedItem2 = removedItem("r_dos")

	modTimeTests = []modTimeTest{
		{
			name:          "ValidModTimes",
			validModTimes: true,
		},
		{
			name: "InvalidModTimes",
		},
	}

	nilPager = func(*testing.T, bool) NonDeltaHandler[testItem] {
		return nil
	}
)

func (suite *PagerUnitSuite) TestGetAddedAndRemovedItemIDs() {
	pagerTypeTests := []struct {
		name        string
		prevDelta   string
		canUseDelta bool
		pagersFunc  func(
			p *testIDsNonDeltaMultiPager,
		) (NonDeltaHandler[testItem], DeltaHandler[testItem])
		expectDeltaReset bool
		expectDeltaLink  assert.ValueAssertionFunc
	}{
		{
			name:        "NoPrevDelta",
			canUseDelta: true,
			pagersFunc: func(
				p *testIDsNonDeltaMultiPager,
			) (NonDeltaHandler[testItem], DeltaHandler[testItem]) {
				return nil, newDeltaPager(p)
			},
			expectDeltaReset: true,
			expectDeltaLink:  assert.NotEmpty,
		},
		{
			name:        "PrevDelta",
			prevDelta:   "a",
			canUseDelta: true,
			pagersFunc: func(
				p *testIDsNonDeltaMultiPager,
			) (NonDeltaHandler[testItem], DeltaHandler[testItem]) {
				return nil, newDeltaPager(p)
			},
			expectDeltaLink: assert.NotEmpty,
		},
		{
			name:      "DeltaNotAllowed",
			prevDelta: "a",
			pagersFunc: func(
				p *testIDsNonDeltaMultiPager,
			) (NonDeltaHandler[testItem], DeltaHandler[testItem]) {
				return p, nil
			},
			expectDeltaReset: true,
			expectDeltaLink:  assert.Empty,
		},
	}

	type expected struct {
		errCheck assert.ErrorAssertionFunc
		added    []testItem
		removed  []testItem
	}

	table := []struct {
		name        string
		pagerGetter func(
			t *testing.T,
			validModTimes bool,
		) *testIDsNonDeltaMultiPager
		filter func(a testItem) bool
		expect expected
	}{
		{
			name: "OnePage",
			pagerGetter: func(t *testing.T, validModTime bool) *testIDsNonDeltaMultiPager {
				return &testIDsNonDeltaMultiPager{
					t: t,
					pages: []pageResult{
						{
							items: []testItem{
								addedItem1,
								addedItem2,
								removedItem1,
								removedItem2,
							},
						},
					},
					validModTimes: validModTime,
				}
			},
			expect: expected{
				errCheck: assert.NoError,
				added: []testItem{
					addedItem1,
					addedItem2,
				},
				removed: []testItem{removedItem1, removedItem2},
			},
		},
		{
			name: "TwoPages",
			pagerGetter: func(t *testing.T, validModTime bool) *testIDsNonDeltaMultiPager {
				return &testIDsNonDeltaMultiPager{
					t: t,
					pages: []pageResult{
						{
							items: []testItem{
								addedItem1,
								removedItem1,
							},
						},
						{
							items: []testItem{
								addedItem2,
								removedItem2,
							},
						},
					},
					validModTimes: validModTime,
				}
			},
			expect: expected{
				errCheck: assert.NoError,
				added: []testItem{
					addedItem1,
					addedItem2,
				},
				removed: []testItem{removedItem1, removedItem2},
			},
		},
		{
			name: "OnePage FilterFailsAll",
			pagerGetter: func(t *testing.T, validModTimes bool) *testIDsNonDeltaMultiPager {
				return &testIDsNonDeltaMultiPager{
					t: t,
					pages: []pageResult{
						{
							items: []testItem{
								addedItem1,
								addedItem2,
								removedItem1,
								removedItem2,
							},
						},
					},
					validModTimes: validModTimes,
				}
			},
			filter: func(testItem) bool { return false },
			expect: expected{
				errCheck: assert.NoError,
			},
		},
		{
			name: "TwoPages FilterFailsAll",
			pagerGetter: func(t *testing.T, validModTimes bool) *testIDsNonDeltaMultiPager {
				return &testIDsNonDeltaMultiPager{
					t: t,
					pages: []pageResult{
						{
							items: []testItem{
								addedItem1,
								removedItem1,
							},
						},
						{
							items: []testItem{
								addedItem2,
								removedItem2,
							},
						},
					},
					validModTimes: validModTimes,
				}
			},
			filter: func(testItem) bool { return false },
			expect: expected{
				errCheck: assert.NoError,
			},
		},
		{
			name: "Error",
			pagerGetter: func(t *testing.T, validModTimes bool) *testIDsNonDeltaMultiPager {
				return &testIDsNonDeltaMultiPager{
					t: t,
					pages: []pageResult{
						{
							err: assert.AnError,
						},
					},
					validModTimes: validModTimes,
				}
			},
			expect: expected{
				errCheck: assert.Error,
			},
		},
	}

	for _, modTimeTest := range modTimeTests {
		suite.Run(modTimeTest.name, func() {
			for _, pagerTypeTest := range pagerTypeTests {
				suite.Run(pagerTypeTest.name, func() {
					for _, test := range table {
						suite.Run(test.name, func() {
							t := suite.T()

							ctx, flush := tester.NewContext(t)
							defer flush()

							filters := []func(testItem) bool{}
							if test.filter != nil {
								filters = append(filters, test.filter)
							}

							basePager := test.pagerGetter(t, modTimeTest.validModTimes)
							getter, deltaGetter := pagerTypeTest.pagersFunc(basePager)

							added, validModTimes, removed, deltaUpdate, err := GetAddedAndRemovedItemIDs[testItem](
								ctx,
								getter,
								deltaGetter,
								pagerTypeTest.prevDelta,
								pagerTypeTest.canUseDelta,
								AddedAndRemovedByAddtlData[testItem],
								filters...)
							test.expect.errCheck(t, err, "getting added and removed item IDs: %+v", clues.ToCore(err))

							if err != nil {
								return
							}

							assert.Equal(t, modTimeTest.validModTimes, validModTimes, "valid mod times")
							assert.Equal(t, pagerTypeTest.expectDeltaReset, deltaUpdate.Reset, "delta update")
							pagerTypeTest.expectDeltaLink(t, deltaUpdate.URL, "delta link")

							assertAddedAndRemoved(
								t,
								modTimeTest.validModTimes,
								test.expect.added,
								added,
								test.expect.removed,
								removed)
						})
					}
				})
			}
		})
	}
}

// TestGetAddedAndRemovedItemIDs_FallbackPagers tests that when pagers get reset
// or need to fallback to other pager types things work as expected. This only
// tests for basic cases where we enumerate everything with the fallback pager.
// These are here mostly to ensure we clear the results from the invalid pager
// properly. If we can ensure that then other tests will ensure the fallback
// pager properly handles all the other things like item filtering, item limits,
// cancellation, etc.
func (suite *PagerUnitSuite) TestGetAddedAndRemovedItemIDs_FallbackPagers() {
	type expected struct {
		errCheck  assert.ErrorAssertionFunc
		added     []testItem
		removed   []testItem
		deltaLink assert.ValueAssertionFunc
	}

	tests := []struct {
		name        string
		pagerGetter func(
			t *testing.T,
			validModTimes bool,
		) NonDeltaHandler[testItem]
		deltaPagerGetter func(
			t *testing.T,
			validModTimes bool,
		) DeltaHandler[testItem]
		expect expected
	}{
		{
			name:        "TwoValidPages DeltaReset",
			pagerGetter: nilPager,
			deltaPagerGetter: func(
				t *testing.T,
				validModTimes bool,
			) DeltaHandler[testItem] {
				return newDeltaPager(
					&testIDsNonDeltaMultiPager{
						t: t,
						pages: []pageResult{
							{
								items: []testItem{
									addedItem1,
									removedItem1,
								},
							},
							{
								errCode:    "SyncStateNotFound",
								needsReset: true,
							},
							{
								items: []testItem{
									removedItem2,
									addedItem2,
								},
							},
						},
						validModTimes: validModTimes,
					})
			},
			expect: expected{
				errCheck: assert.NoError,
				added: []testItem{
					addedItem2,
				},
				removed: []testItem{
					removedItem2,
				},
				deltaLink: assert.NotEmpty,
			},
		},
		{
			name:        "TwoPages DeltaResetAtEnd",
			pagerGetter: nilPager,
			deltaPagerGetter: func(
				t *testing.T,
				validModTimes bool,
			) DeltaHandler[testItem] {
				return newDeltaPager(
					&testIDsNonDeltaMultiPager{
						t: t,
						pages: []pageResult{
							{
								items: []testItem{
									addedItem("uno", time.Now()),
									removedItem("tres"),
								},
							},
							{
								items: []testItem{
									removedItem("quatro"),
									addedItem("dos", time.Now()),
								},
							},
							{
								errCode:    "SyncStateNotFound",
								needsReset: true,
							},
							// Return an empty page to show no new results after reset.
							{},
						},
						validModTimes: validModTimes,
					})
			},
			expect: expected{
				errCheck:  assert.NoError,
				deltaLink: assert.NotEmpty,
			},
		},
		{
			name: "TwoValidPages DeltaNotSupported",
			pagerGetter: func(
				t *testing.T,
				validModTimes bool,
			) NonDeltaHandler[testItem] {
				return &testIDsNonDeltaMultiPager{
					t: t,
					pages: []pageResult{
						{
							items: []testItem{
								addedItem1,
								removedItem1,
							},
						},
						{
							items: []testItem{
								removedItem2,
								addedItem2,
							},
						},
					},
					validModTimes: validModTimes,
				}
			},
			deltaPagerGetter: func(
				t *testing.T,
				validModTimes bool,
			) DeltaHandler[testItem] {
				return newDeltaPager(
					&testIDsNonDeltaMultiPager{
						t: t,
						pages: []pageResult{
							{
								err:        graph.ErrDeltaNotSupported,
								needsReset: true,
							},
						},
						validModTimes: validModTimes,
					})
			},
			expect: expected{
				errCheck: assert.NoError,
				added: []testItem{
					addedItem1,
					addedItem2,
				},
				removed: []testItem{
					removedItem1,
					removedItem2,
				},
				deltaLink: assert.Empty,
			},
		},
		{
			name: "TwoPages DeltaNotSupportedAtEnd",
			pagerGetter: func(
				t *testing.T,
				validModTimes bool,
			) NonDeltaHandler[testItem] {
				return &testIDsNonDeltaMultiPager{
					t: t,
					pages: []pageResult{
						// Return an empty page.
						{},
					},
					validModTimes: validModTimes,
				}
			},
			deltaPagerGetter: func(
				t *testing.T,
				validModTimes bool,
			) DeltaHandler[testItem] {
				return newDeltaPager(
					&testIDsNonDeltaMultiPager{
						t: t,
						pages: []pageResult{
							{
								items: []testItem{
									addedItem1,
									removedItem1,
								},
							},
							{
								items: []testItem{
									removedItem2,
									addedItem2,
								},
							},
							{
								err:        graph.ErrDeltaNotSupported,
								needsReset: true,
							},
						},
						validModTimes: validModTimes,
					})
			},
			expect: expected{
				errCheck:  assert.NoError,
				deltaLink: assert.Empty,
			},
		},
	}

	for _, modTimeTest := range modTimeTests {
		suite.Run(modTimeTest.name, func() {
			for _, test := range tests {
				suite.Run(test.name, func() {
					t := suite.T()

					ctx, flush := tester.NewContext(t)
					defer flush()

					added, validModTimes, removed, deltaUpdate, err := GetAddedAndRemovedItemIDs[testItem](
						ctx,
						test.pagerGetter(t, modTimeTest.validModTimes),
						test.deltaPagerGetter(t, modTimeTest.validModTimes),
						"a",
						true,
						AddedAndRemovedByAddtlData[testItem])
					require.NoError(
						t,
						err,
						"getting added and removed item IDs: %+v",
						clues.ToCore(err))

					assert.Equal(t, modTimeTest.validModTimes, validModTimes, "valid mod times")
					assert.True(t, deltaUpdate.Reset, "delta update")
					test.expect.deltaLink(t, deltaUpdate.URL, "delta link")

					assertAddedAndRemoved(
						t,
						modTimeTest.validModTimes,
						test.expect.added,
						added,
						test.expect.removed,
						removed)
				})
			}
		})
	}
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

func (suite *PagerUnitSuite) TestNextAndDeltaLink() {
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

				l := deltaNextLink{
					nextLink:  nextLink{nextLink: next.inputLink},
					deltaLink: delta.inputLink,
				}
				gotNext, gotDelta := NextAndDeltaLink(l)

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
func (suite *PagerUnitSuite) TestIsLinkValid() {
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
			got := IsNextLinkValid(test.inputString)
			test.isValid(suite.T(), got)
		})
	}
}
