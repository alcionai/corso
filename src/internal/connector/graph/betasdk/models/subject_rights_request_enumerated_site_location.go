package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SubjectRightsRequestEnumeratedSiteLocation 
type SubjectRightsRequestEnumeratedSiteLocation struct {
    SubjectRightsRequestSiteLocation
    // Collection of site URLs that should be included. Includes the URL of each site, for example, https://www.contoso.com/site1.
    urls []string
}
// NewSubjectRightsRequestEnumeratedSiteLocation instantiates a new SubjectRightsRequestEnumeratedSiteLocation and sets the default values.
func NewSubjectRightsRequestEnumeratedSiteLocation()(*SubjectRightsRequestEnumeratedSiteLocation) {
    m := &SubjectRightsRequestEnumeratedSiteLocation{
        SubjectRightsRequestSiteLocation: *NewSubjectRightsRequestSiteLocation(),
    }
    odataTypeValue := "#microsoft.graph.subjectRightsRequestEnumeratedSiteLocation";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSubjectRightsRequestEnumeratedSiteLocationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSubjectRightsRequestEnumeratedSiteLocationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSubjectRightsRequestEnumeratedSiteLocation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SubjectRightsRequestEnumeratedSiteLocation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SubjectRightsRequestSiteLocation.GetFieldDeserializers()
    res["urls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetUrls(res)
        }
        return nil
    }
    return res
}
// GetUrls gets the urls property value. Collection of site URLs that should be included. Includes the URL of each site, for example, https://www.contoso.com/site1.
func (m *SubjectRightsRequestEnumeratedSiteLocation) GetUrls()([]string) {
    return m.urls
}
// Serialize serializes information the current object
func (m *SubjectRightsRequestEnumeratedSiteLocation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SubjectRightsRequestSiteLocation.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetUrls() != nil {
        err = writer.WriteCollectionOfStringValues("urls", m.GetUrls())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUrls sets the urls property value. Collection of site URLs that should be included. Includes the URL of each site, for example, https://www.contoso.com/site1.
func (m *SubjectRightsRequestEnumeratedSiteLocation) SetUrls(value []string)() {
    m.urls = value
}
