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
	"bytes"
	"crypto/sha256"
	"fmt"
	"strings"

	"github.com/alcionai/clues"
	"github.com/pkg/errors"
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
	errMissingSegment = errors.New("missing required path element")
	errParsingPath    = errors.New("parsing resource path")
)

// For now, adding generic functions to pull information from segments.
// Resources that don't have the requested information should return an empty
// string.
type Path interface {
	String() string
	Service() ServiceType
	Category() CategoryType
	Tenant() string
	ResourceOwner() string
	Folder(bool) string
	Folders() []string
	Item() string
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
	Elements() []string
	// Append returns a new Path object with the given element added to the end of
	// the old Path if possible. If the old Path is an item Path then Append
	// returns an error.
	Append(element string, isItem bool) (Path, error)
	// ShortRef returns a short reference representing this path. The short
	// reference is guaranteed to be unique. No guarantees are made about whether
	// a short reference can be converted back into the Path that generated it.
	ShortRef() string
	// ToBuilder returns a Builder instance that represents the current Path.
	ToBuilder() *Builder
}

// Builder is a simple path representation that only tracks path elements. It
// can join, escape, and unescape elements. Higher-level packages are expected
// to wrap this struct to build resource-speicific contexts (e.x. an
// ExchangeMailPath).
// Resource-specific paths allow access to more information like segments in the
// path. Builders that are turned into resource paths later on do not need to
// manually add prefixes for items that normally appear in the data layer (ex.
// tenant ID, service, user ID, etc).
type Builder struct {
	// Unescaped version of elements.
	elements []string
}

// UnescapeAndAppend creates a copy of this Builder and adds one or more already
// escaped path elements to the end of the new Builder. Elements are added in
// the order they are passed.
func (pb Builder) UnescapeAndAppend(elements ...string) (*Builder, error) {
	res := &Builder{elements: make([]string, 0, len(pb.elements))}
	copy(res.elements, pb.elements)

	if err := res.appendElements(true, elements); err != nil {
		return nil, err
	}

	return res, nil
}

// SplitUnescapeAppend takes in an escaped string representing a directory
// path, splits the string, and appends it to the current builder.
func (pb Builder) SplitUnescapeAppend(s string) (*Builder, error) {
	elems := Split(TrimTrailingSlash(s))

	return pb.UnescapeAndAppend(elems...)
}

// Append creates a copy of this Builder and adds the given elements them to the
// end of the new Builder. Elements are added in the order they are passed.
func (pb Builder) Append(elements ...string) *Builder {
	res := &Builder{elements: make([]string, len(pb.elements))}
	copy(res.elements, pb.elements)

	// Unescaped elements can't fail validation.
	//nolint:errcheck
	res.appendElements(false, elements)

	return res
}

func (pb *Builder) appendElements(escaped bool, elements []string) error {
	for _, e := range elements {
		if len(e) == 0 {
			continue
		}

		tmp := e

		if escaped {
			tmp = TrimTrailingSlash(tmp)
			// If tmp was just the path separator then it will be empty now.
			if len(tmp) == 0 {
				continue
			}

			if err := validateEscapedElement(tmp); err != nil {
				return err
			}

			tmp = unescape(tmp)
		}

		pb.elements = append(pb.elements, tmp)
	}

	return nil
}

func (pb Builder) PopFront() *Builder {
	if len(pb.elements) <= 1 {
		return &Builder{}
	}

	elements := make([]string, len(pb.elements)-1)
	copy(elements, pb.elements[1:])

	return &Builder{
		elements: elements,
	}
}

func (pb Builder) Dir() *Builder {
	if len(pb.elements) <= 1 {
		return &Builder{}
	}

	return &Builder{
		// Safe to use the same elements because Builders are immutable.
		elements: pb.elements[:len(pb.elements)-1],
	}
}

// String returns a string that contains all path elements joined together.
// Elements of the path that need escaping are escaped.
func (pb Builder) String() string {
	escaped := make([]string, 0, len(pb.elements))

	for _, e := range pb.elements {
		escaped = append(escaped, escapeElement(e))
	}

	return join(escaped)
}

func (pb Builder) ShortRef() string {
	if len(pb.elements) == 0 {
		return ""
	}

	data := bytes.Buffer{}

	for _, element := range pb.elements {
		data.WriteString(element)
	}

	sum := sha256.Sum256(data.Bytes())

	// Some conversions to get the right number of characters in the output. This
	// outputs hex, so we need to take the target number of characters and do the
	// equivalent of (shortRefCharacters * 4) / 8. This is
	// <number of bits represented> / <bits per byte> which gets us how many bytes
	// to give to our format command.
	numBytes := shortRefCharacters / 2

	return fmt.Sprintf("%x", sum[:numBytes])
}

// Elements returns all the elements in the path. This is a temporary function
// and will likely be updated to handle encoded elements instead of clear-text
// elements in the future.
func (pb Builder) Elements() []string {
	return append([]string{}, pb.elements...)
}

func verifyInputValues(tenant, resourceOwner string) error {
	if len(tenant) == 0 {
		return clues.Stack(errMissingSegment, errors.New("tenant"))
	}

	if len(resourceOwner) == 0 {
		return clues.Stack(errMissingSegment, errors.New("resourceOwner"))
	}

	return nil
}

func (pb Builder) verifyPrefix(tenant, resourceOwner string) error {
	if err := verifyInputValues(tenant, resourceOwner); err != nil {
		return err
	}

	if len(pb.elements) == 0 {
		return errors.New("missing path beyond prefix")
	}

	return nil
}

func (pb Builder) withPrefix(elements ...string) *Builder {
	res := Builder{}.Append(elements...)
	res.elements = append(res.elements, pb.elements...)

	return res
}

func (pb Builder) ToStreamStorePath(
	tenant, purpose string,
	service ServiceType,
	isItem bool,
) (Path, error) {
	if err := verifyInputValues(tenant, purpose); err != nil {
		return nil, err
	}

	if isItem && len(pb.elements) == 0 {
		return nil, errors.New("missing path beyond prefix")
	}

	metadataService := UnknownService

	switch service {
	case ExchangeService:
		metadataService = ExchangeMetadataService
	case OneDriveService:
		metadataService = OneDriveMetadataService
	case SharePointService:
		metadataService = SharePointMetadataService
	}

	return &dataLayerResourcePath{
		Builder: *pb.withPrefix(
			tenant,
			metadataService.String(),
			purpose,
			DetailsCategory.String(),
		),
		service:  metadataService,
		category: DetailsCategory,
		hasItem:  isItem,
	}, nil
}

func (pb Builder) ToServiceCategoryMetadataPath(
	tenant, user string,
	service ServiceType,
	category CategoryType,
	isItem bool,
) (Path, error) {
	if err := validateServiceAndCategory(service, category); err != nil {
		return nil, err
	}

	if err := verifyInputValues(tenant, user); err != nil {
		return nil, err
	}

	if isItem && len(pb.elements) == 0 {
		return nil, errors.New("missing path beyond prefix")
	}

	metadataService := UnknownService

	switch service {
	case ExchangeService:
		metadataService = ExchangeMetadataService
	case OneDriveService:
		metadataService = OneDriveMetadataService
	case SharePointService:
		metadataService = SharePointMetadataService
	}

	return &dataLayerResourcePath{
		Builder: *pb.withPrefix(
			tenant,
			metadataService.String(),
			user,
			category.String(),
		),
		service:  metadataService,
		category: category,
		hasItem:  isItem,
	}, nil
}

func (pb Builder) ToDataLayerPath(
	tenant, user string,
	service ServiceType,
	category CategoryType,
	isItem bool,
) (Path, error) {
	if err := validateServiceAndCategory(service, category); err != nil {
		return nil, err
	}

	if err := pb.verifyPrefix(tenant, user); err != nil {
		return nil, err
	}

	return &dataLayerResourcePath{
		Builder: *pb.withPrefix(
			tenant,
			service.String(),
			user,
			category.String(),
		),
		service:  service,
		category: category,
		hasItem:  isItem,
	}, nil
}

func (pb Builder) ToDataLayerExchangePathForCategory(
	tenant, user string,
	category CategoryType,
	isItem bool,
) (Path, error) {
	return pb.ToDataLayerPath(tenant, user, ExchangeService, category, isItem)
}

func (pb Builder) ToDataLayerOneDrivePath(
	tenant, user string,
	isItem bool,
) (Path, error) {
	return pb.ToDataLayerPath(tenant, user, OneDriveService, FilesCategory, isItem)
}

func (pb Builder) ToDataLayerSharePointPath(
	tenant, site string,
	category CategoryType,
	isItem bool,
) (Path, error) {
	return pb.ToDataLayerPath(tenant, site, SharePointService, category, isItem)
}

// FromDataLayerPath parses the escaped path p, validates the elements in p
// match a resource-specific path format, and returns a Path struct for that
// resource-specific type. If p does not match any resource-specific paths or
// is malformed returns an error.
func FromDataLayerPath(p string, isItem bool) (Path, error) {
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

	if len(pb.elements) < 5 {
		return nil, clues.New("path has too few segments").With("path_string", p)
	}

	service, category, err := validateServiceAndCategoryStrings(
		pb.elements[1],
		pb.elements[3],
	)
	if err != nil {
		return nil, clues.Stack(errParsingPath, err).With("path_string", p)
	}

	return &dataLayerResourcePath{
		Builder:  *pb,
		service:  service,
		category: category,
		hasItem:  isItem,
	}, nil
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
		return errors.New("trailing escape character")
	}

	return nil
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

// join returns a string containing the given elements joined by the path
// separator '/'.
func join(elements []string) string {
	// Have to use strings because path package does not handle escaped '/' and
	// '\' according to the escaping rules.
	return strings.Join(elements, string(PathSeparator))
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
