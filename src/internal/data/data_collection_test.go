package data

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CollectionSuite struct {
	suite.Suite
}

func TestDataCollectionSuite(t *testing.T) {
	suite.Run(t, new(CollectionSuite))
}
