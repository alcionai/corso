package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CrossTenantIdentitySyncPolicyPartner 
type CrossTenantIdentitySyncPolicyPartner struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Display name for the cross-tenant user synchronization policy. Use the name of the partner Azure AD tenant to easily identify the policy. Optional.
    displayName *string
    // The OdataType property
    odataType *string
    // Tenant identifier for the partner Azure AD organization. Read-only.
    tenantId *string
}
// NewCrossTenantIdentitySyncPolicyPartner instantiates a new crossTenantIdentitySyncPolicyPartner and sets the default values.
func NewCrossTenantIdentitySyncPolicyPartner()(*CrossTenantIdentitySyncPolicyPartner) {
    m := &CrossTenantIdentitySyncPolicyPartner{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCrossTenantIdentitySyncPolicyPartnerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCrossTenantIdentitySyncPolicyPartnerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCrossTenantIdentitySyncPolicyPartner(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CrossTenantIdentitySyncPolicyPartner) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. Display name for the cross-tenant user synchronization policy. Use the name of the partner Azure AD tenant to easily identify the policy. Optional.
func (m *CrossTenantIdentitySyncPolicyPartner) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CrossTenantIdentitySyncPolicyPartner) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CrossTenantIdentitySyncPolicyPartner) GetOdataType()(*string) {
    return m.odataType
}
// GetTenantId gets the tenantId property value. Tenant identifier for the partner Azure AD organization. Read-only.
func (m *CrossTenantIdentitySyncPolicyPartner) GetTenantId()(*string) {
    return m.tenantId
}
// Serialize serializes information the current object
func (m *CrossTenantIdentitySyncPolicyPartner) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
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
        err := writer.WriteStringValue("tenantId", m.GetTenantId())
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
func (m *CrossTenantIdentitySyncPolicyPartner) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. Display name for the cross-tenant user synchronization policy. Use the name of the partner Azure AD tenant to easily identify the policy. Optional.
func (m *CrossTenantIdentitySyncPolicyPartner) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CrossTenantIdentitySyncPolicyPartner) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTenantId sets the tenantId property value. Tenant identifier for the partner Azure AD organization. Read-only.
func (m *CrossTenantIdentitySyncPolicyPartner) SetTenantId(value *string)() {
    m.tenantId = value
}
