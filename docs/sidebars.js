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
    'install',
    'tutorial',
    {
      type: 'category',
      label: 'Initial Configuration',
      items: ['configuration/concepts', 'configuration/m365_access', 'configuration/repos'],
    },
    {
      type: 'category',
      label: 'Command Line Reference',
      items: [
        'cli/corso', 'cli/corso_repo_init_s3', 'cli/corso_repo_connect_s3',
        'cli/corso_backup_create_exchange', 'cli/corso_backup_list_exchange', 'cli/corso_backup_details_exchange',
        'cli/corso_restore_exchange', 'cli/corso_env'
      ]
    }, 
    {
      type: 'category',
      label: 'Developer Guide',
      items: [
        'developers/architecture', 'developers/build', 'developers/testing', 'developers/linters'
      ],
    }, 
  
  ],
};

module.exports = sidebars;
