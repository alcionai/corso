package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestorSettings 
type RequestorSettings struct {
    // Indicates whether new requests are accepted on this policy.
    acceptRequests *bool
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The users who are allowed to request on this policy, which can be singleUser, groupMembers, and connectedOrganizationMembers.
    allowedRequestors []UserSetable
    // The OdataType property
    odataType *string
    // Who can request. One of NoSubjects, SpecificDirectorySubjects, SpecificConnectedOrganizationSubjects, AllConfiguredConnectedOrganizationSubjects, AllExistingConnectedOrganizationSubjects, AllExistingDirectoryMemberUsers, AllExistingDirectorySubjects or AllExternalSubjects.
    scopeType *string
}
// NewRequestorSettings instantiates a new requestorSettings and sets the default values.
func NewRequestorSettings()(*RequestorSettings) {
    m := &RequestorSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRequestorSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequestorSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestorSettings(), nil
}
// GetAcceptRequests gets the acceptRequests property value. Indicates whether new requests are accepted on this policy.
func (m *RequestorSettings) GetAcceptRequests()(*bool) {
    return m.acceptRequests
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RequestorSettings) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAllowedRequestors gets the allowedRequestors property value. The users who are allowed to request on this policy, which can be singleUser, groupMembers, and connectedOrganizationMembers.
func (m *RequestorSettings) GetAllowedRequestors()([]UserSetable) {
    return m.allowedRequestors
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestorSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["acceptRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcceptRequests(val)
        }
        return nil
    }
    res["allowedRequestors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserSetable, len(val))
            for i, v := range val {
                res[i] = v.(UserSetable)
            }
            m.SetAllowedRequestors(res)
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
    res["scopeType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScopeType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RequestorSettings) GetOdataType()(*string) {
    return m.odataType
}
// GetScopeType gets the scopeType property value. Who can request. One of NoSubjects, SpecificDirectorySubjects, SpecificConnectedOrganizationSubjects, AllConfiguredConnectedOrganizationSubjects, AllExistingConnectedOrganizationSubjects, AllExistingDirectoryMemberUsers, AllExistingDirectorySubjects or AllExternalSubjects.
func (m *RequestorSettings) GetScopeType()(*string) {
    return m.scopeType
}
// Serialize serializes information the current object
func (m *RequestorSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("acceptRequests", m.GetAcceptRequests())
        if err != nil {
            return err
        }
    }
    if m.GetAllowedRequestors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAllowedRequestors()))
        for i, v := range m.GetAllowedRequestors() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("allowedRequestors", cast)
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
        err := writer.WriteStringValue("scopeType", m.GetScopeType())
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
// SetAcceptRequests sets the acceptRequests property value. Indicates whether new requests are accepted on this policy.
func (m *RequestorSettings) SetAcceptRequests(value *bool)() {
    m.acceptRequests = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RequestorSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAllowedRequestors sets the allowedRequestors property value. The users who are allowed to request on this policy, which can be singleUser, groupMembers, and connectedOrganizationMembers.
func (m *RequestorSettings) SetAllowedRequestors(value []UserSetable)() {
    m.allowedRequestors = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RequestorSettings) SetOdataType(value *string)() {
    m.odataType = value
}
// SetScopeType sets the scopeType property value. Who can request. One of NoSubjects, SpecificDirectorySubjects, SpecificConnectedOrganizationSubjects, AllConfiguredConnectedOrganizationSubjects, AllExistingConnectedOrganizationSubjects, AllExistingDirectoryMemberUsers, AllExistingDirectorySubjects or AllExternalSubjects.
func (m *RequestorSettings) SetScopeType(value *string)() {
    m.scopeType = value
}
