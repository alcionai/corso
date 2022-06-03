package dataTransport

// GraphData is an interface that encapsulates serialized data
// from the M365 to the BackupWriter
type GraphData interface {
	//Provides file name for BackupWriter
	UUID() string
}
