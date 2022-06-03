package dataTransport

import "io"

// GraphData is an interface that encapsulates serialized data
// from the M365 to the BackupWriter
type GraphData interface{
    //Provides file data to BackupWriter
    ToReader()  io.Reader
    //Provides file name for BackupWriter
    UUID() string
}
