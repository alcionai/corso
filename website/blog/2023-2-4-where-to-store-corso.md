---
slug: where-to-store-corso
title: "Where to store your Corso Repository"
description: "Storage Options for Corso"
authors: nica
tags: [corso, microsoft 365, backups, S3]
date: 2023-2-4
image: ./images/boxes_web.jpeg
---

![image of a large number of packing boxes](./images/boxes_web.jpeg)

We all know that Corso is a free and open-source tool for creating backups of your Microsoft 365 data. But where does
that data go?

Corso creates a repository to store your backups, and the default in our documentation is to send that data to AWS S3.
It is possible however to back up to any object storage system that has an S3-compatible API. Let’s talk about some options.

## **S3-Compatible Object Storage**

A number of other cloud providers aren’t the 500-pound gorilla of AWS but still offer an S3-compatible API.
Some of them include:

- Google Cloud: One of the largest cloud providers in the world, Google offers
[an S3-compatible API](https://cloud.google.com/storage/docs/interoperability) on top of its Google Cloud Storage (GCS) offering.
- Backblaze: Known for its deep analysis of hard drive failure statistics, Backblaze offers an S3-compatible API for its
B2 Cloud Storage product. They also make the bold claim of costing [significantly less than AWS S3](https://www.backblaze.com/b2/cloud-storage-pricing.html)
(I haven’t evaluated these claims) but Glacier is still cheaper (see below for more details)
- HPE: HPE Greenlake offers S3 compatibility and claims superior performance over S3. If you want to get a sense of how
‘Enterprise’ HPE is, the best writeup I could find of their offerings is [available only as a PDF](https://www.hpe.com/us/en/collaterals/collateral.a50006216.Create-value-from-data-2C-at-scale-E2-80-93-HPE-GreenLake-for-Scality-solution-brief.html).
- Wasabi: Another very popular offering, Wasabi has very good integration with existing AWS components at a reduced
cost but watch out for the minimum monthly storage charge and the minimum storage duration policy!

This is an incomplete list, but any S3-compliant storage with immediate retrieval is expected to work with Corso today.

## Local S3 Testing

In my own testing, I use [MinIO](https://min.io/) to create a local S3 server and bucket. This has some great advantages
including extremely low latency for testing. Unless you have a significant hardware and software investment to ensure
reliable storage and compute infrastructure, you probably do not want to rely on a simple MinIO setup as your primary
backup location, but it’s a great way to do a zero-cost test backup that you totally control.

While there are a number of in-depth tutorials on how to use
[MinIO to run a local S3 server](https://simonjcarr.medium.com/running-s3-object-storage-locally-with-minio-f50540ffc239),
here’s the single script that can run a non-production instance of MinIO within a Docker container (you’ll need Docker
and the AWS CLI as prerequisites) and get you started with Corso quickly:

```bash
mkdir -p ~\s/minio/data

docker run \
   -p 9000:9000 \
   -p 9090:9090 \
   --name minio \
   -v ~/minio/data:/data \
   -e "MINIO_ROOT_USER=ROOTNAME" \
   -e "MINIO_ROOT_PASSWORD=CHANGEME123" \
   quay.io/minio/minio server /data --console-address ":9090"
```

In a separate window, create a bucket (`corso-backup`) for use with Corso.

```bash
export AWS_ACCESS_KEY_ID=ROOTNAME
export AWS_SECRET_ACCESS_KEY=CHANGEME123

aws s3api create-bucket --bucket corso-backup --endpoint=http://127.0.0.1:9000
```

To connect Corso to a local MinIO server with `[corso repo init](https://corsobackup.io/docs/cli/corso-repo-init-s3/)`
you’ll want to pass the `--disable-tls` flag so that it will accept an `http` connection

## Reducing Cost With S3 Storage Classes

AWS S3 offers [storage classes](https://aws.amazon.com/s3/storage-classes/) for a variety of different use cases and
Corso can leverage a number of them, but not all, to reduce the cost of storing data in the cloud.

By default, Corso works hard to reduce its data footprint. It will compress and deduplicate data at source to reduce the
amount of storage used as well as the amount of network traffic when writing to object storage. Corso also combines
different emails, attachments, etc. into larger objects to make it more cost-effective by reducing the number of API
calls and increasing network throughput as well as making Corso data eligible and cost-effective for some of the other
storage classes described below.

Stepping away from the default S3 offering (S3 Standard), S3 offers a number of different Glacier (cheap and deep)
storage classes that can help to further reduce the cost for backup and archival workloads. Within the storage classes,
Corso today supports Glacier Instant Retrieval but, because of user responsiveness and metadata requirements, not the
other Glacier variants.

Glacier Instant Retrieval should provide the best price performance for a  backup workload as backup data blobs are
typically written once, with occasional re-compacting, and read infrequently in the case of restore. One should note
that recommendations such as these are always workload dependent and should be verified for your use case. For example,
we would not recommend Glacier Instant Retrieval if you are constantly testing large restores or have heavy churn in your backups and
limited retention. However, for most typical backup workloads (write mostly, read rarely), Glacier Instant Retrieval
should work just fine and deliver the best price-performance ratio. 

You can configure your storage to use Glacier Instant Retrieval by adding a `.storageconfig` file to the root of your
bucket. If you have configured Corso to store the repository in a subfolder within your bucket by adding a
`prefix = '[folder name]'` configuration, the `.storageconfig` should go within that folder in the bucket.

Here’s an example:

```json
{
   "blobOptions": [
     { "prefix": "p", "storageClass": "GLACIER_IR" },
     { "storageClass": "STANDARD" }
  ]
}
```

The `"prefix": "p"` parameter is unrelated to the subfolder `prefix` setting mentioned above. It simply tells Corso to
use the selected storage class for data blobs (named with a `p` prefix). By default, all other objects including
metadata and indices will use the standard storage tier.

We would love to hear from you if you’ve deployed Corso with a different storage class, an object storage provider not
listed above, or have questions about how to best cost-optimize your setup. Come [find us on Discord](https://discord.gg/63DTTSnuhT)!
