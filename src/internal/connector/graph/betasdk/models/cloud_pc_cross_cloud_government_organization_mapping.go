package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcCrossCloudGovernmentOrganizationMapping 
type CloudPcCrossCloudGovernmentOrganizationMapping struct {
    Entity
    // The tenant ID in the Azure Government cloud corresponding to the GCC tenant in the public cloud. Currently, 1:1 mappings are supported, so this collection can only contain one tenant ID.
    organizationIdsInUSGovCloud []string
}
// NewCloudPcCrossCloudGovernmentOrganizationMapping instantiates a new CloudPcCrossCloudGovernmentOrganizationMapping and sets the default values.
func NewCloudPcCrossCloudGovernmentOrganizationMapping()(*CloudPcCrossCloudGovernmentOrganizationMapping) {
    m := &CloudPcCrossCloudGovernmentOrganizationMapping{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPcCrossCloudGovernmentOrganizationMappingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcCrossCloudGovernmentOrganizationMappingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcCrossCloudGovernmentOrganizationMapping(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcCrossCloudGovernmentOrganizationMapping) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["organizationIdsInUSGovCloud"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetOrganizationIdsInUSGovCloud(res)
        }
        return nil
    }
    return res
}
// GetOrganizationIdsInUSGovCloud gets the organizationIdsInUSGovCloud property value. The tenant ID in the Azure Government cloud corresponding to the GCC tenant in the public cloud. Currently, 1:1 mappings are supported, so this collection can only contain one tenant ID.
func (m *CloudPcCrossCloudGovernmentOrganizationMapping) GetOrganizationIdsInUSGovCloud()([]string) {
    return m.organizationIdsInUSGovCloud
}
// Serialize serializes information the current object
func (m *CloudPcCrossCloudGovernmentOrganizationMapping) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetOrganizationIdsInUSGovCloud() != nil {
        err = writer.WriteCollectionOfStringValues("organizationIdsInUSGovCloud", m.GetOrganizationIdsInUSGovCloud())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOrganizationIdsInUSGovCloud sets the organizationIdsInUSGovCloud property value. The tenant ID in the Azure Government cloud corresponding to the GCC tenant in the public cloud. Currently, 1:1 mappings are supported, so this collection can only contain one tenant ID.
func (m *CloudPcCrossCloudGovernmentOrganizationMapping) SetOrganizationIdsInUSGovCloud(value []string)() {
    m.organizationIdsInUSGovCloud = value
}
