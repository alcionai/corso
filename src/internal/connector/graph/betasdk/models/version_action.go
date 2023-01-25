package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VersionAction 
type VersionAction struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The name of the new version that was created by this action.
    newVersion *string
    // The OdataType property
    odataType *string
}
// NewVersionAction instantiates a new versionAction and sets the default values.
func NewVersionAction()(*VersionAction) {
    m := &VersionAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateVersionActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVersionActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVersionAction(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VersionAction) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VersionAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["newVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewVersion(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetNewVersion gets the newVersion property value. The name of the new version that was created by this action.
func (m *VersionAction) GetNewVersion()(*string) {
    return m.newVersion
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *VersionAction) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *VersionAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("newVersion", m.GetNewVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VersionAction) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetNewVersion sets the newVersion property value. The name of the new version that was created by this action.
func (m *VersionAction) SetNewVersion(value *string)() {
    m.newVersion = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *VersionAction) SetOdataType(value *string)() {
    m.odataType = value
}
