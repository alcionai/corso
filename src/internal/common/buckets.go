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

// NormalizePrefix ensures that a bucket prefix is always treated as
// object store folder prefix.
func NormalizePrefix(p string) string {
	tp := strings.TrimRight(p, "/")

	if len(tp) > 0 {
		tp = tp + "/"
	}

	return tp
}
