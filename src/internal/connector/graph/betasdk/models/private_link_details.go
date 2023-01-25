package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivateLinkDetails 
type PrivateLinkDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // The unique identifier for the Private Link policy.
    policyId *string
    // The name of the Private Link policy in Azure AD.
    policyName *string
    // The tenant identifier of the Azure AD tenant the Private Link policy belongs to.
    policyTenantId *string
    // The Azure Resource Manager (ARM) path for the Private Link policy resource.
    resourceId *string
}
// NewPrivateLinkDetails instantiates a new privateLinkDetails and sets the default values.
func NewPrivateLinkDetails()(*PrivateLinkDetails) {
    m := &PrivateLinkDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePrivateLinkDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivateLinkDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivateLinkDetails(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrivateLinkDetails) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivateLinkDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["policyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyId(val)
        }
        return nil
    }
    res["policyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyName(val)
        }
        return nil
    }
    res["policyTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyTenantId(val)
        }
        return nil
    }
    res["resourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PrivateLinkDetails) GetOdataType()(*string) {
    return m.odataType
}
// GetPolicyId gets the policyId property value. The unique identifier for the Private Link policy.
func (m *PrivateLinkDetails) GetPolicyId()(*string) {
    return m.policyId
}
// GetPolicyName gets the policyName property value. The name of the Private Link policy in Azure AD.
func (m *PrivateLinkDetails) GetPolicyName()(*string) {
    return m.policyName
}
// GetPolicyTenantId gets the policyTenantId property value. The tenant identifier of the Azure AD tenant the Private Link policy belongs to.
func (m *PrivateLinkDetails) GetPolicyTenantId()(*string) {
    return m.policyTenantId
}
// GetResourceId gets the resourceId property value. The Azure Resource Manager (ARM) path for the Private Link policy resource.
func (m *PrivateLinkDetails) GetResourceId()(*string) {
    return m.resourceId
}
// Serialize serializes information the current object
func (m *PrivateLinkDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("policyId", m.GetPolicyId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("policyName", m.GetPolicyName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("policyTenantId", m.GetPolicyTenantId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resourceId", m.GetResourceId())
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
func (m *PrivateLinkDetails) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PrivateLinkDetails) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPolicyId sets the policyId property value. The unique identifier for the Private Link policy.
func (m *PrivateLinkDetails) SetPolicyId(value *string)() {
    m.policyId = value
}
// SetPolicyName sets the policyName property value. The name of the Private Link policy in Azure AD.
func (m *PrivateLinkDetails) SetPolicyName(value *string)() {
    m.policyName = value
}
// SetPolicyTenantId sets the policyTenantId property value. The tenant identifier of the Azure AD tenant the Private Link policy belongs to.
func (m *PrivateLinkDetails) SetPolicyTenantId(value *string)() {
    m.policyTenantId = value
}
// SetResourceId sets the resourceId property value. The Azure Resource Manager (ARM) path for the Private Link policy resource.
func (m *PrivateLinkDetails) SetResourceId(value *string)() {
    m.resourceId = value
}
