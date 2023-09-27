package api

import (
	"context"
	"fmt"
	"strings"

	"github.com/alcionai/clues"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/groups"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/internal/common/tform"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
)

const (
	teamsAdditionalDataLabel    = "Team"
	ResourceProvisioningOptions = "resourceProvisioningOptions"
)

// ---------------------------------------------------------------------------
// controller
// ---------------------------------------------------------------------------

func (c Client) Groups() Groups {
	return Groups{c}
}

// Groups is an interface-compliant provider of the client.
type Groups struct {
	Client
}

func (c Groups) GetAll(
	ctx context.Context,
	errs *fault.Bus,
) ([]models.Groupable, error) {
	service, err := c.Service()
	if err != nil {
		return nil, err
	}

	return getGroups(ctx, errs, service)
}

// GetAll retrieves all groups.
func getGroups(
	ctx context.Context,
	errs *fault.Bus,
	service graph.Servicer,
) ([]models.Groupable, error) {
	resp, err := service.Client().Groups().Get(ctx, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting all groups")
	}

	iter, err := msgraphgocore.NewPageIterator[models.Groupable](
		resp,
		service.Adapter(),
		models.CreateGroupCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating groups iterator")
	}

	var (
		results = make([]models.Groupable, 0)
		el      = errs.Local()
	)

	iterator := func(item models.Groupable) bool {
		if el.Failure() != nil {
			return false
		}

		err := ValidateGroup(item)
		if err != nil {
			el.AddRecoverable(ctx, graph.Wrap(ctx, err, "validating groups"))
		} else {
			results = append(results, item)
		}

		return true
	}

	if err := iter.Iterate(ctx, iterator); err != nil {
		return nil, graph.Wrap(ctx, err, "iterating all groups")
	}

	return results, el.Failure()
}

const filterGroupByDisplayNameQueryTmpl = "displayName eq '%s'"

// GetID can look up a group by either its canonical id (a uuid)
// or by the group's display name.  If looking up the display name
// an error will be returned if more than one group gets returned
// in the results.
func (c Groups) GetByID(
	ctx context.Context,
	identifier string,
	_ CallConfig, // matching standards
) (models.Groupable, error) {
	service, err := c.Service()
	if err != nil {
		return nil, err
	}

	ctx = clues.Add(ctx, "resource_identifier", identifier)

	var group models.Groupable

	// prefer lookup by id, but fallback to lookup by display name,
	// even in the case of a uuid, just in case the display name itself
	// is a uuid.
	if uuidRE.MatchString(identifier) {
		group, err = service.
			Client().
			Groups().
			ByGroupId(identifier).
			Get(ctx, nil)
		if err == nil {
			return group, nil
		}

		logger.CtxErr(ctx, err).Info("finding group by id, falling back to display name")
	}

	opts := &groups.GroupsRequestBuilderGetRequestConfiguration{
		Headers: newEventualConsistencyHeaders(),
		QueryParameters: &groups.GroupsRequestBuilderGetQueryParameters{
			Filter: ptr.To(fmt.Sprintf(filterGroupByDisplayNameQueryTmpl, identifier)),
		},
	}

	resp, err := service.Client().Groups().Get(ctx, opts)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "finding group by display name")
	}

	vs := resp.GetValue()

	if len(vs) == 0 {
		return nil, clues.Stack(graph.ErrResourceOwnerNotFound).WithClues(ctx)
	} else if len(vs) > 1 {
		return nil, clues.Stack(graph.ErrMultipleResultsMatchIdentifier).WithClues(ctx)
	}

	group = vs[0]

	return group, nil
}

// GetAllSites gets all the sites that belong to a group. This is
// necessary as private and shared channels gets their on individual
// sites. All the other channels make use of the root site.
func (c Groups) GetAllSites(
	ctx context.Context,
	identifier string,
	errs *fault.Bus,
) ([]models.Siteable, error) {
	root, err := c.GetRootSite(ctx, identifier)
	if err != nil {
		return nil, clues.Wrap(err, "getting root site")
	}

	sites := []models.Siteable{root}

	channels, err := Channels(c).GetChannels(ctx, identifier)
	if err != nil {
		return nil, clues.Wrap(err, "getting channels")
	}

	service, err := c.Service()
	if err != nil {
		return nil, err
	}

	for _, ch := range channels {
		if ptr.Val(ch.GetMembershipType()) == models.STANDARD_CHANNELMEMBERSHIPTYPE {
			// Standard channels use root site
			continue
		}

		resp, err := service.
			Client().
			Teams().
			ByTeamId(identifier).
			Channels().
			ByChannelId(ptr.Val(ch.GetId())).
			FilesFolder().
			Get(ctx, nil)
		if err != nil {
			return nil, clues.Wrap(err, "getting files folder for channel").
				With("channel_id", ptr.Val(ch.GetId()))
		}

		// WebURL returned here is the url to the documents folder, we
		// have to trim that out to get the actual site's webURL
		// https://example.sharepoint.com/sites/<site-name>/Shared%20Documents/<channelName>
		documentWebURL := ptr.Val(resp.GetWebUrl())

		// TODO(meain): is this hacky? is there a better way?
		dwurlSplits := strings.Split(documentWebURL, "/")
		siteWebURL := strings.Join(dwurlSplits[:5], "/")

		site, err := Sites(c).GetByID(ctx, siteWebURL, CallConfig{})
		if err != nil {
			// TODO(meain): should this be a recoverable error?
			return nil, clues.Wrap(err, "getting site").
				With(
					"channel_id", ptr.Val(ch.GetId()),
					"document_web_url", documentWebURL,
					"site_web_url", siteWebURL)
		}

		sites = append(sites, site)
	}

	return sites, nil
}

func (c Groups) GetRootSite(
	ctx context.Context,
	identifier string,
) (models.Siteable, error) {
	service, err := c.Service()
	if err != nil {
		return nil, err
	}

	resp, err := service.
		Client().
		Groups().
		ByGroupId(identifier).
		Sites().
		BySiteId("root").
		Get(ctx, nil)
	if err != nil {
		return nil, clues.Wrap(err, "getting root site for group")
	}

	return resp, graph.Stack(ctx, err).OrNil()
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

// ValidateGroup ensures the item is a Groupable, and contains the necessary
// identifiers that we handle with all groups.
func ValidateGroup(item models.Groupable) error {
	if item.GetId() == nil {
		return clues.New("missing ID")
	}

	if item.GetDisplayName() == nil {
		return clues.New("missing display name")
	}

	return nil
}

func OnlyTeams(ctx context.Context, gs []models.Groupable) []models.Groupable {
	var teams []models.Groupable

	for _, g := range gs {
		if IsTeam(ctx, g) {
			teams = append(teams, g)
		}
	}

	return teams
}

func IsTeam(ctx context.Context, mg models.Groupable) bool {
	log := logger.Ctx(ctx)

	if mg.GetAdditionalData()[ResourceProvisioningOptions] == nil {
		return false
	}

	val, _ := tform.AnyValueToT[[]any](ResourceProvisioningOptions, mg.GetAdditionalData())
	for _, v := range val {
		s, err := str.AnyToString(v)
		if err != nil {
			log.Debug("could not be converted to string value: ", ResourceProvisioningOptions)
			continue
		}

		if s == teamsAdditionalDataLabel {
			return true
		}
	}

	return false
}

// GetIDAndName looks up the group matching the given ID, and returns
// its canonical ID and the name.
func (c Groups) GetIDAndName(
	ctx context.Context,
	groupID string,
	cc CallConfig,
) (string, string, error) {
	s, err := c.GetByID(ctx, groupID, cc)
	if err != nil {
		return "", "", err
	}

	return ptr.Val(s.GetId()), ptr.Val(s.GetDisplayName()), nil
}
