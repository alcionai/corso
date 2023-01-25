package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SearchSettings 
type SearchSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // Enables the developer to define the appearance of the content and configure conditions that dictate when the template should be displayed. Maximum of 2 search result templates per connection.
    searchResultTemplates []DisplayTemplateable
}
// NewSearchSettings instantiates a new searchSettings and sets the default values.
func NewSearchSettings()(*SearchSettings) {
    m := &SearchSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSearchSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSearchSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSearchSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchSettings) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["searchResultTemplates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDisplayTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DisplayTemplateable, len(val))
            for i, v := range val {
                res[i] = v.(DisplayTemplateable)
            }
            m.SetSearchResultTemplates(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SearchSettings) GetOdataType()(*string) {
    return m.odataType
}
// GetSearchResultTemplates gets the searchResultTemplates property value. Enables the developer to define the appearance of the content and configure conditions that dictate when the template should be displayed. Maximum of 2 search result templates per connection.
func (m *SearchSettings) GetSearchResultTemplates()([]DisplayTemplateable) {
    return m.searchResultTemplates
}
// Serialize serializes information the current object
func (m *SearchSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetSearchResultTemplates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSearchResultTemplates()))
        for i, v := range m.GetSearchResultTemplates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("searchResultTemplates", cast)
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
func (m *SearchSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SearchSettings) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSearchResultTemplates sets the searchResultTemplates property value. Enables the developer to define the appearance of the content and configure conditions that dictate when the template should be displayed. Maximum of 2 search result templates per connection.
func (m *SearchSettings) SetSearchResultTemplates(value []DisplayTemplateable)() {
    m.searchResultTemplates = value
}
