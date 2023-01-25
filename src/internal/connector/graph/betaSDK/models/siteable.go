package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Siteable 
type Siteable interface {
    BaseItemable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnalytics()(ItemAnalyticsable)
    GetColumns()([]ColumnDefinitionable)
    GetContentTypes()([]ContentTypeable)
    GetDeleted()(Deletedable)
    GetDisplayName()(*string)
    GetDrive()(Driveable)
    GetDrives()([]Driveable)
    GetExternalColumns()([]ColumnDefinitionable)
    GetInformationProtection()(InformationProtectionable)
    GetItems()([]BaseItemable)
    GetLists()([]Listable)
    GetOnenote()(Onenoteable)
    GetOperations()([]RichLongRunningOperationable)
    GetPages()([]SitePageable)
    GetPermissions()([]Permissionable)
    GetRoot()(Rootable)
    GetSettings()(SiteSettingsable)
    GetSharepointIds()(SharepointIdsable)
    GetSiteCollection()(SiteCollectionable)
    GetSites()([]Siteable)
    SetAnalytics(value ItemAnalyticsable)()
    SetColumns(value []ColumnDefinitionable)()
    SetContentTypes(value []ContentTypeable)()
    SetDeleted(value Deletedable)()
    SetDisplayName(value *string)()
    SetDrive(value Driveable)()
    SetDrives(value []Driveable)()
    SetExternalColumns(value []ColumnDefinitionable)()
    SetInformationProtection(value InformationProtectionable)()
    SetItems(value []BaseItemable)()
    SetLists(value []Listable)()
    SetOnenote(value Onenoteable)()
    SetOperations(value []RichLongRunningOperationable)()
    SetPages(value []SitePageable)()
    SetPermissions(value []Permissionable)()
    SetRoot(value Rootable)()
    SetSettings(value SiteSettingsable)()
    SetSharepointIds(value SharepointIdsable)()
    SetSiteCollection(value SiteCollectionable)()
    SetSites(value []Siteable)()
}
