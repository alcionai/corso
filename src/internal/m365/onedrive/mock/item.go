package mock

import (
	"bytes"
	"context"
	"io"
	"time"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/backup/details"
)

// ---------------------------------------------------------------------------
// data.Stream
// ---------------------------------------------------------------------------

var _ data.Stream = &Data{}

// TODO: move to data/mock for service-agnostic mocking
// Data represents a single item retrieved from, or restored to, onedrive
type Data struct {
	ID           string
	Reader       io.ReadCloser
	ReadErr      error
	size         int64
	modifiedTime time.Time
	deleted      bool
}

func (d *Data) UUID() string       { return d.ID }
func (d *Data) Deleted() bool      { return d.deleted }
func (d *Data) Size() int64        { return d.size }
func (d *Data) ModTime() time.Time { return d.modifiedTime }

func (d *Data) ToReader() io.ReadCloser {
	if d.ReadErr != nil {
		return io.NopCloser(errReader{d.ReadErr})
	}

	return d.Reader
}

func (d *Data) Info() details.ItemInfo {
	return details.ItemInfo{
		OneDrive: &details.OneDriveInfo{
			ItemType: details.OneDriveItem,
			ItemName: "test.txt",
			Size:     1,
		},
	}
}

type errReader struct {
	readErr error
}

func (er errReader) Read([]byte) (int, error) {
	return 0, er.readErr
}

// ---------------------------------------------------------------------------
// FetchItemByNamer
// ---------------------------------------------------------------------------

var _ data.FetchItemByNamer = &FetchItemByName{}

type FetchItemByName struct {
	Item data.Stream
	Err  error
}

func (f FetchItemByName) FetchItemByName(context.Context, string) (data.Stream, error) {
	return f.Item, f.Err
}

// ---------------------------------------------------------------------------
// stub payload
// ---------------------------------------------------------------------------

func FileRespReadCloser(pl string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(pl)))
}

const (
	DriveItemFileName = "fnords.txt"
	DriveFileMetaData = `{"fileName": "` + DriveItemFileName + `"}`
)

//nolint:lll
const DriveFilePayloadData = `{
    "@odata.context": "https://graph.microsoft.com/v1.0/$metadata#drives('b%22-8wC6Jt04EWvKr1fQUDOyw5Gk8jIUJdEjzqonlSRf48i67LJdwopT4-6kiycJ5AV')/items/$entity",
    "@microsoft.graph.downloadUrl": "https://test-my.sharepoint.com/personal/brunhilda_test_onmicrosoft_com/_layouts/15/download.aspx?UniqueId=deadbeef-1b6a-4d13-aae6-bf5f9b07d424&Translate=false&tempauth=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOiIwMDAwMDAwMy0wMDAwLTBmZjEtY2UwMC0wMDAwMDAwMDAwMDAvMTBycWMyLW15LnNoYXJlcG9pbnQuY29tQGZiOGFmYmFhLWU5NGMtNGVhNS04YThhLTI0YWZmMDRkNzg3NCIsImlzcyI6IjAwMDAwMDAzLTAwMDAtMGZmMS1jZTAwLTAwMDAwMDAwMDAwMCIsIm5iZiI6IjE2ODUxMjk1MzIiLCJleHAiOiIxNjg1MTMzMTMyIiwiZW5kcG9pbnR1cmwiOiJkTStxblBIQitkNDMzS0ErTHVTUVZMRi9IaVliSkI2eHJWN0tuYk45aXQ0PSIsImVuZHBvaW50dXJsTGVuZ3RoIjoiMTYxIiwiaXNsb29wYmFjayI6IlRydWUiLCJjaWQiOiJOVFl4TXpNMFkyWXRZVFk0TVMwMFpXUmxMVGt5TjJZdFlXVmpNVGMwTldWbU16TXgiLCJ2ZXIiOiJoYXNoZWRwcm9vZnRva2VuIiwic2l0ZWlkIjoiWlRnd01tTmpabUl0TnpRNVlpMDBOV1V3TFdGbU1tRXRZbVExWmpReE5EQmpaV05pIiwiYXBwX2Rpc3BsYXluYW1lIjoiS2VlcGVyc19Mb2NhbCIsIm5hbWVpZCI6ImFkYjk3MTQ2LTcxYTctNDkxYS05YWMwLWUzOGFkNzdkZWViNkBmYjhhZmJhYS1lOTRjLTRlYTUtOGE4YS0yNGFmZjA0ZDc4NzQiLCJyb2xlcyI6ImFsbHNpdGVzLndyaXRlIGFsbHNpdGVzLm1hbmFnZSBhbGxmaWxlcy53cml0ZSBhbGxzaXRlcy5mdWxsY29udHJvbCBhbGxwcm9maWxlcy5yZWFkIiwidHQiOiIxIiwidXNlUGVyc2lzdGVudENvb2tpZSI6bnVsbCwiaXBhZGRyIjoiMjA1MTkwLjE1Ny4zMCJ9.lN7Vpfzk1abEyE0M3gyRyZXEaGQ3JMXCyaXUBNbD5Vo&ApiVersion=2.0",
    "createdDateTime": "2023-04-25T21:32:58Z",
    "eTag": "\"{DEADBEEF-1B6A-4D13-AAE6-BF5F9B07D424},1\"",
    "id": "017W47IH3FQVEFI23QCNG2VZV7L6NQPVBE",
    "lastModifiedDateTime": "2023-04-25T21:32:58Z",
    "name": "huehuehue.GIF",
    "webUrl": "https://test-my.sharepoint.com/personal/brunhilda_test_onmicrosoft_com/Documents/test/huehuehue.GIF",
    "cTag": "\"c:{DEADBEEF-1B6A-4D13-AAE6-BF5F9B07D424},1\"",
    "size": 88843,
    "createdBy": {
        "user": {
            "email": "brunhilda@test.onmicrosoft.com",
            "id": "DEADBEEF-4c80-4da4-86ef-a08d8d6f0f94",
            "displayName": "BrunHilda"
        }
    },
    "lastModifiedBy": {
        "user": {
            "email": "brunhilda@10rqc2.onmicrosoft.com",
            "id": "DEADBEEF-4c80-4da4-86ef-a08d8d6f0f94",
            "displayName": "BrunHilda"
        }
    },
    "parentReference": {
        "driveType": "business",
        "driveId": "b!-8wC6Jt04EWvKr1fQUDOyw5Gk8jIUJdEjzqonlSRf48i67LJdwopT4-6kiycJ5VA",
        "id": "017W47IH6DRQF2GS2N6NGWLZRS7RUJ2DIP",
        "path": "/drives/b!-8wC6Jt04EWvKr1fQUDOyw5Gk8jIUJdEjzqonlSRf48i67LJdwopT4-6kiycJ5VA/root:/test",
        "siteId": "DEADBEEF-749b-45e0-af2a-bd5f4140cecb"
    },
    "file": {
        "mimeType": "image/gif",
        "hashes": {
            "quickXorHash": "sU5rmXOvVFn6zJHpCPro9cYaK+Q="
        }
    },
    "fileSystemInfo": {
        "createdDateTime": "2023-04-25T21:32:58Z",
        "lastModifiedDateTime": "2023-04-25T21:32:58Z"
    },
    "image": {}
}`
