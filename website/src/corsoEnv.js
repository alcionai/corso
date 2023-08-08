import React from 'react';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';

export const Version = () => {
  const {siteConfig} = useDocusaurusContext();
  return siteConfig.customFields.corsoVersion;
}
