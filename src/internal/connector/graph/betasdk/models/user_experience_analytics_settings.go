package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsSettings the user experience analytics insight is the recomendation to improve the user experience analytics score.
type UserExperienceAnalyticsSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // True if Tenant attach is configured. If configured then SCCM tenant attached devices will show up in UXA reporting.
    configurationManagerDataConnectorConfigured *bool
    // The OdataType property
    odataType *string
}
// NewUserExperienceAnalyticsSettings instantiates a new userExperienceAnalyticsSettings and sets the default values.
func NewUserExperienceAnalyticsSettings()(*UserExperienceAnalyticsSettings) {
    m := &UserExperienceAnalyticsSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUserExperienceAnalyticsSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsSettings) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetConfigurationManagerDataConnectorConfigured gets the configurationManagerDataConnectorConfigured property value. True if Tenant attach is configured. If configured then SCCM tenant attached devices will show up in UXA reporting.
func (m *UserExperienceAnalyticsSettings) GetConfigurationManagerDataConnectorConfigured()(*bool) {
    return m.configurationManagerDataConnectorConfigured
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["configurationManagerDataConnectorConfigured"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationManagerDataConnectorConfigured(val)
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
func (m *UserExperienceAnalyticsSettings) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("configurationManagerDataConnectorConfigured", m.GetConfigurationManagerDataConnectorConfigured())
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
func (m *UserExperienceAnalyticsSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetConfigurationManagerDataConnectorConfigured sets the configurationManagerDataConnectorConfigured property value. True if Tenant attach is configured. If configured then SCCM tenant attached devices will show up in UXA reporting.
func (m *UserExperienceAnalyticsSettings) SetConfigurationManagerDataConnectorConfigured(value *bool)() {
    m.configurationManagerDataConnectorConfigured = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserExperienceAnalyticsSettings) SetOdataType(value *string)() {
    m.odataType = value
}
