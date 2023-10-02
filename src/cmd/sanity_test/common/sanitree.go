package common

import (
	"context"

	"golang.org/x/exp/maps"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/pkg/path"
)

type Sanileaf[T any] struct {
	Parent *Sanitree[T]
	Self   T
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
type Sanitree[T any] struct {
	Parent *Sanitree[T]

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
	Leaves map[string]*Sanileaf[T]
	// Children holds all child containers
	// name -> node
	Children map[string]*Sanitree[T]

	// Expand is an arbitrary k:v map of any data that is
	// uniquely scrutinized by a given service.
	Expand map[string]any
}

func (s *Sanitree[T]) Path() path.Elements {
	if s.Parent == nil {
		return path.NewElements(s.Name)
	}

	fp := s.Parent.Path()

	return append(fp, s.Name)
}

func (s *Sanitree[T]) NodeAt(
	ctx context.Context,
	relPath string,
) *Sanitree[T] {
	var (
		elems = path.Split(relPath)
		node  = s
	)

	Assert(
		ctx,
		func() bool { return elems[0] == s.Name },
		"relative path root should match initial sanitree node",
		relPath,
		s.Name)

	for _, e := range elems[1:] {
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
	ContainerComparatorFn[T, R any] func(
		ctx context.Context,
		expect *Sanitree[T],
		result *Sanitree[R])
	LeafComparatorFn[T, R any] func(
		ctx context.Context,
		expect *Sanileaf[T],
		result *Sanileaf[R])
)

func AssertEqualTrees[T, R any](
	ctx context.Context,
	expect *Sanitree[T],
	result *Sanitree[R],
	customContainerCheck ContainerComparatorFn[T, R],
	customLeafCheck LeafComparatorFn[T, R],
) {
	if expect == nil && result == nil {
		return
	}

	checkChildrenAndLeaves(ctx, expect, result)
	ctx = clues.Add(ctx, "container_name", expect.Name)

	if customContainerCheck != nil {
		customContainerCheck(ctx, expect, result)
	}

	CompareLeaves[T, R](
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

type NodeComparator[T, R any] func(
	ctx context.Context,
	expect *Sanitree[T],
	result *Sanitree[R],
)

// CompareDiffTrees recursively compares two sanitrees that have
// different data types.  The two trees are expected to represent
// a common hierarchy.
//
// Additional comparisons besides the tre hierarchy are optionally
// left to the caller by population of the NodeComparator func.
func CompareDiffTrees[T, R any](
	ctx context.Context,
	expect *Sanitree[T],
	result *Sanitree[R],
	comparator NodeComparator[T, R],
) {
	if expect == nil && result == nil {
		return
	}

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
// Checking hierarcy likeness
// ---------------------------------------------------------------------------

func checkChildrenAndLeaves[T, R any](
	ctx context.Context,
	expect *Sanitree[T],
	result *Sanitree[R],
) {
	Assert(
		ctx,
		func() bool { return expect != nil && result != nil },
		"non nil nodes",
		expect,
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

func CompareLeaves[T, R any](
	ctx context.Context,
	expect map[string]*Sanileaf[T],
	result map[string]*Sanileaf[R],
	customLeafCheck LeafComparatorFn[T, R],
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
