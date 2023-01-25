package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyUploadedLanguageFile the entity represents an ADML (Administrative Template language) XML file uploaded by Administrator.
type GroupPolicyUploadedLanguageFile struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The contents of the uploaded ADML file.
    content []byte
    // The file name of the uploaded ADML file.
    fileName *string
    // Key of the entity.
    id *string
    // The language code of the uploaded ADML file.
    languageCode *string
    // The date and time the entity was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
}
// NewGroupPolicyUploadedLanguageFile instantiates a new groupPolicyUploadedLanguageFile and sets the default values.
func NewGroupPolicyUploadedLanguageFile()(*GroupPolicyUploadedLanguageFile) {
    m := &GroupPolicyUploadedLanguageFile{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGroupPolicyUploadedLanguageFileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyUploadedLanguageFileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyUploadedLanguageFile(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupPolicyUploadedLanguageFile) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetContent gets the content property value. The contents of the uploaded ADML file.
func (m *GroupPolicyUploadedLanguageFile) GetContent()([]byte) {
    return m.content
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyUploadedLanguageFile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["fileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileName(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["languageCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguageCode(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
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
// GetFileName gets the fileName property value. The file name of the uploaded ADML file.
func (m *GroupPolicyUploadedLanguageFile) GetFileName()(*string) {
    return m.fileName
}
// GetId gets the id property value. Key of the entity.
func (m *GroupPolicyUploadedLanguageFile) GetId()(*string) {
    return m.id
}
// GetLanguageCode gets the languageCode property value. The language code of the uploaded ADML file.
func (m *GroupPolicyUploadedLanguageFile) GetLanguageCode()(*string) {
    return m.languageCode
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyUploadedLanguageFile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *GroupPolicyUploadedLanguageFile) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *GroupPolicyUploadedLanguageFile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fileName", m.GetFileName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("languageCode", m.GetLanguageCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
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
func (m *GroupPolicyUploadedLanguageFile) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetContent sets the content property value. The contents of the uploaded ADML file.
func (m *GroupPolicyUploadedLanguageFile) SetContent(value []byte)() {
    m.content = value
}
// SetFileName sets the fileName property value. The file name of the uploaded ADML file.
func (m *GroupPolicyUploadedLanguageFile) SetFileName(value *string)() {
    m.fileName = value
}
// SetId sets the id property value. Key of the entity.
func (m *GroupPolicyUploadedLanguageFile) SetId(value *string)() {
    m.id = value
}
// SetLanguageCode sets the languageCode property value. The language code of the uploaded ADML file.
func (m *GroupPolicyUploadedLanguageFile) SetLanguageCode(value *string)() {
    m.languageCode = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyUploadedLanguageFile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *GroupPolicyUploadedLanguageFile) SetOdataType(value *string)() {
    m.odataType = value
}
