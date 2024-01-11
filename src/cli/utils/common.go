package utils

import (
	"errors"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/pkg/selectors"
)

func includeFolderAndPageSelectors(
	sel, opts any,
	folderPaths, fileNames, pageFolders, pageItems, lists []string,
) {
	var (
		pages               []string
		files               []string
		containsFolders     []string
		prefixFolders       []string
		containsPageFolders []string
		prefixPageFolders   []string
	)

	files = fileNames
	if len(files) == 0 {
		files = selectors.Any()
	}

	pages = pageItems
	if len(pages) == 0 {
		pages = selectors.Any()
	}

	if len(folderPaths)+len(fileNames) > 0 {
		trimmedFolderPaths := trimFolderSlash(folderPaths)
		containsFolders, prefixFolders = splitFoldersIntoContainsAndPrefix(trimmedFolderPaths)
	}

	if len(pageFolders)+len(pageItems) > 0 {
		trimmedPageFolderPaths := trimFolderSlash(pageFolders)
		containsPageFolders, prefixPageFolders = splitFoldersIntoContainsAndPrefix(trimmedPageFolderPaths)
	}

	switch s := sel.(type) {
	case *selectors.SharePointRestore:
		if len(containsFolders) > 0 {
			s.Include(s.LibraryItems(containsFolders, files))
		}

		if len(prefixFolders) > 0 {
			s.Include(s.LibraryItems(prefixFolders, files, selectors.PrefixMatch()))
		}

		configureSharepointListsSelector(sel, lists)

		if len(containsPageFolders) > 0 {
			s.Include(s.PageItems(containsPageFolders, pages))
		}

		if len(prefixPageFolders) > 0 {
			s.Include(s.PageItems(prefixPageFolders, pages, selectors.PrefixMatch()))
		}
	case *selectors.GroupsRestore:
		if len(containsFolders) > 0 {
			s.Include(s.LibraryItems(containsFolders, files))
		}

		if len(prefixFolders) > 0 {
			s.Include(s.LibraryItems(prefixFolders, files, selectors.PrefixMatch()))
		}

		configureSharepointListsSelector(sel, lists)

		if len(containsPageFolders) > 0 {
			s.Include(s.PageItems(containsPageFolders, pages))
		}

		if len(prefixPageFolders) > 0 {
			s.Include(s.PageItems(prefixPageFolders, pages, selectors.PrefixMatch()))
		}
	}

	switch opts := opts.(type) {
	case GroupsOpts:
		opts.Page = pages
		opts.FileName = files
	case SharePointOpts:
		opts.Page = pages
		opts.FileName = files
	}
}

func validateCommonTimeFlags(opts any) error {
	timeFlags := []string{
		flags.FileCreatedAfterFN,
		flags.FileCreatedBeforeFN,
		flags.FileModifiedAfterFN,
		flags.FileModifiedBeforeFN,
	}

	isFlagPopulated := func(opts any, flag string) bool {
		switch opts := opts.(type) {
		case GroupsOpts:
			_, ok := opts.Populated[flag]
			return ok
		case SharePointOpts:
			_, ok := opts.Populated[flag]
			return ok
		default:
			return false
		}
	}

	getTimeField := func(opts any, flag string) (string, error) {
		switch opts := opts.(type) {
		case GroupsOpts:
			return opts.FileTimeField(flag), nil
		case SharePointOpts:
			return opts.FileTimeField(flag), nil
		default:
			return "", errors.New("unsupported type")
		}
	}

	for _, flag := range timeFlags {
		if populated := isFlagPopulated(opts, flag); populated {
			timeField, err := getTimeField(opts, flag)
			if err != nil {
				return err
			}

			if !IsValidTimeFormat(timeField) {
				return clues.New("invalid time format for " + flag)
			}
		}
	}

	return nil
}

func configureSharepointListsSelector(sel any, optsList []string) {
	if len(optsList) > 0 {
		optsList = trimFolderSlash(optsList)
		switch s := sel.(type) {
		case *selectors.SharePointRestore:
			s.Include(s.ListItems(optsList, optsList, selectors.StrictEqualMatch()))
			s.Configure(selectors.Config{OnlyMatchItemNames: true})
		case *selectors.GroupsRestore:
			s.Include(s.ListItems(optsList, optsList, selectors.StrictEqualMatch()))
			s.Configure(selectors.Config{OnlyMatchItemNames: true})
		}
	}
}
