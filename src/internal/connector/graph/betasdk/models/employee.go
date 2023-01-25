package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Employee provides operations to call the add method.
type Employee struct {
    Entity
    // The address property
    address PostalAddressTypeable
    // The birthDate property
    birthDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The displayName property
    displayName *string
    // The email property
    email *string
    // The employmentDate property
    employmentDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The givenName property
    givenName *string
    // The jobTitle property
    jobTitle *string
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The middleName property
    middleName *string
    // The mobilePhone property
    mobilePhone *string
    // The number property
    number *string
    // The personalEmail property
    personalEmail *string
    // The phoneNumber property
    phoneNumber *string
    // The picture property
    picture []Pictureable
    // The statisticsGroupCode property
    statisticsGroupCode *string
    // The status property
    status *string
    // The surname property
    surname *string
    // The terminationDate property
    terminationDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
}
// NewEmployee instantiates a new employee and sets the default values.
func NewEmployee()(*Employee) {
    m := &Employee{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEmployeeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEmployeeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEmployee(), nil
}
// GetAddress gets the address property value. The address property
func (m *Employee) GetAddress()(PostalAddressTypeable) {
    return m.address
}
// GetBirthDate gets the birthDate property value. The birthDate property
func (m *Employee) GetBirthDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.birthDate
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *Employee) GetDisplayName()(*string) {
    return m.displayName
}
// GetEmail gets the email property value. The email property
func (m *Employee) GetEmail()(*string) {
    return m.email
}
// GetEmploymentDate gets the employmentDate property value. The employmentDate property
func (m *Employee) GetEmploymentDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.employmentDate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Employee) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePostalAddressTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val.(PostalAddressTypeable))
        }
        return nil
    }
    res["birthDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBirthDate(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["employmentDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmploymentDate(val)
        }
        return nil
    }
    res["givenName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGivenName(val)
        }
        return nil
    }
    res["jobTitle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobTitle(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["middleName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMiddleName(val)
        }
        return nil
    }
    res["mobilePhone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobilePhone(val)
        }
        return nil
    }
    res["number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumber(val)
        }
        return nil
    }
    res["personalEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersonalEmail(val)
        }
        return nil
    }
    res["phoneNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneNumber(val)
        }
        return nil
    }
    res["picture"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePictureFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Pictureable, len(val))
            for i, v := range val {
                res[i] = v.(Pictureable)
            }
            m.SetPicture(res)
        }
        return nil
    }
    res["statisticsGroupCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatisticsGroupCode(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["surname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSurname(val)
        }
        return nil
    }
    res["terminationDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTerminationDate(val)
        }
        return nil
    }
    return res
}
// GetGivenName gets the givenName property value. The givenName property
func (m *Employee) GetGivenName()(*string) {
    return m.givenName
}
// GetJobTitle gets the jobTitle property value. The jobTitle property
func (m *Employee) GetJobTitle()(*string) {
    return m.jobTitle
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Employee) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetMiddleName gets the middleName property value. The middleName property
func (m *Employee) GetMiddleName()(*string) {
    return m.middleName
}
// GetMobilePhone gets the mobilePhone property value. The mobilePhone property
func (m *Employee) GetMobilePhone()(*string) {
    return m.mobilePhone
}
// GetNumber gets the number property value. The number property
func (m *Employee) GetNumber()(*string) {
    return m.number
}
// GetPersonalEmail gets the personalEmail property value. The personalEmail property
func (m *Employee) GetPersonalEmail()(*string) {
    return m.personalEmail
}
// GetPhoneNumber gets the phoneNumber property value. The phoneNumber property
func (m *Employee) GetPhoneNumber()(*string) {
    return m.phoneNumber
}
// GetPicture gets the picture property value. The picture property
func (m *Employee) GetPicture()([]Pictureable) {
    return m.picture
}
// GetStatisticsGroupCode gets the statisticsGroupCode property value. The statisticsGroupCode property
func (m *Employee) GetStatisticsGroupCode()(*string) {
    return m.statisticsGroupCode
}
// GetStatus gets the status property value. The status property
func (m *Employee) GetStatus()(*string) {
    return m.status
}
// GetSurname gets the surname property value. The surname property
func (m *Employee) GetSurname()(*string) {
    return m.surname
}
// GetTerminationDate gets the terminationDate property value. The terminationDate property
func (m *Employee) GetTerminationDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.terminationDate
}
// Serialize serializes information the current object
func (m *Employee) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("birthDate", m.GetBirthDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("employmentDate", m.GetEmploymentDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("givenName", m.GetGivenName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("jobTitle", m.GetJobTitle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("middleName", m.GetMiddleName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mobilePhone", m.GetMobilePhone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("number", m.GetNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("personalEmail", m.GetPersonalEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("phoneNumber", m.GetPhoneNumber())
        if err != nil {
            return err
        }
    }
    if m.GetPicture() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPicture()))
        for i, v := range m.GetPicture() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("picture", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("statisticsGroupCode", m.GetStatisticsGroupCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("surname", m.GetSurname())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("terminationDate", m.GetTerminationDate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddress sets the address property value. The address property
func (m *Employee) SetAddress(value PostalAddressTypeable)() {
    m.address = value
}
// SetBirthDate sets the birthDate property value. The birthDate property
func (m *Employee) SetBirthDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.birthDate = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *Employee) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEmail sets the email property value. The email property
func (m *Employee) SetEmail(value *string)() {
    m.email = value
}
// SetEmploymentDate sets the employmentDate property value. The employmentDate property
func (m *Employee) SetEmploymentDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.employmentDate = value
}
// SetGivenName sets the givenName property value. The givenName property
func (m *Employee) SetGivenName(value *string)() {
    m.givenName = value
}
// SetJobTitle sets the jobTitle property value. The jobTitle property
func (m *Employee) SetJobTitle(value *string)() {
    m.jobTitle = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Employee) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetMiddleName sets the middleName property value. The middleName property
func (m *Employee) SetMiddleName(value *string)() {
    m.middleName = value
}
// SetMobilePhone sets the mobilePhone property value. The mobilePhone property
func (m *Employee) SetMobilePhone(value *string)() {
    m.mobilePhone = value
}
// SetNumber sets the number property value. The number property
func (m *Employee) SetNumber(value *string)() {
    m.number = value
}
// SetPersonalEmail sets the personalEmail property value. The personalEmail property
func (m *Employee) SetPersonalEmail(value *string)() {
    m.personalEmail = value
}
// SetPhoneNumber sets the phoneNumber property value. The phoneNumber property
func (m *Employee) SetPhoneNumber(value *string)() {
    m.phoneNumber = value
}
// SetPicture sets the picture property value. The picture property
func (m *Employee) SetPicture(value []Pictureable)() {
    m.picture = value
}
// SetStatisticsGroupCode sets the statisticsGroupCode property value. The statisticsGroupCode property
func (m *Employee) SetStatisticsGroupCode(value *string)() {
    m.statisticsGroupCode = value
}
// SetStatus sets the status property value. The status property
func (m *Employee) SetStatus(value *string)() {
    m.status = value
}
// SetSurname sets the surname property value. The surname property
func (m *Employee) SetSurname(value *string)() {
    m.surname = value
}
// SetTerminationDate sets the terminationDate property value. The terminationDate property
func (m *Employee) SetTerminationDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.terminationDate = value
}
