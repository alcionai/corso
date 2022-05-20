package s3

import (
	"context"

	"github.com/kopia/kopia/repo/blob"
	kopiaS3 "github.com/kopia/kopia/repo/blob/s3"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/kopia"
)

func Initializer(
	ctx context.Context,
	bucket, accessKey, secretKey string,
) (kopia.KopiaWrapper, error) {
	bst, err := blobStorage(ctx, bucket, accessKey, secretKey)
	if err != nil {
		return kopia.KopiaWrapper{}, errors.Wrap(err, "making s3 storage initializer")
	}
	return kopia.NewInitializer(bst, nil, nil), nil
}

func Connector(
	ctx context.Context,
	bucket, accessKey, secretKey string,
) (kopia.KopiaWrapper, error) {
	bst, err := blobStorage(ctx, bucket, accessKey, secretKey)
	if err != nil {
		return kopia.KopiaWrapper{}, errors.Wrap(err, "making s3 storage connector")
	}
	return kopia.NewConnector(bst, nil), nil
}

func blobStorage(
	ctx context.Context,
	bucket, accessKey, secretKey string,
) (blob.Storage, error) {
	opts := kopiaS3.Options{
		BucketName:      bucket,
		AccessKeyID:     accessKey,
		SecretAccessKey: secretKey,
	}
	return kopiaS3.New(ctx, &opts)
}
