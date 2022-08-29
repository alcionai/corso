// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

const lightCodeTheme = require('prism-react-renderer/themes/github');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: 'Corso Documentation',
  tagline: 'Because Backups Are Cool',
  url: 'https://corsobackup.io',
  baseUrl: '/',
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'throw',
  favicon: 'img/corso_logo.svg',

  // GitHub pages deployment config.
  // If you aren't using GitHub pages, you don't need these.
  organizationName: 'alcionai', // Usually your GitHub org/user name.
  projectName: 'corso', // Usually your repo name.

  // Even if you don't use internalization, you can use this field to set useful
  // metadata like html lang. For example, if your site is Chinese, you may want
  // to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },
  staticDirectories: ['static', 'public'],

  plugins: [
    require.resolve('docusaurus-plugin-image-zoom')
  ],

  presets: [
    [
      'classic',
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          routeBasePath: '/',
          sidebarPath: require.resolve('./sidebars.js'),
          remarkPlugins: [require('mdx-mermaid')],
          editUrl:
            'https://github.com/alcionai/corso/tree/main/docs',
        },
        blog: false,
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      }),
    ],
  ],

  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      navbar: {
        title: 'Corso',
        logo: {
          alt: 'Corso Logo',
          src: '/img/corso_logo.svg',
          srcDark: 'img/corso_logo_white.svg',
        },
        items: [
          {
            type: 'doc',
            docId: 'intro',
            position: 'left',
            label: 'Docs',
          },
          {
            href: 'https://github.com/alcionai/corso',
            label: 'GitHub',
            position: 'right',
          },
        ],
      },
      footer: {
        style: 'dark',
        links: [
          {
            title: 'Resources',
            items: [
              {
                label: 'Docs',
                to: '/intro',
              },
            ],
          },
          {
            title: 'Community',
            items: [
              {
                label: 'Discussions',
                href: 'https://github.com/alcionai/corso/discussions',
              },
              {
                label: 'Twitter',
                href: 'https://twitter.com/alcionhq',
              },
            ],
          },
          {
            title: 'More',
            items: [
              {
                label: 'GitHub',
                href: 'https://github.com/alcionai/corso',
              },
            ],
          },
        ],
        copyright: `Copyright © ${new Date().getFullYear()}`,
      },
      colorMode: {
        defaultMode: 'dark',
        disableSwitch: true,
        respectPrefersColorScheme: true,
      },

      zoom: {
        selector: '.markdown img',
        background: {
          light: 'rgb(255, 255, 255)',
          dark: 'rgb(50, 50, 50)'
        },
        // options you can specify via https://github.com/francoischalifour/medium-zoom#usage
        config: {
          margin: 24,
          background: '#242526',
          scrollOffset: 0,
        },
      },

      prism: {
        theme: lightCodeTheme,
        darkTheme: darkCodeTheme,
      },
    }),
};

module.exports = config;
