---
description: "Core Corso concepts."
---

# Concepts

Before using Corso, it's important to familiarize yourself with some key concepts.

## Microsoft 365 concepts {#m365-concepts}

* **M365 Tenant** is typically associated with a unique domain (for example, `contoso.com`) and represents a dedicated
and logically segregated instance of the Microsoft 365 services plus associated data available to your organization.

* **M365 Service** refer to a cloud-applications available through the Microsoft 365 platform. Corso supports
backup and recovery for Exchange Online, OneDrive, SharePoint, and Teams.

* **Azure AD Application** represents an Azure AD digital identity/service principal and associated configuration which
define the accessible resources and permitted actions on these resources through the application. Corso uses an Azure AD
application to connect to your *M365 tenant* and transfer data during backup and restore operations.

## Corso concepts {#corso-concepts}

* **Repository** refers to the storage location where Corso securely and efficiently stores encrypted *backups* of your
*M365 Services* data. See [Repositories](../repos) for more information.

* **Backup** is a copy of your *M365 Services* data to be used for restores in case of deletion, loss, or corruption of the
original data. Corso performs backups incrementally, and each backup only captures data that has changed between backup iterations.
