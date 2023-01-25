package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UrlThreatSubmission 
type UrlThreatSubmission struct {
    ThreatSubmission
    // Denotes the webUrl that needs to be submitted.
    webUrl *string
}
// NewUrlThreatSubmission instantiates a new UrlThreatSubmission and sets the default values.
func NewUrlThreatSubmission()(*UrlThreatSubmission) {
    m := &UrlThreatSubmission{
        ThreatSubmission: *NewThreatSubmission(),
    }
    odataTypeValue := "#microsoft.graph.security.urlThreatSubmission";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUrlThreatSubmissionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUrlThreatSubmissionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUrlThreatSubmission(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UrlThreatSubmission) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ThreatSubmission.GetFieldDeserializers()
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetWebUrl gets the webUrl property value. Denotes the webUrl that needs to be submitted.
func (m *UrlThreatSubmission) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *UrlThreatSubmission) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ThreatSubmission.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetWebUrl sets the webUrl property value. Denotes the webUrl that needs to be submitted.
func (m *UrlThreatSubmission) SetWebUrl(value *string)() {
    m.webUrl = value
}
