package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Profileable 
type Profileable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccount()([]UserAccountInformationable)
    GetAddresses()([]ItemAddressable)
    GetAnniversaries()([]PersonAnnualEventable)
    GetAwards()([]PersonAwardable)
    GetCertifications()([]PersonCertificationable)
    GetEducationalActivities()([]EducationalActivityable)
    GetEmails()([]ItemEmailable)
    GetInterests()([]PersonInterestable)
    GetLanguages()([]LanguageProficiencyable)
    GetNames()([]PersonNameable)
    GetNotes()([]PersonAnnotationable)
    GetPatents()([]ItemPatentable)
    GetPhones()([]ItemPhoneable)
    GetPositions()([]WorkPositionable)
    GetProjects()([]ProjectParticipationable)
    GetPublications()([]ItemPublicationable)
    GetSkills()([]SkillProficiencyable)
    GetWebAccounts()([]WebAccountable)
    GetWebsites()([]PersonWebsiteable)
    SetAccount(value []UserAccountInformationable)()
    SetAddresses(value []ItemAddressable)()
    SetAnniversaries(value []PersonAnnualEventable)()
    SetAwards(value []PersonAwardable)()
    SetCertifications(value []PersonCertificationable)()
    SetEducationalActivities(value []EducationalActivityable)()
    SetEmails(value []ItemEmailable)()
    SetInterests(value []PersonInterestable)()
    SetLanguages(value []LanguageProficiencyable)()
    SetNames(value []PersonNameable)()
    SetNotes(value []PersonAnnotationable)()
    SetPatents(value []ItemPatentable)()
    SetPhones(value []ItemPhoneable)()
    SetPositions(value []WorkPositionable)()
    SetProjects(value []ProjectParticipationable)()
    SetPublications(value []ItemPublicationable)()
    SetSkills(value []SkillProficiencyable)()
    SetWebAccounts(value []WebAccountable)()
    SetWebsites(value []PersonWebsiteable)()
}
