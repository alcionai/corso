package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementReports 
type DeviceManagementReports struct {
    Entity
    // Entity representing the configuration of a cached report
    cachedReportConfigurations []DeviceManagementCachedReportConfigurationable
    // Entity representing a job to export a report
    exportJobs []DeviceManagementExportJobable
}
// NewDeviceManagementReports instantiates a new deviceManagementReports and sets the default values.
func NewDeviceManagementReports()(*DeviceManagementReports) {
    m := &DeviceManagementReports{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementReportsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementReportsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementReports(), nil
}
// GetCachedReportConfigurations gets the cachedReportConfigurations property value. Entity representing the configuration of a cached report
func (m *DeviceManagementReports) GetCachedReportConfigurations()([]DeviceManagementCachedReportConfigurationable) {
    return m.cachedReportConfigurations
}
// GetExportJobs gets the exportJobs property value. Entity representing a job to export a report
func (m *DeviceManagementReports) GetExportJobs()([]DeviceManagementExportJobable) {
    return m.exportJobs
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementReports) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["cachedReportConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementCachedReportConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementCachedReportConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementCachedReportConfigurationable)
            }
            m.SetCachedReportConfigurations(res)
        }
        return nil
    }
    res["exportJobs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementExportJobFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementExportJobable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementExportJobable)
            }
            m.SetExportJobs(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementReports) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCachedReportConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCachedReportConfigurations()))
        for i, v := range m.GetCachedReportConfigurations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("cachedReportConfigurations", cast)
        if err != nil {
            return err
        }
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
// SetCachedReportConfigurations sets the cachedReportConfigurations property value. Entity representing the configuration of a cached report
func (m *DeviceManagementReports) SetCachedReportConfigurations(value []DeviceManagementCachedReportConfigurationable)() {
    m.cachedReportConfigurations = value
}
// SetExportJobs sets the exportJobs property value. Entity representing a job to export a report
func (m *DeviceManagementReports) SetExportJobs(value []DeviceManagementExportJobable)() {
    m.exportJobs = value
}
