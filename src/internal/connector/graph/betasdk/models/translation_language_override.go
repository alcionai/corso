package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TranslationLanguageOverride 
type TranslationLanguageOverride struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The language to apply the override.Returned by default. Not nullable.
    languageTag *string
    // The OdataType property
    odataType *string
    // The translation override behavior for the language, if any.Returned by default. Not nullable.
    translationBehavior *TranslationBehavior
}
// NewTranslationLanguageOverride instantiates a new translationLanguageOverride and sets the default values.
func NewTranslationLanguageOverride()(*TranslationLanguageOverride) {
    m := &TranslationLanguageOverride{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTranslationLanguageOverrideFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTranslationLanguageOverrideFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTranslationLanguageOverride(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TranslationLanguageOverride) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TranslationLanguageOverride) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["translationBehavior"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTranslationBehavior)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTranslationBehavior(val.(*TranslationBehavior))
        }
        return nil
    }
    return res
}
// GetLanguageTag gets the languageTag property value. The language to apply the override.Returned by default. Not nullable.
func (m *TranslationLanguageOverride) GetLanguageTag()(*string) {
    return m.languageTag
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TranslationLanguageOverride) GetOdataType()(*string) {
    return m.odataType
}
// GetTranslationBehavior gets the translationBehavior property value. The translation override behavior for the language, if any.Returned by default. Not nullable.
func (m *TranslationLanguageOverride) GetTranslationBehavior()(*TranslationBehavior) {
    return m.translationBehavior
}
// Serialize serializes information the current object
func (m *TranslationLanguageOverride) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetTranslationBehavior() != nil {
        cast := (*m.GetTranslationBehavior()).String()
        err := writer.WriteStringValue("translationBehavior", &cast)
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
func (m *TranslationLanguageOverride) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetLanguageTag sets the languageTag property value. The language to apply the override.Returned by default. Not nullable.
func (m *TranslationLanguageOverride) SetLanguageTag(value *string)() {
    m.languageTag = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TranslationLanguageOverride) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTranslationBehavior sets the translationBehavior property value. The translation override behavior for the language, if any.Returned by default. Not nullable.
func (m *TranslationLanguageOverride) SetTranslationBehavior(value *TranslationBehavior)() {
    m.translationBehavior = value
}
