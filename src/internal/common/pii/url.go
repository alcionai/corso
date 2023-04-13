package pii

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/alcionai/clues"
	"golang.org/x/exp/maps"
)

// SafeURL complies with the clues.Concealer and fmt.Stringer
// interfaces to produce a safely loggable version of the URL.
// Path elements that equal a SafePathWords entry will show in
// plain text.  All other path elements will get hashed by clues.
// Query parameters that match a key in SafeQueryParams will have
// their values displayed in plain text.  All other query parames
// will get hashed by clues.
type SafeURL struct {
	// the original URL
	URL string
	// path elements that do not need to be hidden
	// keys should be lower-cased
	SafePathElems map[string]struct{}
	// query parameters that do not need to be hidden
	// keys should be lower-cased
	SafeQueryKeys map[string]struct{}
}

var _ clues.Concealer = &SafeURL{}

// Conceal produces a string of the url with the sensitive info
// obscured (hashed or replaced).
func (u SafeURL) Conceal() string {
	if len(u.URL) == 0 {
		return ""
	}

	p, err := url.Parse(u.URL)
	if err != nil {
		return "malformed-URL"
	}

	elems := ConcealElements(strings.Split(p.EscapedPath(), "/"), u.SafePathElems)
	qry := maps.Clone(p.Query())

	// conceal any non-safe query param values
	for k, v := range p.Query() {
		if _, ok := u.SafeQueryKeys[strings.ToLower(k)]; ok {
			continue
		}

		for i := range v {
			v[i] = clues.Conceal(v[i])
		}

		qry[k] = v
	}

	je := strings.Join(elems, "/")
	esc := p.Scheme + "://" + p.Hostname() + je

	if len(qry) > 0 {
		esc += "?" + qry.Encode()
	}

	unesc, err := url.QueryUnescape(esc)
	if err != nil {
		return esc
	}

	return unesc
}

// Format ensures the safeURL will output the Conceal() version
// even when used in a PrintF.
func (u SafeURL) Format(fs fmt.State, _ rune) {
	fmt.Fprint(fs, u.Conceal())
}

// String complies with Stringer to ensure the Conceal() version
// of the url is printed anytime it gets transformed to a string.
func (u SafeURL) String() string {
	return u.Conceal()
}
