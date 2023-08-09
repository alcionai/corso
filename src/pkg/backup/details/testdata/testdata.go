package testdata

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
)

// mustParsePath takes a string representing a resource path and returns a path
// instance. Panics if the path cannot be parsed. Useful for simple variable
// assignments.
func mustParsePath(ref string, isItem bool) path.Path {
	p, err := path.FromDataLayerPath(ref, isItem)
	if err != nil {
		panic(err)
	}

	return p
}

// mustAppendPath takes a Path, string representing a path element, and whether
// the element is an item and returns a path instance representing the original
// path with the element appended to it. Panics if the path cannot be parsed.
// Useful for simple variable assignments.
func mustAppendPath(p path.Path, newElement string, isItem bool) path.Path {
	newP, err := p.Append(isItem, newElement)
	if err != nil {
		panic(err)
	}

	return newP
}

func locFromRepo(rr path.Path, isItem bool) *path.Builder {
	loc := &path.Builder{}

	for _, e := range rr.Folders() {
		loc = loc.Append(strings.TrimSuffix(e, folderSuffix))
	}

	if rr.Service() == path.OneDriveService || rr.Category() == path.LibrariesCategory {
		loc = loc.PopFront()
	}

	// Folders don't have their final element in the location.
	if !isItem {
		loc = loc.Dir()
	}

	return loc
}

type repoRefAndLocRef struct {
	RR  path.Path
	Loc *path.Builder
}

func (p repoRefAndLocRef) MustAppend(newElement string, isItem bool) repoRefAndLocRef {
	e := newElement + folderSuffix

	if isItem {
		e = newElement + fileSuffix
	}

	res := repoRefAndLocRef{
		RR: mustAppendPath(p.RR, e, isItem),
	}

	res.Loc = locFromRepo(res.RR, isItem)

	return res
}

func (p repoRefAndLocRef) ItemLocation() string {
	return strings.TrimSuffix(p.RR.Item(), fileSuffix)
}

func (p repoRefAndLocRef) FolderLocation() string {
	lastElem := p.RR.ToBuilder().LastElem()

	if len(p.RR.Item()) > 0 {
		f := p.RR.Folders()
		lastElem = f[len(f)-2]
	}

	return p.Loc.Append(strings.TrimSuffix(lastElem, folderSuffix)).String()
}

// locationAsRepoRef returns a path.Path where the LocationRef is used for the
// folder path instead of the id-based path elements. This is useful for
// generating paths for older versions of Corso.
func (p repoRefAndLocRef) locationAsRepoRef() path.Path {
	tmp := p.Loc
	if len(p.ItemLocation()) > 0 {
		tmp = tmp.Append(p.ItemLocation())
	}

	res, err := tmp.ToDataLayerPath(
		p.RR.Tenant(),
		p.RR.ProtectedResource(),
		p.RR.Service(),
		p.RR.Category(),
		len(p.ItemLocation()) > 0)
	if err != nil {
		panic(err)
	}

	return res
}

func mustPathRep(ref string, isItem bool) repoRefAndLocRef {
	res := repoRefAndLocRef{}
	tmp := mustParsePath(ref, isItem)

	// Now append stuff to the RepoRef elements so we have distinct LocationRef
	// and RepoRef elements to simulate using IDs in the path instead of display
	// names.
	rrPB := &path.Builder{}
	for _, e := range tmp.Folders() {
		rrPB = rrPB.Append(e + folderSuffix)
	}

	if isItem {
		rrPB = rrPB.Append(tmp.Item() + fileSuffix)
	}

	rr, err := rrPB.ToDataLayerPath(
		tmp.Tenant(),
		tmp.ProtectedResource(),
		tmp.Service(),
		tmp.Category(),
		isItem)
	if err != nil {
		panic(err)
	}

	res.RR = rr
	res.Loc = locFromRepo(rr, isItem)

	return res
}

const (
	folderSuffix = ".d"
	fileSuffix   = ".f"

	ItemName1  = "item1"
	ItemName2  = "item2"
	ItemName3  = "item3"
	UserEmail1 = "user1@email.com"
	UserEmail2 = "user2@email.com"
)

var (
	Time1 = time.Date(2022, 9, 21, 10, 0, 0, 0, time.UTC)
	Time2 = time.Date(2022, 10, 21, 10, 0, 0, 0, time.UTC)
	Time3 = time.Date(2023, 9, 21, 10, 0, 0, 0, time.UTC)
	Time4 = time.Date(2023, 10, 21, 10, 0, 0, 0, time.UTC)

	ExchangeEmailInboxPath = mustPathRep("tenant-id/exchange/user-id/email/Inbox", false)
	ExchangeEmailBasePath  = ExchangeEmailInboxPath.MustAppend("subfolder", false)
	ExchangeEmailBasePath2 = ExchangeEmailInboxPath.MustAppend("othersubfolder/", false)
	ExchangeEmailBasePath3 = ExchangeEmailBasePath2.MustAppend("subsubfolder", false)
	ExchangeEmailItemPath1 = ExchangeEmailBasePath.MustAppend(ItemName1, true)
	ExchangeEmailItemPath2 = ExchangeEmailBasePath2.MustAppend(ItemName2, true)
	ExchangeEmailItemPath3 = ExchangeEmailBasePath3.MustAppend(ItemName3, true)

	// These all represent the same set of items however, the different versions
	// have varying amounts of information.
	exchangeEmailItemsByVersion = map[int][]details.Entry{
		version.All8MigrateUserPNToID: {
			{
				RepoRef:     ExchangeEmailItemPath1.RR.String(),
				ShortRef:    ExchangeEmailItemPath1.RR.ShortRef(),
				ParentRef:   ExchangeEmailItemPath1.RR.ToBuilder().Dir().ShortRef(),
				ItemRef:     ExchangeEmailItemPath1.ItemLocation(),
				LocationRef: ExchangeEmailItemPath1.Loc.String(),
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						ItemType: details.ExchangeMail,
						Sender:   "a-person",
						Subject:  "foo",
						Received: Time1,
					},
				},
			},
			{
				RepoRef:     ExchangeEmailItemPath2.RR.String(),
				ShortRef:    ExchangeEmailItemPath2.RR.ShortRef(),
				ParentRef:   ExchangeEmailItemPath2.RR.ToBuilder().Dir().ShortRef(),
				ItemRef:     ExchangeEmailItemPath2.ItemLocation(),
				LocationRef: ExchangeEmailItemPath2.Loc.String(),
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						ItemType: details.ExchangeMail,
						Sender:   "a-person",
						Subject:  "bar",
						Received: Time2,
					},
				},
			},
			{
				RepoRef:     ExchangeEmailItemPath3.RR.String(),
				ShortRef:    ExchangeEmailItemPath3.RR.ShortRef(),
				ParentRef:   ExchangeEmailItemPath3.RR.ToBuilder().Dir().ShortRef(),
				ItemRef:     ExchangeEmailItemPath3.ItemLocation(),
				LocationRef: ExchangeEmailItemPath3.Loc.String(),
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						ItemType: details.ExchangeMail,
						Sender:   "another-person",
						Subject:  "baz",
						Received: Time2,
					},
				},
			},
		},
		version.OneDrive7LocationRef: {
			{
				RepoRef:     ExchangeEmailItemPath1.locationAsRepoRef().String(),
				ShortRef:    ExchangeEmailItemPath1.locationAsRepoRef().ShortRef(),
				ParentRef:   ExchangeEmailItemPath1.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemRef:     ExchangeEmailItemPath1.ItemLocation(),
				LocationRef: ExchangeEmailItemPath1.Loc.String(),
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						ItemType: details.ExchangeMail,
						Sender:   "a-person",
						Subject:  "foo",
						Received: Time1,
					},
				},
			},
			{
				RepoRef:     ExchangeEmailItemPath2.locationAsRepoRef().String(),
				ShortRef:    ExchangeEmailItemPath2.locationAsRepoRef().ShortRef(),
				ParentRef:   ExchangeEmailItemPath2.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemRef:     ExchangeEmailItemPath2.ItemLocation(),
				LocationRef: ExchangeEmailItemPath2.Loc.String(),
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						ItemType: details.ExchangeMail,
						Sender:   "a-person",
						Subject:  "bar",
						Received: Time2,
					},
				},
			},
			{
				RepoRef:     ExchangeEmailItemPath3.locationAsRepoRef().String(),
				ShortRef:    ExchangeEmailItemPath3.locationAsRepoRef().ShortRef(),
				ParentRef:   ExchangeEmailItemPath3.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemRef:     ExchangeEmailItemPath3.ItemLocation(),
				LocationRef: ExchangeEmailItemPath3.Loc.String(),
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						ItemType: details.ExchangeMail,
						Sender:   "another-person",
						Subject:  "baz",
						Received: Time2,
					},
				},
			},
		},
		0: {
			{
				RepoRef:   ExchangeEmailItemPath1.locationAsRepoRef().String(),
				ShortRef:  ExchangeEmailItemPath1.locationAsRepoRef().ShortRef(),
				ParentRef: ExchangeEmailItemPath1.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						ItemType: details.ExchangeMail,
						Sender:   "a-person",
						Subject:  "foo",
						Received: Time1,
					},
				},
			},
			{
				RepoRef:   ExchangeEmailItemPath2.locationAsRepoRef().String(),
				ShortRef:  ExchangeEmailItemPath2.locationAsRepoRef().ShortRef(),
				ParentRef: ExchangeEmailItemPath2.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						ItemType: details.ExchangeMail,
						Sender:   "a-person",
						Subject:  "bar",
						Received: Time2,
					},
				},
			},
			{
				RepoRef:   ExchangeEmailItemPath3.locationAsRepoRef().String(),
				ShortRef:  ExchangeEmailItemPath3.locationAsRepoRef().ShortRef(),
				ParentRef: ExchangeEmailItemPath3.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						ItemType: details.ExchangeMail,
						Sender:   "another-person",
						Subject:  "baz",
						Received: Time2,
					},
				},
			},
		},
	}

	ExchangeContactsRootPath  = mustPathRep("tenant-id/exchange/user-id/contacts/contacts", false)
	ExchangeContactsBasePath  = ExchangeContactsRootPath.MustAppend("contacts", false)
	ExchangeContactsBasePath2 = ExchangeContactsRootPath.MustAppend("morecontacts", false)
	ExchangeContactsItemPath1 = ExchangeContactsBasePath.MustAppend(ItemName1, true)
	ExchangeContactsItemPath2 = ExchangeContactsBasePath2.MustAppend(ItemName2, true)

	exchangeContactsItemsByVersion = map[int][]details.Entry{
		version.All8MigrateUserPNToID: {
			{
				RepoRef:     ExchangeContactsItemPath1.RR.String(),
				ShortRef:    ExchangeContactsItemPath1.RR.ShortRef(),
				ParentRef:   ExchangeContactsItemPath1.RR.ToBuilder().Dir().ShortRef(),
				ItemRef:     ExchangeContactsItemPath1.ItemLocation(),
				LocationRef: ExchangeContactsItemPath1.Loc.String(),
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						ItemType:    details.ExchangeContact,
						ContactName: "a-person",
					},
				},
			},
			{
				RepoRef:     ExchangeContactsItemPath2.RR.String(),
				ShortRef:    ExchangeContactsItemPath2.RR.ShortRef(),
				ParentRef:   ExchangeContactsItemPath2.RR.ToBuilder().Dir().ShortRef(),
				ItemRef:     ExchangeContactsItemPath2.ItemLocation(),
				LocationRef: ExchangeContactsItemPath2.Loc.String(),
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						ItemType:    details.ExchangeContact,
						ContactName: "another-person",
					},
				},
			},
		},
		version.OneDrive7LocationRef: {
			{
				RepoRef:     ExchangeContactsItemPath1.locationAsRepoRef().String(),
				ShortRef:    ExchangeContactsItemPath1.locationAsRepoRef().ShortRef(),
				ParentRef:   ExchangeContactsItemPath1.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemRef:     ExchangeContactsItemPath1.ItemLocation(),
				LocationRef: ExchangeContactsItemPath1.Loc.String(),
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						ItemType:    details.ExchangeContact,
						ContactName: "a-person",
					},
				},
			},
			{
				RepoRef:     ExchangeContactsItemPath2.locationAsRepoRef().String(),
				ShortRef:    ExchangeContactsItemPath2.locationAsRepoRef().ShortRef(),
				ParentRef:   ExchangeContactsItemPath2.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemRef:     ExchangeContactsItemPath2.ItemLocation(),
				LocationRef: ExchangeContactsItemPath2.Loc.String(),
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						ItemType:    details.ExchangeContact,
						ContactName: "another-person",
					},
				},
			},
		},
		0: {
			{
				RepoRef:   ExchangeContactsItemPath1.locationAsRepoRef().String(),
				ShortRef:  ExchangeContactsItemPath1.locationAsRepoRef().ShortRef(),
				ParentRef: ExchangeContactsItemPath1.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						ItemType:    details.ExchangeContact,
						ContactName: "a-person",
					},
				},
			},
			{
				RepoRef:   ExchangeContactsItemPath2.locationAsRepoRef().String(),
				ShortRef:  ExchangeContactsItemPath2.locationAsRepoRef().ShortRef(),
				ParentRef: ExchangeContactsItemPath2.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						ItemType:    details.ExchangeContact,
						ContactName: "another-person",
					},
				},
			},
		},
	}

	ExchangeEventsBasePath  = mustPathRep("tenant-id/exchange/user-id/events/holidays", false)
	ExchangeEventsBasePath2 = mustPathRep("tenant-id/exchange/user-id/events/moreholidays", false)
	ExchangeEventsItemPath1 = ExchangeEventsBasePath.MustAppend(ItemName1, true)
	ExchangeEventsItemPath2 = ExchangeEventsBasePath2.MustAppend(ItemName2, true)

	exchangeEventsItemsByVersion = map[int][]details.Entry{
		version.All8MigrateUserPNToID: {
			{
				RepoRef:     ExchangeEventsItemPath1.RR.String(),
				ShortRef:    ExchangeEventsItemPath1.RR.ShortRef(),
				ParentRef:   ExchangeEventsItemPath1.RR.ToBuilder().Dir().ShortRef(),
				ItemRef:     ExchangeEventsItemPath1.ItemLocation(),
				LocationRef: ExchangeEventsItemPath1.Loc.String(),
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						ItemType:    details.ExchangeEvent,
						Organizer:   "a-person",
						Subject:     "foo",
						EventStart:  Time1,
						EventRecurs: false,
					},
				},
			},
			{
				RepoRef:     ExchangeEventsItemPath2.RR.String(),
				ShortRef:    ExchangeEventsItemPath2.RR.ShortRef(),
				ParentRef:   ExchangeEventsItemPath2.RR.ToBuilder().Dir().ShortRef(),
				ItemRef:     ExchangeEventsItemPath2.ItemLocation(),
				LocationRef: ExchangeEventsItemPath2.Loc.String(),
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						ItemType:    details.ExchangeEvent,
						Organizer:   "a-person",
						Subject:     "foo",
						EventStart:  Time2,
						EventRecurs: true,
					},
				},
			},
		},
		2: {
			{
				RepoRef:     ExchangeEventsItemPath1.RR.String(),
				ShortRef:    ExchangeEventsItemPath1.RR.ShortRef(),
				ParentRef:   ExchangeEventsItemPath1.RR.ToBuilder().Dir().ShortRef(),
				LocationRef: ExchangeEventsItemPath1.Loc.String(),
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						ItemType:    details.ExchangeEvent,
						Organizer:   "a-person",
						Subject:     "foo",
						EventStart:  Time1,
						EventRecurs: false,
					},
				},
			},
			{
				RepoRef:     ExchangeEventsItemPath2.RR.String(),
				ShortRef:    ExchangeEventsItemPath2.RR.ShortRef(),
				ParentRef:   ExchangeEventsItemPath2.RR.ToBuilder().Dir().ShortRef(),
				LocationRef: ExchangeEventsItemPath2.Loc.String(),
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						ItemType:    details.ExchangeEvent,
						Organizer:   "a-person",
						Subject:     "foo",
						EventStart:  Time2,
						EventRecurs: true,
					},
				},
			},
		},
		0: {
			{
				RepoRef:   ExchangeEventsItemPath1.locationAsRepoRef().String(),
				ShortRef:  ExchangeEventsItemPath1.locationAsRepoRef().ShortRef(),
				ParentRef: ExchangeEventsItemPath1.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						ItemType:    details.ExchangeEvent,
						Organizer:   "a-person",
						Subject:     "foo",
						EventStart:  Time1,
						EventRecurs: false,
					},
				},
			},
			{
				RepoRef:   ExchangeEventsItemPath2.locationAsRepoRef().String(),
				ShortRef:  ExchangeEventsItemPath2.locationAsRepoRef().ShortRef(),
				ParentRef: ExchangeEventsItemPath2.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						ItemType:    details.ExchangeEvent,
						Organizer:   "a-person",
						Subject:     "foo",
						EventStart:  Time2,
						EventRecurs: true,
					},
				},
			},
		},
	}

	OneDriveRootPath   = mustPathRep("tenant-id/onedrive/user-id/files/drives/foo/root:", false)
	OneDriveFolderPath = OneDriveRootPath.MustAppend("folder", false)
	OneDriveBasePath1  = OneDriveFolderPath.MustAppend("a", false)
	OneDriveBasePath2  = OneDriveFolderPath.MustAppend("b", false)

	OneDriveItemPath1 = OneDriveFolderPath.MustAppend(ItemName1, true)
	OneDriveItemPath2 = OneDriveBasePath1.MustAppend(ItemName2, true)
	OneDriveItemPath3 = OneDriveBasePath2.MustAppend(ItemName3, true)

	OneDriveFolderFolder  = OneDriveFolderPath.Loc.PopFront().String()
	OneDriveParentFolder1 = OneDriveBasePath1.Loc.PopFront().String()
	OneDriveParentFolder2 = OneDriveBasePath2.Loc.PopFront().String()

	oneDriveItemsByVersion = map[int][]details.Entry{
		version.All8MigrateUserPNToID: {
			{
				RepoRef:     OneDriveItemPath1.locationAsRepoRef().String(),
				ShortRef:    OneDriveItemPath1.locationAsRepoRef().ShortRef(),
				ParentRef:   OneDriveItemPath1.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemRef:     OneDriveItemPath1.ItemLocation(),
				LocationRef: OneDriveItemPath1.Loc.String(),
				ItemInfo: details.ItemInfo{
					OneDrive: &details.OneDriveInfo{
						ItemType:   details.OneDriveItem,
						ParentPath: OneDriveFolderFolder,
						ItemName:   OneDriveItemPath1.ItemLocation() + "name",
						Size:       int64(23),
						Owner:      UserEmail1,
						Created:    Time2,
						Modified:   Time4,
					},
				},
			},
			{
				RepoRef:     OneDriveItemPath2.locationAsRepoRef().String(),
				ShortRef:    OneDriveItemPath2.locationAsRepoRef().ShortRef(),
				ParentRef:   OneDriveItemPath2.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemRef:     OneDriveItemPath2.ItemLocation(),
				LocationRef: OneDriveItemPath2.Loc.String(),
				ItemInfo: details.ItemInfo{
					OneDrive: &details.OneDriveInfo{
						ItemType:   details.OneDriveItem,
						ParentPath: OneDriveParentFolder1,
						ItemName:   OneDriveItemPath2.ItemLocation() + "name",
						Size:       int64(42),
						Owner:      UserEmail1,
						Created:    Time1,
						Modified:   Time3,
					},
				},
			},
			{
				RepoRef:     OneDriveItemPath3.locationAsRepoRef().String(),
				ShortRef:    OneDriveItemPath3.locationAsRepoRef().ShortRef(),
				ParentRef:   OneDriveItemPath3.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemRef:     OneDriveItemPath3.ItemLocation(),
				LocationRef: OneDriveItemPath3.Loc.String(),
				ItemInfo: details.ItemInfo{
					OneDrive: &details.OneDriveInfo{
						ItemType:   details.OneDriveItem,
						ParentPath: OneDriveParentFolder2,
						ItemName:   OneDriveItemPath3.ItemLocation() + "name",
						Size:       int64(19),
						Owner:      UserEmail2,
						Created:    Time2,
						Modified:   Time4,
					},
				},
			},
		},
		version.OneDrive7LocationRef: {
			{
				RepoRef:     OneDriveItemPath1.locationAsRepoRef().String(),
				ShortRef:    OneDriveItemPath1.locationAsRepoRef().ShortRef(),
				ParentRef:   OneDriveItemPath1.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				LocationRef: OneDriveItemPath1.Loc.String(),
				ItemInfo: details.ItemInfo{
					OneDrive: &details.OneDriveInfo{
						ItemType:   details.OneDriveItem,
						ParentPath: OneDriveFolderFolder,
						ItemName:   OneDriveItemPath1.ItemLocation() + "name",
						Size:       int64(23),
						Owner:      UserEmail1,
						Created:    Time2,
						Modified:   Time4,
					},
				},
			},
			{
				RepoRef:     OneDriveItemPath2.locationAsRepoRef().String(),
				ShortRef:    OneDriveItemPath2.locationAsRepoRef().ShortRef(),
				ParentRef:   OneDriveItemPath2.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				LocationRef: OneDriveItemPath2.Loc.String(),
				ItemInfo: details.ItemInfo{
					OneDrive: &details.OneDriveInfo{
						ItemType:   details.OneDriveItem,
						ParentPath: OneDriveParentFolder1,
						ItemName:   OneDriveItemPath2.ItemLocation() + "name",
						Size:       int64(42),
						Owner:      UserEmail1,
						Created:    Time1,
						Modified:   Time3,
					},
				},
			},
			{
				RepoRef:     OneDriveItemPath3.locationAsRepoRef().String(),
				ShortRef:    OneDriveItemPath3.locationAsRepoRef().ShortRef(),
				ParentRef:   OneDriveItemPath3.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				LocationRef: OneDriveItemPath3.Loc.String(),
				ItemInfo: details.ItemInfo{
					OneDrive: &details.OneDriveInfo{
						ItemType:   details.OneDriveItem,
						ParentPath: OneDriveParentFolder2,
						ItemName:   OneDriveItemPath3.ItemLocation() + "name",
						Size:       int64(19),
						Owner:      UserEmail2,
						Created:    Time2,
						Modified:   Time4,
					},
				},
			},
		},
		version.OneDrive6NameInMeta: {
			{
				RepoRef:   OneDriveItemPath1.locationAsRepoRef().String(),
				ShortRef:  OneDriveItemPath1.locationAsRepoRef().ShortRef(),
				ParentRef: OneDriveItemPath1.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemInfo: details.ItemInfo{
					OneDrive: &details.OneDriveInfo{
						ItemType:   details.OneDriveItem,
						ParentPath: OneDriveFolderFolder,
						ItemName:   OneDriveItemPath1.ItemLocation() + "name",
						Size:       int64(23),
						Owner:      UserEmail1,
						Created:    Time2,
						Modified:   Time4,
					},
				},
			},
			{
				RepoRef:   OneDriveItemPath2.locationAsRepoRef().String(),
				ShortRef:  OneDriveItemPath2.locationAsRepoRef().ShortRef(),
				ParentRef: OneDriveItemPath2.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemInfo: details.ItemInfo{
					OneDrive: &details.OneDriveInfo{
						ItemType:   details.OneDriveItem,
						ParentPath: OneDriveParentFolder1,
						ItemName:   OneDriveItemPath2.ItemLocation() + "name",
						Size:       int64(42),
						Owner:      UserEmail1,
						Created:    Time1,
						Modified:   Time3,
					},
				},
			},
			{
				RepoRef:   OneDriveItemPath3.locationAsRepoRef().String(),
				ShortRef:  OneDriveItemPath3.locationAsRepoRef().ShortRef(),
				ParentRef: OneDriveItemPath3.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemInfo: details.ItemInfo{
					OneDrive: &details.OneDriveInfo{
						ItemType:   details.OneDriveItem,
						ParentPath: OneDriveParentFolder2,
						ItemName:   OneDriveItemPath3.ItemLocation() + "name",
						Size:       int64(19),
						Owner:      UserEmail2,
						Created:    Time2,
						Modified:   Time4,
					},
				},
			},
		},
		0: {
			{
				RepoRef:   OneDriveItemPath1.locationAsRepoRef().String() + "name",
				ShortRef:  OneDriveItemPath1.locationAsRepoRef().ShortRef(),
				ParentRef: OneDriveItemPath1.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemInfo: details.ItemInfo{
					OneDrive: &details.OneDriveInfo{
						ItemType:   details.OneDriveItem,
						ParentPath: OneDriveFolderFolder,
						ItemName:   OneDriveItemPath1.ItemLocation() + "name",
						Size:       int64(23),
						Owner:      UserEmail1,
						Created:    Time2,
						Modified:   Time4,
					},
				},
			},
			{
				RepoRef:   OneDriveItemPath2.locationAsRepoRef().String() + "name",
				ShortRef:  OneDriveItemPath2.locationAsRepoRef().ShortRef(),
				ParentRef: OneDriveItemPath2.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemInfo: details.ItemInfo{
					OneDrive: &details.OneDriveInfo{
						ItemType:   details.OneDriveItem,
						ParentPath: OneDriveParentFolder1,
						ItemName:   OneDriveItemPath2.ItemLocation() + "name",
						Size:       int64(42),
						Owner:      UserEmail1,
						Created:    Time1,
						Modified:   Time3,
					},
				},
			},
			{
				RepoRef:   OneDriveItemPath3.locationAsRepoRef().String() + "name",
				ShortRef:  OneDriveItemPath3.locationAsRepoRef().ShortRef(),
				ParentRef: OneDriveItemPath3.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemInfo: details.ItemInfo{
					OneDrive: &details.OneDriveInfo{
						ItemType:   details.OneDriveItem,
						ParentPath: OneDriveParentFolder2,
						ItemName:   OneDriveItemPath3.ItemLocation() + "name",
						Size:       int64(19),
						Owner:      UserEmail2,
						Created:    Time2,
						Modified:   Time4,
					},
				},
			},
		},
	}

	SharePointRootPath    = mustPathRep("tenant-id/sharepoint/site-id/libraries/drives/foo/root:", false)
	SharePointLibraryPath = SharePointRootPath.MustAppend("library", false)
	SharePointBasePath1   = SharePointLibraryPath.MustAppend("a", false)
	SharePointBasePath2   = SharePointLibraryPath.MustAppend("b", false)

	SharePointLibraryItemPath1 = SharePointLibraryPath.MustAppend(ItemName1, true)
	SharePointLibraryItemPath2 = SharePointBasePath1.MustAppend(ItemName2, true)
	SharePointLibraryItemPath3 = SharePointBasePath2.MustAppend(ItemName3, true)

	SharePointLibraryFolder  = SharePointLibraryPath.Loc.PopFront().String()
	SharePointParentLibrary1 = SharePointBasePath1.Loc.PopFront().String()
	SharePointParentLibrary2 = SharePointBasePath2.Loc.PopFront().String()

	sharePointLibraryItemsByVersion = map[int][]details.Entry{
		version.All8MigrateUserPNToID: {
			{
				RepoRef:     SharePointLibraryItemPath1.locationAsRepoRef().String(),
				ShortRef:    SharePointLibraryItemPath1.locationAsRepoRef().ShortRef(),
				ParentRef:   SharePointLibraryItemPath1.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemRef:     SharePointLibraryItemPath1.ItemLocation(),
				LocationRef: SharePointLibraryItemPath1.Loc.String(),
				ItemInfo: details.ItemInfo{
					SharePoint: &details.SharePointInfo{
						ItemType:   details.SharePointLibrary,
						ParentPath: SharePointLibraryFolder,
						ItemName:   SharePointLibraryItemPath1.ItemLocation() + "name",
						Size:       int64(23),
						Owner:      UserEmail1,
						Created:    Time2,
						Modified:   Time4,
					},
				},
			},
			{
				RepoRef:     SharePointLibraryItemPath2.locationAsRepoRef().String(),
				ShortRef:    SharePointLibraryItemPath2.locationAsRepoRef().ShortRef(),
				ParentRef:   SharePointLibraryItemPath2.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemRef:     SharePointLibraryItemPath2.ItemLocation(),
				LocationRef: SharePointLibraryItemPath2.Loc.String(),
				ItemInfo: details.ItemInfo{
					SharePoint: &details.SharePointInfo{
						ItemType:   details.SharePointLibrary,
						ParentPath: SharePointParentLibrary1,
						ItemName:   SharePointLibraryItemPath2.ItemLocation() + "name",
						Size:       int64(42),
						Owner:      UserEmail1,
						Created:    Time1,
						Modified:   Time3,
					},
				},
			},
			{
				RepoRef:     SharePointLibraryItemPath3.locationAsRepoRef().String(),
				ShortRef:    SharePointLibraryItemPath3.locationAsRepoRef().ShortRef(),
				ParentRef:   SharePointLibraryItemPath3.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemRef:     SharePointLibraryItemPath3.ItemLocation(),
				LocationRef: SharePointLibraryItemPath3.Loc.String(),
				ItemInfo: details.ItemInfo{
					SharePoint: &details.SharePointInfo{
						ItemType:   details.SharePointLibrary,
						ParentPath: SharePointParentLibrary2,
						ItemName:   SharePointLibraryItemPath3.ItemLocation() + "name",
						Size:       int64(19),
						Owner:      UserEmail2,
						Created:    Time2,
						Modified:   Time4,
					},
				},
			},
		},
		version.OneDrive7LocationRef: {
			{
				RepoRef:     SharePointLibraryItemPath1.locationAsRepoRef().String(),
				ShortRef:    SharePointLibraryItemPath1.locationAsRepoRef().ShortRef(),
				ParentRef:   SharePointLibraryItemPath1.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				LocationRef: SharePointLibraryItemPath1.Loc.String(),
				ItemInfo: details.ItemInfo{
					SharePoint: &details.SharePointInfo{
						ItemType:   details.OneDriveItem,
						ParentPath: SharePointLibraryFolder,
						ItemName:   SharePointLibraryItemPath1.ItemLocation() + "name",
						Size:       int64(23),
						Owner:      UserEmail1,
						Created:    Time2,
						Modified:   Time4,
					},
				},
			},
			{
				RepoRef:     SharePointLibraryItemPath2.locationAsRepoRef().String(),
				ShortRef:    SharePointLibraryItemPath2.locationAsRepoRef().ShortRef(),
				ParentRef:   SharePointLibraryItemPath2.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				LocationRef: SharePointLibraryItemPath2.Loc.String(),
				ItemInfo: details.ItemInfo{
					SharePoint: &details.SharePointInfo{
						ItemType:   details.OneDriveItem,
						ParentPath: SharePointParentLibrary1,
						ItemName:   SharePointLibraryItemPath2.ItemLocation() + "name",
						Size:       int64(42),
						Owner:      UserEmail1,
						Created:    Time1,
						Modified:   Time3,
					},
				},
			},
			{
				RepoRef:     SharePointLibraryItemPath3.locationAsRepoRef().String(),
				ShortRef:    SharePointLibraryItemPath3.locationAsRepoRef().ShortRef(),
				ParentRef:   SharePointLibraryItemPath3.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				LocationRef: SharePointLibraryItemPath3.Loc.String(),
				ItemInfo: details.ItemInfo{
					SharePoint: &details.SharePointInfo{
						ItemType:   details.OneDriveItem,
						ParentPath: SharePointParentLibrary2,
						ItemName:   SharePointLibraryItemPath3.ItemLocation() + "name",
						Size:       int64(19),
						Owner:      UserEmail2,
						Created:    Time2,
						Modified:   Time4,
					},
				},
			},
		},
		version.OneDrive6NameInMeta: {
			{
				RepoRef:   SharePointLibraryItemPath1.locationAsRepoRef().String(),
				ShortRef:  SharePointLibraryItemPath1.locationAsRepoRef().ShortRef(),
				ParentRef: SharePointLibraryItemPath1.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemInfo: details.ItemInfo{
					SharePoint: &details.SharePointInfo{
						ItemType:   details.OneDriveItem,
						ParentPath: SharePointLibraryFolder,
						ItemName:   SharePointLibraryItemPath1.ItemLocation() + "name",
						Size:       int64(23),
						Owner:      UserEmail1,
						Created:    Time2,
						Modified:   Time4,
					},
				},
			},
			{
				RepoRef:   SharePointLibraryItemPath2.locationAsRepoRef().String(),
				ShortRef:  SharePointLibraryItemPath2.locationAsRepoRef().ShortRef(),
				ParentRef: SharePointLibraryItemPath2.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemInfo: details.ItemInfo{
					SharePoint: &details.SharePointInfo{
						ItemType:   details.OneDriveItem,
						ParentPath: SharePointParentLibrary1,
						ItemName:   SharePointLibraryItemPath2.ItemLocation() + "name",
						Size:       int64(42),
						Owner:      UserEmail1,
						Created:    Time1,
						Modified:   Time3,
					},
				},
			},
			{
				RepoRef:   SharePointLibraryItemPath3.locationAsRepoRef().String(),
				ShortRef:  SharePointLibraryItemPath3.locationAsRepoRef().ShortRef(),
				ParentRef: SharePointLibraryItemPath3.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemInfo: details.ItemInfo{
					SharePoint: &details.SharePointInfo{
						ItemType:   details.OneDriveItem,
						ParentPath: SharePointParentLibrary2,
						ItemName:   SharePointLibraryItemPath3.ItemLocation() + "name",
						Size:       int64(19),
						Owner:      UserEmail2,
						Created:    Time2,
						Modified:   Time4,
					},
				},
			},
		},
		0: {
			{
				RepoRef:   SharePointLibraryItemPath1.locationAsRepoRef().String() + "name",
				ShortRef:  SharePointLibraryItemPath1.locationAsRepoRef().ShortRef(),
				ParentRef: SharePointLibraryItemPath1.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemInfo: details.ItemInfo{
					SharePoint: &details.SharePointInfo{
						ItemType:   details.OneDriveItem,
						ParentPath: SharePointLibraryFolder,
						ItemName:   SharePointLibraryItemPath1.ItemLocation() + "name",
						Size:       int64(23),
						Owner:      UserEmail1,
						Created:    Time2,
						Modified:   Time4,
					},
				},
			},
			{
				RepoRef:   SharePointLibraryItemPath2.locationAsRepoRef().String() + "name",
				ShortRef:  SharePointLibraryItemPath2.locationAsRepoRef().ShortRef(),
				ParentRef: SharePointLibraryItemPath2.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemInfo: details.ItemInfo{
					SharePoint: &details.SharePointInfo{
						ItemType:   details.OneDriveItem,
						ParentPath: SharePointParentLibrary1,
						ItemName:   SharePointLibraryItemPath2.ItemLocation() + "name",
						Size:       int64(42),
						Owner:      UserEmail1,
						Created:    Time1,
						Modified:   Time3,
					},
				},
			},
			{
				RepoRef:   SharePointLibraryItemPath3.locationAsRepoRef().String() + "name",
				ShortRef:  SharePointLibraryItemPath3.locationAsRepoRef().ShortRef(),
				ParentRef: SharePointLibraryItemPath3.locationAsRepoRef().ToBuilder().Dir().ShortRef(),
				ItemInfo: details.ItemInfo{
					SharePoint: &details.SharePointInfo{
						ItemType:   details.OneDriveItem,
						ParentPath: SharePointParentLibrary2,
						ItemName:   SharePointLibraryItemPath3.ItemLocation() + "name",
						Size:       int64(19),
						Owner:      UserEmail2,
						Created:    Time2,
						Modified:   Time4,
					},
				},
			},
		},
	}
)

func GetDetailsSetForVersion(t *testing.T, wantedVersion int) *details.Details {
	entries := []details.Entry{}
	// TODO(ashmrtn): At some point make an exported variable somewhere that has
	// all the valid service/category pairs.
	dataTypes := map[path.ServiceType][]path.CategoryType{
		path.ExchangeService: {
			path.EmailCategory,
			path.EventsCategory,
			path.ContactsCategory,
		},
		path.OneDriveService: {
			path.FilesCategory,
		},
		path.SharePointService: {
			path.LibrariesCategory,
		},
	}

	for s, cats := range dataTypes {
		for _, cat := range cats {
			entries = append(entries, GetDeetsForVersion(t, s, cat, wantedVersion)...)
		}
	}

	return &details.Details{
		DetailsModel: details.DetailsModel{
			Entries: entries,
		},
	}
}

// GetItemsForVersion returns the set of items for the requested
// (service, category, version) tuple that reside at the indicated indices. If
// -1 is the only index provided then returns all items.
func GetItemsForVersion(
	t *testing.T,
	service path.ServiceType,
	cat path.CategoryType,
	wantVersion int,
	indices ...int,
) []details.Entry {
	deets := GetDeetsForVersion(t, service, cat, wantVersion)

	if len(indices) == 1 && indices[0] == -1 {
		return deets
	}

	var res []details.Entry

	for _, i := range indices {
		require.Less(t, i, len(deets), "requested index out of bounds", i, len(deets))
		res = append(res, deets[i])
	}

	return res
}

// GetDeetsForVersion returns the set of details with the highest
// version <= the requested version.
func GetDeetsForVersion(
	t *testing.T,
	service path.ServiceType,
	cat path.CategoryType,
	wantVersion int,
) []details.Entry {
	var input map[int][]details.Entry

	switch service {
	case path.ExchangeService:
		switch cat {
		case path.EmailCategory:
			input = exchangeEmailItemsByVersion

		case path.EventsCategory:
			input = exchangeEventsItemsByVersion

		case path.ContactsCategory:
			input = exchangeContactsItemsByVersion
		}

	case path.OneDriveService:
		if cat == path.FilesCategory {
			input = oneDriveItemsByVersion
		}

	case path.SharePointService:
		if cat == path.LibrariesCategory {
			input = sharePointLibraryItemsByVersion
		}
	}

	require.NotNil(
		t,
		input,
		"unsupported (service, category)",
		service.String(),
		cat.String())

	return getDeetsForVersion(t, wantVersion, input)
}

func getDeetsForVersion(
	t *testing.T,
	wantVersion int,
	deetsSet map[int][]details.Entry,
) []details.Entry {
	var (
		res        []details.Entry
		resVersion = version.NoBackup
	)

	for v, deets := range deetsSet {
		if v <= wantVersion && v > resVersion {
			resVersion = v
			res = deets
		}
	}

	require.NotEmpty(t, res, "unable to find details for version", wantVersion)

	return slices.Clone(res)
}
