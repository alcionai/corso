package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Profile 
type Profile struct {
    Entity
    // The account property
    account []UserAccountInformationable
    // Represents details of addresses associated with the user.
    addresses []ItemAddressable
    // Represents the details of meaningful dates associated with a person.
    anniversaries []PersonAnnualEventable
    // Represents the details of awards or honors associated with a person.
    awards []PersonAwardable
    // Represents the details of certifications associated with a person.
    certifications []PersonCertificationable
    // Represents data that a user has supplied related to undergraduate, graduate, postgraduate or other educational activities.
    educationalActivities []EducationalActivityable
    // Represents detailed information about email addresses associated with the user.
    emails []ItemEmailable
    // Provides detailed information about interests the user has associated with themselves in various services.
    interests []PersonInterestable
    // Represents detailed information about languages that a user has added to their profile.
    languages []LanguageProficiencyable
    // Represents the names a user has added to their profile.
    names []PersonNameable
    // Represents notes that a user has added to their profile.
    notes []PersonAnnotationable
    // Represents patents that a user has added to their profile.
    patents []ItemPatentable
    // Represents detailed information about phone numbers associated with a user in various services.
    phones []ItemPhoneable
    // Represents detailed information about work positions associated with a user's profile.
    positions []WorkPositionable
    // Represents detailed information about projects associated with a user.
    projects []ProjectParticipationable
    // Represents details of any publications a user has added to their profile.
    publications []ItemPublicationable
    // Represents detailed information about skills associated with a user in various services.
    skills []SkillProficiencyable
    // Represents web accounts the user has indicated they use or has added to their user profile.
    webAccounts []WebAccountable
    // Represents detailed information about websites associated with a user in various services.
    websites []PersonWebsiteable
}
// NewProfile instantiates a new profile and sets the default values.
func NewProfile()(*Profile) {
    m := &Profile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProfile(), nil
}
// GetAccount gets the account property value. The account property
func (m *Profile) GetAccount()([]UserAccountInformationable) {
    return m.account
}
// GetAddresses gets the addresses property value. Represents details of addresses associated with the user.
func (m *Profile) GetAddresses()([]ItemAddressable) {
    return m.addresses
}
// GetAnniversaries gets the anniversaries property value. Represents the details of meaningful dates associated with a person.
func (m *Profile) GetAnniversaries()([]PersonAnnualEventable) {
    return m.anniversaries
}
// GetAwards gets the awards property value. Represents the details of awards or honors associated with a person.
func (m *Profile) GetAwards()([]PersonAwardable) {
    return m.awards
}
// GetCertifications gets the certifications property value. Represents the details of certifications associated with a person.
func (m *Profile) GetCertifications()([]PersonCertificationable) {
    return m.certifications
}
// GetEducationalActivities gets the educationalActivities property value. Represents data that a user has supplied related to undergraduate, graduate, postgraduate or other educational activities.
func (m *Profile) GetEducationalActivities()([]EducationalActivityable) {
    return m.educationalActivities
}
// GetEmails gets the emails property value. Represents detailed information about email addresses associated with the user.
func (m *Profile) GetEmails()([]ItemEmailable) {
    return m.emails
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Profile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["account"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserAccountInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserAccountInformationable, len(val))
            for i, v := range val {
                res[i] = v.(UserAccountInformationable)
            }
            m.SetAccount(res)
        }
        return nil
    }
    res["addresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemAddressable, len(val))
            for i, v := range val {
                res[i] = v.(ItemAddressable)
            }
            m.SetAddresses(res)
        }
        return nil
    }
    res["anniversaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePersonAnnualEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PersonAnnualEventable, len(val))
            for i, v := range val {
                res[i] = v.(PersonAnnualEventable)
            }
            m.SetAnniversaries(res)
        }
        return nil
    }
    res["awards"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePersonAwardFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PersonAwardable, len(val))
            for i, v := range val {
                res[i] = v.(PersonAwardable)
            }
            m.SetAwards(res)
        }
        return nil
    }
    res["certifications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePersonCertificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PersonCertificationable, len(val))
            for i, v := range val {
                res[i] = v.(PersonCertificationable)
            }
            m.SetCertifications(res)
        }
        return nil
    }
    res["educationalActivities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationalActivityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationalActivityable, len(val))
            for i, v := range val {
                res[i] = v.(EducationalActivityable)
            }
            m.SetEducationalActivities(res)
        }
        return nil
    }
    res["emails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemEmailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemEmailable, len(val))
            for i, v := range val {
                res[i] = v.(ItemEmailable)
            }
            m.SetEmails(res)
        }
        return nil
    }
    res["interests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePersonInterestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PersonInterestable, len(val))
            for i, v := range val {
                res[i] = v.(PersonInterestable)
            }
            m.SetInterests(res)
        }
        return nil
    }
    res["languages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLanguageProficiencyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LanguageProficiencyable, len(val))
            for i, v := range val {
                res[i] = v.(LanguageProficiencyable)
            }
            m.SetLanguages(res)
        }
        return nil
    }
    res["names"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePersonNameFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PersonNameable, len(val))
            for i, v := range val {
                res[i] = v.(PersonNameable)
            }
            m.SetNames(res)
        }
        return nil
    }
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePersonAnnotationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PersonAnnotationable, len(val))
            for i, v := range val {
                res[i] = v.(PersonAnnotationable)
            }
            m.SetNotes(res)
        }
        return nil
    }
    res["patents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemPatentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemPatentable, len(val))
            for i, v := range val {
                res[i] = v.(ItemPatentable)
            }
            m.SetPatents(res)
        }
        return nil
    }
    res["phones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemPhoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemPhoneable, len(val))
            for i, v := range val {
                res[i] = v.(ItemPhoneable)
            }
            m.SetPhones(res)
        }
        return nil
    }
    res["positions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkPositionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkPositionable, len(val))
            for i, v := range val {
                res[i] = v.(WorkPositionable)
            }
            m.SetPositions(res)
        }
        return nil
    }
    res["projects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProjectParticipationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProjectParticipationable, len(val))
            for i, v := range val {
                res[i] = v.(ProjectParticipationable)
            }
            m.SetProjects(res)
        }
        return nil
    }
    res["publications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemPublicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemPublicationable, len(val))
            for i, v := range val {
                res[i] = v.(ItemPublicationable)
            }
            m.SetPublications(res)
        }
        return nil
    }
    res["skills"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSkillProficiencyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SkillProficiencyable, len(val))
            for i, v := range val {
                res[i] = v.(SkillProficiencyable)
            }
            m.SetSkills(res)
        }
        return nil
    }
    res["webAccounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWebAccountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WebAccountable, len(val))
            for i, v := range val {
                res[i] = v.(WebAccountable)
            }
            m.SetWebAccounts(res)
        }
        return nil
    }
    res["websites"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePersonWebsiteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PersonWebsiteable, len(val))
            for i, v := range val {
                res[i] = v.(PersonWebsiteable)
            }
            m.SetWebsites(res)
        }
        return nil
    }
    return res
}
// GetInterests gets the interests property value. Provides detailed information about interests the user has associated with themselves in various services.
func (m *Profile) GetInterests()([]PersonInterestable) {
    return m.interests
}
// GetLanguages gets the languages property value. Represents detailed information about languages that a user has added to their profile.
func (m *Profile) GetLanguages()([]LanguageProficiencyable) {
    return m.languages
}
// GetNames gets the names property value. Represents the names a user has added to their profile.
func (m *Profile) GetNames()([]PersonNameable) {
    return m.names
}
// GetNotes gets the notes property value. Represents notes that a user has added to their profile.
func (m *Profile) GetNotes()([]PersonAnnotationable) {
    return m.notes
}
// GetPatents gets the patents property value. Represents patents that a user has added to their profile.
func (m *Profile) GetPatents()([]ItemPatentable) {
    return m.patents
}
// GetPhones gets the phones property value. Represents detailed information about phone numbers associated with a user in various services.
func (m *Profile) GetPhones()([]ItemPhoneable) {
    return m.phones
}
// GetPositions gets the positions property value. Represents detailed information about work positions associated with a user's profile.
func (m *Profile) GetPositions()([]WorkPositionable) {
    return m.positions
}
// GetProjects gets the projects property value. Represents detailed information about projects associated with a user.
func (m *Profile) GetProjects()([]ProjectParticipationable) {
    return m.projects
}
// GetPublications gets the publications property value. Represents details of any publications a user has added to their profile.
func (m *Profile) GetPublications()([]ItemPublicationable) {
    return m.publications
}
// GetSkills gets the skills property value. Represents detailed information about skills associated with a user in various services.
func (m *Profile) GetSkills()([]SkillProficiencyable) {
    return m.skills
}
// GetWebAccounts gets the webAccounts property value. Represents web accounts the user has indicated they use or has added to their user profile.
func (m *Profile) GetWebAccounts()([]WebAccountable) {
    return m.webAccounts
}
// GetWebsites gets the websites property value. Represents detailed information about websites associated with a user in various services.
func (m *Profile) GetWebsites()([]PersonWebsiteable) {
    return m.websites
}
// Serialize serializes information the current object
func (m *Profile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccount() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccount()))
        for i, v := range m.GetAccount() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("account", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAddresses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAddresses()))
        for i, v := range m.GetAddresses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("addresses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAnniversaries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAnniversaries()))
        for i, v := range m.GetAnniversaries() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("anniversaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAwards() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAwards()))
        for i, v := range m.GetAwards() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("awards", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCertifications() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCertifications()))
        for i, v := range m.GetCertifications() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("certifications", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEducationalActivities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEducationalActivities()))
        for i, v := range m.GetEducationalActivities() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("educationalActivities", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEmails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEmails()))
        for i, v := range m.GetEmails() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("emails", cast)
        if err != nil {
            return err
        }
    }
    if m.GetInterests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetInterests()))
        for i, v := range m.GetInterests() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("interests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetLanguages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLanguages()))
        for i, v := range m.GetLanguages() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("languages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetNames() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNames()))
        for i, v := range m.GetNames() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("names", cast)
        if err != nil {
            return err
        }
    }
    if m.GetNotes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNotes()))
        for i, v := range m.GetNotes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("notes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPatents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPatents()))
        for i, v := range m.GetPatents() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("patents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPhones() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPhones()))
        for i, v := range m.GetPhones() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("phones", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPositions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPositions()))
        for i, v := range m.GetPositions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("positions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetProjects() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProjects()))
        for i, v := range m.GetProjects() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("projects", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPublications() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPublications()))
        for i, v := range m.GetPublications() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("publications", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSkills() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSkills()))
        for i, v := range m.GetSkills() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("skills", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWebAccounts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWebAccounts()))
        for i, v := range m.GetWebAccounts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("webAccounts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWebsites() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWebsites()))
        for i, v := range m.GetWebsites() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("websites", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccount sets the account property value. The account property
func (m *Profile) SetAccount(value []UserAccountInformationable)() {
    m.account = value
}
// SetAddresses sets the addresses property value. Represents details of addresses associated with the user.
func (m *Profile) SetAddresses(value []ItemAddressable)() {
    m.addresses = value
}
// SetAnniversaries sets the anniversaries property value. Represents the details of meaningful dates associated with a person.
func (m *Profile) SetAnniversaries(value []PersonAnnualEventable)() {
    m.anniversaries = value
}
// SetAwards sets the awards property value. Represents the details of awards or honors associated with a person.
func (m *Profile) SetAwards(value []PersonAwardable)() {
    m.awards = value
}
// SetCertifications sets the certifications property value. Represents the details of certifications associated with a person.
func (m *Profile) SetCertifications(value []PersonCertificationable)() {
    m.certifications = value
}
// SetEducationalActivities sets the educationalActivities property value. Represents data that a user has supplied related to undergraduate, graduate, postgraduate or other educational activities.
func (m *Profile) SetEducationalActivities(value []EducationalActivityable)() {
    m.educationalActivities = value
}
// SetEmails sets the emails property value. Represents detailed information about email addresses associated with the user.
func (m *Profile) SetEmails(value []ItemEmailable)() {
    m.emails = value
}
// SetInterests sets the interests property value. Provides detailed information about interests the user has associated with themselves in various services.
func (m *Profile) SetInterests(value []PersonInterestable)() {
    m.interests = value
}
// SetLanguages sets the languages property value. Represents detailed information about languages that a user has added to their profile.
func (m *Profile) SetLanguages(value []LanguageProficiencyable)() {
    m.languages = value
}
// SetNames sets the names property value. Represents the names a user has added to their profile.
func (m *Profile) SetNames(value []PersonNameable)() {
    m.names = value
}
// SetNotes sets the notes property value. Represents notes that a user has added to their profile.
func (m *Profile) SetNotes(value []PersonAnnotationable)() {
    m.notes = value
}
// SetPatents sets the patents property value. Represents patents that a user has added to their profile.
func (m *Profile) SetPatents(value []ItemPatentable)() {
    m.patents = value
}
// SetPhones sets the phones property value. Represents detailed information about phone numbers associated with a user in various services.
func (m *Profile) SetPhones(value []ItemPhoneable)() {
    m.phones = value
}
// SetPositions sets the positions property value. Represents detailed information about work positions associated with a user's profile.
func (m *Profile) SetPositions(value []WorkPositionable)() {
    m.positions = value
}
// SetProjects sets the projects property value. Represents detailed information about projects associated with a user.
func (m *Profile) SetProjects(value []ProjectParticipationable)() {
    m.projects = value
}
// SetPublications sets the publications property value. Represents details of any publications a user has added to their profile.
func (m *Profile) SetPublications(value []ItemPublicationable)() {
    m.publications = value
}
// SetSkills sets the skills property value. Represents detailed information about skills associated with a user in various services.
func (m *Profile) SetSkills(value []SkillProficiencyable)() {
    m.skills = value
}
// SetWebAccounts sets the webAccounts property value. Represents web accounts the user has indicated they use or has added to their user profile.
func (m *Profile) SetWebAccounts(value []WebAccountable)() {
    m.webAccounts = value
}
// SetWebsites sets the websites property value. Represents detailed information about websites associated with a user in various services.
func (m *Profile) SetWebsites(value []PersonWebsiteable)() {
    m.websites = value
}
