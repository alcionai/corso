---
description: "Core Corso concepts."
---

# Concepts

Before following this guide, it is important to familiarize yourself with some of the key Corso concepts.

## Microsoft 365 {#m365-concepts}

* **M365 Tenant** is typically associated with a unique domain (e.g. `contoso.com) and represents a dedicated
and logically segregated instance of the Microsoft 365 services plus associated data available to your organization.

* **M365 Service** refer to a cloud-applications available through the Microsoft 365 platform. Corso currently supports
backup and recovery for Exchange Online, OneDrive, SharePoint, and Teams.

* **Azure AD Application** represents an Azure AD digital identity/service pricipal and associated configuration which
define the resources that can be accessed and the actions that can be taken on these resources by the application.
Corso uses an Azure AD application to connect to your *M365 tenant* and transfer data during backup and restore operations.

## Corso {#corso-concepts}

* **Repository** refers to the storage location where Corso securely and efficiently stores encrypted *backups* of your
*M365 Services* data.

* **Backup** is a copy of your *M365 Services* data that can be restored if the original data is deleted, lost, or
corrupted. Corso performs backups incrementally; each backup only captures data that has changed between backup iterations.
