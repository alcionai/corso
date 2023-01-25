package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationConditions 
type AuthenticationConditions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The applications property
    applications AuthenticationConditionsApplicationsable
    // The OdataType property
    odataType *string
}
// NewAuthenticationConditions instantiates a new authenticationConditions and sets the default values.
func NewAuthenticationConditions()(*AuthenticationConditions) {
    m := &AuthenticationConditions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAuthenticationConditionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationConditionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationConditions(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationConditions) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetApplications gets the applications property value. The applications property
func (m *AuthenticationConditions) GetApplications()(AuthenticationConditionsApplicationsable) {
    return m.applications
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationConditions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["applications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationConditionsApplicationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplications(val.(AuthenticationConditionsApplicationsable))
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
func (m *AuthenticationConditions) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AuthenticationConditions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("applications", m.GetApplications())
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
func (m *AuthenticationConditions) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetApplications sets the applications property value. The applications property
func (m *AuthenticationConditions) SetApplications(value AuthenticationConditionsApplicationsable)() {
    m.applications = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AuthenticationConditions) SetOdataType(value *string)() {
    m.odataType = value
}
