package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationConditionApplication 
type AuthenticationConditionApplication struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The appId property
    appId *string
    // The OdataType property
    odataType *string
}
// NewAuthenticationConditionApplication instantiates a new authenticationConditionApplication and sets the default values.
func NewAuthenticationConditionApplication()(*AuthenticationConditionApplication) {
    m := &AuthenticationConditionApplication{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAuthenticationConditionApplicationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationConditionApplicationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationConditionApplication(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationConditionApplication) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAppId gets the appId property value. The appId property
func (m *AuthenticationConditionApplication) GetAppId()(*string) {
    return m.appId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationConditionApplication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AuthenticationConditionApplication) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AuthenticationConditionApplication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appId", m.GetAppId())
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
func (m *AuthenticationConditionApplication) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAppId sets the appId property value. The appId property
func (m *AuthenticationConditionApplication) SetAppId(value *string)() {
    m.appId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AuthenticationConditionApplication) SetOdataType(value *string)() {
    m.odataType = value
}
