package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Setting 
type Setting struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The display name for the setting. Required. Read-only.
    displayName *string
    // The value for the setting serialized as string of JSON. Required. Read-only.
    jsonValue *string
    // The OdataType property
    odataType *string
    // A flag indicating whether the setting can be override existing configurations when applied. Required. Read-only.
    overwriteAllowed *bool
    // The settingId property
    settingId *string
    // The valueType property
    valueType *ManagementParameterValueType
}
// NewSetting instantiates a new setting and sets the default values.
func NewSetting()(*Setting) {
    m := &Setting{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSetting(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Setting) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. The display name for the setting. Required. Read-only.
func (m *Setting) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Setting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["jsonValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJsonValue(val)
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
    res["overwriteAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOverwriteAllowed(val)
        }
        return nil
    }
    res["settingId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingId(val)
        }
        return nil
    }
    res["valueType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementParameterValueType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueType(val.(*ManagementParameterValueType))
        }
        return nil
    }
    return res
}
// GetJsonValue gets the jsonValue property value. The value for the setting serialized as string of JSON. Required. Read-only.
func (m *Setting) GetJsonValue()(*string) {
    return m.jsonValue
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Setting) GetOdataType()(*string) {
    return m.odataType
}
// GetOverwriteAllowed gets the overwriteAllowed property value. A flag indicating whether the setting can be override existing configurations when applied. Required. Read-only.
func (m *Setting) GetOverwriteAllowed()(*bool) {
    return m.overwriteAllowed
}
// GetSettingId gets the settingId property value. The settingId property
func (m *Setting) GetSettingId()(*string) {
    return m.settingId
}
// GetValueType gets the valueType property value. The valueType property
func (m *Setting) GetValueType()(*ManagementParameterValueType) {
    return m.valueType
}
// Serialize serializes information the current object
func (m *Setting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("jsonValue", m.GetJsonValue())
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
        err := writer.WriteBoolValue("overwriteAllowed", m.GetOverwriteAllowed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("settingId", m.GetSettingId())
        if err != nil {
            return err
        }
    }
    if m.GetValueType() != nil {
        cast := (*m.GetValueType()).String()
        err := writer.WriteStringValue("valueType", &cast)
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
func (m *Setting) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. The display name for the setting. Required. Read-only.
func (m *Setting) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetJsonValue sets the jsonValue property value. The value for the setting serialized as string of JSON. Required. Read-only.
func (m *Setting) SetJsonValue(value *string)() {
    m.jsonValue = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Setting) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOverwriteAllowed sets the overwriteAllowed property value. A flag indicating whether the setting can be override existing configurations when applied. Required. Read-only.
func (m *Setting) SetOverwriteAllowed(value *bool)() {
    m.overwriteAllowed = value
}
// SetSettingId sets the settingId property value. The settingId property
func (m *Setting) SetSettingId(value *string)() {
    m.settingId = value
}
// SetValueType sets the valueType property value. The valueType property
func (m *Setting) SetValueType(value *ManagementParameterValueType)() {
    m.valueType = value
}
