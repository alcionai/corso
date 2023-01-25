package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationSourceFilter 
type AuthenticationSourceFilter struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Applications to include for evaluation of the authenticationListener. These applications trigger the associated action when used as the client application in the authentication flow. The application identifer is the application's client id.
    includeApplications []string
    // The OdataType property
    odataType *string
}
// NewAuthenticationSourceFilter instantiates a new authenticationSourceFilter and sets the default values.
func NewAuthenticationSourceFilter()(*AuthenticationSourceFilter) {
    m := &AuthenticationSourceFilter{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAuthenticationSourceFilterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationSourceFilterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationSourceFilter(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationSourceFilter) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationSourceFilter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["includeApplications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
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
// GetIncludeApplications gets the includeApplications property value. Applications to include for evaluation of the authenticationListener. These applications trigger the associated action when used as the client application in the authentication flow. The application identifer is the application's client id.
func (m *AuthenticationSourceFilter) GetIncludeApplications()([]string) {
    return m.includeApplications
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AuthenticationSourceFilter) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AuthenticationSourceFilter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetIncludeApplications() != nil {
        err := writer.WriteCollectionOfStringValues("includeApplications", m.GetIncludeApplications())
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
func (m *AuthenticationSourceFilter) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetIncludeApplications sets the includeApplications property value. Applications to include for evaluation of the authenticationListener. These applications trigger the associated action when used as the client application in the authentication flow. The application identifer is the application's client id.
func (m *AuthenticationSourceFilter) SetIncludeApplications(value []string)() {
    m.includeApplications = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AuthenticationSourceFilter) SetOdataType(value *string)() {
    m.odataType = value
}
