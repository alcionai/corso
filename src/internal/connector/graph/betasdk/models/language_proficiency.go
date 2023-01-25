package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LanguageProficiency 
type LanguageProficiency struct {
    ItemFacet
    // Contains the long-form name for the language.
    displayName *string
    // The proficiency property
    proficiency *LanguageProficiencyLevel
    // Represents the users reading comprehension for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
    reading *LanguageProficiencyLevel
    // Represents the users spoken proficiency for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
    spoken *LanguageProficiencyLevel
    // Contains the four-character BCP47 name for the language (en-US, no-NB, en-AU).
    tag *string
    // The thumbnailUrl property
    thumbnailUrl *string
    // Represents the users written proficiency for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
    written *LanguageProficiencyLevel
}
// NewLanguageProficiency instantiates a new LanguageProficiency and sets the default values.
func NewLanguageProficiency()(*LanguageProficiency) {
    m := &LanguageProficiency{
        ItemFacet: *NewItemFacet(),
    }
    odataTypeValue := "#microsoft.graph.languageProficiency";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateLanguageProficiencyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLanguageProficiencyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLanguageProficiency(), nil
}
// GetDisplayName gets the displayName property value. Contains the long-form name for the language.
func (m *LanguageProficiency) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LanguageProficiency) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
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
    res["proficiency"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLanguageProficiencyLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProficiency(val.(*LanguageProficiencyLevel))
        }
        return nil
    }
    res["reading"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLanguageProficiencyLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReading(val.(*LanguageProficiencyLevel))
        }
        return nil
    }
    res["spoken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLanguageProficiencyLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpoken(val.(*LanguageProficiencyLevel))
        }
        return nil
    }
    res["tag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTag(val)
        }
        return nil
    }
    res["thumbnailUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThumbnailUrl(val)
        }
        return nil
    }
    res["written"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLanguageProficiencyLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWritten(val.(*LanguageProficiencyLevel))
        }
        return nil
    }
    return res
}
// GetProficiency gets the proficiency property value. The proficiency property
func (m *LanguageProficiency) GetProficiency()(*LanguageProficiencyLevel) {
    return m.proficiency
}
// GetReading gets the reading property value. Represents the users reading comprehension for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
func (m *LanguageProficiency) GetReading()(*LanguageProficiencyLevel) {
    return m.reading
}
// GetSpoken gets the spoken property value. Represents the users spoken proficiency for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
func (m *LanguageProficiency) GetSpoken()(*LanguageProficiencyLevel) {
    return m.spoken
}
// GetTag gets the tag property value. Contains the four-character BCP47 name for the language (en-US, no-NB, en-AU).
func (m *LanguageProficiency) GetTag()(*string) {
    return m.tag
}
// GetThumbnailUrl gets the thumbnailUrl property value. The thumbnailUrl property
func (m *LanguageProficiency) GetThumbnailUrl()(*string) {
    return m.thumbnailUrl
}
// GetWritten gets the written property value. Represents the users written proficiency for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
func (m *LanguageProficiency) GetWritten()(*LanguageProficiencyLevel) {
    return m.written
}
// Serialize serializes information the current object
func (m *LanguageProficiency) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetProficiency() != nil {
        cast := (*m.GetProficiency()).String()
        err = writer.WriteStringValue("proficiency", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetReading() != nil {
        cast := (*m.GetReading()).String()
        err = writer.WriteStringValue("reading", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSpoken() != nil {
        cast := (*m.GetSpoken()).String()
        err = writer.WriteStringValue("spoken", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tag", m.GetTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("thumbnailUrl", m.GetThumbnailUrl())
        if err != nil {
            return err
        }
    }
    if m.GetWritten() != nil {
        cast := (*m.GetWritten()).String()
        err = writer.WriteStringValue("written", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. Contains the long-form name for the language.
func (m *LanguageProficiency) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetProficiency sets the proficiency property value. The proficiency property
func (m *LanguageProficiency) SetProficiency(value *LanguageProficiencyLevel)() {
    m.proficiency = value
}
// SetReading sets the reading property value. Represents the users reading comprehension for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
func (m *LanguageProficiency) SetReading(value *LanguageProficiencyLevel)() {
    m.reading = value
}
// SetSpoken sets the spoken property value. Represents the users spoken proficiency for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
func (m *LanguageProficiency) SetSpoken(value *LanguageProficiencyLevel)() {
    m.spoken = value
}
// SetTag sets the tag property value. Contains the four-character BCP47 name for the language (en-US, no-NB, en-AU).
func (m *LanguageProficiency) SetTag(value *string)() {
    m.tag = value
}
// SetThumbnailUrl sets the thumbnailUrl property value. The thumbnailUrl property
func (m *LanguageProficiency) SetThumbnailUrl(value *string)() {
    m.thumbnailUrl = value
}
// SetWritten sets the written property value. Represents the users written proficiency for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
func (m *LanguageProficiency) SetWritten(value *LanguageProficiencyLevel)() {
    m.written = value
}
