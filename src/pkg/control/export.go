package control

// ExportConfig contains config for exports
type ExportConfig struct {
	// Archive decides if we should create an archive from the data
	// instead of just returning all the files. If Archive is set to
	// true, we return a single collection with a single file which is
	// the archive.
	Archive bool

	// DataFormat
	// TODO: Enable once we support outlook exports
	// DataFormat string

	// Format decides the format in which we return the data.
	// ex: html vs pst vs other.
	// Default format is decided on a per-service or per-data basis.
	Format FormatType
}

type FormatType string

var (
	// Follow whatever format is the default for the service or data type.
	DefaultFormat FormatType
	// export the data as raw, unmodified json
	JSONFormat FormatType = "json"
)

func DefaultExportConfig() ExportConfig {
	return ExportConfig{
		Archive: false,
	}
}
