package api

// DeltaUpdate holds the results of a current delta token.  It normally
// gets produced when aggregating the addition and removal of items in
// a delta-queryable folder.
type DeltaUpdate struct {
	// the deltaLink itself
	URL string
	// true if the old delta was marked as invalid
	Reset bool
}
