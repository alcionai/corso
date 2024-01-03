package testdata

import _ "embed"

//go:embed contacts-input.json
var ContactsInput string

//go:embed contacts-output.eml
var ContactsOutput string
