package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AzureADRegistrationPolicy 
type AzureADRegistrationPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The identifiers of the groups that are in the scope of the policy. Either this property or allowedUsers is required when the appliesTo property is set to selected.
    allowedGroups []string
    // The identifiers of users that are in the scope of the policy. Either this property or allowedGroups is required when the appliesTo property is set to selected.
    allowedUsers []string
    // Specifies whether to block or allow fine-grained control of the policy scope. The possible values are: 0 (meaning none), 1 (meaning all), 2 (meaning selected), 3 (meaning unknownFutureValue). The default value is 1. When set to 2, at least one user or group identifier must be specified in either allowedUsers or allowedGroups.  Setting this property to 0 or 1 removes all identifiers in both allowedUsers and allowedGroups.
    appliesTo *PolicyScope
    // Specifies whether this policy scope is configurable by the admin. The default value is false. When an admin has enabled Intune (MEM) to manage devices, this property is set to false and appliesTo defaults to 1 (meaning all).
    isAdminConfigurable *bool
    // The OdataType property
    odataType *string
}
// NewAzureADRegistrationPolicy instantiates a new azureADRegistrationPolicy and sets the default values.
func NewAzureADRegistrationPolicy()(*AzureADRegistrationPolicy) {
    m := &AzureADRegistrationPolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAzureADRegistrationPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAzureADRegistrationPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureADRegistrationPolicy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AzureADRegistrationPolicy) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAllowedGroups gets the allowedGroups property value. The identifiers of the groups that are in the scope of the policy. Either this property or allowedUsers is required when the appliesTo property is set to selected.
func (m *AzureADRegistrationPolicy) GetAllowedGroups()([]string) {
    return m.allowedGroups
}
// GetAllowedUsers gets the allowedUsers property value. The identifiers of users that are in the scope of the policy. Either this property or allowedGroups is required when the appliesTo property is set to selected.
func (m *AzureADRegistrationPolicy) GetAllowedUsers()([]string) {
    return m.allowedUsers
}
// GetAppliesTo gets the appliesTo property value. Specifies whether to block or allow fine-grained control of the policy scope. The possible values are: 0 (meaning none), 1 (meaning all), 2 (meaning selected), 3 (meaning unknownFutureValue). The default value is 1. When set to 2, at least one user or group identifier must be specified in either allowedUsers or allowedGroups.  Setting this property to 0 or 1 removes all identifiers in both allowedUsers and allowedGroups.
func (m *AzureADRegistrationPolicy) GetAppliesTo()(*PolicyScope) {
    return m.appliesTo
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AzureADRegistrationPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowedGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAllowedGroups(res)
        }
        return nil
    }
    res["allowedUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAllowedUsers(res)
        }
        return nil
    }
    res["appliesTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePolicyScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppliesTo(val.(*PolicyScope))
        }
        return nil
    }
    res["isAdminConfigurable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAdminConfigurable(val)
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
// GetIsAdminConfigurable gets the isAdminConfigurable property value. Specifies whether this policy scope is configurable by the admin. The default value is false. When an admin has enabled Intune (MEM) to manage devices, this property is set to false and appliesTo defaults to 1 (meaning all).
func (m *AzureADRegistrationPolicy) GetIsAdminConfigurable()(*bool) {
    return m.isAdminConfigurable
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AzureADRegistrationPolicy) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AzureADRegistrationPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAllowedGroups() != nil {
        err := writer.WriteCollectionOfStringValues("allowedGroups", m.GetAllowedGroups())
        if err != nil {
            return err
        }
    }
    if m.GetAllowedUsers() != nil {
        err := writer.WriteCollectionOfStringValues("allowedUsers", m.GetAllowedUsers())
        if err != nil {
            return err
        }
    }
    if m.GetAppliesTo() != nil {
        cast := (*m.GetAppliesTo()).String()
        err := writer.WriteStringValue("appliesTo", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isAdminConfigurable", m.GetIsAdminConfigurable())
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
func (m *AzureADRegistrationPolicy) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAllowedGroups sets the allowedGroups property value. The identifiers of the groups that are in the scope of the policy. Either this property or allowedUsers is required when the appliesTo property is set to selected.
func (m *AzureADRegistrationPolicy) SetAllowedGroups(value []string)() {
    m.allowedGroups = value
}
// SetAllowedUsers sets the allowedUsers property value. The identifiers of users that are in the scope of the policy. Either this property or allowedGroups is required when the appliesTo property is set to selected.
func (m *AzureADRegistrationPolicy) SetAllowedUsers(value []string)() {
    m.allowedUsers = value
}
// SetAppliesTo sets the appliesTo property value. Specifies whether to block or allow fine-grained control of the policy scope. The possible values are: 0 (meaning none), 1 (meaning all), 2 (meaning selected), 3 (meaning unknownFutureValue). The default value is 1. When set to 2, at least one user or group identifier must be specified in either allowedUsers or allowedGroups.  Setting this property to 0 or 1 removes all identifiers in both allowedUsers and allowedGroups.
func (m *AzureADRegistrationPolicy) SetAppliesTo(value *PolicyScope)() {
    m.appliesTo = value
}
// SetIsAdminConfigurable sets the isAdminConfigurable property value. Specifies whether this policy scope is configurable by the admin. The default value is false. When an admin has enabled Intune (MEM) to manage devices, this property is set to false and appliesTo defaults to 1 (meaning all).
func (m *AzureADRegistrationPolicy) SetIsAdminConfigurable(value *bool)() {
    m.isAdminConfigurable = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AzureADRegistrationPolicy) SetOdataType(value *string)() {
    m.odataType = value
}
