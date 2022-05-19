package repository_test

import (
	"testing"

	"github.com/alcionai/corso/pkg/repository"
)

func TestNewS3(t *testing.T) {
	r := repository.NewS3("bkt", "ak", "sk", "tid", "cid", "sec")
	if _, ok := r.Config.(repository.S3Config); !ok {
		t.Errorf("expected repository config to be of type repository.S3Config, got [%T]", r.Config)
	}
	if r.Provider != repository.S3Provider {
		t.Errorf("expected repository provider type to be %s, got %s", repository.S3Provider.Name(), r.Provider.Name())
	}
}
