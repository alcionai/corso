package control

// ExportConfig contains config for exports
type ExportConfig struct {
	// Archive decides if we should create an archive from the data
	// instead of just returning all the files. If Archive is set to
	// true, we return a single collection with a single file which is
	// the archive.
	Archive bool

	// DataFormat decides the format in which we return the data. This is
	// only useful for outlook exports, for example they can be in eml
	// or pst for emails.
	// TODO: Enable once we support outlook exports
	// DataFormat string
}

func DefaultExportConfig() ExportConfig {
	return ExportConfig{
		Archive: false,
	}
}
