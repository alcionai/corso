package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// InformationProtectionPolicySetting 
type InformationProtectionPolicySetting struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // The defaultLabelId property
    defaultLabelId *string
    // Exposes whether justification input is required on label downgrade.
    isDowngradeJustificationRequired *bool
    // Exposes whether mandatory labeling is enabled.
    isMandatory *bool
    // Exposes the more information URL that can be configured by the administrator.
    moreInfoUrl *string
}
// NewInformationProtectionPolicySetting instantiates a new informationProtectionPolicySetting and sets the default values.
func NewInformationProtectionPolicySetting()(*InformationProtectionPolicySetting) {
    m := &InformationProtectionPolicySetting{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateInformationProtectionPolicySettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInformationProtectionPolicySettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInformationProtectionPolicySetting(), nil
}
// GetDefaultLabelId gets the defaultLabelId property value. The defaultLabelId property
func (m *InformationProtectionPolicySetting) GetDefaultLabelId()(*string) {
    return m.defaultLabelId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InformationProtectionPolicySetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["defaultLabelId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultLabelId(val)
        }
        return nil
    }
    res["isDowngradeJustificationRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDowngradeJustificationRequired(val)
        }
        return nil
    }
    res["isMandatory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMandatory(val)
        }
        return nil
    }
    res["moreInfoUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMoreInfoUrl(val)
        }
        return nil
    }
    return res
}
// GetIsDowngradeJustificationRequired gets the isDowngradeJustificationRequired property value. Exposes whether justification input is required on label downgrade.
func (m *InformationProtectionPolicySetting) GetIsDowngradeJustificationRequired()(*bool) {
    return m.isDowngradeJustificationRequired
}
// GetIsMandatory gets the isMandatory property value. Exposes whether mandatory labeling is enabled.
func (m *InformationProtectionPolicySetting) GetIsMandatory()(*bool) {
    return m.isMandatory
}
// GetMoreInfoUrl gets the moreInfoUrl property value. Exposes the more information URL that can be configured by the administrator.
func (m *InformationProtectionPolicySetting) GetMoreInfoUrl()(*string) {
    return m.moreInfoUrl
}
// Serialize serializes information the current object
func (m *InformationProtectionPolicySetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("defaultLabelId", m.GetDefaultLabelId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDowngradeJustificationRequired", m.GetIsDowngradeJustificationRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMandatory", m.GetIsMandatory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("moreInfoUrl", m.GetMoreInfoUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefaultLabelId sets the defaultLabelId property value. The defaultLabelId property
func (m *InformationProtectionPolicySetting) SetDefaultLabelId(value *string)() {
    m.defaultLabelId = value
}
// SetIsDowngradeJustificationRequired sets the isDowngradeJustificationRequired property value. Exposes whether justification input is required on label downgrade.
func (m *InformationProtectionPolicySetting) SetIsDowngradeJustificationRequired(value *bool)() {
    m.isDowngradeJustificationRequired = value
}
// SetIsMandatory sets the isMandatory property value. Exposes whether mandatory labeling is enabled.
func (m *InformationProtectionPolicySetting) SetIsMandatory(value *bool)() {
    m.isMandatory = value
}
// SetMoreInfoUrl sets the moreInfoUrl property value. Exposes the more information URL that can be configured by the administrator.
func (m *InformationProtectionPolicySetting) SetMoreInfoUrl(value *string)() {
    m.moreInfoUrl = value
}
