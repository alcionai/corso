package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContentProperties 
type ContentProperties struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The extensions property
    extensions []string
    // The lastModifiedBy property
    lastModifiedBy *string
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The metadata property
    metadata ContentMetadataable
    // The OdataType property
    odataType *string
}
// NewContentProperties instantiates a new contentProperties and sets the default values.
func NewContentProperties()(*ContentProperties) {
    m := &ContentProperties{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateContentPropertiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContentPropertiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.attachmentContentProperties":
                        return NewAttachmentContentProperties(), nil
                }
            }
        }
    }
    return NewContentProperties(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContentProperties) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetExtensions gets the extensions property value. The extensions property
func (m *ContentProperties) GetExtensions()([]string) {
    return m.extensions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContentProperties) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["extensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExtensions(res)
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val)
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
    res["metadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateContentMetadataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetadata(val.(ContentMetadataable))
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
// GetLastModifiedBy gets the lastModifiedBy property value. The lastModifiedBy property
func (m *ContentProperties) GetLastModifiedBy()(*string) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *ContentProperties) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetMetadata gets the metadata property value. The metadata property
func (m *ContentProperties) GetMetadata()(ContentMetadataable) {
    return m.metadata
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ContentProperties) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ContentProperties) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetExtensions() != nil {
        err := writer.WriteCollectionOfStringValues("extensions", m.GetExtensions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("lastModifiedBy", m.GetLastModifiedBy())
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
        err := writer.WriteObjectValue("metadata", m.GetMetadata())
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
func (m *ContentProperties) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetExtensions sets the extensions property value. The extensions property
func (m *ContentProperties) SetExtensions(value []string)() {
    m.extensions = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The lastModifiedBy property
func (m *ContentProperties) SetLastModifiedBy(value *string)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *ContentProperties) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetMetadata sets the metadata property value. The metadata property
func (m *ContentProperties) SetMetadata(value ContentMetadataable)() {
    m.metadata = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ContentProperties) SetOdataType(value *string)() {
    m.odataType = value
}
