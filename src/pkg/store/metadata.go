package store

// MetadataFile holds a standard representation of a
// metadata file. Primarily used for debugging purposes.
type MetadataFile struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Data any    `json:"data"`
}

// TODO: printable support
// var _ print.Printable = &MetadataFile{}
