package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationContext 
type AuthenticationContext struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Describes how the conditional access authentication context was triggered. A value of previouslySatisfied means the auth context was because the user already satisfied the requirements for that authentication context in some previous authentication event. A value of required means the user had to meet the authentication context requirement as part of the sign-in flow. The possible values are: required, previouslySatisfied, notApplicable, unknownFutureValue.
    detail *AuthenticationContextDetail
    // The identifier of a authentication context in your tenant.
    id *string
    // The OdataType property
    odataType *string
}
// NewAuthenticationContext instantiates a new authenticationContext and sets the default values.
func NewAuthenticationContext()(*AuthenticationContext) {
    m := &AuthenticationContext{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAuthenticationContextFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationContextFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationContext(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationContext) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDetail gets the detail property value. Describes how the conditional access authentication context was triggered. A value of previouslySatisfied means the auth context was because the user already satisfied the requirements for that authentication context in some previous authentication event. A value of required means the user had to meet the authentication context requirement as part of the sign-in flow. The possible values are: required, previouslySatisfied, notApplicable, unknownFutureValue.
func (m *AuthenticationContext) GetDetail()(*AuthenticationContextDetail) {
    return m.detail
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationContext) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["detail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationContextDetail)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetail(val.(*AuthenticationContextDetail))
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
// GetId gets the id property value. The identifier of a authentication context in your tenant.
func (m *AuthenticationContext) GetId()(*string) {
    return m.id
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AuthenticationContext) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AuthenticationContext) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDetail() != nil {
        cast := (*m.GetDetail()).String()
        err := writer.WriteStringValue("detail", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
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
func (m *AuthenticationContext) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDetail sets the detail property value. Describes how the conditional access authentication context was triggered. A value of previouslySatisfied means the auth context was because the user already satisfied the requirements for that authentication context in some previous authentication event. A value of required means the user had to meet the authentication context requirement as part of the sign-in flow. The possible values are: required, previouslySatisfied, notApplicable, unknownFutureValue.
func (m *AuthenticationContext) SetDetail(value *AuthenticationContextDetail)() {
    m.detail = value
}
// SetId sets the id property value. The identifier of a authentication context in your tenant.
func (m *AuthenticationContext) SetId(value *string)() {
    m.id = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AuthenticationContext) SetOdataType(value *string)() {
    m.odataType = value
}
