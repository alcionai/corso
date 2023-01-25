package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationConditionsApplications 
type AuthenticationConditionsApplications struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The includeAllApplications property
    includeAllApplications *bool
    // The includeApplications property
    includeApplications []AuthenticationConditionApplicationable
    // The OdataType property
    odataType *string
}
// NewAuthenticationConditionsApplications instantiates a new authenticationConditionsApplications and sets the default values.
func NewAuthenticationConditionsApplications()(*AuthenticationConditionsApplications) {
    m := &AuthenticationConditionsApplications{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAuthenticationConditionsApplicationsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationConditionsApplicationsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationConditionsApplications(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationConditionsApplications) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationConditionsApplications) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["includeAllApplications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncludeAllApplications(val)
        }
        return nil
    }
    res["includeApplications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationConditionApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationConditionApplicationable, len(val))
            for i, v := range val {
                res[i] = v.(AuthenticationConditionApplicationable)
            }
            m.SetIncludeApplications(res)
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
// GetIncludeAllApplications gets the includeAllApplications property value. The includeAllApplications property
func (m *AuthenticationConditionsApplications) GetIncludeAllApplications()(*bool) {
    return m.includeAllApplications
}
// GetIncludeApplications gets the includeApplications property value. The includeApplications property
func (m *AuthenticationConditionsApplications) GetIncludeApplications()([]AuthenticationConditionApplicationable) {
    return m.includeApplications
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AuthenticationConditionsApplications) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AuthenticationConditionsApplications) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("includeAllApplications", m.GetIncludeAllApplications())
        if err != nil {
            return err
        }
    }
    if m.GetIncludeApplications() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIncludeApplications()))
        for i, v := range m.GetIncludeApplications() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("includeApplications", cast)
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
func (m *AuthenticationConditionsApplications) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetIncludeAllApplications sets the includeAllApplications property value. The includeAllApplications property
func (m *AuthenticationConditionsApplications) SetIncludeAllApplications(value *bool)() {
    m.includeAllApplications = value
}
// SetIncludeApplications sets the includeApplications property value. The includeApplications property
func (m *AuthenticationConditionsApplications) SetIncludeApplications(value []AuthenticationConditionApplicationable)() {
    m.includeApplications = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AuthenticationConditionsApplications) SetOdataType(value *string)() {
    m.odataType = value
}
