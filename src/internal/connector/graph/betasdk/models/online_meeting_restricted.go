package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnlineMeetingRestricted 
type OnlineMeetingRestricted struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The contentSharingDisabled property
    contentSharingDisabled *OnlineMeetingContentSharingDisabledReason
    // The OdataType property
    odataType *string
    // The videoDisabled property
    videoDisabled *OnlineMeetingVideoDisabledReason
}
// NewOnlineMeetingRestricted instantiates a new onlineMeetingRestricted and sets the default values.
func NewOnlineMeetingRestricted()(*OnlineMeetingRestricted) {
    m := &OnlineMeetingRestricted{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOnlineMeetingRestrictedFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnlineMeetingRestrictedFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnlineMeetingRestricted(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnlineMeetingRestricted) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetContentSharingDisabled gets the contentSharingDisabled property value. The contentSharingDisabled property
func (m *OnlineMeetingRestricted) GetContentSharingDisabled()(*OnlineMeetingContentSharingDisabledReason) {
    return m.contentSharingDisabled
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnlineMeetingRestricted) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contentSharingDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnlineMeetingContentSharingDisabledReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentSharingDisabled(val.(*OnlineMeetingContentSharingDisabledReason))
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
    res["videoDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnlineMeetingVideoDisabledReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVideoDisabled(val.(*OnlineMeetingVideoDisabledReason))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OnlineMeetingRestricted) GetOdataType()(*string) {
    return m.odataType
}
// GetVideoDisabled gets the videoDisabled property value. The videoDisabled property
func (m *OnlineMeetingRestricted) GetVideoDisabled()(*OnlineMeetingVideoDisabledReason) {
    return m.videoDisabled
}
// Serialize serializes information the current object
func (m *OnlineMeetingRestricted) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetContentSharingDisabled() != nil {
        cast := (*m.GetContentSharingDisabled()).String()
        err := writer.WriteStringValue("contentSharingDisabled", &cast)
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
    if m.GetVideoDisabled() != nil {
        cast := (*m.GetVideoDisabled()).String()
        err := writer.WriteStringValue("videoDisabled", &cast)
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
func (m *OnlineMeetingRestricted) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetContentSharingDisabled sets the contentSharingDisabled property value. The contentSharingDisabled property
func (m *OnlineMeetingRestricted) SetContentSharingDisabled(value *OnlineMeetingContentSharingDisabledReason)() {
    m.contentSharingDisabled = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OnlineMeetingRestricted) SetOdataType(value *string)() {
    m.odataType = value
}
// SetVideoDisabled sets the videoDisabled property value. The videoDisabled property
func (m *OnlineMeetingRestricted) SetVideoDisabled(value *OnlineMeetingVideoDisabledReason)() {
    m.videoDisabled = value
}
