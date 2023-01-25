package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProfileCardAnnotation 
type ProfileCardAnnotation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // If present, the value of this field is used by the profile card as the default property label in the experience (for example, 'Cost Center').
    displayName *string
    // Each resource in this collection represents the localized value of the attribute name for a given language, used as the default label for that locale. For example, a user with a no-NB client gets 'Kostnads Senter' as the attribute label, rather than 'Cost Center.'
    localizations []DisplayNameLocalizationable
    // The OdataType property
    odataType *string
}
// NewProfileCardAnnotation instantiates a new profileCardAnnotation and sets the default values.
func NewProfileCardAnnotation()(*ProfileCardAnnotation) {
    m := &ProfileCardAnnotation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateProfileCardAnnotationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProfileCardAnnotationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProfileCardAnnotation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProfileCardAnnotation) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. If present, the value of this field is used by the profile card as the default property label in the experience (for example, 'Cost Center').
func (m *ProfileCardAnnotation) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProfileCardAnnotation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["localizations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDisplayNameLocalizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DisplayNameLocalizationable, len(val))
            for i, v := range val {
                res[i] = v.(DisplayNameLocalizationable)
            }
            m.SetLocalizations(res)
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
// GetLocalizations gets the localizations property value. Each resource in this collection represents the localized value of the attribute name for a given language, used as the default label for that locale. For example, a user with a no-NB client gets 'Kostnads Senter' as the attribute label, rather than 'Cost Center.'
func (m *ProfileCardAnnotation) GetLocalizations()([]DisplayNameLocalizationable) {
    return m.localizations
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ProfileCardAnnotation) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ProfileCardAnnotation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetLocalizations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLocalizations()))
        for i, v := range m.GetLocalizations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("localizations", cast)
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
func (m *ProfileCardAnnotation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. If present, the value of this field is used by the profile card as the default property label in the experience (for example, 'Cost Center').
func (m *ProfileCardAnnotation) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLocalizations sets the localizations property value. Each resource in this collection represents the localized value of the attribute name for a given language, used as the default label for that locale. For example, a user with a no-NB client gets 'Kostnads Senter' as the attribute label, rather than 'Cost Center.'
func (m *ProfileCardAnnotation) SetLocalizations(value []DisplayNameLocalizationable)() {
    m.localizations = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ProfileCardAnnotation) SetOdataType(value *string)() {
    m.odataType = value
}
