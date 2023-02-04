package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SiteSettings
type SiteSettings struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]interface{}
	// The language tag for the language used on this site.
	languageTag *string
	// The OdataType property
	odataType *string
	// Indicates the time offset for the time zone of the site from Coordinated Universal Time (UTC).
	timeZone *string
}

// NewSiteSettings instantiates a new siteSettings and sets the default values.
func NewSiteSettings() *SiteSettings {
	m := &SiteSettings{}
	m.SetAdditionalData(make(map[string]interface{}))
	return m
}

// CreateSiteSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSiteSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSiteSettings(), nil
}

// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SiteSettings) GetAdditionalData() map[string]interface{} {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
func (m *SiteSettings) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["languageTag"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLanguageTag(val)
		}
		return nil
	}
	res["@odata.type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOdataType(val)
		}
		return nil
	}
	res["timeZone"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTimeZone(val)
		}
		return nil
	}
	return res
}

// GetLanguageTag gets the languageTag property value. The language tag for the language used on this site.
func (m *SiteSettings) GetLanguageTag() *string {
	return m.languageTag
}

// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SiteSettings) GetOdataType() *string {
	return m.odataType
}

// GetTimeZone gets the timeZone property value. Indicates the time offset for the time zone of the site from Coordinated Universal Time (UTC).
func (m *SiteSettings) GetTimeZone() *string {
	return m.timeZone
}

// Serialize serializes information the current object
func (m *SiteSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("languageTag", m.GetLanguageTag())
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
		err := writer.WriteStringValue("timeZone", m.GetTimeZone())
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
func (m *SiteSettings) SetAdditionalData(value map[string]interface{}) {
	m.additionalData = value
}

// SetLanguageTag sets the languageTag property value. The language tag for the language used on this site.
func (m *SiteSettings) SetLanguageTag(value *string) {
	m.languageTag = value
}

// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SiteSettings) SetOdataType(value *string) {
	m.odataType = value
}

// SetTimeZone sets the timeZone property value. Indicates the time offset for the time zone of the site from Coordinated Universal Time (UTC).
func (m *SiteSettings) SetTimeZone(value *string) {
	m.timeZone = value
}
