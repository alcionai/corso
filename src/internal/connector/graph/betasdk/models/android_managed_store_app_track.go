package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidManagedStoreAppTrack contains track information for Android Managed Store apps.
type AndroidManagedStoreAppTrack struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // Friendly name for track.
    trackAlias *string
    // Unique track identifier.
    trackId *string
}
// NewAndroidManagedStoreAppTrack instantiates a new androidManagedStoreAppTrack and sets the default values.
func NewAndroidManagedStoreAppTrack()(*AndroidManagedStoreAppTrack) {
    m := &AndroidManagedStoreAppTrack{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAndroidManagedStoreAppTrackFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidManagedStoreAppTrackFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidManagedStoreAppTrack(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidManagedStoreAppTrack) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidManagedStoreAppTrack) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["trackAlias"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrackAlias(val)
        }
        return nil
    }
    res["trackId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrackId(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AndroidManagedStoreAppTrack) GetOdataType()(*string) {
    return m.odataType
}
// GetTrackAlias gets the trackAlias property value. Friendly name for track.
func (m *AndroidManagedStoreAppTrack) GetTrackAlias()(*string) {
    return m.trackAlias
}
// GetTrackId gets the trackId property value. Unique track identifier.
func (m *AndroidManagedStoreAppTrack) GetTrackId()(*string) {
    return m.trackId
}
// Serialize serializes information the current object
func (m *AndroidManagedStoreAppTrack) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("trackAlias", m.GetTrackAlias())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("trackId", m.GetTrackId())
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
func (m *AndroidManagedStoreAppTrack) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AndroidManagedStoreAppTrack) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTrackAlias sets the trackAlias property value. Friendly name for track.
func (m *AndroidManagedStoreAppTrack) SetTrackAlias(value *string)() {
    m.trackAlias = value
}
// SetTrackId sets the trackId property value. Unique track identifier.
func (m *AndroidManagedStoreAppTrack) SetTrackId(value *string)() {
    m.trackId = value
}
