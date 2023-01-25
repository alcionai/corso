package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationSynchronizationCustomizationsable 
type EducationSynchronizationCustomizationsable interface {
    EducationSynchronizationCustomizationsBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSchool()(EducationSynchronizationCustomizationable)
    GetSection()(EducationSynchronizationCustomizationable)
    GetStudent()(EducationSynchronizationCustomizationable)
    GetStudentEnrollment()(EducationSynchronizationCustomizationable)
    GetTeacher()(EducationSynchronizationCustomizationable)
    GetTeacherRoster()(EducationSynchronizationCustomizationable)
    SetSchool(value EducationSynchronizationCustomizationable)()
    SetSection(value EducationSynchronizationCustomizationable)()
    SetStudent(value EducationSynchronizationCustomizationable)()
    SetStudentEnrollment(value EducationSynchronizationCustomizationable)()
    SetTeacher(value EducationSynchronizationCustomizationable)()
    SetTeacherRoster(value EducationSynchronizationCustomizationable)()
}
