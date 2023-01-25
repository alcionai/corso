package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSPrivacyAccessControlItem represents per-process privacy preferences.
type MacOSPrivacyAccessControlItem struct {
    // Possible values of a property
    accessibility *Enablement
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Possible values of a property
    addressBook *Enablement
    // Allow or deny the app or process to send a restricted Apple event to another app or process. You will need to know the identifier, identifier type, and code requirement of the receiving app or process. This collection can contain a maximum of 500 elements.
    appleEventsAllowedReceivers []MacOSAppleEventReceiverable
    // Block access to camera app.
    blockCamera *bool
    // Block the app or process from listening to events from input devices such as mouse, keyboard, and trackpad.Requires macOS 10.15 or later.
    blockListenEvent *bool
    // Block access to microphone.
    blockMicrophone *bool
    // Block app from capturing contents of system display. Requires macOS 10.15 or later.
    blockScreenCapture *bool
    // Possible values of a property
    calendar *Enablement
    // Enter the code requirement, which can be obtained with the command 'codesign –display -r –' in the Terminal app. Include everything after '=>'.
    codeRequirement *string
    // The display name of the app, process, or executable.
    displayName *string
    // Possible values of a property
    fileProviderPresence *Enablement
    // The bundle ID or path of the app, process, or executable.
    identifier *string
    // Process identifier types for MacOS Privacy Preferences
    identifierType *MacOSProcessIdentifierType
    // Possible values of a property
    mediaLibrary *Enablement
    // The OdataType property
    odataType *string
    // Possible values of a property
    photos *Enablement
    // Possible values of a property
    postEvent *Enablement
    // Possible values of a property
    reminders *Enablement
    // Possible values of a property
    speechRecognition *Enablement
    // Statically validates the code requirement. Use this setting if the process invalidates its dynamic code signature.
    staticCodeValidation *bool
    // Possible values of a property
    systemPolicyAllFiles *Enablement
    // Possible values of a property
    systemPolicyDesktopFolder *Enablement
    // Possible values of a property
    systemPolicyDocumentsFolder *Enablement
    // Possible values of a property
    systemPolicyDownloadsFolder *Enablement
    // Possible values of a property
    systemPolicyNetworkVolumes *Enablement
    // Possible values of a property
    systemPolicyRemovableVolumes *Enablement
    // Possible values of a property
    systemPolicySystemAdminFiles *Enablement
}
// NewMacOSPrivacyAccessControlItem instantiates a new macOSPrivacyAccessControlItem and sets the default values.
func NewMacOSPrivacyAccessControlItem()(*MacOSPrivacyAccessControlItem) {
    m := &MacOSPrivacyAccessControlItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMacOSPrivacyAccessControlItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSPrivacyAccessControlItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSPrivacyAccessControlItem(), nil
}
// GetAccessibility gets the accessibility property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) GetAccessibility()(*Enablement) {
    return m.accessibility
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MacOSPrivacyAccessControlItem) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAddressBook gets the addressBook property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) GetAddressBook()(*Enablement) {
    return m.addressBook
}
// GetAppleEventsAllowedReceivers gets the appleEventsAllowedReceivers property value. Allow or deny the app or process to send a restricted Apple event to another app or process. You will need to know the identifier, identifier type, and code requirement of the receiving app or process. This collection can contain a maximum of 500 elements.
func (m *MacOSPrivacyAccessControlItem) GetAppleEventsAllowedReceivers()([]MacOSAppleEventReceiverable) {
    return m.appleEventsAllowedReceivers
}
// GetBlockCamera gets the blockCamera property value. Block access to camera app.
func (m *MacOSPrivacyAccessControlItem) GetBlockCamera()(*bool) {
    return m.blockCamera
}
// GetBlockListenEvent gets the blockListenEvent property value. Block the app or process from listening to events from input devices such as mouse, keyboard, and trackpad.Requires macOS 10.15 or later.
func (m *MacOSPrivacyAccessControlItem) GetBlockListenEvent()(*bool) {
    return m.blockListenEvent
}
// GetBlockMicrophone gets the blockMicrophone property value. Block access to microphone.
func (m *MacOSPrivacyAccessControlItem) GetBlockMicrophone()(*bool) {
    return m.blockMicrophone
}
// GetBlockScreenCapture gets the blockScreenCapture property value. Block app from capturing contents of system display. Requires macOS 10.15 or later.
func (m *MacOSPrivacyAccessControlItem) GetBlockScreenCapture()(*bool) {
    return m.blockScreenCapture
}
// GetCalendar gets the calendar property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) GetCalendar()(*Enablement) {
    return m.calendar
}
// GetCodeRequirement gets the codeRequirement property value. Enter the code requirement, which can be obtained with the command 'codesign –display -r –' in the Terminal app. Include everything after '=>'.
func (m *MacOSPrivacyAccessControlItem) GetCodeRequirement()(*string) {
    return m.codeRequirement
}
// GetDisplayName gets the displayName property value. The display name of the app, process, or executable.
func (m *MacOSPrivacyAccessControlItem) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSPrivacyAccessControlItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accessibility"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessibility(val.(*Enablement))
        }
        return nil
    }
    res["addressBook"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddressBook(val.(*Enablement))
        }
        return nil
    }
    res["appleEventsAllowedReceivers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMacOSAppleEventReceiverFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MacOSAppleEventReceiverable, len(val))
            for i, v := range val {
                res[i] = v.(MacOSAppleEventReceiverable)
            }
            m.SetAppleEventsAllowedReceivers(res)
        }
        return nil
    }
    res["blockCamera"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockCamera(val)
        }
        return nil
    }
    res["blockListenEvent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockListenEvent(val)
        }
        return nil
    }
    res["blockMicrophone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockMicrophone(val)
        }
        return nil
    }
    res["blockScreenCapture"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockScreenCapture(val)
        }
        return nil
    }
    res["calendar"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalendar(val.(*Enablement))
        }
        return nil
    }
    res["codeRequirement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCodeRequirement(val)
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
    res["fileProviderPresence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileProviderPresence(val.(*Enablement))
        }
        return nil
    }
    res["identifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentifier(val)
        }
        return nil
    }
    res["identifierType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMacOSProcessIdentifierType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentifierType(val.(*MacOSProcessIdentifierType))
        }
        return nil
    }
    res["mediaLibrary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaLibrary(val.(*Enablement))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["photos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhotos(val.(*Enablement))
        }
        return nil
    }
    res["postEvent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostEvent(val.(*Enablement))
        }
        return nil
    }
    res["reminders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReminders(val.(*Enablement))
        }
        return nil
    }
    res["speechRecognition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpeechRecognition(val.(*Enablement))
        }
        return nil
    }
    res["staticCodeValidation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStaticCodeValidation(val)
        }
        return nil
    }
    res["systemPolicyAllFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemPolicyAllFiles(val.(*Enablement))
        }
        return nil
    }
    res["systemPolicyDesktopFolder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemPolicyDesktopFolder(val.(*Enablement))
        }
        return nil
    }
    res["systemPolicyDocumentsFolder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemPolicyDocumentsFolder(val.(*Enablement))
        }
        return nil
    }
    res["systemPolicyDownloadsFolder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemPolicyDownloadsFolder(val.(*Enablement))
        }
        return nil
    }
    res["systemPolicyNetworkVolumes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemPolicyNetworkVolumes(val.(*Enablement))
        }
        return nil
    }
    res["systemPolicyRemovableVolumes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemPolicyRemovableVolumes(val.(*Enablement))
        }
        return nil
    }
    res["systemPolicySystemAdminFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemPolicySystemAdminFiles(val.(*Enablement))
        }
        return nil
    }
    return res
}
// GetFileProviderPresence gets the fileProviderPresence property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) GetFileProviderPresence()(*Enablement) {
    return m.fileProviderPresence
}
// GetIdentifier gets the identifier property value. The bundle ID or path of the app, process, or executable.
func (m *MacOSPrivacyAccessControlItem) GetIdentifier()(*string) {
    return m.identifier
}
// GetIdentifierType gets the identifierType property value. Process identifier types for MacOS Privacy Preferences
func (m *MacOSPrivacyAccessControlItem) GetIdentifierType()(*MacOSProcessIdentifierType) {
    return m.identifierType
}
// GetMediaLibrary gets the mediaLibrary property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) GetMediaLibrary()(*Enablement) {
    return m.mediaLibrary
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MacOSPrivacyAccessControlItem) GetOdataType()(*string) {
    return m.odataType
}
// GetPhotos gets the photos property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) GetPhotos()(*Enablement) {
    return m.photos
}
// GetPostEvent gets the postEvent property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) GetPostEvent()(*Enablement) {
    return m.postEvent
}
// GetReminders gets the reminders property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) GetReminders()(*Enablement) {
    return m.reminders
}
// GetSpeechRecognition gets the speechRecognition property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) GetSpeechRecognition()(*Enablement) {
    return m.speechRecognition
}
// GetStaticCodeValidation gets the staticCodeValidation property value. Statically validates the code requirement. Use this setting if the process invalidates its dynamic code signature.
func (m *MacOSPrivacyAccessControlItem) GetStaticCodeValidation()(*bool) {
    return m.staticCodeValidation
}
// GetSystemPolicyAllFiles gets the systemPolicyAllFiles property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) GetSystemPolicyAllFiles()(*Enablement) {
    return m.systemPolicyAllFiles
}
// GetSystemPolicyDesktopFolder gets the systemPolicyDesktopFolder property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) GetSystemPolicyDesktopFolder()(*Enablement) {
    return m.systemPolicyDesktopFolder
}
// GetSystemPolicyDocumentsFolder gets the systemPolicyDocumentsFolder property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) GetSystemPolicyDocumentsFolder()(*Enablement) {
    return m.systemPolicyDocumentsFolder
}
// GetSystemPolicyDownloadsFolder gets the systemPolicyDownloadsFolder property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) GetSystemPolicyDownloadsFolder()(*Enablement) {
    return m.systemPolicyDownloadsFolder
}
// GetSystemPolicyNetworkVolumes gets the systemPolicyNetworkVolumes property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) GetSystemPolicyNetworkVolumes()(*Enablement) {
    return m.systemPolicyNetworkVolumes
}
// GetSystemPolicyRemovableVolumes gets the systemPolicyRemovableVolumes property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) GetSystemPolicyRemovableVolumes()(*Enablement) {
    return m.systemPolicyRemovableVolumes
}
// GetSystemPolicySystemAdminFiles gets the systemPolicySystemAdminFiles property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) GetSystemPolicySystemAdminFiles()(*Enablement) {
    return m.systemPolicySystemAdminFiles
}
// Serialize serializes information the current object
func (m *MacOSPrivacyAccessControlItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAccessibility() != nil {
        cast := (*m.GetAccessibility()).String()
        err := writer.WriteStringValue("accessibility", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAddressBook() != nil {
        cast := (*m.GetAddressBook()).String()
        err := writer.WriteStringValue("addressBook", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppleEventsAllowedReceivers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppleEventsAllowedReceivers()))
        for i, v := range m.GetAppleEventsAllowedReceivers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("appleEventsAllowedReceivers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("blockCamera", m.GetBlockCamera())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("blockListenEvent", m.GetBlockListenEvent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("blockMicrophone", m.GetBlockMicrophone())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("blockScreenCapture", m.GetBlockScreenCapture())
        if err != nil {
            return err
        }
    }
    if m.GetCalendar() != nil {
        cast := (*m.GetCalendar()).String()
        err := writer.WriteStringValue("calendar", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("codeRequirement", m.GetCodeRequirement())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetFileProviderPresence() != nil {
        cast := (*m.GetFileProviderPresence()).String()
        err := writer.WriteStringValue("fileProviderPresence", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("identifier", m.GetIdentifier())
        if err != nil {
            return err
        }
    }
    if m.GetIdentifierType() != nil {
        cast := (*m.GetIdentifierType()).String()
        err := writer.WriteStringValue("identifierType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMediaLibrary() != nil {
        cast := (*m.GetMediaLibrary()).String()
        err := writer.WriteStringValue("mediaLibrary", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetPhotos() != nil {
        cast := (*m.GetPhotos()).String()
        err := writer.WriteStringValue("photos", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPostEvent() != nil {
        cast := (*m.GetPostEvent()).String()
        err := writer.WriteStringValue("postEvent", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetReminders() != nil {
        cast := (*m.GetReminders()).String()
        err := writer.WriteStringValue("reminders", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSpeechRecognition() != nil {
        cast := (*m.GetSpeechRecognition()).String()
        err := writer.WriteStringValue("speechRecognition", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("staticCodeValidation", m.GetStaticCodeValidation())
        if err != nil {
            return err
        }
    }
    if m.GetSystemPolicyAllFiles() != nil {
        cast := (*m.GetSystemPolicyAllFiles()).String()
        err := writer.WriteStringValue("systemPolicyAllFiles", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSystemPolicyDesktopFolder() != nil {
        cast := (*m.GetSystemPolicyDesktopFolder()).String()
        err := writer.WriteStringValue("systemPolicyDesktopFolder", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSystemPolicyDocumentsFolder() != nil {
        cast := (*m.GetSystemPolicyDocumentsFolder()).String()
        err := writer.WriteStringValue("systemPolicyDocumentsFolder", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSystemPolicyDownloadsFolder() != nil {
        cast := (*m.GetSystemPolicyDownloadsFolder()).String()
        err := writer.WriteStringValue("systemPolicyDownloadsFolder", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSystemPolicyNetworkVolumes() != nil {
        cast := (*m.GetSystemPolicyNetworkVolumes()).String()
        err := writer.WriteStringValue("systemPolicyNetworkVolumes", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSystemPolicyRemovableVolumes() != nil {
        cast := (*m.GetSystemPolicyRemovableVolumes()).String()
        err := writer.WriteStringValue("systemPolicyRemovableVolumes", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSystemPolicySystemAdminFiles() != nil {
        cast := (*m.GetSystemPolicySystemAdminFiles()).String()
        err := writer.WriteStringValue("systemPolicySystemAdminFiles", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessibility sets the accessibility property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetAccessibility(value *Enablement)() {
    m.accessibility = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MacOSPrivacyAccessControlItem) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAddressBook sets the addressBook property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetAddressBook(value *Enablement)() {
    m.addressBook = value
}
// SetAppleEventsAllowedReceivers sets the appleEventsAllowedReceivers property value. Allow or deny the app or process to send a restricted Apple event to another app or process. You will need to know the identifier, identifier type, and code requirement of the receiving app or process. This collection can contain a maximum of 500 elements.
func (m *MacOSPrivacyAccessControlItem) SetAppleEventsAllowedReceivers(value []MacOSAppleEventReceiverable)() {
    m.appleEventsAllowedReceivers = value
}
// SetBlockCamera sets the blockCamera property value. Block access to camera app.
func (m *MacOSPrivacyAccessControlItem) SetBlockCamera(value *bool)() {
    m.blockCamera = value
}
// SetBlockListenEvent sets the blockListenEvent property value. Block the app or process from listening to events from input devices such as mouse, keyboard, and trackpad.Requires macOS 10.15 or later.
func (m *MacOSPrivacyAccessControlItem) SetBlockListenEvent(value *bool)() {
    m.blockListenEvent = value
}
// SetBlockMicrophone sets the blockMicrophone property value. Block access to microphone.
func (m *MacOSPrivacyAccessControlItem) SetBlockMicrophone(value *bool)() {
    m.blockMicrophone = value
}
// SetBlockScreenCapture sets the blockScreenCapture property value. Block app from capturing contents of system display. Requires macOS 10.15 or later.
func (m *MacOSPrivacyAccessControlItem) SetBlockScreenCapture(value *bool)() {
    m.blockScreenCapture = value
}
// SetCalendar sets the calendar property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetCalendar(value *Enablement)() {
    m.calendar = value
}
// SetCodeRequirement sets the codeRequirement property value. Enter the code requirement, which can be obtained with the command 'codesign –display -r –' in the Terminal app. Include everything after '=>'.
func (m *MacOSPrivacyAccessControlItem) SetCodeRequirement(value *string)() {
    m.codeRequirement = value
}
// SetDisplayName sets the displayName property value. The display name of the app, process, or executable.
func (m *MacOSPrivacyAccessControlItem) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetFileProviderPresence sets the fileProviderPresence property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetFileProviderPresence(value *Enablement)() {
    m.fileProviderPresence = value
}
// SetIdentifier sets the identifier property value. The bundle ID or path of the app, process, or executable.
func (m *MacOSPrivacyAccessControlItem) SetIdentifier(value *string)() {
    m.identifier = value
}
// SetIdentifierType sets the identifierType property value. Process identifier types for MacOS Privacy Preferences
func (m *MacOSPrivacyAccessControlItem) SetIdentifierType(value *MacOSProcessIdentifierType)() {
    m.identifierType = value
}
// SetMediaLibrary sets the mediaLibrary property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetMediaLibrary(value *Enablement)() {
    m.mediaLibrary = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MacOSPrivacyAccessControlItem) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPhotos sets the photos property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetPhotos(value *Enablement)() {
    m.photos = value
}
// SetPostEvent sets the postEvent property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetPostEvent(value *Enablement)() {
    m.postEvent = value
}
// SetReminders sets the reminders property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetReminders(value *Enablement)() {
    m.reminders = value
}
// SetSpeechRecognition sets the speechRecognition property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetSpeechRecognition(value *Enablement)() {
    m.speechRecognition = value
}
// SetStaticCodeValidation sets the staticCodeValidation property value. Statically validates the code requirement. Use this setting if the process invalidates its dynamic code signature.
func (m *MacOSPrivacyAccessControlItem) SetStaticCodeValidation(value *bool)() {
    m.staticCodeValidation = value
}
// SetSystemPolicyAllFiles sets the systemPolicyAllFiles property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetSystemPolicyAllFiles(value *Enablement)() {
    m.systemPolicyAllFiles = value
}
// SetSystemPolicyDesktopFolder sets the systemPolicyDesktopFolder property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetSystemPolicyDesktopFolder(value *Enablement)() {
    m.systemPolicyDesktopFolder = value
}
// SetSystemPolicyDocumentsFolder sets the systemPolicyDocumentsFolder property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetSystemPolicyDocumentsFolder(value *Enablement)() {
    m.systemPolicyDocumentsFolder = value
}
// SetSystemPolicyDownloadsFolder sets the systemPolicyDownloadsFolder property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetSystemPolicyDownloadsFolder(value *Enablement)() {
    m.systemPolicyDownloadsFolder = value
}
// SetSystemPolicyNetworkVolumes sets the systemPolicyNetworkVolumes property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetSystemPolicyNetworkVolumes(value *Enablement)() {
    m.systemPolicyNetworkVolumes = value
}
// SetSystemPolicyRemovableVolumes sets the systemPolicyRemovableVolumes property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetSystemPolicyRemovableVolumes(value *Enablement)() {
    m.systemPolicyRemovableVolumes = value
}
// SetSystemPolicySystemAdminFiles sets the systemPolicySystemAdminFiles property value. Possible values of a property
func (m *MacOSPrivacyAccessControlItem) SetSystemPolicySystemAdminFiles(value *Enablement)() {
    m.systemPolicySystemAdminFiles = value
}
