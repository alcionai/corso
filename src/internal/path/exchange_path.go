package path

var _ Path = &ExchangeMail{}

// ExchangeMail represents resource paths in corso for M365 Exchange mail emails
// and folders. Paths have the general for of:
// <tenant ID>/exchange/<user ID>/email/<folder display name>/<item ID>
//
// There can be one or more folders in the path, but the item, if present, is
// always the final path element.
type ExchangeMail struct {
	Builder
	hasItem bool
}

// Tenant returns the tenant ID for the referenced email resource.
func (emp ExchangeMail) Tenant() string {
	return emp.Builder.elements[0]
}

// Cateory returns an identifier noting this is a path for an email resource.
func (emp ExchangeMail) Category() string {
	return emp.Builder.elements[3]
}

// User returns the user ID for the referenced email resource.
func (emp ExchangeMail) User() string {
	return emp.Builder.elements[2]
}

// Folder returns the folder segment for the referenced email resource.
func (emp ExchangeMail) Folder() string {
	if emp.hasItem {
		return emp.Builder.join(4, len(emp.Builder.elements)-1)
	}

	return emp.Builder.join(4, len(emp.Builder.elements))
}

// Mail returns the email ID for the referenced email resource.
func (emp ExchangeMail) Item() string {
	if emp.hasItem {
		return emp.Builder.elements[len(emp.Builder.elements)-1]
	}

	return ""
}
