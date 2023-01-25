package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TranslationPreferences 
type TranslationPreferences struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Translation override behavior for languages, if any.Returned by default.
    languageOverrides []TranslationLanguageOverrideable
    // The OdataType property
    odataType *string
    // The user's preferred translation behavior.Returned by default. Not nullable.
    translationBehavior *TranslationBehavior
    // The list of languages the user does not need translated. This is computed from the authoringLanguages collection in regionalAndLanguageSettings, and the languageOverrides collection in translationPreferences. The list specifies neutral culture values that include the language code without any country or region association. For example, it would specify 'fr' for the neutral French culture, but not 'fr-FR' for the French culture in France. Returned by default. Read only.
    untranslatedLanguages []string
}
// NewTranslationPreferences instantiates a new translationPreferences and sets the default values.
func NewTranslationPreferences()(*TranslationPreferences) {
    m := &TranslationPreferences{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTranslationPreferencesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTranslationPreferencesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTranslationPreferences(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TranslationPreferences) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TranslationPreferences) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["languageOverrides"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTranslationLanguageOverrideFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TranslationLanguageOverrideable, len(val))
            for i, v := range val {
                res[i] = v.(TranslationLanguageOverrideable)
            }
            m.SetLanguageOverrides(res)
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
    res["untranslatedLanguages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetUntranslatedLanguages(res)
        }
        return nil
    }
    return res
}
// GetLanguageOverrides gets the languageOverrides property value. Translation override behavior for languages, if any.Returned by default.
func (m *TranslationPreferences) GetLanguageOverrides()([]TranslationLanguageOverrideable) {
    return m.languageOverrides
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TranslationPreferences) GetOdataType()(*string) {
    return m.odataType
}
// GetTranslationBehavior gets the translationBehavior property value. The user's preferred translation behavior.Returned by default. Not nullable.
func (m *TranslationPreferences) GetTranslationBehavior()(*TranslationBehavior) {
    return m.translationBehavior
}
// GetUntranslatedLanguages gets the untranslatedLanguages property value. The list of languages the user does not need translated. This is computed from the authoringLanguages collection in regionalAndLanguageSettings, and the languageOverrides collection in translationPreferences. The list specifies neutral culture values that include the language code without any country or region association. For example, it would specify 'fr' for the neutral French culture, but not 'fr-FR' for the French culture in France. Returned by default. Read only.
func (m *TranslationPreferences) GetUntranslatedLanguages()([]string) {
    return m.untranslatedLanguages
}
// Serialize serializes information the current object
func (m *TranslationPreferences) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetLanguageOverrides() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLanguageOverrides()))
        for i, v := range m.GetLanguageOverrides() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("languageOverrides", cast)
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
    if m.GetUntranslatedLanguages() != nil {
        err := writer.WriteCollectionOfStringValues("untranslatedLanguages", m.GetUntranslatedLanguages())
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
func (m *TranslationPreferences) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetLanguageOverrides sets the languageOverrides property value. Translation override behavior for languages, if any.Returned by default.
func (m *TranslationPreferences) SetLanguageOverrides(value []TranslationLanguageOverrideable)() {
    m.languageOverrides = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TranslationPreferences) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTranslationBehavior sets the translationBehavior property value. The user's preferred translation behavior.Returned by default. Not nullable.
func (m *TranslationPreferences) SetTranslationBehavior(value *TranslationBehavior)() {
    m.translationBehavior = value
}
// SetUntranslatedLanguages sets the untranslatedLanguages property value. The list of languages the user does not need translated. This is computed from the authoringLanguages collection in regionalAndLanguageSettings, and the languageOverrides collection in translationPreferences. The list specifies neutral culture values that include the language code without any country or region association. For example, it would specify 'fr' for the neutral French culture, but not 'fr-FR' for the French culture in France. Returned by default. Read only.
func (m *TranslationPreferences) SetUntranslatedLanguages(value []string)() {
    m.untranslatedLanguages = value
}
