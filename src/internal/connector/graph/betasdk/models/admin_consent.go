package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AdminConsent admin consent information.
type AdminConsent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // Admin consent state.
    shareAPNSData *AdminConsentState
    // Admin consent state.
    shareUserExperienceAnalyticsData *AdminConsentState
}
// NewAdminConsent instantiates a new adminConsent and sets the default values.
func NewAdminConsent()(*AdminConsent) {
    m := &AdminConsent{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAdminConsentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdminConsentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdminConsent(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AdminConsent) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AdminConsent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["shareAPNSData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAdminConsentState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShareAPNSData(val.(*AdminConsentState))
        }
        return nil
    }
    res["shareUserExperienceAnalyticsData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAdminConsentState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShareUserExperienceAnalyticsData(val.(*AdminConsentState))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AdminConsent) GetOdataType()(*string) {
    return m.odataType
}
// GetShareAPNSData gets the shareAPNSData property value. Admin consent state.
func (m *AdminConsent) GetShareAPNSData()(*AdminConsentState) {
    return m.shareAPNSData
}
// GetShareUserExperienceAnalyticsData gets the shareUserExperienceAnalyticsData property value. Admin consent state.
func (m *AdminConsent) GetShareUserExperienceAnalyticsData()(*AdminConsentState) {
    return m.shareUserExperienceAnalyticsData
}
// Serialize serializes information the current object
func (m *AdminConsent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetShareAPNSData() != nil {
        cast := (*m.GetShareAPNSData()).String()
        err := writer.WriteStringValue("shareAPNSData", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetShareUserExperienceAnalyticsData() != nil {
        cast := (*m.GetShareUserExperienceAnalyticsData()).String()
        err := writer.WriteStringValue("shareUserExperienceAnalyticsData", &cast)
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
func (m *AdminConsent) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AdminConsent) SetOdataType(value *string)() {
    m.odataType = value
}
// SetShareAPNSData sets the shareAPNSData property value. Admin consent state.
func (m *AdminConsent) SetShareAPNSData(value *AdminConsentState)() {
    m.shareAPNSData = value
}
// SetShareUserExperienceAnalyticsData sets the shareUserExperienceAnalyticsData property value. Admin consent state.
func (m *AdminConsent) SetShareUserExperienceAnalyticsData(value *AdminConsentState)() {
    m.shareUserExperienceAnalyticsData = value
}
