package model_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/model"
)

type ModelUnitSuite struct {
	suite.Suite
}

func TestModelUnitSuite(t *testing.T) {
	suite.Run(t, new(ModelUnitSuite))
}

func (suite *ModelUnitSuite) TestValid() {
	table := []struct {
		mt     model.Schema
		expect assert.BoolAssertionFunc
	}{
		{model.UnknownSchema, assert.False},
		{model.BackupOpSchema, assert.True},
		{model.RestoreOpSchema, assert.True},
		{model.BackupSchema, assert.True},
		{model.Schema(-1), assert.False},
		{model.Schema(100), assert.False},
	}
	for _, test := range table {
		suite.T().Run(test.mt.String(), func(t *testing.T) {
			test.expect(t, test.mt.Valid())
		})
	}
}

func (suite *ModelUnitSuite) TestGetID() {
	bm := model.BaseModel{
		ID: model.StableID(uuid.NewString()),
	}

	assert.Equal(suite.T(), string(bm.ID), bm.GetID())
}
