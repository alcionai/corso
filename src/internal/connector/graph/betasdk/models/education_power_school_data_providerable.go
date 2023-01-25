package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationPowerSchoolDataProviderable 
type EducationPowerSchoolDataProviderable interface {
    EducationSynchronizationDataProviderable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowTeachersInMultipleSchools()(*bool)
    GetClientId()(*string)
    GetClientSecret()(*string)
    GetConnectionUrl()(*string)
    GetCustomizations()(EducationSynchronizationCustomizationsable)
    GetSchoolsIds()([]string)
    GetSchoolYear()(*string)
    SetAllowTeachersInMultipleSchools(value *bool)()
    SetClientId(value *string)()
    SetClientSecret(value *string)()
    SetConnectionUrl(value *string)()
    SetCustomizations(value EducationSynchronizationCustomizationsable)()
    SetSchoolsIds(value []string)()
    SetSchoolYear(value *string)()
}
