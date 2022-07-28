package path

import (
	"strings"

	"github.com/pkg/errors"
)

const (
	emailCategory = "email"
)

var _ Path = &ExchangeMail{}

type ExchangeMail struct {
	Base
}

// NewExchangeEmailPath creates and returns a new ExchangeEmailPath struct after
// verifying the path is properly escaped and contains information for the
// required segments. The provided segments and folder elements should not be
// escaped prior to calling this.
func NewExchangeMail(
	tenant string,
	user string,
	folder []string,
	item string,
) (*ExchangeMail, error) {
	tmpFolder := strings.Join(folder, "")
	if err := validateExchangeMailSegments(tenant, user, tmpFolder, item); err != nil {
		return nil, err
	}

	p := newPath([][]string{
		{tenant},
		{emailCategory},
		{user},
		folder,
		{item},
	})

	return &ExchangeMail{p}, nil
}

// NewExchangeMailFromEscapedSegments takes a series of already escaped segments
// representing the tenant, user, folder, and item validates them and returns a
// *ExchangeMail. The caller is expected to concatenate of all folders
// into a single string like `some/subfolder/structure`. Any special characters
// in the folder path need to be escaped.
func NewExchangeMailFromEscapedSegments(tenant, user, folder, item string) (*ExchangeMail, error) {
	if err := validateExchangeMailSegments(tenant, user, folder, item); err != nil {
		return nil, err
	}

	p, err := newPathFromEscapedSegments([]string{tenant, emailCategory, user, folder, item})
	if err != nil {
		return nil, err
	}

	return &ExchangeMail{p}, nil
}

func validateExchangeMailSegments(tenant, user, folder, item string) error {
	if len(tenant) == 0 {
		return errors.Wrap(errMissingSegment, "tenant")
	}

	if len(user) == 0 {
		return errors.Wrap(errMissingSegment, "user")
	}

	if len(folder) == 0 {
		return errors.Wrap(errMissingSegment, "mail folder")
	}

	if len(item) == 0 {
		return errors.Wrap(errMissingSegment, "mail item")
	}

	return nil
}

// Tenant returns the tenant ID for the referenced email resource.
func (emp ExchangeMail) Tenant() string {
	return emp.segment(0)
}

// Cateory returns an identifier noting this is a path for an email resource.
func (emp ExchangeMail) Category() string {
	return emp.segment(1)
}

// User returns the user ID for the referenced email resource.
func (emp ExchangeMail) User() string {
	return emp.segment(2)
}

// Folder returns the folder segment for the referenced email resource.
func (emp ExchangeMail) Folder() string {
	return emp.segment(3)
}

func (emp ExchangeMail) FolderElements() []string {
	return emp.unescapedSegmentElements(3)
}

// Mail returns the email ID for the referenced email resource.
func (emp ExchangeMail) Item() string {
	return emp.segment(4)
}
