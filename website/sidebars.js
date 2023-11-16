/**
 * Creating a sidebar enables you to:
 - create an ordered group of docs
 - render a sidebar for each doc of that group
 - provide next/previous navigation

 The sidebars can be generated from the filesystem, or explicitly defined here.

 Create as many sidebars as you want.
 */

// @ts-check

/** @type {import('@docusaurus/plugin-content-docs').SidebarsConfig} */
const sidebars = {
  // By default, Docusaurus generates a sidebar from the docs folder structure
  docsSidebar: [
    'intro',
    'quickstart',
    {
      type: 'category',
      label: 'Usage',
      items: [
        'setup/concepts',
        'setup/download',
        'setup/m365-access',
        'setup/configuration',
        'setup/repos',
        'setup/fault-tolerance',
        'setup/restore-options',
        'setup/maintenance'
      ],
    },
    {
      type: 'category',
      label: 'Command line reference',
      link: {
        slug: 'cli/corso',
        description: 'Explore the commonly-used Corso CLI commands',
        type: 'generated-index',
      },
      items: [
        {
          type: 'category',
          label: 'Setup and maintenance',
          link: {
            slug: 'cli/setup',
            description: 'Documentation for Corso setup and maintenance commands',
          },
          items: [
            'cli/corso-repo-init-s3',
            'cli/corso-repo-connect-s3',
            'cli/corso-repo-init-filesystem',
            'cli/corso-repo-connect-filesystem',
            'cli/corso-repo-maintenance',
            'cli/corso-repo-update-passphrase',
            'cli/corso-env']
        },
        {
          type: 'category',
          label: 'Exchange',
          link: {
            slug: 'cli/exchange',
            description: 'Documentation for commonly-used Corso Exchange CLI commands',
          },
          items: [
            'cli/corso-backup-create-exchange',
            'cli/corso-backup-list-exchange',
            'cli/corso-backup-details-exchange',
            'cli/corso-backup-delete-exchange',
            'cli/corso-restore-exchange',
            'cli/corso-export-exchange']
        },
        {
          type: 'category',
          label: 'Groups & Teams',
          link: {
            slug: 'cli/groups',
            description: 'Documentation for commonly-used Corso Groups & Teams CLI commands',
          },
          items: [
            'cli/corso-backup-create-groups',
            'cli/corso-backup-list-groups',
            'cli/corso-backup-details-groups',
            'cli/corso-backup-delete-groups',
            'cli/corso-restore-groups',
            'cli/corso-export-groups']
        },
        {
          type: 'category',
          label: 'OneDrive',
          link: {
            slug: 'cli/onedrive',
            description: 'Documentation for commonly-used Corso OneDrive CLI commands',
          },
          items: [
            'cli/corso-backup-create-onedrive',
            'cli/corso-backup-list-onedrive',
            'cli/corso-backup-details-onedrive',
            'cli/corso-backup-delete-onedrive',
            'cli/corso-restore-onedrive',
            'cli/corso-export-onedrive']
        },
        {
          type: 'category',
          label: 'SharePoint',
          link: {
            slug: 'cli/sharepoint',
            description: 'Documentation for commonly-used Corso SharePoint CLI commands',
          },
          items: [
            'cli/corso-backup-create-sharepoint',
            'cli/corso-backup-list-sharepoint',
            'cli/corso-backup-details-sharepoint',
            'cli/corso-backup-delete-sharepoint',
            'cli/corso-restore-sharepoint',
            'cli/corso-export-sharepoint']
        }
      ]
    },
    {
      type: 'category',
      label: 'Support',
      items: [
        'support/bugs-and-features', 'support/known-issues', 'support/faq'
      ],
    },
    {
      type: 'category',
      label: 'Developer guide',
      items: [
        'developers/build', 'developers/testing', 'developers/linters'
      ],
    },
  ],
};

module.exports = sidebars;
