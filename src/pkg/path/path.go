// Package path provides a set of functions for wrangling paths from the outside
// world into paths that corso can understand. Paths use the standard Unix path
// separator character '/'. If for some reason an individual element in a raw
// path contains the '/' character, it should be escaped with '\'. If the path
// contains '\' it should be escaped by turning it into '\\'.
//
// Paths can be split into elements by splitting on '/' if the '/' is not
// escaped. Additionally, corso may operate on segments in a path. Segments are
// made up of one or more path elements.
//
// Examples of paths splitting by elements and canonicalization with escaping:
// 1.
//
//	input path: `this/is/a/path`
//	elements of path: `this`, `is`, `a`, `path`
//
// 2.
//
//	input path: `this/is\/a/path`
//	elements of path: `this`, `is/a`, `path`
//
// 3.
//
//	input path: `this/is\\/a/path`
//	elements of path: `this`, `is\`, `a`, `path`
//
// 4.
//
//	input path: `this/is\\\/a/path`
//	elements of path: `this`, `is\/a`, `path`
//
// 5.
//
//	input path: `this/is//a/path`
//	elements of path: `this`, `is`, `a`, `path`
//
// 6.
//
//	input path: `this/is\//a/path`
//	elements of path: `this`, `is/`, `a`, `path`
//
// 7.
//
//	input path: `this/is/a/path/`
//	elements of path: `this`, `is`, `a`, `path`
//
// 8.
//
//	input path: `this/is/a/path\/`
//	elements of path: `this`, `is`, `a`, `path/`
package path

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/alcionai/clues"
)

const (
	escapeCharacter = '\\'
	PathSeparator   = '/'

	shortRefCharacters = 12
)

var charactersToEscape = map[rune]struct{}{
	PathSeparator:   {},
	escapeCharacter: {},
}

var (
	errMissingSegment = clues.New("missing required path element")
	errParsingPath    = clues.New("parsing resource path")
)

// For now, adding generic functions to pull information from segments.
// Resources that don't have the requested information should return an empty
// string.
type Path interface {
	String() string
	Service() ServiceType
	Category() CategoryType
	Tenant() string
	ProtectedResource() string
	Folder(escaped bool) string
	Folders() Elements
	Item() string
	// UpdateParent updates parent from old to new if the item/folder was
	// parented by old path
	UpdateParent(prev, cur Path) bool
	// PopFront returns a Builder object with the first element (left-side)
	// removed. As the resulting set of elements is no longer a valid resource
	// path a Builder is returned instead.
	PopFront() *Builder
	// Dir returns a Path object with the right-most element removed if possible.
	// If removing the right-most element would discard one of the required prefix
	// elements then an error is returned.
	Dir() (Path, error)
	// Elements returns all the elements in the path. This is a temporary function
	// and will likely be updated to handle encoded elements instead of clear-text
	// elements in the future.
	Elements() Elements
	Equal(other Path) bool
	// Append returns a new Path object with the given element added to the end of
	// the old Path if possible. If the old Path is an item Path then Append
	// returns an error.
	Append(isItem bool, elems ...string) (Path, error)
	// AppendItem is a shorthand for Append(true, someItem)
	AppendItem(item string) (Path, error)
	// ShortRef returns a short reference representing this path. The short
	// reference is guaranteed to be unique. No guarantees are made about whether
	// a short reference can be converted back into the Path that generated it.
	ShortRef() string
	// ToBuilder returns a Builder instance that represents the current Path.
	ToBuilder() *Builder

	// Every path needs to comply with these funcs to ensure that PII
	// is appropriately hidden from logging, errors, and other outputs.
	clues.Concealer
	fmt.Stringer
}

// RestorePaths denotes the location to find an item in kopia and the path of
// the collection to place the item in for restore.
type RestorePaths struct {
	StoragePath Path
	RestorePath Path
}

// ---------------------------------------------------------------------------
// Exported Helpers
// ---------------------------------------------------------------------------

func Build(
	tenant, resourceOwner string,
	service ServiceType,
	category CategoryType,
	hasItem bool,
	elements ...string,
) (Path, error) {
	b := Builder{}.Append(elements...)

	return b.ToDataLayerPath(
		tenant, resourceOwner,
		service, category,
		hasItem)
}

// BuildMetadata is a shorthand for Builder{}.Append(...).ToServiceCategoryMetadataPath(...)
func BuildMetadata(
	tenant, resourceOwner string,
	service ServiceType,
	category CategoryType,
	hasItem bool,
	elements ...string,
) (Path, error) {
	return Builder{}.
		Append(elements...).
		ToServiceCategoryMetadataPath(
			tenant, resourceOwner,
			service, category,
			hasItem)
}

// BuildOrPrefix is the same as Build, but allows for 0-len folders
// (ie: only builds the prefix).
func BuildOrPrefix(
	tenant, resourceOwner string,
	service ServiceType,
	category CategoryType,
	hasItem bool,
	elements ...string,
) (Path, error) {
	pb := Builder{}

	if err := ValidateServiceAndCategory(service, category); err != nil {
		return nil, err
	}

	if err := verifyInputValues(tenant, resourceOwner); err != nil {
		return nil, err
	}

	prefixItems := append(Elements{
		tenant,
		service.String(),
		resourceOwner,
		category.String(),
	}, elements...)

	return &dataLayerResourcePath{
		Builder:  *pb.withPrefix(prefixItems...),
		service:  service,
		category: category,
		hasItem:  hasItem,
	}, nil
}

func BuildPrefix(
	tenant, resourceOwner string,
	s ServiceType,
	c CategoryType,
) (Path, error) {
	pb := Builder{}

	if err := ValidateServiceAndCategory(s, c); err != nil {
		return nil, err
	}

	if err := verifyInputValues(tenant, resourceOwner); err != nil {
		return nil, err
	}

	return &dataLayerResourcePath{
		Builder:  *pb.withPrefix(tenant, s.String(), resourceOwner, c.String()),
		service:  s,
		category: c,
		hasItem:  false,
	}, nil
}

func pathFromDataLayerPath(p string, isItem bool, minElements int) (Path, error) {
	p = TrimTrailingSlash(p)
	// If p was just the path separator then it will be empty now.
	if len(p) == 0 {
		return nil, clues.New("logically empty path given").With("path_string", p)
	}

	// Turn into a Builder to reuse code that ignores empty elements.
	pb, err := Builder{}.UnescapeAndAppend(Split(p)...)
	if err != nil {
		return nil, clues.Stack(errParsingPath, err).With("path_string", p)
	}

	if len(pb.elements) < minElements {
		return nil, clues.New(
			"missing required tenant, service, category, protected resource ID, or non-prefix segment").
			With("path_string", pb)
	}

	service, category, err := validateServiceAndCategoryStrings(
		pb.elements[1],
		pb.elements[3])
	if err != nil {
		return nil, clues.Stack(errParsingPath, err).With("path_string", pb)
	}

	if err := verifyInputValues(pb.elements[0], pb.elements[2]); err != nil {
		return nil, clues.Stack(err).With("path_string", pb)
	}

	return &dataLayerResourcePath{
		Builder:  *pb,
		service:  service,
		category: category,
		hasItem:  isItem,
	}, nil
}

func PrefixOrPathFromDataLayerPath(p string, isItem bool) (Path, error) {
	res, err := pathFromDataLayerPath(p, isItem, 4)
	return res, clues.Stack(err).OrNil()
}

// FromDataLayerPath parses the escaped path p, validates the elements in p
// match a resource-specific path format, and returns a Path struct for that
// resource-specific type. If p does not match any resource-specific paths or
// is malformed returns an error.
func FromDataLayerPath(p string, isItem bool) (Path, error) {
	res, err := pathFromDataLayerPath(p, isItem, 5)
	return res, clues.Stack(err).OrNil()
}

// TrimTrailingSlash takes an escaped path element and returns an escaped path
// element with the trailing path separator character(s) removed if they were not
// escaped. If there were no trailing path separator character(s) or the separator(s)
// were escaped the input is returned unchanged.
func TrimTrailingSlash(element string) string {
	for len(element) > 0 && element[len(element)-1] == PathSeparator {
		lastIdx := len(element) - 1
		numSlashes := 0

		for i := lastIdx - 1; i >= 0; i-- {
			if element[i] != escapeCharacter {
				break
			}

			numSlashes++
		}

		if numSlashes%2 != 0 {
			break
		}

		element = element[:lastIdx]
	}

	return element
}

// split takes an escaped string and returns a slice of path elements. The
// string is split on the path separator according to the escaping rules. The
// provided string must not contain an unescaped trailing path separator.
func Split(segment string) []string {
	res := make([]string, 0)
	numEscapes := 0
	startIdx := 0
	// Start with true to ignore leading separator.
	prevWasSeparator := true

	for i, c := range segment {
		if c == escapeCharacter {
			prevWasSeparator = false
			numEscapes++

			continue
		}

		if c != PathSeparator {
			prevWasSeparator = false
			numEscapes = 0

			continue
		}

		// Remaining is just path separator handling.
		if numEscapes%2 != 0 {
			// This is an escaped separator.
			prevWasSeparator = false
			numEscapes = 0

			continue
		}

		// Ignore leading separator characters and don't add elements that would
		// be empty.
		if !prevWasSeparator {
			res = append(res, segment[startIdx:i])
		}

		// We don't want to include the path separator in the result.
		startIdx = i + 1
		prevWasSeparator = true
		numEscapes = 0
	}

	// Add the final segment because the loop above won't catch it. There should
	// be no trailing separator character.
	res = append(res, segment[startIdx:])

	return res
}

func ArePathsEquivalent(path1, path2 string) bool {
	normalizedPath1 := strings.TrimSpace(filepath.Clean(path1))
	normalizedPath2 := strings.TrimSpace(filepath.Clean(path2))

	normalizedPath1 = strings.TrimSuffix(normalizedPath1, string(filepath.Separator))
	normalizedPath2 = strings.TrimSuffix(normalizedPath2, string(filepath.Separator))

	return normalizedPath1 == normalizedPath2
}

// ---------------------------------------------------------------------------
// Unexported Helpers
// ---------------------------------------------------------------------------

func verifyInputValues(tenant, resourceOwner string) error {
	if len(tenant) == 0 {
		return clues.Stack(errMissingSegment, clues.New("tenant"))
	}

	if len(resourceOwner) == 0 {
		return clues.Stack(errMissingSegment, clues.New("resourceOwner"))
	}

	return nil
}

// escapeElement takes a single path element and escapes all characters that
// require an escape sequence. If there are no characters that need escaping,
// the input is returned unchanged.
func escapeElement(element string) string {
	escapeIdx := make([]int, 0)

	for i, c := range element {
		if _, ok := charactersToEscape[c]; ok {
			escapeIdx = append(escapeIdx, i)
		}
	}

	if len(escapeIdx) == 0 {
		return element
	}

	startIdx := 0
	b := strings.Builder{}
	b.Grow(len(element) + len(escapeIdx))

	for _, idx := range escapeIdx {
		b.WriteString(element[startIdx:idx])
		b.WriteRune(escapeCharacter)

		startIdx = idx
	}

	// Add the end of the element after the last escape character.
	b.WriteString(element[startIdx:])

	return b.String()
}

// unescape returns the given element and converts it into a "raw"
// element that does not have escape characters before characters that need
// escaping. Using this function on segments that contain escaped path
// separators will result in an ambiguous or incorrect segment.
func unescape(element string) string {
	b := strings.Builder{}
	startIdx := 0
	prevWasEscape := false

	for i, c := range element {
		if c != escapeCharacter || prevWasEscape {
			prevWasEscape = false
			continue
		}

		// This is an escape character, remove it from the output.
		b.WriteString(element[startIdx:i])
		startIdx = i + 1
		prevWasEscape = true
	}

	b.WriteString(element[startIdx:])

	return b.String()
}

// validateEscapedElement takes an escaped element that has had trailing
// separators trimmed and ensures that no characters requiring escaping are
// unescaped and that no escape characters are combined with characters that
// don't need escaping.
func validateEscapedElement(element string) error {
	prevWasEscape := false

	for _, c := range element {
		switch prevWasEscape {
		case true:
			prevWasEscape = false

			if _, ok := charactersToEscape[c]; !ok {
				return clues.New("bad escape sequence in path").
					With("escape_sequence", fmt.Sprintf("'%c%c'", escapeCharacter, c))
			}

		case false:
			if c == escapeCharacter {
				prevWasEscape = true
				continue
			}

			if _, ok := charactersToEscape[c]; ok {
				return clues.New("unescaped character in path").With("character", c)
			}
		}
	}

	if prevWasEscape {
		return clues.New("trailing escape character")
	}

	return nil
}

// join returns a string containing the given elements joined by the path
// separator '/'.
func join(elements []string) string {
	// Have to use strings because path package does not handle escaped '/' and
	// '\' according to the escaping rules.
	return strings.Join(elements, string(PathSeparator))
}
