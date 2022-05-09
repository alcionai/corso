package repository_test

import (
	"testing"

	"github.com/alcionai/corso/repository"
)

func TestS3(t *testing.T) {
	s3 := repository.NewS3("this", "is a", "test", "doesn't", "matter", "yet")
	if len(s3.ID) == 0 {
		t.Fatalf("Generated s3 repo should have had an ID")
	}
}

func TestInitialize(t *testing.T) {
	s3 := repository.S3Repository{}
	err := s3.Initialize()
	if err != nil {
		t.Fatalf("Didn't expect initialize to error, got [%v]", err)
	}
}

func TestConnect(t *testing.T) {
	s3 := repository.S3Repository{}
	err := s3.Connect()
	if err != nil {
		t.Fatalf("Didn't expect connect to error, got [%v]", err)
	}
}
