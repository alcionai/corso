package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FileThreatSubmission 
type FileThreatSubmission struct {
    ThreatSubmission
    // It specifies the file name to be submitted.
    fileName *string
}
// NewFileThreatSubmission instantiates a new FileThreatSubmission and sets the default values.
func NewFileThreatSubmission()(*FileThreatSubmission) {
    m := &FileThreatSubmission{
        ThreatSubmission: *NewThreatSubmission(),
    }
    odataTypeValue := "#microsoft.graph.security.fileThreatSubmission";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateFileThreatSubmissionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFileThreatSubmissionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.security.fileContentThreatSubmission":
                        return NewFileContentThreatSubmission(), nil
                    case "#microsoft.graph.security.fileUrlThreatSubmission":
                        return NewFileUrlThreatSubmission(), nil
                }
            }
        }
    }
    return NewFileThreatSubmission(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FileThreatSubmission) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ThreatSubmission.GetFieldDeserializers()
    res["fileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileName(val)
        }
        return nil
    }
    return res
}
// GetFileName gets the fileName property value. It specifies the file name to be submitted.
func (m *FileThreatSubmission) GetFileName()(*string) {
    return m.fileName
}
// Serialize serializes information the current object
func (m *FileThreatSubmission) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ThreatSubmission.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("fileName", m.GetFileName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFileName sets the fileName property value. It specifies the file name to be submitted.
func (m *FileThreatSubmission) SetFileName(value *string)() {
    m.fileName = value
}
