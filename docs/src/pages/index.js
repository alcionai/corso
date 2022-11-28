import React from "react";
import Layout from "@theme/Layout";
import { MainComp } from "@site/src/components/parts/MainComp";

export default function Home() {
  return (
    <Layout
      title="Free, Secure, and Open-Source Backup for Microsoft 365"
      description="Intro, docs, and blog for Corso, an open-source tool, that protects Microsoft 365 data by securely and efficiently backing up all business-critical data to object storage."
    >
      <MainComp />
    </Layout>
  );
}
