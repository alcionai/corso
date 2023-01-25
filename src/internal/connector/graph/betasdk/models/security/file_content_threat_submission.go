package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FileContentThreatSubmission 
type FileContentThreatSubmission struct {
    FileThreatSubmission
    // It specifies the file content in base 64 format.
    fileContent *string
}
// NewFileContentThreatSubmission instantiates a new FileContentThreatSubmission and sets the default values.
func NewFileContentThreatSubmission()(*FileContentThreatSubmission) {
    m := &FileContentThreatSubmission{
        FileThreatSubmission: *NewFileThreatSubmission(),
    }
    odataTypeValue := "#microsoft.graph.security.fileContentThreatSubmission";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateFileContentThreatSubmissionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFileContentThreatSubmissionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileContentThreatSubmission(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FileContentThreatSubmission) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.FileThreatSubmission.GetFieldDeserializers()
    res["fileContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileContent(val)
        }
        return nil
    }
    return res
}
// GetFileContent gets the fileContent property value. It specifies the file content in base 64 format.
func (m *FileContentThreatSubmission) GetFileContent()(*string) {
    return m.fileContent
}
// Serialize serializes information the current object
func (m *FileContentThreatSubmission) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.FileThreatSubmission.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("fileContent", m.GetFileContent())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFileContent sets the fileContent property value. It specifies the file content in base 64 format.
func (m *FileContentThreatSubmission) SetFileContent(value *string)() {
    m.fileContent = value
}
