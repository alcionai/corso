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
    "intro",
    "quickstart",
    // {
    //   type: 'category',
    //   label: 'Corso setup',
    //   items: ['setup/concepts', 'setup/download', 'setup/m365_access', 'setup/configuration', 'setup/repos'],
    // },
    // {
    //   type: 'category',
    //   label: 'Command line reference',
    //   link: {
    //     slug: 'cli/corso',
    //     description: 'Explore the commonly used Corso CLI commands',
    //     type: 'generated-index',
    //   },
    //   items: [
    //     'cli/corso_repo_init_s3', 'cli/corso_repo_connect_s3',
    //     'cli/corso_backup_create_exchange', 'cli/corso_backup_list_exchange', 'cli/corso_backup_details_exchange',
    //     'cli/corso_backup_create_onedrive', 'cli/corso_backup_list_onedrive', 'cli/corso_backup_details_onedrive',
    //     'cli/corso_restore_exchange', 'cli/corso_restore_onedrive',
    //     'cli/corso_env'
    //   ]
    // },
    // {
    //   type: 'category',
    //   label: 'Support',
    //   items: [
    //     'support/bugs_and_features', 'support/known_issues', 'support/faq'
    //   ],
    // },
    // {
    //   type: 'category',
    //   label: 'Developer guide',
    //   items: [
    //     'developers/build', 'developers/testing', 'developers/linters'
    //   ],
    // },
  ],
};

module.exports = sidebars;
