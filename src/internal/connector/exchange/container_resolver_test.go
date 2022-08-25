package exchange

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/path"
)

type mockContainerInfo struct {
	parentID    string
	id          string
	displayName string
	returned    bool
	// String respresentation of the item's path for test assertions.
	path string
}

func (mci mockContainerInfo) ID() string {
	return mci.id
}

func (mci mockContainerInfo) ParentID() string {
	return mci.parentID
}

func (mci mockContainerInfo) DisplayName() string {
	return mci.displayName
}

// newMockContainerInfo returns a mockContainerInfo with the given displayName
// and random parentID and id fields. This container can be used for a test.
// This is a separate function because tests modify some values in
// mockContainerInfo during the test, making it unsafe to share them if tests
// are ever run in parallel or if they are not reset between serial tests.
func newMockContainerInfo(displayName string) *mockContainerInfo {
	return &mockContainerInfo{
		parentID:    uuid.NewString(),
		id:          uuid.NewString(),
		displayName: displayName,
	}
}

type mockContainerResolver struct {
	mappings map[string]*mockContainerInfo
	// ID of the root container. Must be one of the items in mappings or an error
	// will be returned on FetchRoot.
	rootID string
	// returnOnce disables returning the same container multiple times. This can
	// be used to ensure caching is working properly.
	returnOnce bool
}

// Fetch looks for the ID in mappings. If the container is present and not
// already returned (returnOnce == true) or returnOnce == false it is returned.
// If the container is not found or the above is not true, an error is returned.
func (mcr *mockContainerResolver) Fetch(
	ctx context.Context,
	id string,
) (container, error) {
	c, ok := mcr.mappings[id]
	if !ok {
		return nil, errors.New("no such container")
	}

	if mcr.returnOnce && c.returned {
		return nil, errors.Errorf("container with ID: %v requested multiple times", id)
	}

	c.returned = true

	return c, nil
}

func (mcr *mockContainerResolver) FetchRoot(context.Context) (container, error) {
	if mcr.rootID == "" {
		return nil, errors.New("no root configured")
	}

	if c, ok := mcr.mappings[mcr.rootID]; ok {
		// Make sure this isn't also requested for a normal lookup.
		c.returned = true
		return c, nil
	}

	return nil, errors.Errorf("no root container with ID: %v", mcr.rootID)
}

type CachingContainerResolverUnitSuite struct {
	suite.Suite
}

func TestCachingContainerResolverUnitSuite(t *testing.T) {
	suite.Run(t, new(CachingContainerResolverUnitSuite))
}

func (suite *CachingContainerResolverUnitSuite) TestFailsInitWithNoRoot() {
	ctx := context.Background()
	mcr := &mockContainerResolver{}

	ccr := cachingContainerResolver{
		cached:   map[string]cachedContainer{},
		resolver: mcr,
		prefix:   &path.Builder{},
	}

	assert.Error(suite.T(), ccr.Initialize(ctx))
}

func (suite *CachingContainerResolverUnitSuite) TestLookupFailsWithNoMapping() {
	t := suite.T()
	ctx := context.Background()
	root := newMockContainerInfo("root")
	mcr := &mockContainerResolver{
		mappings: map[string]*mockContainerInfo{
			root.ID(): root,
		},
		rootID: root.ID(),
	}

	ccr := cachingContainerResolver{
		cached:   map[string]cachedContainer{},
		resolver: mcr,
		prefix:   &path.Builder{},
	}

	require.NoError(t, ccr.Initialize(ctx))

	_, err := ccr.Lookup(ctx, "other")
	assert.Error(t, err)
}

// This suite cannot run its tests in parallel.
type ConfiguredCachingContainerResolverUnitSuite struct {
	suite.Suite

	root   *mockContainerInfo
	first  *mockContainerInfo
	second *mockContainerInfo
	third  *mockContainerInfo
	mcr    *mockContainerResolver
}

func (suite *ConfiguredCachingContainerResolverUnitSuite) SetupSuite() {
	suite.root = newMockContainerInfo("root")
	suite.first = newMockContainerInfo("folder")
	suite.second = newMockContainerInfo("subfolder")
	suite.third = newMockContainerInfo("subsubfolder")

	suite.root.path = "root"
	suite.first.path = "root/folder"
	suite.first.parentID = suite.root.ID()
	suite.second.path = "root/folder/subfolder"
	suite.second.parentID = suite.first.ID()
	suite.third.path = "root/folder/subfolder/subsubfolder"
	suite.third.parentID = suite.second.ID()

	suite.mcr = &mockContainerResolver{
		mappings: map[string]*mockContainerInfo{
			suite.root.ID():   suite.root,
			suite.first.ID():  suite.first,
			suite.second.ID(): suite.second,
			suite.third.ID():  suite.third,
		},
		rootID: suite.root.ID(),
	}
}

func (suite *ConfiguredCachingContainerResolverUnitSuite) SetupTest() {
	for _, c := range suite.mcr.mappings {
		c.returned = false
	}
}

func TestConfiguredCachingContainerResolverUnitSuite(t *testing.T) {
	suite.Run(t, new(ConfiguredCachingContainerResolverUnitSuite))
}

func (suite *ConfiguredCachingContainerResolverUnitSuite) TestLookup() {
	ctx := context.Background()
	suite.mcr.returnOnce = false
	ccr := cachingContainerResolver{
		cached:   map[string]cachedContainer{},
		resolver: suite.mcr,
		prefix:   path.Builder{}.Append(suite.root.DisplayName()),
	}

	require.NoError(suite.T(), ccr.Initialize(ctx))

	for _, c := range suite.mcr.mappings {
		suite.T().Run(fmt.Sprintf("Lookup-%s", c.DisplayName()), func(t *testing.T) {
			p, err := ccr.Lookup(ctx, c.ID())
			require.NoError(t, err)

			assert.Equal(t, c.path, p.String())
		})
	}
}

func (suite *ConfiguredCachingContainerResolverUnitSuite) TestLookupCaching() {
	ctx := context.Background()
	suite.mcr.returnOnce = true
	ccr := cachingContainerResolver{
		cached:   map[string]cachedContainer{},
		resolver: suite.mcr,
		prefix:   path.Builder{}.Append(suite.root.DisplayName()),
	}

	require.NoError(suite.T(), ccr.Initialize(ctx))

	for _, c := range suite.mcr.mappings {
		suite.T().Run(fmt.Sprintf("Lookup-%s", c.DisplayName()), func(t *testing.T) {
			p, err := ccr.Lookup(ctx, c.ID())
			require.NoError(t, err)

			assert.Equal(t, c.path, p.String())
		})
	}
}
