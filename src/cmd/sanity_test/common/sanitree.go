package common

import (
	"context"

	"github.com/alcionai/clues"
	"golang.org/x/exp/maps"

	"github.com/alcionai/canario/src/pkg/path"
)

type Sanileaf[T, L any] struct {
	Parent *Sanitree[T, L]
	Self   L
	ID     string
	Name   string
	Size   int64

	// Expand is an arbitrary k:v map of any data that is
	// uniquely scrutinized by a given service.
	Expand map[string]any
}

// Sanitree is used to build out a hierarchical tree of items
// for comparison against each other.  Primarily so that a restore
// can compare two subtrees easily.
type Sanitree[T, L any] struct {
	Parent *Sanitree[T, L]

	Self T
	ID   string
	Name string

	// CountLeaves is the number of non-container child items.
	// Used for services that don't need full item metadata, and
	// just want a count of children.
	CountLeaves int
	// leaves are non-container child items.  Used by services
	// that need more than just a count of items.
	// name (or equivalent) -> leaf
	Leaves map[string]*Sanileaf[T, L]
	// Children holds all child containers
	// name -> node
	Children map[string]*Sanitree[T, L]

	// Expand is an arbitrary k:v map of any data that is
	// uniquely scrutinized by a given service.
	Expand map[string]any
}

func (s *Sanitree[T, L]) Path() path.Elements {
	if s.Parent == nil {
		return path.NewElements(s.Name)
	}

	fp := s.Parent.Path()

	return append(fp, s.Name)
}

func (s *Sanitree[T, L]) NodeAt(
	ctx context.Context,
	elems []string,
) *Sanitree[T, L] {
	node := s

	for _, e := range elems {
		child, ok := node.Children[e]

		Assert(
			ctx,
			func() bool { return ok },
			"tree node should contain next child",
			s.Path(),
			maps.Keys(s.Children))

		node = child
	}

	return node
}

// ---------------------------------------------------------------------------
// Comparing trees
// ---------------------------------------------------------------------------

type (
	ContainerComparatorFn[ET, EL, RT, RL any] func(
		ctx context.Context,
		expect *Sanitree[ET, EL],
		result *Sanitree[RT, RL])
	LeafComparatorFn[ET, EL, RT, RL any] func(
		ctx context.Context,
		expect *Sanileaf[ET, EL],
		result *Sanileaf[RT, RL])
)

func AssertEqualTrees[ET, EL, RT, RL any](
	ctx context.Context,
	expect *Sanitree[ET, EL],
	result *Sanitree[RT, RL],
	customContainerCheck ContainerComparatorFn[ET, EL, RT, RL],
	customLeafCheck LeafComparatorFn[ET, EL, RT, RL],
) {
	if expect == nil && result == nil {
		return
	}

	Debugf(ctx, "comparing trees at path: %+v", expect.Path())

	checkChildrenAndLeaves(ctx, expect, result)
	ctx = clues.Add(ctx, "container_name", expect.Name)

	if customContainerCheck != nil {
		customContainerCheck(ctx, expect, result)
	}

	CompareLeaves[ET, EL, RT, RL](
		ctx,
		expect.Leaves,
		result.Leaves,
		customLeafCheck)

	// recurse
	for name, s := range expect.Children {
		r, ok := result.Children[name]
		Assert(
			ctx,
			func() bool { return ok },
			"found matching child container",
			name,
			maps.Keys(result.Children))

		AssertEqualTrees(ctx, s, r, customContainerCheck, customLeafCheck)
	}
}

// ---------------------------------------------------------------------------
// Comparing differently typed trees.
// ---------------------------------------------------------------------------

type NodeComparator[ET, EL, RT, RL any] func(
	ctx context.Context,
	expect *Sanitree[ET, EL],
	result *Sanitree[RT, RL],
)

// CompareDiffTrees recursively compares two sanitrees that have
// different data types.  The two trees are expected to represent
// a common hierarchy.
//
// Additional comparisons besides the tree hierarchy are optionally
// left to the caller by population of the NodeComparator func.
func CompareDiffTrees[ET, EL, RT, RL any](
	ctx context.Context,
	expect *Sanitree[ET, EL],
	result *Sanitree[RT, RL],
	comparator NodeComparator[ET, EL, RT, RL],
) {
	if expect == nil && result == nil {
		return
	}

	Debugf(ctx, "comparing tree at path: %+v", expect.Path())

	checkChildrenAndLeaves(ctx, expect, result)
	ctx = clues.Add(ctx, "container_name", expect.Name)

	if comparator != nil {
		comparator(ctx, expect, result)
	}

	// recurse
	for name, s := range expect.Children {
		r, ok := result.Children[name]
		Assert(
			ctx,
			func() bool { return ok },
			"found matching child container",
			name,
			maps.Keys(result.Children))

		CompareDiffTrees(ctx, s, r, comparator)
	}
}

// ---------------------------------------------------------------------------
// Checking hierarchy likeness
// ---------------------------------------------------------------------------

func checkChildrenAndLeaves[ET, EL, RT, RL any](
	ctx context.Context,
	expect *Sanitree[ET, EL],
	result *Sanitree[RT, RL],
) {
	Assert(
		ctx,
		func() bool { return expect != nil },
		"expected stree is nil",
		"not nil",
		expect)

	Assert(
		ctx,
		func() bool { return result != nil },
		"result stree is nil",
		"not nil",
		result)

	ctx = clues.Add(ctx, "container_name", expect.Name)

	Assert(
		ctx,
		func() bool { return expect.Name == result.Name },
		"container names match",
		expect.Name,
		result.Name)

	Assert(
		ctx,
		func() bool { return expect.CountLeaves == result.CountLeaves },
		"count of leaves in container matches",
		expect.CountLeaves,
		result.CountLeaves)

	Assert(
		ctx,
		func() bool { return len(expect.Leaves) == len(result.Leaves) },
		"len of leaves in container matches",
		len(expect.Leaves),
		len(result.Leaves))

	Assert(
		ctx,
		func() bool { return len(expect.Children) == len(result.Children) },
		"count of child containers matches",
		len(expect.Children),
		len(result.Children))
}

func CompareLeaves[ET, EL, RT, RL any](
	ctx context.Context,
	expect map[string]*Sanileaf[ET, EL],
	result map[string]*Sanileaf[RT, RL],
	customLeafCheck LeafComparatorFn[ET, EL, RT, RL],
) {
	for name, l := range expect {
		ictx := clues.Add(ctx, "leaf_name", l.Name)

		r, ok := result[name]
		Assert(
			ictx,
			func() bool { return ok },
			"found matching leaf item",
			name,
			maps.Keys(result))

		Assert(
			ictx,
			func() bool { return l.Size == r.Size },
			"leaf sizes match",
			l.Size,
			r.Size)

		if customLeafCheck != nil {
			customLeafCheck(ictx, l, r)
		}
	}
}
