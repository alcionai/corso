package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProviderTenantSetting provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ProviderTenantSetting struct {
    Entity
    // The azureTenantId property
    azureTenantId *string
    // The enabled property
    enabled *bool
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The provider property
    provider *string
    // The vendor property
    vendor_escaped *string
}
// NewProviderTenantSetting instantiates a new providerTenantSetting and sets the default values.
func NewProviderTenantSetting()(*ProviderTenantSetting) {
    m := &ProviderTenantSetting{
        Entity: *NewEntity(),
    }
    return m
}
// CreateProviderTenantSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProviderTenantSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProviderTenantSetting(), nil
}
// GetAzureTenantId gets the azureTenantId property value. The azureTenantId property
func (m *ProviderTenantSetting) GetAzureTenantId()(*string) {
    return m.azureTenantId
}
// GetEnabled gets the enabled property value. The enabled property
func (m *ProviderTenantSetting) GetEnabled()(*bool) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProviderTenantSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["azureTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureTenantId(val)
        }
        return nil
    }
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["provider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvider(val)
        }
        return nil
    }
    res["vendor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendor(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *ProviderTenantSetting) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetProvider gets the provider property value. The provider property
func (m *ProviderTenantSetting) GetProvider()(*string) {
    return m.provider
}
// GetVendor gets the vendor property value. The vendor property
func (m *ProviderTenantSetting) GetVendor()(*string) {
    return m.vendor_escaped
}
// Serialize serializes information the current object
func (m *ProviderTenantSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("azureTenantId", m.GetAzureTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("provider", m.GetProvider())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("vendor", m.GetVendor())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAzureTenantId sets the azureTenantId property value. The azureTenantId property
func (m *ProviderTenantSetting) SetAzureTenantId(value *string)() {
    m.azureTenantId = value
}
// SetEnabled sets the enabled property value. The enabled property
func (m *ProviderTenantSetting) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *ProviderTenantSetting) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetProvider sets the provider property value. The provider property
func (m *ProviderTenantSetting) SetProvider(value *string)() {
    m.provider = value
}
// SetVendor sets the vendor property value. The vendor property
func (m *ProviderTenantSetting) SetVendor(value *string)() {
    m.vendor_escaped = value
}
