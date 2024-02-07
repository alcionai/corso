package testdata

import _ "embed"

//go:embed email-with-attachments.json
var EmailWithAttachments string

//go:embed email-with-event-info.json
var EmailWithEventInfo string

//go:embed email-with-event-object.json
var EmailWithEventObject string

//go:embed email-within-email.json
var EmailWithinEmail string
