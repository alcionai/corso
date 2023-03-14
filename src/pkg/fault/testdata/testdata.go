package testdata

import (
	"github.com/stretchr/testify/assert"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/pkg/fault"
)

func MakeErrors(failure, recovered, skipped bool) fault.Errors {
	fe := fault.Errors{}

	if failure {
		fe.Failure = assert.AnError
	}

	if recovered {
		fe.Recovered = []error{clues.New("recoverable")}
	}

	if skipped {
		fe.Skipped = []fault.Skipped{*fault.FileSkip(fault.SkipMalware, "id", "name", nil)}
	}

	return fe
}
