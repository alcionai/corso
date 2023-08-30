package path

import (
	"bytes"
	"crypto/sha256"
	"fmt"

	"github.com/alcionai/clues"
)

// interface compliance required for handling PII
var (
	_ clues.Concealer = &Builder{}
	_ fmt.Stringer    = &Builder{}
)

// Builder is a simple path representation that only tracks path elements. It
// can join, escape, and unescape elements. Higher-level packages are expected
// to wrap this struct to build resource-specific contexts (e.x. an
// ExchangeMailPath).
// Resource-specific paths allow access to more information like segments in the
// path. Builders that are turned into resource paths later on do not need to
// manually add prefixes for items that normally appear in the data layer (ex.
// tenant ID, service, user ID, etc).
type Builder struct {
	// Unescaped version of elements.
	elements Elements
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

// Dir removes the last element from the builder.
func (pb Builder) Dir() *Builder {
	if len(pb.elements) <= 1 {
		return &Builder{}
	}

	return &Builder{
		// Safe to use the same elements because Builders are immutable.
		elements: pb.elements[:len(pb.elements)-1],
	}
}

// HeadElem returns the first element in the Builder.
func (pb Builder) HeadElem() string {
	if len(pb.elements) == 0 {
		return ""
	}

	return pb.elements[0]
}

// LastElem returns the last element in the Builder.
func (pb Builder) LastElem() string {
	if len(pb.elements) == 0 {
		return ""
	}

	return pb.elements[len(pb.elements)-1]
}

// UpdateParent updates leading elements matching prev to be cur and returns
// true if it was updated. If prev is not a prefix of this Builder changes
// nothing and returns false. If either prev or cur is nil does nothing and
// returns false.
func (pb *Builder) UpdateParent(prev, cur *Builder) bool {
	if prev == cur || prev == nil || cur == nil || len(prev.Elements()) > len(pb.Elements()) {
		return false
	}

	parent := true

	for i, e := range prev.Elements() {
		if pb.elements[i] != e {
			parent = false
			break
		}
	}

	if !parent {
		return false
	}

	pb.elements = append(cur.Elements(), pb.elements[len(prev.Elements()):]...)

	return true
}

// ShortRef produces a truncated hash of the builder that
// acts as a unique identifier.
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
func (pb Builder) Elements() Elements {
	return append(Elements{}, pb.elements...)
}

// withPrefix creates a Builder prefixed with the parameter values, and
// concatenated with the current builder elements.
func (pb Builder) withPrefix(elements ...string) *Builder {
	res := Builder{}.Append(elements...)
	res.elements = append(res.elements, pb.elements...)

	return res
}

// verifyPrefix ensures that the tenant and resourceOwner are valid
// values, and that the builder has some directory structure.
func (pb Builder) verifyPrefix(tenant, resourceOwner string) error {
	if err := verifyInputValues(tenant, resourceOwner); err != nil {
		return err
	}

	if len(pb.elements) == 0 {
		return clues.New("missing path beyond prefix")
	}

	return nil
}

// ---------------------------------------------------------------------------
// Data Layer Path Transformers
// ---------------------------------------------------------------------------

func (pb Builder) ToStreamStorePath(
	tenant, purpose string,
	service ServiceType,
	isItem bool,
) (Path, error) {
	if err := verifyInputValues(tenant, purpose); err != nil {
		return nil, err
	}

	if isItem && len(pb.elements) == 0 {
		return nil, clues.New("missing path beyond prefix")
	}

	metadataService := UnknownService

	switch service {
	case ExchangeService:
		metadataService = ExchangeMetadataService
	case OneDriveService:
		metadataService = OneDriveMetadataService
	case SharePointService:
		metadataService = SharePointMetadataService
	case GroupsService:
		metadataService = GroupsMetadataService
	}

	return &dataLayerResourcePath{
		Builder: *pb.withPrefix(
			tenant,
			metadataService.String(),
			purpose,
			DetailsCategory.String()),
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
	if err := ValidateServiceAndCategory(service, category); err != nil {
		return nil, err
	}

	if err := verifyInputValues(tenant, user); err != nil {
		return nil, err
	}

	if isItem && len(pb.elements) == 0 {
		return nil, clues.New("missing path beyond prefix")
	}

	metadataService := UnknownService

	switch service {
	case ExchangeService:
		metadataService = ExchangeMetadataService
	case OneDriveService:
		metadataService = OneDriveMetadataService
	case SharePointService:
		metadataService = SharePointMetadataService
	case GroupsService:
		metadataService = GroupsMetadataService
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
	elems ...string,
) (Path, error) {
	if err := ValidateServiceAndCategory(service, category); err != nil {
		return nil, err
	}

	if err := pb.verifyPrefix(tenant, user); err != nil {
		return nil, err
	}

	prefixItems := append([]string{
		tenant,
		service.String(),
		user,
		category.String(),
	}, elems...)

	return &dataLayerResourcePath{
		Builder:  *pb.withPrefix(prefixItems...),
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

// ---------------------------------------------------------------------------
// Stringers and PII Concealer Compliance
// ---------------------------------------------------------------------------

// Conceal produces a concealed representation of the builder, suitable for
// logging, storing in errors, and other output.
func (pb Builder) Conceal() string {
	return pb.elements.Conceal()
}

// Format produces a concealed representation of the builder, even when
// used within a PrintF, suitable for logging, storing in errors,
// and other output.
func (pb Builder) Format(fs fmt.State, _ rune) {
	fmt.Fprint(fs, pb.Conceal())
}

// String returns a string that contains all path elements joined together.
// Elements of the path that need escaping are escaped.
// The result is not concealed, and is not suitable for logging or structured
// errors.
func (pb Builder) String() string {
	return pb.elements.String()
}

// PlainString returns an unescaped, unmodified string of the builder.
// The result is not concealed, and is not suitable for logging or structured
// errors.
func (pb Builder) PlainString() string {
	return pb.elements.PlainString()
}
