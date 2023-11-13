package testdata

import "github.com/google/uuid"

const hashLength = 7

func NewHashForRepoConfigName() string {
	return uuid.NewString()[:hashLength]
}
