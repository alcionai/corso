package search

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// AnswerVariant 
type AnswerVariant struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Answer variation description shown on search results page.
    description *string
    // Answer variation name displayed in search results.
    displayName *string
    // The languageTag property
    languageTag *string
    // The OdataType property
    odataType *string
    // The platform property
    platform *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.DevicePlatformType
    // Answer variation URL link. When users click this answer variation in search results, they will go to this URL.
    webUrl *string
}
// NewAnswerVariant instantiates a new answerVariant and sets the default values.
func NewAnswerVariant()(*AnswerVariant) {
    m := &AnswerVariant{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAnswerVariantFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAnswerVariantFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAnswerVariant(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AnswerVariant) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDescription gets the description property value. Answer variation description shown on search results page.
func (m *AnswerVariant) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Answer variation name displayed in search results.
func (m *AnswerVariant) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AnswerVariant) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["languageTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguageTag(val)
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
    res["platform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.ParseDevicePlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val.(*ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.DevicePlatformType))
        }
        return nil
    }
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetLanguageTag gets the languageTag property value. The languageTag property
func (m *AnswerVariant) GetLanguageTag()(*string) {
    return m.languageTag
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AnswerVariant) GetOdataType()(*string) {
    return m.odataType
}
// GetPlatform gets the platform property value. The platform property
func (m *AnswerVariant) GetPlatform()(*ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.DevicePlatformType) {
    return m.platform
}
// GetWebUrl gets the webUrl property value. Answer variation URL link. When users click this answer variation in search results, they will go to this URL.
func (m *AnswerVariant) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *AnswerVariant) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
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
    if m.GetPlatform() != nil {
        cast := (*m.GetPlatform()).String()
        err := writer.WriteStringValue("platform", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("webUrl", m.GetWebUrl())
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
func (m *AnswerVariant) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDescription sets the description property value. Answer variation description shown on search results page.
func (m *AnswerVariant) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Answer variation name displayed in search results.
func (m *AnswerVariant) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLanguageTag sets the languageTag property value. The languageTag property
func (m *AnswerVariant) SetLanguageTag(value *string)() {
    m.languageTag = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AnswerVariant) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPlatform sets the platform property value. The platform property
func (m *AnswerVariant) SetPlatform(value *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.DevicePlatformType)() {
    m.platform = value
}
// SetWebUrl sets the webUrl property value. Answer variation URL link. When users click this answer variation in search results, they will go to this URL.
func (m *AnswerVariant) SetWebUrl(value *string)() {
    m.webUrl = value
}
