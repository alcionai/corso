package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WebPartPosition
type WebPartPosition struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]interface{}
	// Indicates the identifier of the column where the web part is located.
	columnId *float64
	// Indicates the horizontal section where the web part is located.
	horizontalSectionId *float64
	// Indicates whether the web part is located in the vertical section.
	isInVerticalSection *bool
	// The OdataType property
	odataType *string
	// Index of the current web part. Represents the order of the web part in this column or section.
	webPartIndex *float64
}

// NewWebPartPosition instantiates a new webPartPosition and sets the default values.
func NewWebPartPosition() *WebPartPosition {
	m := &WebPartPosition{}
	m.SetAdditionalData(make(map[string]interface{}))
	return m
}

// CreateWebPartPositionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWebPartPositionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewWebPartPosition(), nil
}

// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WebPartPosition) GetAdditionalData() map[string]interface{} {
	return m.additionalData
}

// GetColumnId gets the columnId property value. Indicates the identifier of the column where the web part is located.
func (m *WebPartPosition) GetColumnId() *float64 {
	return m.columnId
}

// GetFieldDeserializers the deserialization information for the current model
func (m *WebPartPosition) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["columnId"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetColumnId(val)
		}
		return nil
	}
	res["horizontalSectionId"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetHorizontalSectionId(val)
		}
		return nil
	}
	res["isInVerticalSection"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIsInVerticalSection(val)
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
	res["webPartIndex"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetWebPartIndex(val)
		}
		return nil
	}
	return res
}

// GetHorizontalSectionId gets the horizontalSectionId property value. Indicates the horizontal section where the web part is located.
func (m *WebPartPosition) GetHorizontalSectionId() *float64 {
	return m.horizontalSectionId
}

// GetIsInVerticalSection gets the isInVerticalSection property value. Indicates whether the web part is located in the vertical section.
func (m *WebPartPosition) GetIsInVerticalSection() *bool {
	return m.isInVerticalSection
}

// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WebPartPosition) GetOdataType() *string {
	return m.odataType
}

// GetWebPartIndex gets the webPartIndex property value. Index of the current web part. Represents the order of the web part in this column or section.
func (m *WebPartPosition) GetWebPartIndex() *float64 {
	return m.webPartIndex
}

// Serialize serializes information the current object
func (m *WebPartPosition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteFloat64Value("columnId", m.GetColumnId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteFloat64Value("horizontalSectionId", m.GetHorizontalSectionId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("isInVerticalSection", m.GetIsInVerticalSection())
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
		err := writer.WriteFloat64Value("webPartIndex", m.GetWebPartIndex())
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
func (m *WebPartPosition) SetAdditionalData(value map[string]interface{}) {
	m.additionalData = value
}

// SetColumnId sets the columnId property value. Indicates the identifier of the column where the web part is located.
func (m *WebPartPosition) SetColumnId(value *float64) {
	m.columnId = value
}

// SetHorizontalSectionId sets the horizontalSectionId property value. Indicates the horizontal section where the web part is located.
func (m *WebPartPosition) SetHorizontalSectionId(value *float64) {
	m.horizontalSectionId = value
}

// SetIsInVerticalSection sets the isInVerticalSection property value. Indicates whether the web part is located in the vertical section.
func (m *WebPartPosition) SetIsInVerticalSection(value *bool) {
	m.isInVerticalSection = value
}

// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WebPartPosition) SetOdataType(value *string) {
	m.odataType = value
}

// SetWebPartIndex sets the webPartIndex property value. Index of the current web part. Represents the order of the web part in this column or section.
func (m *WebPartPosition) SetWebPartIndex(value *float64) {
	m.webPartIndex = value
}
