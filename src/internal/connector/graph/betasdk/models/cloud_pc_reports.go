package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcReports 
type CloudPcReports struct {
    Entity
    // The export jobs created for downloading reports.
    exportJobs []CloudPcExportJobable
}
// NewCloudPcReports instantiates a new CloudPcReports and sets the default values.
func NewCloudPcReports()(*CloudPcReports) {
    m := &CloudPcReports{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPcReportsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcReportsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcReports(), nil
}
// GetExportJobs gets the exportJobs property value. The export jobs created for downloading reports.
func (m *CloudPcReports) GetExportJobs()([]CloudPcExportJobable) {
    return m.exportJobs
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcReports) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["exportJobs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcExportJobFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcExportJobable, len(val))
            for i, v := range val {
                res[i] = v.(CloudPcExportJobable)
            }
            m.SetExportJobs(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CloudPcReports) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetExportJobs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExportJobs()))
        for i, v := range m.GetExportJobs() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("exportJobs", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExportJobs sets the exportJobs property value. The export jobs created for downloading reports.
func (m *CloudPcReports) SetExportJobs(value []CloudPcExportJobable)() {
    m.exportJobs = value
}
