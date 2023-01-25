package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OmaSettingInteger 
type OmaSettingInteger struct {
    OmaSetting
    // By setting to true, the CSP (configuration service provider) specified in the OMA-URI will perform a get, instead of set
    isReadOnly *bool
    // Value.
    value *int32
}
// NewOmaSettingInteger instantiates a new OmaSettingInteger and sets the default values.
func NewOmaSettingInteger()(*OmaSettingInteger) {
    m := &OmaSettingInteger{
        OmaSetting: *NewOmaSetting(),
    }
    odataTypeValue := "#microsoft.graph.omaSettingInteger";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOmaSettingIntegerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOmaSettingIntegerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOmaSettingInteger(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OmaSettingInteger) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OmaSetting.GetFieldDeserializers()
    res["isReadOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsReadOnly(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetIsReadOnly gets the isReadOnly property value. By setting to true, the CSP (configuration service provider) specified in the OMA-URI will perform a get, instead of set
func (m *OmaSettingInteger) GetIsReadOnly()(*bool) {
    return m.isReadOnly
}
// GetValue gets the value property value. Value.
func (m *OmaSettingInteger) GetValue()(*int32) {
    return m.value
}
// Serialize serializes information the current object
func (m *OmaSettingInteger) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OmaSetting.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isReadOnly", m.GetIsReadOnly())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsReadOnly sets the isReadOnly property value. By setting to true, the CSP (configuration service provider) specified in the OMA-URI will perform a get, instead of set
func (m *OmaSettingInteger) SetIsReadOnly(value *bool)() {
    m.isReadOnly = value
}
// SetValue sets the value property value. Value.
func (m *OmaSettingInteger) SetValue(value *int32)() {
    m.value = value
}
