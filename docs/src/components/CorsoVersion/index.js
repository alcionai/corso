import React from 'react';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';

export default function CorsoVersion() {
  const { siteConfig } = useDocusaurusContext();

  return (
    <span>{siteConfig.customFields.corsoVersion}</span>
  );
}
