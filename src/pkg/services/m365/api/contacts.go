package api

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kjson "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
)

// ---------------------------------------------------------------------------
// controller
// ---------------------------------------------------------------------------

func (c Client) Contacts() Contacts {
	return Contacts{c}
}

// Contacts is an interface-compliant provider of the client.
type Contacts struct {
	Client
}

// ---------------------------------------------------------------------------
// containers
// ---------------------------------------------------------------------------

// CreateContainer makes a contact folder with the displayName of folderName.
// If successful, returns the created folder object.
func (c Contacts) CreateContainer(
	ctx context.Context,
	// parentContainerID needed for iface, doesn't apply to contacts
	userID, _, containerName string,
) (graph.Container, error) {
	body := models.NewContactFolder()
	body.SetDisplayName(ptr.To(containerName))

	mdl, err := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		ContactFolders().
		Post(ctx, body, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating contact folder")
	}

	return mdl, nil
}

// DeleteContainer deletes the ContactFolder associated with the M365 ID if permissions are valid.
func (c Contacts) DeleteContainer(
	ctx context.Context,
	userID, containerID string,
) error {
	// deletes require unique http clients
	// https://github.com/alcionai/corso/issues/2707
	srv, err := NewService(c.Credentials)
	if err != nil {
		return graph.Stack(ctx, err)
	}

	err = srv.
		Client().
		Users().
		ByUserId(userID).
		ContactFolders().
		ByContactFolderId(containerID).
		Delete(ctx, nil)
	if err != nil {
		return graph.Stack(ctx, err)
	}

	return nil
}

// prefer GetContainerByID where possible.
// use this only in cases where the models.ContactFolderable
// is required.
func (c Contacts) GetFolder(
	ctx context.Context,
	userID, containerID string,
) (models.ContactFolderable, error) {
	config := &users.ItemContactFoldersContactFolderItemRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemContactFoldersContactFolderItemRequestBuilderGetQueryParameters{
			Select: idAnd(displayName, parentFolderID),
		},
	}

	resp, err := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		ContactFolders().
		ByContactFolderId(containerID).
		Get(ctx, config)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return resp, nil
}

// interface-compliant wrapper of GetFolder
func (c Contacts) GetContainerByID(
	ctx context.Context,
	userID, containerID string,
) (graph.Container, error) {
	return c.GetFolder(ctx, userID, containerID)
}

// GetContainerByName fetches a folder by name
func (c Contacts) GetContainerByName(
	ctx context.Context,
	// parentContainerID needed for iface, doesn't apply to contacts
	userID, _, containerName string,
) (graph.Container, error) {
	filter := fmt.Sprintf("displayName eq '%s'", containerName)
	options := &users.ItemContactFoldersRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemContactFoldersRequestBuilderGetQueryParameters{
			Filter: &filter,
		},
	}

	ctx = clues.Add(ctx, "container_name", containerName)

	resp, err := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		ContactFolders().
		Get(ctx, options)
	if err != nil {
		return nil, graph.Stack(ctx, err).WithClues(ctx)
	}

	gv := resp.GetValue()

	if len(gv) == 0 {
		return nil, clues.New("container not found").WithClues(ctx)
	}

	// We only allow the api to match one container with the provided name.
	// Return an error if multiple container exist (unlikely) or if no container
	// is found.
	if len(gv) != 1 {
		return nil, clues.New("unexpected number of folders returned").
			With("returned_container_count", len(gv)).
			WithClues(ctx)
	}

	// Sanity check ID and name
	container := gv[0]

	if err := graph.CheckIDAndName(container); err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	return container, nil
}

func (c Contacts) PatchFolder(
	ctx context.Context,
	userID, containerID string,
	body models.ContactFolderable,
) error {
	_, err := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		ContactFolders().
		ByContactFolderId(containerID).
		Patch(ctx, body, nil)
	if err != nil {
		return graph.Wrap(ctx, err, "patching contact folder")
	}

	return nil
}

// ---------------------------------------------------------------------------
// items
// ---------------------------------------------------------------------------

// GetItem retrieves a Contactable item.
func (c Contacts) GetItem(
	ctx context.Context,
	userID, itemID string,
	immutableIDs bool,
	_ *fault.Bus, // no attachments to iterate over, so this goes unused
) (serialization.Parsable, *details.ExchangeInfo, error) {
	options := &users.ItemContactsContactItemRequestBuilderGetRequestConfiguration{
		Headers: newPreferHeaders(preferImmutableIDs(immutableIDs)),
	}

	cont, err := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		Contacts().
		ByContactId(itemID).
		Get(ctx, options)
	if err != nil {
		return nil, nil, graph.Stack(ctx, err)
	}

	return cont, ContactInfo(cont), nil
}

func (c Contacts) PostItem(
	ctx context.Context,
	userID, containerID string,
	body models.Contactable,
) (models.Contactable, error) {
	itm, err := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		ContactFolders().
		ByContactFolderId(containerID).
		Contacts().
		Post(ctx, body, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating contact")
	}

	return itm, nil
}

func (c Contacts) DeleteItem(
	ctx context.Context,
	userID, itemID string,
) error {
	// deletes require unique http clients
	// https://github.com/alcionai/corso/issues/2707
	srv, err := c.Service()
	if err != nil {
		return graph.Stack(ctx, err)
	}

	err = srv.
		Client().
		Users().
		ByUserId(userID).
		Contacts().
		ByContactId(itemID).
		Delete(ctx, nil)
	if err != nil {
		return graph.Wrap(ctx, err, "deleting contact")
	}

	return nil
}

// ---------------------------------------------------------------------------
// Serialization
// ---------------------------------------------------------------------------

func BytesToContactable(bytes []byte) (models.Contactable, error) {
	v, err := createFromBytes(bytes, models.CreateContactFromDiscriminatorValue)
	if err != nil {
		return nil, clues.Wrap(err, "deserializing bytes to contact")
	}

	return v.(models.Contactable), nil
}

func (c Contacts) Serialize(
	ctx context.Context,
	item serialization.Parsable,
	userID, itemID string,
) ([]byte, error) {
	contact, ok := item.(models.Contactable)
	if !ok {
		return nil, clues.New(fmt.Sprintf("item is not a Contactable: %T", item))
	}

	ctx = clues.Add(ctx, "item_id", ptr.Val(contact.GetId()))
	writer := kjson.NewJsonSerializationWriter()

	defer writer.Close()

	if err := writer.WriteObjectValue("", contact); err != nil {
		return nil, graph.Stack(ctx, err)
	}

	bs, err := writer.GetSerializedContent()
	if err != nil {
		return nil, graph.Wrap(ctx, err, "serializing contact")
	}

	return bs, nil
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func ContactInfo(contact models.Contactable) *details.ExchangeInfo {
	name := ptr.Val(contact.GetDisplayName())
	created := ptr.Val(contact.GetCreatedDateTime())

	return &details.ExchangeInfo{
		ItemType:    details.ExchangeContact,
		ContactName: name,
		Created:     created,
		Modified:    ptr.OrNow(contact.GetLastModifiedDateTime()),
	}
}

func contactCollisionKeyProps() []string {
	return idAnd(givenName, surname, emailAddresses, mobilePhone)
}

// ContactCollisionKey constructs a key from the contactable's creation time and either displayName or given+surname.
// collision keys are used to identify duplicate item conflicts for handling advanced restoration config.
func ContactCollisionKey(item models.Contactable) string {
	if item == nil {
		return ""
	}

	var (
		given  = ptr.Val(item.GetGivenName())
		sur    = ptr.Val(item.GetSurname())
		emails = item.GetEmailAddresses()
		email  string
		phone  = ptr.Val(item.GetMobilePhone())
	)

	for _, em := range emails {
		email += ptr.Val(em.GetAddress())
	}

	return given + sur + email + phone
}
