package utils

import "errors"

// RequireProps validates the existence of the properties
//  in the map.  Expects the format map[propName]propVal.
func RequireProps(props map[string]string) error {
	for name, val := range props {
		if len(val) == 0 {
			return errors.New(name + " is required to perform this command")
		}
	}
	return nil
}
