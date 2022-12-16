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
      label: 'Corso setup',
      items: ['setup/concepts', 'setup/download', 'setup/m365-access', 'setup/configuration', 'setup/repos'],
    },
    {
      type: 'category',
      label: 'Command line reference',
      link: {
        slug: 'cli/corso',
        description: 'Explore the commonly used Corso CLI commands',
        type: 'generated-index',
      },
      items: [
        'cli/corso-repo-init-s3', 'cli/corso-repo-connect-s3',
        'cli/corso-backup-create-exchange', 'cli/corso-backup-list-exchange', 'cli/corso-backup-details-exchange',
        'cli/corso-backup-create-onedrive', 'cli/corso-backup-list-onedrive', 'cli/corso-backup-details-onedrive',
        'cli/corso-restore-exchange', 'cli/corso-restore-onedrive',
        'cli/corso-env'
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