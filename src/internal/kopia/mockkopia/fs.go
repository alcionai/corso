package mockkopia

import (
	"bytes"
	"context"
	"io"
	"os"
	"time"

	"github.com/kopia/kopia/fs"
)

const (
	DefaultPermissions os.FileMode = 0o777
)

type MockEntry struct {
	EntryName    string
	EntryMode    os.FileMode
	EntrySize    int64
	EntryModTime time.Time
	EntryOwner   fs.OwnerInfo
	EntryDevice  fs.DeviceInfo
}

func (me MockEntry) Name() string {
	return me.EntryName
}

func (me MockEntry) IsDir() bool {
	return me.EntryMode.IsDir()
}

func (me MockEntry) Mode() os.FileMode {
	return me.EntryMode
}

func (me MockEntry) ModTime() time.Time {
	return me.EntryModTime
}

func (me MockEntry) Size() int64 {
	return me.EntrySize
}

func (me MockEntry) Sys() any {
	return nil
}

func (me MockEntry) Owner() fs.OwnerInfo {
	return me.EntryOwner
}

func (me MockEntry) Device() fs.DeviceInfo {
	return me.EntryDevice
}

func (me MockEntry) LocalFilesystemPath() string {
	return ""
}

func (me *MockEntry) Close() {
}

type mockReader struct {
	e fs.Entry
	io.ReadSeeker
}

func (mr mockReader) Entry() (fs.Entry, error) {
	return mr.e, nil
}

func (mr *mockReader) Close() error {
	return nil
}

type MockFile struct {
	fs.Entry
	OpenErr error
	Data    []byte
}

func (mf *MockFile) Open(context.Context) (fs.Reader, error) {
	if mf.OpenErr != nil {
		return nil, mf.OpenErr
	}

	return &mockReader{
		e:          mf.Entry,
		ReadSeeker: bytes.NewReader(mf.Data),
	}, nil
}
