package testdata

import "github.com/google/uuid"

const hashLength = 7

func NewHashForRepoConfigName() string {
	_ = uuid.NewString()[:hashLength]
	return "constant"
}
