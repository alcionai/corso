package metadata

// ConversationPostMetadata stores metadata for a given conversation post,
// stored as a .meta file in kopia.
type ConversationPostMetadata struct {
	Recipients []string `json:"recipients,omitempty"`
	Topic      string   `json:"topic,omitempty"`
}
