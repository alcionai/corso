package tenantadmin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdleSessionSignOut 
type IdleSessionSignOut struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Indicates whether the idle session sign-out policy is enabled.
    isEnabled *bool
    // The OdataType property
    odataType *string
    // Number of seconds of inactivity after which a user is signed out.
    signOutAfterInSeconds *int64
    // Number of seconds of inactivity after which a user is notified that they'll be signed out.
    warnAfterInSeconds *int64
}
// NewIdleSessionSignOut instantiates a new idleSessionSignOut and sets the default values.
func NewIdleSessionSignOut()(*IdleSessionSignOut) {
    m := &IdleSessionSignOut{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIdleSessionSignOutFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdleSessionSignOutFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdleSessionSignOut(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdleSessionSignOut) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdleSessionSignOut) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
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
    res["signOutAfterInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignOutAfterInSeconds(val)
        }
        return nil
    }
    res["warnAfterInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWarnAfterInSeconds(val)
        }
        return nil
    }
    return res
}
// GetIsEnabled gets the isEnabled property value. Indicates whether the idle session sign-out policy is enabled.
func (m *IdleSessionSignOut) GetIsEnabled()(*bool) {
    return m.isEnabled
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *IdleSessionSignOut) GetOdataType()(*string) {
    return m.odataType
}
// GetSignOutAfterInSeconds gets the signOutAfterInSeconds property value. Number of seconds of inactivity after which a user is signed out.
func (m *IdleSessionSignOut) GetSignOutAfterInSeconds()(*int64) {
    return m.signOutAfterInSeconds
}
// GetWarnAfterInSeconds gets the warnAfterInSeconds property value. Number of seconds of inactivity after which a user is notified that they'll be signed out.
func (m *IdleSessionSignOut) GetWarnAfterInSeconds()(*int64) {
    return m.warnAfterInSeconds
}
// Serialize serializes information the current object
func (m *IdleSessionSignOut) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
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
        err := writer.WriteInt64Value("signOutAfterInSeconds", m.GetSignOutAfterInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("warnAfterInSeconds", m.GetWarnAfterInSeconds())
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
func (m *IdleSessionSignOut) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetIsEnabled sets the isEnabled property value. Indicates whether the idle session sign-out policy is enabled.
func (m *IdleSessionSignOut) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *IdleSessionSignOut) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSignOutAfterInSeconds sets the signOutAfterInSeconds property value. Number of seconds of inactivity after which a user is signed out.
func (m *IdleSessionSignOut) SetSignOutAfterInSeconds(value *int64)() {
    m.signOutAfterInSeconds = value
}
// SetWarnAfterInSeconds sets the warnAfterInSeconds property value. Number of seconds of inactivity after which a user is notified that they'll be signed out.
func (m *IdleSessionSignOut) SetWarnAfterInSeconds(value *int64)() {
    m.warnAfterInSeconds = value
}
