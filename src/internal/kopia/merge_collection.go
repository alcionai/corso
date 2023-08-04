package kopia

import (
	"context"
	"errors"

	"github.com/alcionai/clues"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

var _ data.RestoreCollection = &mergeCollection{}

type col struct {
	storagePath string
	data.RestoreCollection
}

type mergeCollection struct {
	cols []col
	// Technically don't need to track this but it can help detect errors.
	fullPath path.Path
}

func (mc *mergeCollection) addCollection(
	storagePath string,
	c data.RestoreCollection,
) error {
	if c == nil {
		return clues.New("adding nil collection").
			With("current_path", mc.FullPath())
	} else if mc.FullPath().String() != c.FullPath().String() {
		return clues.New("attempting to merge collection with different path").
			With("current_path", mc.FullPath(), "new_path", c.FullPath())
	}

	mc.cols = append(mc.cols, col{storagePath: storagePath, RestoreCollection: c})

	// Keep a stable sorting of this merged collection set so we can say there's
	// some deterministic behavior when Fetch is called. We don't expect to have
	// to merge many collections.
	slices.SortStableFunc(mc.cols, func(a, b col) int {
		switch true {
		case a.storagePath < b.storagePath:
			return -1
		case a.storagePath > b.storagePath:
			return 1
		default:
			return 0
		}
	})

	return nil
}

func (mc mergeCollection) FullPath() path.Path {
	return mc.fullPath
}

func (mc *mergeCollection) Items(
	ctx context.Context,
	errs *fault.Bus,
) <-chan data.Stream {
	res := make(chan data.Stream)

	go func() {
		defer close(res)

		logger.Ctx(ctx).Infow(
			"getting items for merged collection",
			"merged_collection_count", len(mc.cols))

		for _, c := range mc.cols {
			// Unfortunately doesn't seem to be a way right now to see if the
			// iteration failed and we should be exiting early.
			ictx := clues.Add(
				ctx,
				"merged_collection_storage_path", path.LoggableDir(c.storagePath))
			logger.Ctx(ictx).Debug("sending items from merged collection")

			for item := range c.Items(ictx, errs) {
				res <- item
			}
		}
	}()

	return res
}

// Fetch goes through all the collections in this one and returns the first
// match found or the first error that is not data.ErrNotFound. If multiple
// collections have the requested item, the instance in the collection with the
// lexicographically smallest storage path is returned.
func (mc *mergeCollection) FetchItemByName(
	ctx context.Context,
	name string,
) (data.Stream, error) {
	logger.Ctx(ctx).Infow(
		"fetching item in merged collection",
		"merged_collection_count", len(mc.cols))

	for _, c := range mc.cols {
		ictx := clues.Add(
			ctx,
			"merged_collection_storage_path", path.LoggableDir(c.storagePath))

		logger.Ctx(ictx).Debug("looking for item in merged collection")

		s, err := c.FetchItemByName(ictx, name)
		if err == nil {
			return s, nil
		} else if err != nil && !errors.Is(err, data.ErrNotFound) {
			return nil, clues.Wrap(err, "fetching from merged collection").
				WithClues(ictx)
		}
	}

	return nil, clues.Wrap(data.ErrNotFound, "merged collection fetch")
}
