package path

import (
	"github.com/alcionai/clues"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/internal/common/tform"
)

// ---------------------------------------------------------------------------
// Tuple
// ---------------------------------------------------------------------------

// ServiceResource holds a service + resource tuple.  The tuple implies
// that the resource owns some data in the given service.
type ServiceResource struct {
	Service           ServiceType
	ProtectedResource string
}

func MakeServiceResource(
	st ServiceType,
	protectedResource string,
) ServiceResource {
	return ServiceResource{
		Service:           st,
		ProtectedResource: protectedResource,
	}
}

func (sr ServiceResource) validate() error {
	if len(sr.ProtectedResource) == 0 {
		return clues.Stack(errMissingSegment, clues.New("protected resource"))
	}

	return nil
}

// ---------------------------------------------------------------------------
// Exported Helpers
// ---------------------------------------------------------------------------

// NewServiceResources is a lenient constructor for building a
// new []ServiceResource.  It allows the caller to pass in any
// number of arbitrary values, but will require the following:
// 1. even values must be path.ServiceType typed
// 2. odd values must be string typed
// 3. a non-zero, even number of values must be provided
func NewServiceResources(elems ...any) ([]ServiceResource, error) {
	if len(elems) == 0 {
		return nil, clues.New("missing service resources")
	}

	if len(elems)%2 == 1 {
		return nil, clues.New("odd number of service resources")
	}

	srs := make([]ServiceResource, 0, len(elems)/2)

	for i, j := 0, 1; i < len(elems); i, j = i+2, j+2 {
		srv, err := tform.AnyToT[ServiceType](elems[i])
		if err != nil {
			return nil, clues.Wrap(err, "service")
		}

		pr, err := str.AnyToString(elems[j])
		if err != nil {
			return nil, clues.Wrap(err, "protected resource")
		}

		srs = append(srs, MakeServiceResource(srv, pr))
	}

	return srs, nil
}

func ServiceResourcesToResources(srs []ServiceResource) []string {
	prs := make([]string, len(srs))

	for i := range srs {
		prs[i] = srs[i].ProtectedResource
	}

	return prs
}

func ServiceResourcesToServices(srs []ServiceResource) []ServiceType {
	sts := make([]ServiceType, len(srs))

	for i := range srs {
		sts[i] = srs[i].Service
	}

	return sts
}

func ServiceResourcesMatchServices(srs []ServiceResource, sts []ServiceType) bool {
	return slices.EqualFunc(srs, sts, func(sr ServiceResource, st ServiceType) bool {
		return sr.Service == st
	})
}

func ServiceResourcesToElements(srs []ServiceResource) Elements {
	es := make(Elements, 0, len(srs)*2)

	for _, tuple := range srs {
		es = append(es, tuple.Service.String())
		es = append(es, tuple.ProtectedResource)
	}

	return es
}

// ---------------------------------------------------------------------------
// Unexported Helpers
// ---------------------------------------------------------------------------

// elementsToServiceResources turns as many pairs of elems as possible
// into ServiceResource tuples.  Elems must begin with a service, but
// may contain more entries than there are service/resource pairs.
// This transformer will continue consuming elements until it finds an
// even-numbered index that cannot be cast to a ServiceType.
// Returns the serviceResource pairs, the first index containing element
// that is not part of a service/resource pair, and an error if elems is
// len==0 or contains no services.
func elementsToServiceResources(elems Elements) ([]ServiceResource, int, error) {
	if len(elems) == 0 {
		return nil, -1, clues.Wrap(errMissingSegment, "service")
	}

	var (
		srs = make([]ServiceResource, 0)
		i   int
	)

	for j := 1; i < len(elems); i, j = i+2, j+2 {
		service := toServiceType(elems[i])
		if service == UnknownService {
			if i == 0 {
				return nil, -1, clues.Wrap(errMissingSegment, "service")
			}

			break
		}

		srs = append(srs, ServiceResource{service, elems[j]})
	}

	return srs, i, nil
}

// checks for the following:
// 1. each ServiceResource is valid
// 2. if len(srs) > 1, srs[i], srs[i+1] pass subservice checks.
func validateServiceResources(srs []ServiceResource) error {
	switch len(srs) {
	case 0:
		return clues.Stack(errMissingSegment, clues.New("service"))
	case 1:
		return srs[0].validate()
	}

	for i, tuple := range srs {
		if err := tuple.validate(); err != nil {
			return err
		}

		if i+1 >= len(srs) {
			continue
		}

		if err := ValidateServiceAndSubService(tuple.Service, srs[i+1].Service); err != nil {
			return err
		}
	}

	return nil
}

// makes a copy of the slice with all of the Services swapped for their
// metadata service countterparts.
func toMetadataServices(srs []ServiceResource) []ServiceResource {
	msrs := make([]ServiceResource, 0, len(srs))

	for _, sr := range srs {
		msr := sr
		metadataService := UnknownService

		switch sr.Service {
		// TODO: add groups
		case ExchangeService:
			metadataService = ExchangeMetadataService
		case OneDriveService:
			metadataService = OneDriveMetadataService
		case SharePointService:
			metadataService = SharePointMetadataService
		}

		msr.Service = metadataService
		msrs = append(msrs, msr)
	}

	return msrs
}
