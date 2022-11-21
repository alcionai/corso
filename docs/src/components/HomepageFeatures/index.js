import React from 'react';
import clsx from 'clsx';
import styles from './styles.module.css';

const FeatureList = [
  {
    title: 'Secure',
    Svg: require('@site/static/img/security.svg').default,
    description: (
      <>
        Corso provides secure data backup that protects customers against accidental data loss, service provider downtime, and malicious threats including ransomware attacks.
      </>
    ),
  },
  {
    title: 'Robust',
    Svg: require('@site/static/img/data.svg').default,
    description: (
      <>
       Corso, purpose-built for M365 protection, provides easy-to-use comprehensive backup and restore workflows that reduce backup time, improve time-to-recovery, reduce admin overhead, and replace unreliable scripts or workarounds.
      </>
    ),
  },
  {
    title: 'Low Cost',
    Svg: require('@site/static/img/savings.svg').default,
    description: (
      <>
       Corso, a 100% open-source tool, provides a free alternative for cost-conscious teams. It further reduces storage costs by supporting flexible retention policies and efficiently compressing and deduplicating data before storing it in low-cost cloud object storage.
      </>
    ),
  },
];

function Feature({Svg, title, description}) {
  return (
    <div className={clsx('col col--4')}>
      <div className="text--center">
        <Svg className={styles.featureSvg} role="img" />
      </div>
      <div className="text--center padding-horiz--md">
        <h3>{title}</h3>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures() {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
