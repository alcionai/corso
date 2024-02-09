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
        slug: 'cli/canario',
        description: 'Explore the commonly-used Canario CLI commands',
        type: 'generated-index',
      },
      items: [
        {
          type: 'category',
          label: 'Setup and maintenance',
          link: {
            slug: 'cli/setup',
            description: 'Documentation for Canario setup and maintenance commands',
          },
          items: [
            'cli/canario-repo-init-s3',
            'cli/canario-repo-connect-s3',
            'cli/canario-repo-init-filesystem',
            'cli/canario-repo-connect-filesystem',
            'cli/canario-repo-maintenance',
            'cli/canario-repo-update-passphrase',
            'cli/canario-env']
        },
        {
          type: 'category',
          label: 'Exchange',
          link: {
            slug: 'cli/exchange',
            description: 'Documentation for commonly-used Canario Exchange CLI commands',
          },
          items: [
            'cli/canario-backup-create-exchange',
            'cli/canario-backup-list-exchange',
            'cli/canario-backup-details-exchange',
            'cli/canario-backup-delete-exchange',
            'cli/canario-restore-exchange',
            'cli/canario-export-exchange']
        },
        {
          type: 'category',
          label: 'Groups & Teams',
          link: {
            slug: 'cli/groups',
            description: 'Documentation for commonly-used Canario Groups & Teams CLI commands',
          },
          items: [
            'cli/canario-backup-create-groups',
            'cli/canario-backup-list-groups',
            'cli/canario-backup-details-groups',
            'cli/canario-backup-delete-groups',
            'cli/canario-restore-groups',
            'cli/canario-export-groups']
        },
        {
          type: 'category',
          label: 'OneDrive',
          link: {
            slug: 'cli/onedrive',
            description: 'Documentation for commonly-used Canario OneDrive CLI commands',
          },
          items: [
            'cli/canario-backup-create-onedrive',
            'cli/canario-backup-list-onedrive',
            'cli/canario-backup-details-onedrive',
            'cli/canario-backup-delete-onedrive',
            'cli/canario-restore-onedrive',
            'cli/canario-export-onedrive']
        },
        {
          type: 'category',
          label: 'SharePoint',
          link: {
            slug: 'cli/sharepoint',
            description: 'Documentation for commonly-used Canario SharePoint CLI commands',
          },
          items: [
            'cli/canario-backup-create-sharepoint',
            'cli/canario-backup-list-sharepoint',
            'cli/canario-backup-details-sharepoint',
            'cli/canario-backup-delete-sharepoint',
            'cli/canario-restore-sharepoint',
            'cli/canario-export-sharepoint']
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
