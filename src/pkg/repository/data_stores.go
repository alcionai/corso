package repository

import (
	"github.com/alcionai/corso/src/internal/kopia/inject"
)

type DataStorer interface {
	inject.RestoreProducer
}

type DataStoreConnector interface {
	DataStore() DataStorer
}

func (r *repository) DataStore() DataStorer {
	return r.dataLayer
}
