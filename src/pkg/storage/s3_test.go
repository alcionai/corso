package storage_test

import (
	"testing"

	"github.com/alcionai/corso/pkg/storage"
)

func TestS3Config_Config(t *testing.T) {
	s3 := storage.S3Config{"ak", "bkt", "end", "sk", "tkn"}
	c := s3.Config()
	table := []struct {
		key    string
		expect string
	}{
		{"s3_bucket", s3.Bucket},
		{"s3_accessKey", s3.AccessKey},
		{"s3_endpoint", s3.Endpoint},
		{"s3_secretKey", s3.SecretKey},
		{"s3_sessionToken", s3.SessionToken},
	}
	for _, test := range table {
		key := test.key
		expect := test.expect
		if c[key] != expect {
			t.Errorf("expected config key [%s] to hold value [%s], got [%s]", key, expect, c[key])
		}
	}
}

func TestStorage_S3Config(t *testing.T) {
	in := storage.S3Config{"ak", "bkt", "end", "sk", "tkn"}
	s := storage.NewStorage(storage.ProviderS3, in)
	out := s.S3Config()
	if in.Bucket != out.Bucket {
		t.Errorf("expected S3Config.Bucket to be [%s], got [%s]", in.Bucket, out.Bucket)
	}
	if in.AccessKey != out.AccessKey {
		t.Errorf("expected S3Config.AccessKey to be [%s], got [%s]", in.AccessKey, out.AccessKey)
	}
	if in.Endpoint != out.Endpoint {
		t.Errorf("expected S3Config.Endpoint to be [%s], got [%s]", in.Endpoint, out.Endpoint)
	}
	if in.SecretKey != out.SecretKey {
		t.Errorf("expected S3Config.SecretKey to be [%s], got [%s]", in.SecretKey, out.SecretKey)
	}
	if in.SessionToken != out.SessionToken {
		t.Errorf("expected S3Config.SessionToken to be [%s], got [%s]", in.SessionToken, out.SessionToken)
	}
}
