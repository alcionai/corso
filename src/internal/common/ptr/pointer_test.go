package ptr_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
)

type PointerSuite struct {
	suite.Suite
}

func TestPointerSuite(t *testing.T) {
	suite.Run(t, new(PointerSuite))
}

// TestValue  checks to ptr derefencing for the
// following types:
// - *string
// - *bool
// - *time.Time
func (suite *PointerSuite) TestValue() {
	var (
		t          = suite.T()
		created    *time.Time
		testString *string
		testBool   *bool
		testInt    *int
		testInt32  *int32
		testInt64  *int64
	)

	// String Checks
	subject := ptr.Value(testString)
	assert.Empty(t, subject)

	hello := "Hello World"
	testString = &hello
	subject = ptr.Value(testString)

	t.Logf("Received: %s", subject)
	assert.NotEmpty(t, subject)

	// Time Checks

	myTime := ptr.Value(created)
	assert.Empty(t, myTime)
	assert.NotNil(t, myTime)

	now := time.Now()
	created = &now
	myTime = ptr.Value(created)
	assert.NotEmpty(t, myTime)

	// Bool Checks
	truth := true
	myBool := ptr.Value(testBool)
	assert.NotNil(t, myBool)
	assert.False(t, myBool)

	testBool = &truth
	myBool = ptr.Value(testBool)
	assert.NotNil(t, myBool)
	assert.True(t, myBool)

	// Int checks
	myInt := ptr.Value(testInt)
	myInt32 := ptr.Value(testInt32)
	myInt64 := ptr.Value(testInt64)

	assert.NotNil(t, myInt)
	assert.NotNil(t, myInt32)
	assert.NotNil(t, myInt64)
	assert.Empty(t, myInt)
	assert.Empty(t, myInt32)
	assert.Empty(t, myInt64)

	num := 4071
	num32 := int32(num * 32)
	num64 := int64(num * 2048)
	testInt = &num
	testInt32 = &num32
	testInt64 = &num64

	myInt = ptr.Value(testInt)
	myInt32 = ptr.Value(testInt32)
	myInt64 = ptr.Value(testInt64)

	assert.NotNil(t, myInt)
	assert.NotNil(t, myInt32)
	assert.NotNil(t, myInt64)
	assert.NotEmpty(t, myInt)
	assert.NotEmpty(t, myInt32)
	assert.NotEmpty(t, myInt64)
}
