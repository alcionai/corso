package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagementActionInfo 
type ManagementActionInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The identifier for the management action. Required. Read-only.
    managementActionId *string
    // The identifier for the management template. Required. Read-only.
    managementTemplateId *string
    // The managementTemplateVersion property
    managementTemplateVersion *int32
    // The OdataType property
    odataType *string
}
// NewManagementActionInfo instantiates a new managementActionInfo and sets the default values.
func NewManagementActionInfo()(*ManagementActionInfo) {
    m := &ManagementActionInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateManagementActionInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagementActionInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagementActionInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagementActionInfo) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementActionInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["managementActionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementActionId(val)
        }
        return nil
    }
    res["managementTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateId(val)
        }
        return nil
    }
    res["managementTemplateVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateVersion(val)
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
// GetManagementActionId gets the managementActionId property value. The identifier for the management action. Required. Read-only.
func (m *ManagementActionInfo) GetManagementActionId()(*string) {
    return m.managementActionId
}
// GetManagementTemplateId gets the managementTemplateId property value. The identifier for the management template. Required. Read-only.
func (m *ManagementActionInfo) GetManagementTemplateId()(*string) {
    return m.managementTemplateId
}
// GetManagementTemplateVersion gets the managementTemplateVersion property value. The managementTemplateVersion property
func (m *ManagementActionInfo) GetManagementTemplateVersion()(*int32) {
    return m.managementTemplateVersion
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ManagementActionInfo) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ManagementActionInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("managementActionId", m.GetManagementActionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("managementTemplateId", m.GetManagementTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("managementTemplateVersion", m.GetManagementTemplateVersion())
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
func (m *ManagementActionInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetManagementActionId sets the managementActionId property value. The identifier for the management action. Required. Read-only.
func (m *ManagementActionInfo) SetManagementActionId(value *string)() {
    m.managementActionId = value
}
// SetManagementTemplateId sets the managementTemplateId property value. The identifier for the management template. Required. Read-only.
func (m *ManagementActionInfo) SetManagementTemplateId(value *string)() {
    m.managementTemplateId = value
}
// SetManagementTemplateVersion sets the managementTemplateVersion property value. The managementTemplateVersion property
func (m *ManagementActionInfo) SetManagementTemplateVersion(value *int32)() {
    m.managementTemplateVersion = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ManagementActionInfo) SetOdataType(value *string)() {
    m.odataType = value
}
