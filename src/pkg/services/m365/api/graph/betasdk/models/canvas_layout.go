package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	msmodel "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// CanvasLayout
type CanvasLayout struct {
	msmodel.Entity
	// Collection of horizontal sections on the SharePoint page.
	horizontalSections []HorizontalSectionable
	// Vertical section on the SharePoint page.
	verticalSection VerticalSectionable
}

// NewCanvasLayout instantiates a new canvasLayout and sets the default values.
func NewCanvasLayout() *CanvasLayout {
	m := &CanvasLayout{
		Entity: *msmodel.NewEntity(),
	}
	return m
}

// CreateCanvasLayoutFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCanvasLayoutFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCanvasLayout(), nil
}

// GetFieldDeserializers the deserialization information for the current model
func (m *CanvasLayout) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Entity.GetFieldDeserializers()
	res["horizontalSections"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateHorizontalSectionFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]HorizontalSectionable, len(val))
			for i, v := range val {
				res[i] = v.(HorizontalSectionable)
			}
			m.SetHorizontalSections(res)
		}
		return nil
	}
	res["verticalSection"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateVerticalSectionFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetVerticalSection(val.(VerticalSectionable))
		}
		return nil
	}
	return res
}

// GetHorizontalSections gets the horizontalSections property value. Collection of horizontal sections on the SharePoint page.
func (m *CanvasLayout) GetHorizontalSections() []HorizontalSectionable {
	return m.horizontalSections
}

// GetVerticalSection gets the verticalSection property value. Vertical section on the SharePoint page.
func (m *CanvasLayout) GetVerticalSection() VerticalSectionable {
	return m.verticalSection
}

// Serialize serializes information the current object
func (m *CanvasLayout) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	err := m.Entity.Serialize(writer)
	if err != nil {
		return err
	}
	if m.GetHorizontalSections() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHorizontalSections()))
		for i, v := range m.GetHorizontalSections() {
			cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
		}
		err = writer.WriteCollectionOfObjectValues("horizontalSections", cast)
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteObjectValue("verticalSection", m.GetVerticalSection())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetHorizontalSections sets the horizontalSections property value. Collection of horizontal sections on the SharePoint page.
func (m *CanvasLayout) SetHorizontalSections(value []HorizontalSectionable) {
	m.horizontalSections = value
}

// SetVerticalSection sets the verticalSection property value. Vertical section on the SharePoint page.
func (m *CanvasLayout) SetVerticalSection(value VerticalSectionable) {
	m.verticalSection = value
}
