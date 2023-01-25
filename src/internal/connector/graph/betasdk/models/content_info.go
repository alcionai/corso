package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContentInfo 
type ContentInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The format property
    format *ContentFormat
    // Identifier used for Azure Information Protection Analytics.
    identifier *string
    // Existing Microsoft Purview Information Protection metadata is passed as key/value pairs, where the key is the MSIP_Label_GUID_PropName.
    metadata []KeyValuePairable
    // The OdataType property
    odataType *string
    // The state property
    state *ContentState
}
// NewContentInfo instantiates a new contentInfo and sets the default values.
func NewContentInfo()(*ContentInfo) {
    m := &ContentInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateContentInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContentInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContentInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContentInfo) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContentInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["format"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseContentFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(*ContentFormat))
        }
        return nil
    }
    res["identifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentifier(val)
        }
        return nil
    }
    res["metadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(KeyValuePairable)
            }
            m.SetMetadata(res)
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
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseContentState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*ContentState))
        }
        return nil
    }
    return res
}
// GetFormat gets the format property value. The format property
func (m *ContentInfo) GetFormat()(*ContentFormat) {
    return m.format
}
// GetIdentifier gets the identifier property value. Identifier used for Azure Information Protection Analytics.
func (m *ContentInfo) GetIdentifier()(*string) {
    return m.identifier
}
// GetMetadata gets the metadata property value. Existing Microsoft Purview Information Protection metadata is passed as key/value pairs, where the key is the MSIP_Label_GUID_PropName.
func (m *ContentInfo) GetMetadata()([]KeyValuePairable) {
    return m.metadata
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ContentInfo) GetOdataType()(*string) {
    return m.odataType
}
// GetState gets the state property value. The state property
func (m *ContentInfo) GetState()(*ContentState) {
    return m.state
}
// Serialize serializes information the current object
func (m *ContentInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetFormat() != nil {
        cast := (*m.GetFormat()).String()
        err := writer.WriteStringValue("format", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("identifier", m.GetIdentifier())
        if err != nil {
            return err
        }
    }
    if m.GetMetadata() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMetadata()))
        for i, v := range m.GetMetadata() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("metadata", cast)
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
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
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
func (m *ContentInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetFormat sets the format property value. The format property
func (m *ContentInfo) SetFormat(value *ContentFormat)() {
    m.format = value
}
// SetIdentifier sets the identifier property value. Identifier used for Azure Information Protection Analytics.
func (m *ContentInfo) SetIdentifier(value *string)() {
    m.identifier = value
}
// SetMetadata sets the metadata property value. Existing Microsoft Purview Information Protection metadata is passed as key/value pairs, where the key is the MSIP_Label_GUID_PropName.
func (m *ContentInfo) SetMetadata(value []KeyValuePairable)() {
    m.metadata = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ContentInfo) SetOdataType(value *string)() {
    m.odataType = value
}
// SetState sets the state property value. The state property
func (m *ContentInfo) SetState(value *ContentState)() {
    m.state = value
}
