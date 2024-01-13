package testdata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/tester"
)

type LocSetUnitSuite struct {
	tester.Suite
}

func TestLocSetUnitSuite(t *testing.T) {
	suite.Run(t, &LocSetUnitSuite{Suite: tester.NewUnitSuite(t)})
}

const (
	l1  = "lr_1"
	l2  = "lr_2"
	l13 = "lr_1/lr_3"
	l14 = "lr_1/lr_4"
	i1  = "ir_1"
	i2  = "ir_2"
	i3  = "ir_3"
	i4  = "ir_4"
)

func (suite *LocSetUnitSuite) TestAdd() {
	t := suite.T()

	ls := newLocSet()

	ls.AddItem(l1, i1)
	ls.AddLocation(l2)

	assert.ElementsMatch(t, []string{l1, l2}, maps.Keys(ls.Locations))
	assert.ElementsMatch(t, []string{i1}, maps.Keys(ls.Locations[l1]))
	assert.Empty(t, maps.Keys(ls.Locations[l2]))
	assert.Empty(t, maps.Keys(ls.Locations[l13]))
}

func (suite *LocSetUnitSuite) TestRemove() {
	t := suite.T()

	ls := newLocSet()

	ls.AddItem(l1, i1)
	ls.AddItem(l1, i2)
	ls.AddLocation(l13)
	ls.AddItem(l14, i3)
	ls.AddItem(l14, i4)

	assert.ElementsMatch(t, []string{l1, l13, l14}, maps.Keys(ls.Locations))
	assert.ElementsMatch(t, []string{i1, i2}, maps.Keys(ls.Locations[l1]))
	assert.Empty(t, maps.Keys(ls.Locations[l13]))
	assert.ElementsMatch(t, []string{i3, i4}, maps.Keys(ls.Locations[l14]))

	// nop removal
	ls.RemoveItem(l2, i1)
	assert.ElementsMatch(t, []string{i1, i2}, maps.Keys(ls.Locations[l1]))

	// item removal
	ls.RemoveItem(l1, i2)
	assert.ElementsMatch(t, []string{i1}, maps.Keys(ls.Locations[l1]))

	// nop location removal
	ls.RemoveLocation(l2)
	assert.ElementsMatch(t, []string{l1, l13, l14}, maps.Keys(ls.Locations))

	// non-cascading location removal
	ls.RemoveLocation(l13)
	assert.ElementsMatch(t, []string{l1, l14}, maps.Keys(ls.Locations))
	assert.ElementsMatch(t, []string{i1}, maps.Keys(ls.Locations[l1]))
	assert.ElementsMatch(t, []string{i3, i4}, maps.Keys(ls.Locations[l14]))

	// cascading location removal
	ls.RemoveLocation(l1)
	assert.Empty(t, maps.Keys(ls.Locations))
	assert.Empty(t, maps.Keys(ls.Locations[l1]))
	assert.Empty(t, maps.Keys(ls.Locations[l13]))
	assert.Empty(t, maps.Keys(ls.Locations[l14]))
}

func (suite *LocSetUnitSuite) TestSubset() {
	ls := newLocSet()

	ls.AddItem(l1, i1)
	ls.AddItem(l1, i2)
	ls.AddLocation(l13)
	ls.AddItem(l14, i3)
	ls.AddItem(l14, i4)

	table := []struct {
		name   string
		locPfx string
		expect func(*testing.T, *locSet)
	}{
		{
			name:   "nop",
			locPfx: l2,
			expect: func(t *testing.T, ss *locSet) {
				assert.Empty(t, maps.Keys(ss.Locations))
			},
		},
		{
			name:   "no items",
			locPfx: l13,
			expect: func(t *testing.T, ss *locSet) {
				assert.ElementsMatch(t, []string{l13}, maps.Keys(ss.Locations))
				assert.Empty(t, maps.Keys(ss.Locations[l13]))
			},
		},
		{
			name:   "non-cascading",
			locPfx: l14,
			expect: func(t *testing.T, ss *locSet) {
				assert.ElementsMatch(t, []string{l14}, maps.Keys(ss.Locations))
				assert.ElementsMatch(t, []string{i3, i4}, maps.Keys(ss.Locations[l14]))
			},
		},
		{
			name:   "cascading",
			locPfx: l1,
			expect: func(t *testing.T, ss *locSet) {
				assert.ElementsMatch(t, []string{l1, l13, l14}, maps.Keys(ss.Locations))
				assert.ElementsMatch(t, []string{i1, i2}, maps.Keys(ss.Locations[l1]))
				assert.ElementsMatch(t, []string{i3, i4}, maps.Keys(ss.Locations[l14]))
				assert.Empty(t, maps.Keys(ss.Locations[l13]))
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			test.expect(t, ls.Subset(test.locPfx))
		})
	}
}

func (suite *LocSetUnitSuite) TestRename() {
	t := suite.T()

	makeSet := func() *locSet {
		ls := newLocSet()

		ls.AddItem(l1, i1)
		ls.AddItem(l1, i2)
		ls.AddLocation(l13)
		ls.AddItem(l14, i3)
		ls.AddItem(l14, i4)

		return ls
	}

	ts := makeSet()
	assert.ElementsMatch(t, []string{l1, l13, l14}, maps.Keys(ts.Locations))
	assert.ElementsMatch(t, []string{i1, i2}, maps.Keys(ts.Locations[l1]))
	assert.Empty(t, maps.Keys(ts.Locations[l13]))
	assert.ElementsMatch(t, []string{i3, i4}, maps.Keys(ts.Locations[l14]))

	table := []struct {
		name   string
		from   string
		to     string
		expect func(*testing.T, *locSet)
	}{
		{
			name: "nop",
			from: l2,
			to:   "foo",
			expect: func(t *testing.T, ls *locSet) {
				assert.ElementsMatch(t, []string{l1, l13, l14}, maps.Keys(ls.Locations))
				assert.Empty(t, maps.Keys(ls.Locations[l2]))
				assert.Empty(t, maps.Keys(ls.Locations["foo"]))
			},
		},
		{
			name: "no items",
			from: l13,
			to:   "foo",
			expect: func(t *testing.T, ls *locSet) {
				assert.ElementsMatch(t, []string{l1, "foo", l14}, maps.Keys(ls.Locations))
				assert.Empty(t, maps.Keys(ls.Locations[l13]))
				assert.Empty(t, maps.Keys(ls.Locations["foo"]))
			},
		},
		{
			name: "with items",
			from: l14,
			to:   "foo",
			expect: func(t *testing.T, ls *locSet) {
				assert.ElementsMatch(t, []string{l1, l13, "foo"}, maps.Keys(ls.Locations))
				assert.Empty(t, maps.Keys(ls.Locations[l14]))
				assert.ElementsMatch(t, []string{i3, i4}, maps.Keys(ls.Locations["foo"]))
			},
		},
		{
			name: "cascading locations",
			from: l1,
			to:   "foo",
			expect: func(t *testing.T, ls *locSet) {
				assert.ElementsMatch(t, []string{"foo", "foo/lr_3", "foo/lr_4"}, maps.Keys(ls.Locations))
				assert.Empty(t, maps.Keys(ls.Locations[l1]))
				assert.Empty(t, maps.Keys(ls.Locations[l14]))
				assert.Empty(t, maps.Keys(ls.Locations[l13]))
				assert.ElementsMatch(t, []string{i1, i2}, maps.Keys(ls.Locations["foo"]))
				assert.Empty(t, maps.Keys(ls.Locations["foo/lr_3"]))
				assert.ElementsMatch(t, []string{i3, i4}, maps.Keys(ls.Locations["foo/lr_4"]))
			},
		},
		{
			name: "to existing location",
			from: l14,
			to:   l1,
			expect: func(t *testing.T, ls *locSet) {
				assert.ElementsMatch(t, []string{l1, l13}, maps.Keys(ls.Locations))
				assert.Empty(t, maps.Keys(ls.Locations[l14]))
				assert.ElementsMatch(t, []string{i1, i2, i3, i4}, maps.Keys(ls.Locations[l1]))
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			ls := makeSet()

			ls.RenameLocation(test.from, test.to)
			test.expect(t, ls)
		})
	}
}

func (suite *LocSetUnitSuite) TestItem() {
	t := suite.T()
	b4 := "bar/lr_4"

	makeSet := func() *locSet {
		ls := newLocSet()

		ls.AddItem(l1, i1)
		ls.AddItem(l1, i2)
		ls.AddLocation(l13)
		ls.AddItem(l14, i3)
		ls.AddItem(l14, i4)
		ls.AddItem(b4, "fnord")

		return ls
	}

	ts := makeSet()
	assert.ElementsMatch(t, []string{l1, l13, l14, b4}, maps.Keys(ts.Locations))
	assert.ElementsMatch(t, []string{i1, i2}, maps.Keys(ts.Locations[l1]))
	assert.Empty(t, maps.Keys(ts.Locations[l13]))
	assert.ElementsMatch(t, []string{i3, i4}, maps.Keys(ts.Locations[l14]))
	assert.ElementsMatch(t, []string{"fnord"}, maps.Keys(ts.Locations[b4]))

	table := []struct {
		name   string
		item   string
		from   string
		to     string
		expect func(*testing.T, *locSet)
	}{
		{
			name: "nop item",
			item: "floob",
			from: l2,
			to:   l1,
			expect: func(t *testing.T, ls *locSet) {
				assert.ElementsMatch(t, []string{i1, i2, "floob"}, maps.Keys(ls.Locations[l1]))
				assert.Empty(t, maps.Keys(ls.Locations[l2]))
			},
		},
		{
			name: "nop origin",
			item: i1,
			from: "smarf",
			to:   l2,
			expect: func(t *testing.T, ls *locSet) {
				assert.ElementsMatch(t, []string{i1, i2}, maps.Keys(ls.Locations[l1]))
				assert.ElementsMatch(t, []string{i1}, maps.Keys(ls.Locations[l2]))
				assert.Empty(t, maps.Keys(ls.Locations["smarf"]))
			},
		},
		{
			name: "new location",
			item: i1,
			from: l1,
			to:   "fnords",
			expect: func(t *testing.T, ls *locSet) {
				assert.ElementsMatch(t, []string{i2}, maps.Keys(ls.Locations[l1]))
				assert.ElementsMatch(t, []string{i1}, maps.Keys(ls.Locations["fnords"]))
			},
		},
		{
			name: "existing location",
			item: i1,
			from: l1,
			to:   l2,
			expect: func(t *testing.T, ls *locSet) {
				assert.ElementsMatch(t, []string{i2}, maps.Keys(ls.Locations[l1]))
				assert.ElementsMatch(t, []string{i1}, maps.Keys(ls.Locations[l2]))
			},
		},
		{
			name: "same location",
			item: i1,
			from: l1,
			to:   l1,
			expect: func(t *testing.T, ls *locSet) {
				assert.ElementsMatch(t, []string{i1, i2}, maps.Keys(ls.Locations[l1]))
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			ls := makeSet()

			ls.MoveItem(test.from, test.to, test.item)
			test.expect(t, ls)
		})
	}
}

func (suite *LocSetUnitSuite) TestMoveLocation() {
	t := suite.T()
	b4 := "bar/lr_4"

	makeSet := func() *locSet {
		ls := newLocSet()

		ls.AddItem(l1, i1)
		ls.AddItem(l1, i2)
		ls.AddLocation(l13)
		ls.AddItem(l14, i3)
		ls.AddItem(l14, i4)
		ls.AddItem(b4, "fnord")

		return ls
	}

	ts := makeSet()
	assert.ElementsMatch(t, []string{l1, l13, l14, b4}, maps.Keys(ts.Locations))
	assert.ElementsMatch(t, []string{i1, i2}, maps.Keys(ts.Locations[l1]))
	assert.Empty(t, maps.Keys(ts.Locations[l13]))
	assert.ElementsMatch(t, []string{i3, i4}, maps.Keys(ts.Locations[l14]))
	assert.ElementsMatch(t, []string{"fnord"}, maps.Keys(ts.Locations[b4]))

	table := []struct {
		name         string
		from         string
		to           string
		expect       func(*testing.T, *locSet)
		expectNewLoc string
	}{
		{
			name: "nop root",
			from: l2,
			to:   "",
			expect: func(t *testing.T, ls *locSet) {
				assert.ElementsMatch(t, []string{l1, l13, l14, b4}, maps.Keys(ls.Locations))
				assert.Empty(t, maps.Keys(ls.Locations[l2]))
			},
			expectNewLoc: l2,
		},
		{
			name: "nop child",
			from: l2,
			to:   "foo",
			expect: func(t *testing.T, ls *locSet) {
				assert.ElementsMatch(t, []string{l1, l13, l14, b4}, maps.Keys(ls.Locations))
				assert.Empty(t, maps.Keys(ls.Locations["foo"]))
				assert.Empty(t, maps.Keys(ls.Locations["foo/"+l2]))
			},
			expectNewLoc: "foo/" + l2,
		},
		{
			name: "no items",
			from: l13,
			to:   "foo",
			expect: func(t *testing.T, ls *locSet) {
				newLoc := "foo/lr_3"
				assert.ElementsMatch(t, []string{l1, newLoc, l14, b4}, maps.Keys(ls.Locations))
				assert.Empty(t, maps.Keys(ls.Locations[l13]))
				assert.Empty(t, maps.Keys(ls.Locations[newLoc]))
			},
			expectNewLoc: "foo/lr_3",
		},
		{
			name: "with items",
			from: l14,
			to:   "foo",
			expect: func(t *testing.T, ls *locSet) {
				newLoc := "foo/lr_4"
				assert.ElementsMatch(t, []string{l1, l13, newLoc, b4}, maps.Keys(ls.Locations))
				assert.Empty(t, maps.Keys(ls.Locations[l14]))
				assert.ElementsMatch(t, []string{i3, i4}, maps.Keys(ls.Locations[newLoc]))
			},
			expectNewLoc: "foo/lr_4",
		},
		{
			name: "cascading locations",
			from: l1,
			to:   "foo",
			expect: func(t *testing.T, ls *locSet) {
				pfx := "foo/"
				assert.ElementsMatch(t, []string{pfx + l1, pfx + l13, pfx + l14, b4}, maps.Keys(ls.Locations))
				assert.Empty(t, maps.Keys(ls.Locations[l1]))
				assert.Empty(t, maps.Keys(ls.Locations[l14]))
				assert.Empty(t, maps.Keys(ls.Locations[l13]))
				assert.ElementsMatch(t, []string{i1, i2}, maps.Keys(ls.Locations[pfx+l1]))
				assert.Empty(t, maps.Keys(ls.Locations[pfx+l13]))
				assert.ElementsMatch(t, []string{i3, i4}, maps.Keys(ls.Locations[pfx+l14]))
			},
			expectNewLoc: "foo/" + l1,
		},
		{
			name: "to existing location",
			from: l14,
			to:   "bar",
			expect: func(t *testing.T, ls *locSet) {
				assert.ElementsMatch(t, []string{l1, l13, b4}, maps.Keys(ls.Locations))
				assert.Empty(t, maps.Keys(ls.Locations[l14]))
				assert.Empty(t, maps.Keys(ls.Locations["bar"]))
				assert.ElementsMatch(t, []string{"fnord", i3, i4}, maps.Keys(ls.Locations[b4]))
			},
			expectNewLoc: b4,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			ls := makeSet()

			newLoc := ls.MoveLocation(test.from, test.to)
			test.expect(t, ls)
			assert.Equal(t, test.expectNewLoc, newLoc)
		})
	}
}
