package repo

import (
	"fmt"
	"testing"
)

type goodRC struct{}

func (g goodRC) Initialize() error {
	return nil
}

func (g goodRC) Connect() error {
	return nil
}

var (
	initErr = fmt.Errorf("on initialize")
	connErr = fmt.Errorf("on connect")
)

type badRC struct{}

func (b badRC) Initialize() error {
	return initErr
}

func (b badRC) Connect() error {
	return connErr
}

func TestInitializeRepo(t *testing.T) {
	good := goodRC{}
	if err := initializeRepo(good); err != nil {
		t.Errorf("expected no errors, got [%v]", err)
	}

	bad := badRC{}
	if err := initializeRepo(bad); err != initErr {
		t.Errorf("expected the error [%v], got [%v]", initErr, err)
	}
}

func TestConnectRepo(t *testing.T) {
	good := goodRC{}
	if err := connectRepo(good); err != nil {
		t.Errorf("expected no errors, got [%v]", err)
	}

	bad := badRC{}
	if err := connectRepo(bad); err != connErr {
		t.Errorf("expected the error [%v], got [%v]", connErr, err)
	}
}
