package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationAppDeviceDetails 
type AuthenticationAppDeviceDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The version of the client authentication app used during the authentication step.
    appVersion *string
    // The name of the client authentication app used during the authentication step.
    clientApp *string
    // ID of the device used during the authentication step.
    deviceId *string
    // The OdataType property
    odataType *string
    // The operating system running on the device used for the authentication step.
    operatingSystem *string
}
// NewAuthenticationAppDeviceDetails instantiates a new authenticationAppDeviceDetails and sets the default values.
func NewAuthenticationAppDeviceDetails()(*AuthenticationAppDeviceDetails) {
    m := &AuthenticationAppDeviceDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAuthenticationAppDeviceDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationAppDeviceDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationAppDeviceDetails(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationAppDeviceDetails) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAppVersion gets the appVersion property value. The version of the client authentication app used during the authentication step.
func (m *AuthenticationAppDeviceDetails) GetAppVersion()(*string) {
    return m.appVersion
}
// GetClientApp gets the clientApp property value. The name of the client authentication app used during the authentication step.
func (m *AuthenticationAppDeviceDetails) GetClientApp()(*string) {
    return m.clientApp
}
// GetDeviceId gets the deviceId property value. ID of the device used during the authentication step.
func (m *AuthenticationAppDeviceDetails) GetDeviceId()(*string) {
    return m.deviceId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationAppDeviceDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppVersion(val)
        }
        return nil
    }
    res["clientApp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientApp(val)
        }
        return nil
    }
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
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
    res["operatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystem(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AuthenticationAppDeviceDetails) GetOdataType()(*string) {
    return m.odataType
}
// GetOperatingSystem gets the operatingSystem property value. The operating system running on the device used for the authentication step.
func (m *AuthenticationAppDeviceDetails) GetOperatingSystem()(*string) {
    return m.operatingSystem
}
// Serialize serializes information the current object
func (m *AuthenticationAppDeviceDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appVersion", m.GetAppVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("clientApp", m.GetClientApp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceId", m.GetDeviceId())
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
        err := writer.WriteStringValue("operatingSystem", m.GetOperatingSystem())
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
func (m *AuthenticationAppDeviceDetails) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAppVersion sets the appVersion property value. The version of the client authentication app used during the authentication step.
func (m *AuthenticationAppDeviceDetails) SetAppVersion(value *string)() {
    m.appVersion = value
}
// SetClientApp sets the clientApp property value. The name of the client authentication app used during the authentication step.
func (m *AuthenticationAppDeviceDetails) SetClientApp(value *string)() {
    m.clientApp = value
}
// SetDeviceId sets the deviceId property value. ID of the device used during the authentication step.
func (m *AuthenticationAppDeviceDetails) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AuthenticationAppDeviceDetails) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOperatingSystem sets the operatingSystem property value. The operating system running on the device used for the authentication step.
func (m *AuthenticationAppDeviceDetails) SetOperatingSystem(value *string)() {
    m.operatingSystem = value
}
