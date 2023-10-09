package common

import (
	"context"

	"golang.org/x/exp/maps"
)

// Sanitree is used to build out a hierarchical tree of items
// for comparison against each other.  Primarily so that a restore
// can compare two subtrees easily.
type Sanitree[T any] struct {
	Container     T
	ContainerID   string
	ContainerName string
	// non-containers only
	ContainsItems int
	// name -> node
	Children map[string]*Sanitree[T]
}

func AssertEqualTrees[T any](
	ctx context.Context,
	expect, other *Sanitree[T],
) {
	if expect == nil && other == nil {
		return
	}

	Assert(
		ctx,
		func() bool { return expect != nil && other != nil },
		"non nil nodes",
		expect,
		other)

	Assert(
		ctx,
		func() bool { return expect.ContainerName == other.ContainerName },
		"container names match",
		expect.ContainerName,
		other.ContainerName)

	Assert(
		ctx,
		func() bool { return expect.ContainsItems == other.ContainsItems },
		"count of items in container matches",
		expect.ContainsItems,
		other.ContainsItems)

	Assert(
		ctx,
		func() bool { return len(expect.Children) == len(other.Children) },
		"count of child containers matches",
		len(expect.Children),
		len(other.Children))

	for name, s := range expect.Children {
		ch, ok := other.Children[name]
		Assert(
			ctx,
			func() bool { return ok },
			"found matching child container",
			name,
			maps.Keys(other.Children))

		AssertEqualTrees(ctx, s, ch)
	}
}
