package repository

import (
	"context"
	"testing"

	"github.com/kopia/kopia/repo/blob"
)

type testProvider struct{}

func (tp testProvider) KopiaStorage(ctx context.Context, create bool) (blob.Storage, error) {
	return nil, nil
}

func TestNewRepo(t *testing.T) {
	r := newRepo(UnknownProvider, Account{}, testProvider{})
	if _, ok := r.Config.(testProvider); !ok {
		t.Fatalf("Generated repo config interface should be type testProvider")
	}
}

func TestInitialize(t *testing.T) {
	r := newRepo(UnknownProvider, Account{}, testProvider{})
	err := r.Initialize(context.Background())
	if err != nil {
		t.Fatalf("Didn't expect initialize to error, got [%v]", err)
	}
}

func TestConnect(t *testing.T) {
	r := newRepo(UnknownProvider, Account{}, testProvider{})
	err := r.Connect(context.Background())
	if err != nil {
		t.Fatalf("Didn't expect connect to error, got [%v]", err)
	}
}
