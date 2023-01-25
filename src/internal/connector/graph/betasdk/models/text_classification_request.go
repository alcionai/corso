package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TextClassificationRequest 
type TextClassificationRequest struct {
    Entity
    // The fileExtension property
    fileExtension *string
    // The matchTolerancesToInclude property
    matchTolerancesToInclude *MlClassificationMatchTolerance
    // The scopesToRun property
    scopesToRun *SensitiveTypeScope
    // The sensitiveTypeIds property
    sensitiveTypeIds []string
    // The text property
    text *string
}
// NewTextClassificationRequest instantiates a new TextClassificationRequest and sets the default values.
func NewTextClassificationRequest()(*TextClassificationRequest) {
    m := &TextClassificationRequest{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTextClassificationRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTextClassificationRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTextClassificationRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TextClassificationRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["fileExtension"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileExtension(val)
        }
        return nil
    }
    res["matchTolerancesToInclude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMlClassificationMatchTolerance)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMatchTolerancesToInclude(val.(*MlClassificationMatchTolerance))
        }
        return nil
    }
    res["scopesToRun"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensitiveTypeScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScopesToRun(val.(*SensitiveTypeScope))
        }
        return nil
    }
    res["sensitiveTypeIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSensitiveTypeIds(res)
        }
        return nil
    }
    res["text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val)
        }
        return nil
    }
    return res
}
// GetFileExtension gets the fileExtension property value. The fileExtension property
func (m *TextClassificationRequest) GetFileExtension()(*string) {
    return m.fileExtension
}
// GetMatchTolerancesToInclude gets the matchTolerancesToInclude property value. The matchTolerancesToInclude property
func (m *TextClassificationRequest) GetMatchTolerancesToInclude()(*MlClassificationMatchTolerance) {
    return m.matchTolerancesToInclude
}
// GetScopesToRun gets the scopesToRun property value. The scopesToRun property
func (m *TextClassificationRequest) GetScopesToRun()(*SensitiveTypeScope) {
    return m.scopesToRun
}
// GetSensitiveTypeIds gets the sensitiveTypeIds property value. The sensitiveTypeIds property
func (m *TextClassificationRequest) GetSensitiveTypeIds()([]string) {
    return m.sensitiveTypeIds
}
// GetText gets the text property value. The text property
func (m *TextClassificationRequest) GetText()(*string) {
    return m.text
}
// Serialize serializes information the current object
func (m *TextClassificationRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("fileExtension", m.GetFileExtension())
        if err != nil {
            return err
        }
    }
    if m.GetMatchTolerancesToInclude() != nil {
        cast := (*m.GetMatchTolerancesToInclude()).String()
        err = writer.WriteStringValue("matchTolerancesToInclude", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetScopesToRun() != nil {
        cast := (*m.GetScopesToRun()).String()
        err = writer.WriteStringValue("scopesToRun", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSensitiveTypeIds() != nil {
        err = writer.WriteCollectionOfStringValues("sensitiveTypeIds", m.GetSensitiveTypeIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("text", m.GetText())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFileExtension sets the fileExtension property value. The fileExtension property
func (m *TextClassificationRequest) SetFileExtension(value *string)() {
    m.fileExtension = value
}
// SetMatchTolerancesToInclude sets the matchTolerancesToInclude property value. The matchTolerancesToInclude property
func (m *TextClassificationRequest) SetMatchTolerancesToInclude(value *MlClassificationMatchTolerance)() {
    m.matchTolerancesToInclude = value
}
// SetScopesToRun sets the scopesToRun property value. The scopesToRun property
func (m *TextClassificationRequest) SetScopesToRun(value *SensitiveTypeScope)() {
    m.scopesToRun = value
}
// SetSensitiveTypeIds sets the sensitiveTypeIds property value. The sensitiveTypeIds property
func (m *TextClassificationRequest) SetSensitiveTypeIds(value []string)() {
    m.sensitiveTypeIds = value
}
// SetText sets the text property value. The text property
func (m *TextClassificationRequest) SetText(value *string)() {
    m.text = value
}
