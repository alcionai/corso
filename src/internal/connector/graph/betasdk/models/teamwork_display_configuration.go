package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkDisplayConfiguration 
type TeamworkDisplayConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The list of configured displays. Applicable only for Microsoft Teams Rooms devices.
    configuredDisplays []TeamworkConfiguredPeripheralable
    // Total number of connected displays, including the inbuilt display. Applicable only for Teams Rooms devices.
    displayCount *int32
    // Configuration for the inbuilt display. Not applicable for Teams Rooms devices.
    inBuiltDisplayScreenConfiguration TeamworkDisplayScreenConfigurationable
    // True if content duplication is allowed. Applicable only for Teams Rooms devices.
    isContentDuplicationAllowed *bool
    // True if dual display mode is enabled. If isDualDisplayModeEnabled is true, then the content will be displayed on both front of room screens instead of just the one screen, when it is shared via the HDMI ingest module on the Microsoft Teams Rooms device. Applicable only for Teams Rooms devices.
    isDualDisplayModeEnabled *bool
    // The OdataType property
    odataType *string
}
// NewTeamworkDisplayConfiguration instantiates a new teamworkDisplayConfiguration and sets the default values.
func NewTeamworkDisplayConfiguration()(*TeamworkDisplayConfiguration) {
    m := &TeamworkDisplayConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkDisplayConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkDisplayConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkDisplayConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkDisplayConfiguration) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetConfiguredDisplays gets the configuredDisplays property value. The list of configured displays. Applicable only for Microsoft Teams Rooms devices.
func (m *TeamworkDisplayConfiguration) GetConfiguredDisplays()([]TeamworkConfiguredPeripheralable) {
    return m.configuredDisplays
}
// GetDisplayCount gets the displayCount property value. Total number of connected displays, including the inbuilt display. Applicable only for Teams Rooms devices.
func (m *TeamworkDisplayConfiguration) GetDisplayCount()(*int32) {
    return m.displayCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkDisplayConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["configuredDisplays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTeamworkConfiguredPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamworkConfiguredPeripheralable, len(val))
            for i, v := range val {
                res[i] = v.(TeamworkConfiguredPeripheralable)
            }
            m.SetConfiguredDisplays(res)
        }
        return nil
    }
    res["displayCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayCount(val)
        }
        return nil
    }
    res["inBuiltDisplayScreenConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkDisplayScreenConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInBuiltDisplayScreenConfiguration(val.(TeamworkDisplayScreenConfigurationable))
        }
        return nil
    }
    res["isContentDuplicationAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsContentDuplicationAllowed(val)
        }
        return nil
    }
    res["isDualDisplayModeEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDualDisplayModeEnabled(val)
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
// GetInBuiltDisplayScreenConfiguration gets the inBuiltDisplayScreenConfiguration property value. Configuration for the inbuilt display. Not applicable for Teams Rooms devices.
func (m *TeamworkDisplayConfiguration) GetInBuiltDisplayScreenConfiguration()(TeamworkDisplayScreenConfigurationable) {
    return m.inBuiltDisplayScreenConfiguration
}
// GetIsContentDuplicationAllowed gets the isContentDuplicationAllowed property value. True if content duplication is allowed. Applicable only for Teams Rooms devices.
func (m *TeamworkDisplayConfiguration) GetIsContentDuplicationAllowed()(*bool) {
    return m.isContentDuplicationAllowed
}
// GetIsDualDisplayModeEnabled gets the isDualDisplayModeEnabled property value. True if dual display mode is enabled. If isDualDisplayModeEnabled is true, then the content will be displayed on both front of room screens instead of just the one screen, when it is shared via the HDMI ingest module on the Microsoft Teams Rooms device. Applicable only for Teams Rooms devices.
func (m *TeamworkDisplayConfiguration) GetIsDualDisplayModeEnabled()(*bool) {
    return m.isDualDisplayModeEnabled
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkDisplayConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *TeamworkDisplayConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetConfiguredDisplays() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConfiguredDisplays()))
        for i, v := range m.GetConfiguredDisplays() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("configuredDisplays", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("displayCount", m.GetDisplayCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("inBuiltDisplayScreenConfiguration", m.GetInBuiltDisplayScreenConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isContentDuplicationAllowed", m.GetIsContentDuplicationAllowed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isDualDisplayModeEnabled", m.GetIsDualDisplayModeEnabled())
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
func (m *TeamworkDisplayConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetConfiguredDisplays sets the configuredDisplays property value. The list of configured displays. Applicable only for Microsoft Teams Rooms devices.
func (m *TeamworkDisplayConfiguration) SetConfiguredDisplays(value []TeamworkConfiguredPeripheralable)() {
    m.configuredDisplays = value
}
// SetDisplayCount sets the displayCount property value. Total number of connected displays, including the inbuilt display. Applicable only for Teams Rooms devices.
func (m *TeamworkDisplayConfiguration) SetDisplayCount(value *int32)() {
    m.displayCount = value
}
// SetInBuiltDisplayScreenConfiguration sets the inBuiltDisplayScreenConfiguration property value. Configuration for the inbuilt display. Not applicable for Teams Rooms devices.
func (m *TeamworkDisplayConfiguration) SetInBuiltDisplayScreenConfiguration(value TeamworkDisplayScreenConfigurationable)() {
    m.inBuiltDisplayScreenConfiguration = value
}
// SetIsContentDuplicationAllowed sets the isContentDuplicationAllowed property value. True if content duplication is allowed. Applicable only for Teams Rooms devices.
func (m *TeamworkDisplayConfiguration) SetIsContentDuplicationAllowed(value *bool)() {
    m.isContentDuplicationAllowed = value
}
// SetIsDualDisplayModeEnabled sets the isDualDisplayModeEnabled property value. True if dual display mode is enabled. If isDualDisplayModeEnabled is true, then the content will be displayed on both front of room screens instead of just the one screen, when it is shared via the HDMI ingest module on the Microsoft Teams Rooms device. Applicable only for Teams Rooms devices.
func (m *TeamworkDisplayConfiguration) SetIsDualDisplayModeEnabled(value *bool)() {
    m.isDualDisplayModeEnabled = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkDisplayConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
