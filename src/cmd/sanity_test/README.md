# Sanity Tests

A set of high-level blackbox tests designed to ensure the robustness of Corso's restore and export functionality.

## Overview

The purpose of sanity tests is to evaluate the restore and export capabilities of Corso, using OneDrive as an example. Here's a breakdown of how it works:

1. **Setup:** Create a new folder in OneDrive, populate it with subfolders and files, and perform the initial backup.

2. **Restore Testing:**
   - Restore the backup to a different folder.
   - Compare the original folder with the restored folder (eg: checking file names and sizes).
   - Any mismatches are treated as failures.

3. **Export Testing:**
   - Export the backup to a local folder.
   - Compare the exported data with the original data in OneDrive.

## How to Run Locally

Before running the tests locally, ensure your environment is connected to [M365](https://corsobackup.io/docs/setup/m365-access/) and [S3](https://corsobackup.io/docs/setup/repos/). Refer to the [quick start guide](https://corsobackup.io/docs/quickstart/) for detailed instructions.

### Setting up Environment Variables

For sanity tests, configure the following environment variables:

#### For Restore:
- `SANITY_TEST_SOURCE_CONTAINER`: Folder where the test data was created in OneDrive.
- `SANITY_TEST_RESTORE_CONTAINER`: Folder where the test data will be restored in OneDrive.
- `SANITY_BACKUP_ID`: ID of the Corso backup used for restore.

#### For Export:
- `SANITY_TEST_SOURCE_CONTAINER`: Folder where the test data was created.
- `SANITY_TEST_RESTORE_CONTAINER`: Location of the exported data (local folder).
- `SANITY_BACKUP_ID`: ID of the Corso backup used for restore.

### Running Tests

Once your environment is set up, execute the tests with the following command:

```bash
./sanity-test <restore|export> <service>
```

## CI run

CI run is defined in `.github/workflows/sanity-test.yaml`. Check the
[Actions](https://github.com/alcionai/corso/actions/workflows/sanity-test.yaml)
 tab in this repository for detailed CI run logs.
