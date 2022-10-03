package common

import "strings"

// NormalizeBuckets ensures that bucket names are cleaned and
// standardized according to the downstream needs of minio.
//
// Any url prefixing to location the bucket (ex: s3://bckt)
// will be removed, leaving only the bucket name (bckt).
// Corso should only utilize or store the normalized name.
func NormalizeBucket(b string) string {
	return strings.TrimPrefix(b, "s3://")
}
