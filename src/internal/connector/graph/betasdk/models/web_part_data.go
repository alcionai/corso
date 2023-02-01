package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	msmodel "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// WebPartData
type WebPartData struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]interface{}
	// Audience information of the web part. By using this property, specific content will be prioritized to specific audiences.
	audiences []string
	// Data version of the web part. The value is defined by the web part developer. Different dataVersions usually refers to a different property structure.
	dataVersion *string
	// Description of the web part.
	description *string
	// The OdataType property
	odataType *string
	// Properties bag of the web part.
	properties msmodel.Jsonable
	// Contains collections of data that can be processed by server side services like search index and link fixup.
	serverProcessedContent ServerProcessedContentable
	// Title of the web part.
	title *string
}

// NewWebPartData instantiates a new webPartData and sets the default values.
func NewWebPartData() *WebPartData {
	m := &WebPartData{}
	m.SetAdditionalData(make(map[string]interface{}))
	return m
}

// CreateWebPartDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWebPartDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewWebPartData(), nil
}

// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WebPartData) GetAdditionalData() map[string]interface{} {
	return m.additionalData
}

// GetAudiences gets the audiences property value. Audience information of the web part. By using this property, specific content will be prioritized to specific audiences.
func (m *WebPartData) GetAudiences() []string {
	return m.audiences
}

// GetDataVersion gets the dataVersion property value. Data version of the web part. The value is defined by the web part developer. Different dataVersions usually refers to a different property structure.
func (m *WebPartData) GetDataVersion() *string {
	return m.dataVersion
}

// GetDescription gets the description property value. Description of the web part.
func (m *WebPartData) GetDescription() *string {
	return m.description
}

// GetFieldDeserializers the deserialization information for the current model
func (m *WebPartData) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["audiences"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfPrimitiveValues("string")
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]string, len(val))
			for i, v := range val {
				res[i] = *(v.(*string))
			}
			m.SetAudiences(res)
		}
		return nil
	}
	res["dataVersion"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDataVersion(val)
		}
		return nil
	}
	res["description"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDescription(val)
		}
		return nil
	}
	res["@odata.type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOdataType(val)
		}
		return nil
	}
	res["properties"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(msmodel.CreateJsonFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetProperties(val.(msmodel.Jsonable))
		}
		return nil
	}
	res["serverProcessedContent"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateServerProcessedContentFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetServerProcessedContent(val.(ServerProcessedContentable))
		}
		return nil
	}
	res["title"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTitle(val)
		}
		return nil
	}
	return res
}

// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WebPartData) GetOdataType() *string {
	return m.odataType
}

// GetProperties gets the properties property value. Properties bag of the web part.
func (m *WebPartData) GetProperties() msmodel.Jsonable {
	return m.properties
}

// GetServerProcessedContent gets the serverProcessedContent property value. Contains collections of data that can be processed by server side services like search index and link fixup.
func (m *WebPartData) GetServerProcessedContent() ServerProcessedContentable {
	return m.serverProcessedContent
}

// GetTitle gets the title property value. Title of the web part.
func (m *WebPartData) GetTitle() *string {
	return m.title
}

// Serialize serializes information the current object
func (m *WebPartData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetAudiences() != nil {
		err := writer.WriteCollectionOfStringValues("audiences", m.GetAudiences())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("dataVersion", m.GetDataVersion())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("description", m.GetDescription())
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
		err := writer.WriteObjectValue("properties", m.GetProperties())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("serverProcessedContent", m.GetServerProcessedContent())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("title", m.GetTitle())
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
func (m *WebPartData) SetAdditionalData(value map[string]interface{}) {
	m.additionalData = value
}

// SetAudiences sets the audiences property value. Audience information of the web part. By using this property, specific content will be prioritized to specific audiences.
func (m *WebPartData) SetAudiences(value []string) {
	m.audiences = value
}

// SetDataVersion sets the dataVersion property value. Data version of the web part. The value is defined by the web part developer. Different dataVersions usually refers to a different property structure.
func (m *WebPartData) SetDataVersion(value *string) {
	m.dataVersion = value
}

// SetDescription sets the description property value. Description of the web part.
func (m *WebPartData) SetDescription(value *string) {
	m.description = value
}

// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WebPartData) SetOdataType(value *string) {
	m.odataType = value
}

// SetProperties sets the properties property value. Properties bag of the web part.
func (m *WebPartData) SetProperties(value msmodel.Jsonable) {
	m.properties = value
}

// SetServerProcessedContent sets the serverProcessedContent property value. Contains collections of data that can be processed by server side services like search index and link fixup.
func (m *WebPartData) SetServerProcessedContent(value ServerProcessedContentable) {
	m.serverProcessedContent = value
}

// SetTitle sets the title property value. Title of the web part.
func (m *WebPartData) SetTitle(value *string) {
	m.title = value
}
