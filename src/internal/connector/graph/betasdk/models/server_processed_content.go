package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServerProcessedContent 
type ServerProcessedContent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // A key-value map where keys are string identifiers and values are component ids. SharePoint servers might decide to use this hint to preload the script for corresponding components for performance boost.
    componentDependencies []MetaDataKeyStringPairable
    // A key-value map where keys are string identifier and values are object of custom key-value pair.
    customMetadata []MetaDataKeyValuePairable
    // A key-value map where keys are string identifiers and values are rich text with HTML format. SharePoint servers treat the values as HTML content and run services like safety checks, search index and link fixup on them.
    htmlStrings []MetaDataKeyStringPairable
    // A key-value map where keys are string identifiers and values are image sources. SharePoint servers treat the values as image sources and run services like search index and link fixup on them.
    imageSources []MetaDataKeyStringPairable
    // A key-value map where keys are string identifiers and values are links. SharePoint servers treat the values as links and run services like link fixup on them.
    links []MetaDataKeyStringPairable
    // The OdataType property
    odataType *string
    // A key-value map where keys are string identifiers and values are strings that should be search indexed.
    searchablePlainTexts []MetaDataKeyStringPairable
}
// NewServerProcessedContent instantiates a new serverProcessedContent and sets the default values.
func NewServerProcessedContent()(*ServerProcessedContent) {
    m := &ServerProcessedContent{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateServerProcessedContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServerProcessedContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServerProcessedContent(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServerProcessedContent) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetComponentDependencies gets the componentDependencies property value. A key-value map where keys are string identifiers and values are component ids. SharePoint servers might decide to use this hint to preload the script for corresponding components for performance boost.
func (m *ServerProcessedContent) GetComponentDependencies()([]MetaDataKeyStringPairable) {
    return m.componentDependencies
}
// GetCustomMetadata gets the customMetadata property value. A key-value map where keys are string identifier and values are object of custom key-value pair.
func (m *ServerProcessedContent) GetCustomMetadata()([]MetaDataKeyValuePairable) {
    return m.customMetadata
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServerProcessedContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["componentDependencies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMetaDataKeyStringPairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MetaDataKeyStringPairable, len(val))
            for i, v := range val {
                res[i] = v.(MetaDataKeyStringPairable)
            }
            m.SetComponentDependencies(res)
        }
        return nil
    }
    res["customMetadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMetaDataKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MetaDataKeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(MetaDataKeyValuePairable)
            }
            m.SetCustomMetadata(res)
        }
        return nil
    }
    res["htmlStrings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMetaDataKeyStringPairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MetaDataKeyStringPairable, len(val))
            for i, v := range val {
                res[i] = v.(MetaDataKeyStringPairable)
            }
            m.SetHtmlStrings(res)
        }
        return nil
    }
    res["imageSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMetaDataKeyStringPairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MetaDataKeyStringPairable, len(val))
            for i, v := range val {
                res[i] = v.(MetaDataKeyStringPairable)
            }
            m.SetImageSources(res)
        }
        return nil
    }
    res["links"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMetaDataKeyStringPairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MetaDataKeyStringPairable, len(val))
            for i, v := range val {
                res[i] = v.(MetaDataKeyStringPairable)
            }
            m.SetLinks(res)
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
    res["searchablePlainTexts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMetaDataKeyStringPairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MetaDataKeyStringPairable, len(val))
            for i, v := range val {
                res[i] = v.(MetaDataKeyStringPairable)
            }
            m.SetSearchablePlainTexts(res)
        }
        return nil
    }
    return res
}
// GetHtmlStrings gets the htmlStrings property value. A key-value map where keys are string identifiers and values are rich text with HTML format. SharePoint servers treat the values as HTML content and run services like safety checks, search index and link fixup on them.
func (m *ServerProcessedContent) GetHtmlStrings()([]MetaDataKeyStringPairable) {
    return m.htmlStrings
}
// GetImageSources gets the imageSources property value. A key-value map where keys are string identifiers and values are image sources. SharePoint servers treat the values as image sources and run services like search index and link fixup on them.
func (m *ServerProcessedContent) GetImageSources()([]MetaDataKeyStringPairable) {
    return m.imageSources
}
// GetLinks gets the links property value. A key-value map where keys are string identifiers and values are links. SharePoint servers treat the values as links and run services like link fixup on them.
func (m *ServerProcessedContent) GetLinks()([]MetaDataKeyStringPairable) {
    return m.links
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ServerProcessedContent) GetOdataType()(*string) {
    return m.odataType
}
// GetSearchablePlainTexts gets the searchablePlainTexts property value. A key-value map where keys are string identifiers and values are strings that should be search indexed.
func (m *ServerProcessedContent) GetSearchablePlainTexts()([]MetaDataKeyStringPairable) {
    return m.searchablePlainTexts
}
// Serialize serializes information the current object
func (m *ServerProcessedContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetComponentDependencies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetComponentDependencies()))
        for i, v := range m.GetComponentDependencies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("componentDependencies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomMetadata() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomMetadata()))
        for i, v := range m.GetCustomMetadata() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("customMetadata", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHtmlStrings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHtmlStrings()))
        for i, v := range m.GetHtmlStrings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("htmlStrings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetImageSources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetImageSources()))
        for i, v := range m.GetImageSources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("imageSources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetLinks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLinks()))
        for i, v := range m.GetLinks() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("links", cast)
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
    if m.GetSearchablePlainTexts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSearchablePlainTexts()))
        for i, v := range m.GetSearchablePlainTexts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("searchablePlainTexts", cast)
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
func (m *ServerProcessedContent) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetComponentDependencies sets the componentDependencies property value. A key-value map where keys are string identifiers and values are component ids. SharePoint servers might decide to use this hint to preload the script for corresponding components for performance boost.
func (m *ServerProcessedContent) SetComponentDependencies(value []MetaDataKeyStringPairable)() {
    m.componentDependencies = value
}
// SetCustomMetadata sets the customMetadata property value. A key-value map where keys are string identifier and values are object of custom key-value pair.
func (m *ServerProcessedContent) SetCustomMetadata(value []MetaDataKeyValuePairable)() {
    m.customMetadata = value
}
// SetHtmlStrings sets the htmlStrings property value. A key-value map where keys are string identifiers and values are rich text with HTML format. SharePoint servers treat the values as HTML content and run services like safety checks, search index and link fixup on them.
func (m *ServerProcessedContent) SetHtmlStrings(value []MetaDataKeyStringPairable)() {
    m.htmlStrings = value
}
// SetImageSources sets the imageSources property value. A key-value map where keys are string identifiers and values are image sources. SharePoint servers treat the values as image sources and run services like search index and link fixup on them.
func (m *ServerProcessedContent) SetImageSources(value []MetaDataKeyStringPairable)() {
    m.imageSources = value
}
// SetLinks sets the links property value. A key-value map where keys are string identifiers and values are links. SharePoint servers treat the values as links and run services like link fixup on them.
func (m *ServerProcessedContent) SetLinks(value []MetaDataKeyStringPairable)() {
    m.links = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ServerProcessedContent) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSearchablePlainTexts sets the searchablePlainTexts property value. A key-value map where keys are string identifiers and values are strings that should be search indexed.
func (m *ServerProcessedContent) SetSearchablePlainTexts(value []MetaDataKeyStringPairable)() {
    m.searchablePlainTexts = value
}
