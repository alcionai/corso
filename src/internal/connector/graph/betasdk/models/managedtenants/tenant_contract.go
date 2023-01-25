package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TenantContract 
type TenantContract struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The type of relationship that exists between the managing entity and tenant. Optional. Read-only.
    contractType *int32
    // The default domain name for the tenant. Required. Read-only.
    defaultDomainName *string
    // The display name for the tenant. Optional. Read-only.
    displayName *string
    // The OdataType property
    odataType *string
}
// NewTenantContract instantiates a new tenantContract and sets the default values.
func NewTenantContract()(*TenantContract) {
    m := &TenantContract{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTenantContractFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantContractFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenantContract(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TenantContract) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetContractType gets the contractType property value. The type of relationship that exists between the managing entity and tenant. Optional. Read-only.
func (m *TenantContract) GetContractType()(*int32) {
    return m.contractType
}
// GetDefaultDomainName gets the defaultDomainName property value. The default domain name for the tenant. Required. Read-only.
func (m *TenantContract) GetDefaultDomainName()(*string) {
    return m.defaultDomainName
}
// GetDisplayName gets the displayName property value. The display name for the tenant. Optional. Read-only.
func (m *TenantContract) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantContract) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contractType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContractType(val)
        }
        return nil
    }
    res["defaultDomainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultDomainName(val)
        }
        return nil
    }
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
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TenantContract) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *TenantContract) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("contractType", m.GetContractType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("defaultDomainName", m.GetDefaultDomainName())
        if err != nil {
            return err
        }
    }
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TenantContract) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetContractType sets the contractType property value. The type of relationship that exists between the managing entity and tenant. Optional. Read-only.
func (m *TenantContract) SetContractType(value *int32)() {
    m.contractType = value
}
// SetDefaultDomainName sets the defaultDomainName property value. The default domain name for the tenant. Required. Read-only.
func (m *TenantContract) SetDefaultDomainName(value *string)() {
    m.defaultDomainName = value
}
// SetDisplayName sets the displayName property value. The display name for the tenant. Optional. Read-only.
func (m *TenantContract) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TenantContract) SetOdataType(value *string)() {
    m.odataType = value
}
