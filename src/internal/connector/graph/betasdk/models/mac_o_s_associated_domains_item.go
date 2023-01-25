package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSAssociatedDomainsItem a mapping of application identifiers to associated domains.
type MacOSAssociatedDomainsItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The application identifier of the app to associate domains with.
    applicationIdentifier *string
    // Determines whether data should be downloaded directly or via a CDN.
    directDownloadsEnabled *bool
    // The list of domains to associate.
    domains []string
    // The OdataType property
    odataType *string
}
// NewMacOSAssociatedDomainsItem instantiates a new macOSAssociatedDomainsItem and sets the default values.
func NewMacOSAssociatedDomainsItem()(*MacOSAssociatedDomainsItem) {
    m := &MacOSAssociatedDomainsItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMacOSAssociatedDomainsItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSAssociatedDomainsItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSAssociatedDomainsItem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MacOSAssociatedDomainsItem) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetApplicationIdentifier gets the applicationIdentifier property value. The application identifier of the app to associate domains with.
func (m *MacOSAssociatedDomainsItem) GetApplicationIdentifier()(*string) {
    return m.applicationIdentifier
}
// GetDirectDownloadsEnabled gets the directDownloadsEnabled property value. Determines whether data should be downloaded directly or via a CDN.
func (m *MacOSAssociatedDomainsItem) GetDirectDownloadsEnabled()(*bool) {
    return m.directDownloadsEnabled
}
// GetDomains gets the domains property value. The list of domains to associate.
func (m *MacOSAssociatedDomainsItem) GetDomains()([]string) {
    return m.domains
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSAssociatedDomainsItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["applicationIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationIdentifier(val)
        }
        return nil
    }
    res["directDownloadsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectDownloadsEnabled(val)
        }
        return nil
    }
    res["domains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDomains(res)
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MacOSAssociatedDomainsItem) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *MacOSAssociatedDomainsItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("applicationIdentifier", m.GetApplicationIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("directDownloadsEnabled", m.GetDirectDownloadsEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetDomains() != nil {
        err := writer.WriteCollectionOfStringValues("domains", m.GetDomains())
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
func (m *MacOSAssociatedDomainsItem) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetApplicationIdentifier sets the applicationIdentifier property value. The application identifier of the app to associate domains with.
func (m *MacOSAssociatedDomainsItem) SetApplicationIdentifier(value *string)() {
    m.applicationIdentifier = value
}
// SetDirectDownloadsEnabled sets the directDownloadsEnabled property value. Determines whether data should be downloaded directly or via a CDN.
func (m *MacOSAssociatedDomainsItem) SetDirectDownloadsEnabled(value *bool)() {
    m.directDownloadsEnabled = value
}
// SetDomains sets the domains property value. The list of domains to associate.
func (m *MacOSAssociatedDomainsItem) SetDomains(value []string)() {
    m.domains = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MacOSAssociatedDomainsItem) SetOdataType(value *string)() {
    m.odataType = value
}
