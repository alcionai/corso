package s3

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/alcionai/clues"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const noLockCode = "NoSuchObjectLockConfiguration"

var ErrNoRetention = clues.New("missing ObjectLock configuration")

type Options struct {
	BucketName string
	Prefix     string

	Endpoint       string
	DoNotUseTLS    bool
	DoNotVerifyTLS bool

	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string

	Region string
}

type ObjInfo struct {
	Key     string
	Version string
}

type Client struct {
	client     *minio.Client
	bucketName string
	prefix     string
}

// Below init code scavenged from
// https://github.com/kopia/kopia/blob/83b88d8bbf60f4eb17b40dea440c08b594b9e5d3/repo/blob/s3/s3_storage.go

func getCustomTransport(opt *Options) *http.Transport {
	if opt.DoNotVerifyTLS {
		return &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	return http.DefaultTransport.(*http.Transport).Clone()
}

func New(opt *Options) (*Client, error) {
	creds := credentials.NewChainCredentials(
		[]credentials.Provider{
			&credentials.Static{
				Value: credentials.Value{
					AccessKeyID:     opt.AccessKeyID,
					SecretAccessKey: opt.SecretAccessKey,
					SessionToken:    opt.SessionToken,
					SignerType:      credentials.SignatureV4,
				},
			},
			&credentials.EnvAWS{},
			&credentials.IAM{
				Client: &http.Client{
					Transport: http.DefaultTransport,
				},
			},
		})

	if len(opt.BucketName) == 0 {
		return nil, clues.New("empty bucket name")
	}

	minioOpts := &minio.Options{
		Creds:  creds,
		Secure: !opt.DoNotUseTLS,
		Region: opt.Region,
	}

	minioOpts.Transport = getCustomTransport(opt)

	cli, err := minio.New(opt.Endpoint, minioOpts)
	if err != nil {
		return nil, clues.Wrap(err, "creating client")
	}

	return &Client{
		client:     cli,
		bucketName: opt.BucketName,
		prefix:     opt.Prefix,
	}, nil
}

func maybeAddObj(prefix string, obj ObjInfo, m map[string]ObjInfo) {
	// We've already found an item to check.
	if _, ok := m[prefix]; ok {
		fmt.Printf("prefix %s already matched\n", prefix)
		return
	}

	m[prefix] = obj
}

func (c *Client) ListUntilAllFound(
	ctx context.Context,
	wantedPrefixes []string,
	alsoFindDeleted bool,
) (map[string]ObjInfo, map[string]ObjInfo, error) {
	notDeleted := map[string]ObjInfo{}
	deleted := map[string]ObjInfo{}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	objChan := c.client.ListObjects(
		ctx,
		c.bucketName,
		minio.ListObjectsOptions{
			Prefix:       c.prefix,
			WithVersions: alsoFindDeleted,
		})

	for obj := range objChan {
		if err := obj.Err; err != nil {
			return notDeleted, deleted, clues.Wrap(err, "searching object list")
		}

		fmt.Printf("looking at object %q\n", obj.Key)

		var (
			matchesPrefix string
			trimmedKey    = obj.Key[len(c.prefix):]
		)

		for _, prefix := range wantedPrefixes {
			fmt.Printf("checking %q for prefix %q\n", trimmedKey, prefix)
			if strings.HasPrefix(trimmedKey, prefix) {
				fmt.Println("matched prefix")
				matchesPrefix = prefix
				break
			}
		}

		// Item doesn't have a prefix that matches anything we want.
		if len(matchesPrefix) == 0 {
			continue
		}

		objI := ObjInfo{
			Key:     obj.Key,
			Version: obj.VersionID,
		}

		// If the item is the latest version then it's considered not deleted. If
		// it's not the latest then it's likely deleted. Not asking for versions
		// always sets IsLatest to false.
		if !alsoFindDeleted || obj.IsLatest {
			maybeAddObj(matchesPrefix, objI, notDeleted)
		} else {
			maybeAddObj(matchesPrefix, objI, deleted)
		}

		// Break if we've found all the non-deleted items we're looking for and
		// either we're not looking for deleted items or we've found all the deleted
		// items we're looking for.
		if len(notDeleted) == len(wantedPrefixes) &&
			(!alsoFindDeleted || len(deleted) == len(wantedPrefixes)) {
			break
		}
	}

	return notDeleted, deleted, nil
}

func (c *Client) ObjectRetention(
	ctx context.Context,
	obj ObjInfo,
) (*minio.RetentionMode, *time.Time, error) {
	mode, retainUntil, err := c.client.GetObjectRetention(
		ctx,
		c.bucketName,
		obj.Key,
		obj.Version)
	if err != nil {
		// Unfortunately they don't have a sentinel type that we can compare
		// against. We can check the error code though.
		var e minio.ErrorResponse

		if errors.As(err, &e) {
			if e.Code == noLockCode {
				return nil, nil, clues.Stack(ErrNoRetention, err)
			}
		}
	}

	return mode, retainUntil, clues.Wrap(err, fmt.Sprintf(
		"getting object (key) %q (versionID) %q",
		obj.Key,
		obj.Version),
	).OrNil()
}
