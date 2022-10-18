package onedrive

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/tester"
)

func loadTestService(t *testing.T) *oneDriveService {
	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err)

	service, err := NewOneDriveService(m365)
	require.NoError(t, err)

	return service

}
