package readers

import (
	"bytes"
	"encoding/binary"
	"io"
	"os"
	"unsafe"

	"github.com/alcionai/clues"
)

// persistedSerializationVersion is the size of the serialization version in
// storage.
//
// The current on-disk format of this field is written in big endian. The
// highest bit denotes if the item is empty because it was deleted between the
// time we told the storage about it and when we needed to get data for it. The
// lowest two bytes are the version number. All other bits are reserved for
// future use.
//
//	MSB 31         30             16        8       0 LSB
//	     +----------+----+---------+--------+-------+
//	     | del flag |   reserved   | version number |
//	     +----------+----+---------+--------+-------+
type persistedSerializationVersion = uint32

// SerializationVersion is the in-memory size of the version number that gets
// added to the persisted serialization version.
//
// Right now it's only a uint16 but we can expand it to be larger so long as the
// expanded size doesn't clash with the flags in the high-order bits.
type SerializationVersion uint16

// DefaultSerializationVersion is the current (default) version number for all
// services. As services evolve their storage format they should begin tracking
// their own version numbers separate from other services.
const DefaultSerializationVersion SerializationVersion = 1

const (
	VersionFormatSize                               = int(unsafe.Sizeof(persistedSerializationVersion(0)))
	delInFlightMask   persistedSerializationVersion = 1 << ((VersionFormatSize * 8) - 1)
)

// SerializationFormat is a struct describing serialization format versions and
// flags to add for this item.
type SerializationFormat struct {
	Version     SerializationVersion
	DelInFlight bool
}

// NewVersionedBackupReader creates a reader that injects format into the first
// bytes of the returned data. After format has been returned, data is returned
// from baseReaders in the order they're passed in.
func NewVersionedBackupReader(
	format SerializationFormat,
	baseReaders ...io.ReadCloser,
) (io.ReadCloser, error) {
	if format.DelInFlight && len(baseReaders) > 0 {
		// This is a conservative check, but we can always loosen it later on if
		// needed. At the moment we really don't expect any data if the item was
		// deleted.
		return nil, clues.New("item marked deleted but has reader(s)")
	}

	formattedVersion := persistedSerializationVersion(format.Version)
	if format.DelInFlight {
		formattedVersion |= delInFlightMask
	}

	formattedBuf := make([]byte, VersionFormatSize)
	binary.BigEndian.PutUint32(formattedBuf, formattedVersion)

	versionReader := io.NopCloser(bytes.NewReader(formattedBuf))

	// Need to add readers individually because types differ.
	allReaders := make([]io.Reader, 0, len(baseReaders)+1)
	allReaders = append(allReaders, versionReader)

	for _, r := range baseReaders {
		allReaders = append(allReaders, r)
	}

	res := &versionedBackupReader{
		baseReaders: append([]io.ReadCloser{versionReader}, baseReaders...),
		combined:    io.MultiReader(allReaders...),
	}

	return res, nil
}

type versionedBackupReader struct {
	// baseReaders is a reference to the original readers so we can close them.
	baseReaders []io.ReadCloser
	// combined is the reader that will return all data.
	combined io.Reader
}

func (vbr *versionedBackupReader) Read(p []byte) (int, error) {
	if vbr.combined == nil {
		return 0, os.ErrClosed
	}

	n, err := vbr.combined.Read(p)
	if err == io.EOF {
		// Golang doesn't allow wrapping of EOF. If we wrap it other things start
		// thinking it's an actual error.
		return n, err
	}

	return n, clues.Stack(err).OrNil()
}

func (vbr *versionedBackupReader) Close() error {
	if vbr.combined == nil {
		return nil
	}

	vbr.combined = nil

	var errs *clues.Err

	for i, r := range vbr.baseReaders {
		if err := r.Close(); err != nil {
			errs = clues.Stack(
				errs,
				clues.Wrap(err, "closing reader").With("reader_index", i))
		}
	}

	vbr.baseReaders = nil

	return errs.OrNil()
}

// NewVersionedRestoreReader wraps baseReader and provides easy access to the
// SerializationFormat info in the first bytes of the data contained in
// baseReader.
func NewVersionedRestoreReader(
	baseReader io.ReadCloser,
) (*VersionedRestoreReader, error) {
	versionBuf := make([]byte, VersionFormatSize)

	// Loop to account for the unlikely case where we get a short read.
	for read := 0; read < VersionFormatSize; {
		n, err := baseReader.Read(versionBuf[read:])
		if err != nil {
			return nil, clues.Wrap(err, "reading serialization version")
		}

		read += n
	}

	formattedVersion := binary.BigEndian.Uint32(versionBuf)

	return &VersionedRestoreReader{
		baseReader: baseReader,
		format: SerializationFormat{
			Version:     SerializationVersion(formattedVersion),
			DelInFlight: (formattedVersion & delInFlightMask) != 0,
		},
	}, nil
}

type VersionedRestoreReader struct {
	baseReader io.ReadCloser
	format     SerializationFormat
}

func (vrr *VersionedRestoreReader) Read(p []byte) (int, error) {
	n, err := vrr.baseReader.Read(p)
	if err == io.EOF {
		// Golang doesn't allow wrapping of EOF. If we wrap it other things start
		// thinking it's an actual error.
		return n, err
	}

	return n, clues.Stack(err).OrNil()
}

func (vrr *VersionedRestoreReader) Close() error {
	return clues.Stack(vrr.baseReader.Close()).OrNil()
}

func (vrr VersionedRestoreReader) Format() SerializationFormat {
	return vrr.format
}
