package repository

import (
	"time"
)

// PersistentConfig represents configuration info that is persisted in the corso
// repo and can be updated with repository.UpdatePersistentConfig. Leaving a
// field as nil will result in no updates being made to it (i.e. PATCH
// semantics).
type PersistentConfig struct {
	MinEpochDuration *time.Duration
}
