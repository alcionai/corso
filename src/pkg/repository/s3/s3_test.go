package s3_test

import (
	"testing"

	"github.com/alcionai/corso/pkg/repository/s3"
)

func TestNewS3(t *testing.T) {
	cfg := s3.NewConfig("bucket", "access", "secret")
	if cfg.Bucket != "bucket" {
		t.Errorf("expected s3 config bucke to be 'bucket', got '%s'", cfg.Bucket)
	}
}
