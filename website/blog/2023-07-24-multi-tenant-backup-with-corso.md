---
slug: multi-tenant-backup-with-corso
title: "Using Corso to Build a Self-Hosted Multi-Tenant Office 365 Backup Solution"
description: ""
authors:
  - name: meuchels
    title: Corso Community Member, IT Lead
    url: https://github.com/meuchels
    image_url: https://avatars.githubusercontent.com/u/77171293?v=4
tags: [corso, microsoft 365, backups, msp, multi-tenant]
date: 2023-07-24
image: ./images/data-center.jpg
---

![A woman engineer holding a laptop in front of a data center](./images/data-center.jpg)

This community-contributed blog post shows how MSPs in the community are using Corso to build out a multi-tenant backup
solution for their Microsoft 365 customers. If you have questions, come find the author (or us) on
[Discord](https://www.alcion.ai/discord).

<!-- truncate -->

First of all, I offer a fully managed backup solution. My clients have no access to the backup software or the data. I
require them to request recovery in a ticket. For my use case I have a self-hosted instance of MinIO that I won't be
going over but there is [another blog post on it](./2023-2-4-where-to-store-corso.md#local-s3-testing). I will show the
layout and an example of how to backup emails using the exchange option in Corso.

## Organizing the file structure on your storage

I wanted my S3 bucket to be laid out in the following fashion utilizing 1 bucket with prefixes for the tenants. For now,
all I did is create a bucket with access to a user for corso. While it's possible to use a single bucket and use prefix
paths per tenant within it, I didn't do that in my setup. The will be generated later with the backup initialization.

```bash
BUCKET
  tenant1-exchange
  tenant1-onedrive
  tenant1-sharepoint
  tenant2-exchange
  tenant2-onedrive
  tenant2-sharepoint
```

If I don’t backup a particular service for a client, it will be clear by looking at whether the bucket exists or not.

I have a short name for each tenant to differentiate them.

## The backup compute server layout

I utilize Ubuntu Server for this task. In my setup, everything is done as the root user. I have put the corso
executable in `/opt/corso/` and will be building everything under there. Here is the folder layout before I go into
usage.

```bash
# For logs
/opt/corso/logs
# For config files
/opt/corso/toml
# Root of the scripts folder
/opt/corso/scripts
# For building out the environment loaders
/opt/corso/scripts/environments
# For building out the backup scripts
/opt/corso/scripts/back-available
# For adding a link to the backups that will be run
/opt/corso/scripts/back-active
```

## The environment files

For [configuration](../../docs/setup/configuration/), create an environment file
`/opt/corso/scripts/environments/blank-exchange` with the following content for a template. You can copy this template
to `<tenantshortname>-exchange` in the same folder to setup your client exchange backup environment.

```bash
#####################################
#EDIT THIS SECTION TO MEET YOUR NEEDS
#####################################

# this is a shortname for your tenant to setup storage
export tenantshortname=""

# this is your tenant info from the app setup on O365
export AZURE_TENANT_ID=""
export AZURE_CLIENT_ID=""
export AZURE_CLIENT_SECRET=""

# this is your credentials for your s3 storage
export AWS_ACCESS_KEY_ID="<S3-STORAGE-USERNAME>"
export AWS_SECRET_ACCESS_KEY="<S3-STORAGE-PASSWORD"

# this sets your encryption key for your backups
export CORSO_PASSPHRASE="<ENCRYPTION-PASSWORD>"

# this is your s3 storage endpoint
export s3endpoint="<YOUR-S3-STORAGE-SERVER>"
export bucket="<YOUR-BUCKET>"

####################################
#END EDIT
####################################

export configfile=/opt/corso/toml/${tenantshortname}-exchange.toml
```

## The backup scripts

Create a backup script `/opt/corso/scripts/back-available/blank-exchange` with the following content for an exchange
backup template. This can be copied to `tenantshortname-exchange` in the same directory for creating the backup script.

```bash
#!/bin/bash

##############Begin Edit###

# change blank to tenant short name
source /opt/corso/scripts/environments/blank-exchange

##############End Edit###

# create runtime variables
logfilename="/opt/corso/log/${tenantshortname}-exchange/$(date +'%Y-%m-%d-%H%M%S').log" runcorso="/opt/corso/corso"

# init bucket
$runcorso repo init s3 --bucket $bucket --prefix ${tenantshortname}_exchange --endpoint $s3endpoint \
    --log-file $logfilename --config-file $configfile --hide-progress
$runcorso repo connect s3 --bucket $bucket --log-file $logfilename --config-file $configfile --hide- progress

# run Backup
$runcorso backup create exchange --mailbox '*' --log-file $logfilename --config-file $configfile --hide- progress
```

Use this folder for a working directory and create a symbolic link to the scripts that you want to activate in `/opt/corso/scripts/back-active/`.

## The backup runner

To fire it all off, I have a `backuprunner.sh` script that cycles through the `/opt/corso/scripts/back-active` folder
and is scheduled with a `cron` job to run at your interval. You can put it wherever you want but I put it in the scripts
folder as well so I know where everything is. Add your email address. This relies on the Linux mail package, you will
have to accept the email from it.

```bash
#!/bin/bash

# Directory containing the scripts
script_directory="/opt/corso/scripts/back-active"

# Email configuration
recipient="<YOUR-EMAIL-ADDRESS>"
subject_prefix="Backup Job: "

# Iterate over all scripts in the directory
for script_file in "$script_directory"/*; do
  # Run the script and capture the output
  output=$("$script_file")

  # Prepare email subject
  script_name=$(basename "$script_file")
  subject="$subject_prefix$script_name"

  # Send an email with the script output
  echo "$output" | mail -s "$subject" "$recipient"
done
```

Once your backups have completed, you can load the environments using the command
`source /opt/corso/scripts/environments/tenant-exchange` to load the variables and access the backups of that tenant. Be
sure to specify the `–config-file` flag.

```bash
source /opt/corso/scripts/environments/tenant-exchange
/opt/corso/corso backup list exchange --config-file $configfile
```

Don’t forget to backup your /opt/corso folder once in a while to save your scripts!
