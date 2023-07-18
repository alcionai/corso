---
description: "Repository maintenance."
---

# Repository maintenance

Repository maintenance helps optimize the Corso repository as backups are created and possibly deleted by the user.  
Maintenance can also free up space by removing data no longer referenced by any backups from the repository.

It's safe to run maintenance concurrently with backup, restore, and backup deletion operations. However, it's not safe  
to run maintenance operations concurrently on the same repository. Corso uses file locks and the idea of a repository  
owner to try to detect concurrent maintenance operations.

## Repository owner

The repository owner is set to the user and hostname of the machine that runs maintenance on the repo the first time.  

If the user and hostname of the machine running maintenance can change, use either the `--force` flag or the `--user`  
and `--host` flags.

The `--force` flag updates the repository owner and runs maintenance.

The `--user` and `--host` flags act as if the given user/hostname owns the repository for the maintenance operation  
but doesn't update repo owner info.

*If any of these flags are passed the user must make sure no concurrent maintenance operations run on the same  
repository. Concurrent maintenance operations a repository may result in data loss.*

## Maintenance types

Corso allows for two different types of maintenance: `metadata` and `complete`.

Metadata maintenance runs quickly and optimizes indexing data. Complete maintenance takes more time but compacts data  
in backups and removes unreferenced data from the repository.

As Corso allows concurrent backups during maintenance, running complete maintenance immediately after deleting a  
backup may not result in a reduction of objects in the storage service Corso is backing up to.  

Deletion of old objects in the storage service depends on both wall-clock time and running maintenance.

Later maintenance runs on the repository will remove the data.
