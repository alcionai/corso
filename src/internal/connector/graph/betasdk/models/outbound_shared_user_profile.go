package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OutboundSharedUserProfile 
type OutboundSharedUserProfile struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // The tenants property
    tenants []TenantReferenceable
    // The userId property
    userId *string
}
// NewOutboundSharedUserProfile instantiates a new outboundSharedUserProfile and sets the default values.
func NewOutboundSharedUserProfile()(*OutboundSharedUserProfile) {
    m := &OutboundSharedUserProfile{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOutboundSharedUserProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOutboundSharedUserProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOutboundSharedUserProfile(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OutboundSharedUserProfile) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OutboundSharedUserProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["tenants"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTenantReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TenantReferenceable, len(val))
            for i, v := range val {
                res[i] = v.(TenantReferenceable)
            }
            m.SetTenants(res)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OutboundSharedUserProfile) GetOdataType()(*string) {
    return m.odataType
}
// GetTenants gets the tenants property value. The tenants property
func (m *OutboundSharedUserProfile) GetTenants()([]TenantReferenceable) {
    return m.tenants
}
// GetUserId gets the userId property value. The userId property
func (m *OutboundSharedUserProfile) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *OutboundSharedUserProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetTenants() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTenants()))
        for i, v := range m.GetTenants() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("tenants", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userId", m.GetUserId())
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
func (m *OutboundSharedUserProfile) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OutboundSharedUserProfile) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTenants sets the tenants property value. The tenants property
func (m *OutboundSharedUserProfile) SetTenants(value []TenantReferenceable)() {
    m.tenants = value
}
// SetUserId sets the userId property value. The userId property
func (m *OutboundSharedUserProfile) SetUserId(value *string)() {
    m.userId = value
}
