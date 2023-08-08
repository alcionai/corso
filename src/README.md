# SRC Directory

## /pkg
API and Components which are exposed for external usage.

* /pkg/repository  
Control layer for coordinating connections and communication with storage provider repositories.

* /pkg/storage  
Manages compilation and validation of repository configuration and consts.  Both those that are specific to storage providers, and those that are provider-agnostic.

-----

## /cli
Command Line Interface controller.  Utilizes /pkg/repository as an external dependency.

-----

## /internal
Packages which are only intended for use within Corso.