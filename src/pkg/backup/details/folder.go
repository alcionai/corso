package details

import "time"

type FolderInfo struct {
	ItemType    ItemType  `json:"itemType,omitempty"`
	DisplayName string    `json:"displayName"`
	Modified    time.Time `json:"modified,omitempty"`
	Size        int64     `json:"size,omitempty"`
	DataType    ItemType  `json:"dataType,omitempty"`
	DriveName   string    `json:"driveName,omitempty"`
	DriveID     string    `json:"driveID,omitempty"`
}

func (i FolderInfo) Headers() []string {
	return []string{"Display Name"}
}

func (i FolderInfo) Values() []string {
	return []string{i.DisplayName}
}
