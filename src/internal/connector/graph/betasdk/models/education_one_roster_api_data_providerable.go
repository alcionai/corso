package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationOneRosterApiDataProviderable 
type EducationOneRosterApiDataProviderable interface {
    EducationSynchronizationDataProviderable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConnectionSettings()(EducationSynchronizationConnectionSettingsable)
    GetConnectionUrl()(*string)
    GetCustomizations()(EducationSynchronizationCustomizationsable)
    GetProviderName()(*string)
    GetSchoolsIds()([]string)
    GetTermIds()([]string)
    SetConnectionSettings(value EducationSynchronizationConnectionSettingsable)()
    SetConnectionUrl(value *string)()
    SetCustomizations(value EducationSynchronizationCustomizationsable)()
    SetProviderName(value *string)()
    SetSchoolsIds(value []string)()
    SetTermIds(value []string)()
}
