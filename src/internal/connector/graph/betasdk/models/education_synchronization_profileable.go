package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationSynchronizationProfileable 
type EducationSynchronizationProfileable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDataProvider()(EducationSynchronizationDataProviderable)
    GetDisplayName()(*string)
    GetErrors()([]EducationSynchronizationErrorable)
    GetExpirationDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetHandleSpecialCharacterConstraint()(*bool)
    GetIdentitySynchronizationConfiguration()(EducationIdentitySynchronizationConfigurationable)
    GetLicensesToAssign()([]EducationSynchronizationLicenseAssignmentable)
    GetProfileStatus()(EducationSynchronizationProfileStatusable)
    GetState()(*EducationSynchronizationProfileState)
    SetDataProvider(value EducationSynchronizationDataProviderable)()
    SetDisplayName(value *string)()
    SetErrors(value []EducationSynchronizationErrorable)()
    SetExpirationDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetHandleSpecialCharacterConstraint(value *bool)()
    SetIdentitySynchronizationConfiguration(value EducationIdentitySynchronizationConfigurationable)()
    SetLicensesToAssign(value []EducationSynchronizationLicenseAssignmentable)()
    SetProfileStatus(value EducationSynchronizationProfileStatusable)()
    SetState(value *EducationSynchronizationProfileState)()
}
