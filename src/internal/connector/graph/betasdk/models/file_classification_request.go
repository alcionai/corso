package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FileClassificationRequest 
type FileClassificationRequest struct {
    Entity
    // The file property
    file []byte
    // The sensitiveTypeIds property
    sensitiveTypeIds []string
}
// NewFileClassificationRequest instantiates a new FileClassificationRequest and sets the default values.
func NewFileClassificationRequest()(*FileClassificationRequest) {
    m := &FileClassificationRequest{
        Entity: *NewEntity(),
    }
    return m
}
// CreateFileClassificationRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFileClassificationRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileClassificationRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FileClassificationRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["file"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFile(val)
        }
        return nil
    }
    res["sensitiveTypeIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSensitiveTypeIds(res)
        }
        return nil
    }
    return res
}
// GetFile gets the file property value. The file property
func (m *FileClassificationRequest) GetFile()([]byte) {
    return m.file
}
// GetSensitiveTypeIds gets the sensitiveTypeIds property value. The sensitiveTypeIds property
func (m *FileClassificationRequest) GetSensitiveTypeIds()([]string) {
    return m.sensitiveTypeIds
}
// Serialize serializes information the current object
func (m *FileClassificationRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("file", m.GetFile())
        if err != nil {
            return err
        }
    }
    if m.GetSensitiveTypeIds() != nil {
        err = writer.WriteCollectionOfStringValues("sensitiveTypeIds", m.GetSensitiveTypeIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFile sets the file property value. The file property
func (m *FileClassificationRequest) SetFile(value []byte)() {
    m.file = value
}
// SetSensitiveTypeIds sets the sensitiveTypeIds property value. The sensitiveTypeIds property
func (m *FileClassificationRequest) SetSensitiveTypeIds(value []string)() {
    m.sensitiveTypeIds = value
}
